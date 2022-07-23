package chain

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/scroll-dev/eth-faucet/internal/chain/contract"
)

type TxBuilder interface {
	Sender() common.Address
	PackTransfer(ctx context.Context, to string, value *big.Int) (common.Hash, error)
	Transfer(ctx context.Context, to string, value *big.Int) (common.Hash, error)
	TransferERC20Token(ctx context.Context, to string, value *big.Int) (common.Hash, error)
}

type TxBuild struct {
	chainID *big.Int
	client  bind.ContractTransactor
	nonce   uint64
	mutex   sync.Mutex

	tokenAddr common.Address
	auth      *bind.TransactOpts
	token     *contract.ERC20BurnableMockSession
}

func NewTxBuilder(provider, erc20Token string, privateKey *ecdsa.PrivateKey, chainID *big.Int) (TxBuilder, error) {
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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	var (
		tokenAddr common.Address
		token     *contract.ERC20BurnableMock
	)
	if erc20Token != "" {
		tokenAddr = common.HexToAddress(erc20Token)
		token, err = contract.NewERC20BurnableMock(tokenAddr, client)
		if err != nil {
			return nil, err
		}
	} else {
		var tx *types.Transaction
		tokenAddr, tx, token, err = contract.DeployERC20BurnableMock(auth, client, "USDC coin", "USDC", auth.From, big.NewInt(0).Mul(big.NewInt(1e8), big.NewInt(1e18)))
		if err != nil {
			return nil, err
		}
		waitPendingTx(context.Background(), client, tx.Hash())
		log.Infof("Deploy erc20 contract %s successful", tokenAddr.String())
	}

	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, err
	}

	return &TxBuild{
		client:    client,
		nonce:     nonce,
		auth:      auth,
		tokenAddr: tokenAddr,
		token: &contract.ERC20BurnableMockSession{
			Contract: token,
			CallOpts: bind.CallOpts{
				Pending: true,
			},
			TransactOpts: *auth,
		},
	}, nil
}

func (b *TxBuild) Sender() common.Address {
	return b.auth.From
}

func (b *TxBuild) PackTransfer(ctx context.Context, to string, value *big.Int) (common.Hash, error) {
	// serialize transfers because we manually increase the nonce
	b.mutex.Lock()
	defer b.mutex.Unlock()

	txHash, err := b.TransferERC20Token(context.Background(), to, value)
	if err != nil {
		log.Errorf("send ERC20 token failed, err: %v", err.Error())
		return txHash, err
	}
	log.Infof("send ERC20 tx hash: %v", txHash.String())

	txHash, err = b.Transfer(ctx, to, value)
	if err != nil {
		log.Errorf("send ether failed, err: %v", err.Error())
		return txHash, err
	}
	log.Infof("send ether tx hash: %v", txHash.String())

	return txHash, err
}

func (b *TxBuild) Transfer(ctx context.Context, to string, value *big.Int) (common.Hash, error) {
	gasLimit := uint64(21000)
	gasPrice, err := b.client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, err
	}

	toAddress := common.HexToAddress(to)
	unsignedTx := types.NewTx(&types.LegacyTx{
		Nonce:    b.nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
	})

	signedTx, err := b.auth.Signer(b.auth.From, unsignedTx)
	if err != nil {
		return common.Hash{}, err
	}

	err = b.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Hash{}, err
	}
	b.nonce++

	return signedTx.Hash(), nil
}

func (b *TxBuild) TransferERC20Token(ctx context.Context, to string, value *big.Int) (common.Hash, error) {
	// update the nonce in the bind
	b.token.TransactOpts.Nonce = new(big.Int).SetUint64(b.nonce)
	tx, err := b.token.Transfer(common.HexToAddress(to), big.NewInt(0).Mul(big.NewInt(100), value))
	if err != nil {
		return common.Hash{}, err
	}
	b.nonce++
	return tx.Hash(), nil
}
