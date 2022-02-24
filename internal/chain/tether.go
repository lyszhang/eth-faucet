package chain

import (
	"github.com/tidwall/gjson"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type tokenContract struct {
	address common.Address
	abi abi.ABI
}

func newTokenContract(address string, jsondata []byte) (*tokenContract, error) {
    value := gjson.Get(string(jsondata), "abi")
	ai, err := abi.JSON(strings.NewReader(value.String()))
	if err != nil {
		return nil, err
	}
	return &tokenContract{
		address: common.HexToAddress(address),
		abi:     ai,
	}, nil
}

func(c *tokenContract)PackTransfer(to common.Address, value *big.Int) ([]byte, error) {
	return c.abi.Pack("transfer", to, value)
}