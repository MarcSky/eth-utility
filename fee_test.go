package eth_utility

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateFee(t *testing.T) {
	t.Parallel()

	type param struct {
		gasUsed  interface{}
		gasPrice interface{}
		expected *big.Int
	}

	params := []param{
		{0x5208, "0xb2d05e00", new(big.Int).SetInt64(63000000000000)},
		{0x5208, "0x165a0bc00", new(big.Int).SetInt64(126000000000000)},
		{0xea60, "0x1dcd65000", new(big.Int).SetInt64(480000000000000)},
		{0xea60, "8000000000", new(big.Int).SetInt64(480000000000000)},
		{"21000", "13000000000", new(big.Int).SetInt64(273000000000000)},
		{"0xd24a", "0x12a05f2000", new(big.Int).SetInt64(4306720000000000)},
	}

	for _, p := range params {
		res := CalculateFee(p.gasUsed, p.gasPrice)
		assert.Equal(t, p.expected, res)
	}
}
