package chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scroll-dev/eth-faucet/internal/chain/contract"
)

type TxBuilder interface {
	Sender() common.Address
	PackTransfer(ctx context.Context, to string, value *big.Int) (common.Hash, error)
	Transfer(ctx context.Context, to string, value *big.Int) (common.Hash, error)
	TransferERC20Token(ctx context.Context, to string, value *big.Int) (common.Hash, error)
}

type TxBuild struct {
	client      bind.ContractTransactor
	privateKey  *ecdsa.PrivateKey
	signer      types.Signer
	fromAddress common.Address
}

func NewTxBuilder(provider string, privateKey *ecdsa.PrivateKey, chainID *big.Int) (TxBuilder, error) {
	client, err := ethclient.Dial(provider)
	if err != nil {
		return nil, err
	}

	if chainID == nil {
		chainID, err = client.ChainID(context.Background())
		if err != nil {
			return nil, err
		}
	}

	return &TxBuild{
		client:      client,
		privateKey:  privateKey,
		signer:      types.NewEIP2930Signer(chainID),
		fromAddress: crypto.PubkeyToAddress(privateKey.PublicKey),
	}, nil
}

func (b *TxBuild) Sender() common.Address {
	return b.fromAddress
}

func (b *TxBuild) PackTransfer(ctx context.Context, to string, value *big.Int) (common.Hash, error) {
	txHash, err := b.Transfer(ctx, to, value)
	go func() {
		txHash, err := b.TransferERC20Token(context.Background(), to, value)
		if err != nil {
			fmt.Printf("send ERC20 token failed, err: %s", err.Error())
		}
		fmt.Println("send ERC20 tx hash: ", txHash.String())
	}()
	return txHash, err
}

func (b *TxBuild) Transfer(ctx context.Context, to string, value *big.Int) (common.Hash, error) {
	nonce, err := b.client.PendingNonceAt(ctx, b.Sender())
	if err != nil {
		return common.Hash{}, err
	}

	gasLimit := uint64(21000)
	gasPrice, err := b.client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, err
	}

	toAddress := common.HexToAddress(to)
	unsignedTx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
	})

	signedTx, err := types.SignTx(unsignedTx, b.signer, b.privateKey)
	if err != nil {
		return common.Hash{}, err
	}

	return signedTx.Hash(), b.client.SendTransaction(ctx, signedTx)
}

func (b *TxBuild) TransferERC20Token(ctx context.Context, to string, value *big.Int) (common.Hash, error) {
	nonce, err := b.client.PendingNonceAt(ctx, b.Sender())
	if err != nil {
		return common.Hash{}, err
	}

	gasLimit := uint64(860000)
	gasPrice, err := b.client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, err
	}

	// token transfer data
	ctrJSON, _ := contract.Asset("TetherToken.json")
	ctr, err := newTokenContract(getTokenContractAddress(), ctrJSON)
	if err != nil {
		return common.Hash{}, err
	}
	txData, err := ctr.PackTransfer(common.HexToAddress(to), big.NewInt(100*1000))
	if err != nil {
		return common.Hash{}, err
	}

	// tx
	unsignedTx := types.NewTx(&types.AccessListTx{
		ChainID:  b.signer.ChainID(),
		Nonce:    nonce,
		To:       &ctr.address,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     txData,
	})

	signedTx, err := types.SignTx(unsignedTx, b.signer, b.privateKey)
	if err != nil {
		return common.Hash{}, err
	}

	return signedTx.Hash(), b.client.SendTransaction(ctx, signedTx)
}
