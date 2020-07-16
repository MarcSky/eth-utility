package eth_utility

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func getBigIntFromString(value string) *big.Int {
	if strings.HasPrefix(value, "0x") {
		b, _ := hexutil.DecodeBig(value)
		return b
	}
	b, _ := new(big.Int).SetString(value, 10)
	return b
}

func getBigFloatFromString(value string) *big.Float {
	if strings.HasPrefix(value, "0x") {
		b, _ := hexutil.DecodeBig(value)
		return new(big.Float).SetInt(b)
	}
	b, _ := new(big.Float).SetString(value)
	return b
}
