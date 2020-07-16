package eth_utility

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

type convertParams struct {
	unit     Unit
	value    interface{}
	expected *big.Float
}

func TestEth2(t *testing.T) {
	t.Parallel()

	params := []convertParams{
		{Wei, 1, big.NewFloat(1000000000000000000)},
		{Wei, "1", big.NewFloat(1000000000000000000)},
		{Wei, "0x1", big.NewFloat(1000000000000000000)},
		{Wei, uint64(1), big.NewFloat(1000000000000000000)},
		{Wei, int64(1), big.NewFloat(1000000000000000000)},
		{Wei, new(big.Int).SetInt64(1), big.NewFloat(1000000000000000000)},
		{Wei, new(big.Float).SetInt64(1), big.NewFloat(1000000000000000000)},
		{Gwei, int64(1), big.NewFloat(1000000000)},
		{Eth, int64(1), big.NewFloat(1)},
	}

	for _, p := range params {
		res := Eth2(p.unit, p.value)
		v := res.Cmp(p.expected)
		assert.Equal(t, 0, v)
	}
}

func TestGwei2(t *testing.T) {
	t.Parallel()

	params := []convertParams{
		{Wei, 1000000000, big.NewFloat(1000000000000000000)},
		{Gwei, 1000000000, big.NewFloat(1000000000)},
		{Eth, 1000000000, big.NewFloat(1)},
	}

	for _, p := range params {
		res := Gwei2(p.unit, p.value)
		v := res.Cmp(p.expected)
		assert.Equal(t, 0, v)
	}
}

func TestWei2(t *testing.T) {
	t.Parallel()

	params := []convertParams{
		{Wei, 1000000000000000000, big.NewFloat(1000000000000000000)},
		{Gwei, 1000000000000000000, big.NewFloat(1000000000)},
		{Eth, 1000000000000000000, big.NewFloat(1)},
	}

	for _, p := range params {
		res := Wei2(p.unit, p.value)
		v := res.Cmp(p.expected)
		assert.Equal(t, 0, v)
	}
}
