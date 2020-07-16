package eth_utility

import (
	"fmt"
	"math/big"
)

type Unit int

const (
	Eth Unit = iota
	Gwei
	Wei
)

var (
	gwei = big.NewFloat(1000000000)
	wei  = big.NewFloat(1000000000000000000)
)

func convertValueToBigFloat(fValue interface{}) *big.Float {
	switch v := fValue.(type) {
	case int:
		return new(big.Float).SetInt64(int64(v))
	case int64:
		return new(big.Float).SetInt64(v)
	case uint64:
		return new(big.Float).SetUint64(v)
	case string:
		return getBigFloatFromString(v)
	case *big.Float:
		return v
	case *big.Int:
		return new(big.Float).SetInt(v)
	default:
		fmt.Println("type not found")
		return big.NewFloat(0)
	}
}

// Converting Eth2 choose unit (wei, gwei, eth)
// return big.Float with result
func Eth2(unit Unit, value interface{}) *big.Float {
	v := convertValueToBigFloat(value)
	switch unit {
	case Gwei:
		return new(big.Float).Mul(v, gwei)
	case Wei:
		return new(big.Float).Mul(v, wei)
	}
	return v
}

// Converting Gwei2 choose unit (wei, gwei, eth)
// return big.Float with result
func Gwei2(unit Unit, value interface{}) *big.Float {
	v := convertValueToBigFloat(value)
	switch unit {
	case Eth:
		return new(big.Float).Quo(v, gwei)
	case Wei:
		return new(big.Float).Mul(v, gwei)
	}
	return v
}

// Converting Wei2 choose unit (wei, gwei, eth)
// return big.Float with result
func Wei2(unit Unit, value interface{}) *big.Float {
	v := convertValueToBigFloat(value)
	switch unit {
	case Eth:
		return new(big.Float).Quo(v, wei)
	case Gwei:
		return new(big.Float).Quo(v, gwei)
	}
	return v
}
