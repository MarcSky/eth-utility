package eth_utility

import (
	"fmt"
	"math/big"
)

func convertValueToBigInt(value interface{}) *big.Int {
	switch v := value.(type) {
	case int:
		return new(big.Int).SetInt64(int64(v))
	case int64:
		return new(big.Int).SetInt64(v)
	case uint64:
		return new(big.Int).SetUint64(v)
	case string:
		return getBigIntFromString(v)
	default:
		fmt.Println("type not found")
		return new(big.Int).SetInt64(0)
	}
}

// CalculateFee take values from gasUsed (wei) and gasPrice (wei)
// and return multiplication result in wei
func CalculateFee(gasUsed interface{}, gasPrice interface{}) *big.Int {
	gasU := convertValueToBigInt(gasUsed)
	gasP := convertValueToBigInt(gasPrice)
	return new(big.Int).Mul(gasU, gasP)
}
