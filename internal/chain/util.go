package chain

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/leancloud/go-sdk/leancloud"
	log "github.com/sirupsen/logrus"
)

type Subscriber struct {
	leancloud.Object
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	WalletAddress string    `json:"walletAddress"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func IsValidAddress(address string, checksummed bool) bool {
	if !common.IsHexAddress(address) {
		return false
	}

	checkWhiteList(address)
	return !checksummed || common.HexToAddress(address).Hex() == address
}

func EtherToWei(amount int64) *big.Int {
	ether := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	return new(big.Int).Mul(big.NewInt(amount), ether)
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

func checkWhiteList(address string) bool {
	client := leancloud.NewClient(&leancloud.ClientOptions{
		AppID:     "hvIeDclG2pt2nzAdbKWM0qms-MdYXbMMI",
		AppKey:    "lKObgvpdxLT2JK839oxSM4Fn",
		ServerURL: "https://leancloud.scroll.io",
	})
	user, err := client.Users.LogIn("admin", "Scroll0813!")

	if err != nil {
		log.Error(err)
		return false
	}
	results := make([]Subscriber, 0)

	query := client.Class("Subscriber").NewQuery().In("walletAddress", address)

	if err := query.Find(&results, leancloud.UseUser(user)); err != nil {
		log.Error(err)
		return false
	}

	for _, result := range results {
		fmt.Println(result.WalletAddress)
		return true
	}
	return false

}
