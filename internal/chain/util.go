package chain

import (
	"context"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

func IsValidAddress(address string, checksummed bool) bool {
	if !common.IsHexAddress(address) {
		return false
	}
	return !checksummed || common.HexToAddress(address).Hex() == address
}

func EtherToWei(amount int64) *big.Int {
	ether := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	return new(big.Int).Mul(big.NewInt(amount), ether)
}

func getTokenContractAddress() string {
	return os.Getenv("USDC_CONTRACT")
}

func waitPendingTx(ctx context.Context, client *ethclient.Client, hash common.Hash) {
	sleep := 1000
	for {
		if _, ispending, _ := client.TransactionByHash(ctx, hash); !ispending {
			return
		}
		log.Debug("wait tx is not pending", "sleep ms", sleep)
		time.Sleep(time.Millisecond * time.Duration(sleep))
	}
}
