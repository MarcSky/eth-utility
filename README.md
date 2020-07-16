# eth-utility
A set of ready-made functions for convenient work with Ethereum Blockchain

## Calculate Fee
Function CalculateFee accepts parameters on input gasUsed (wei), gasPrice (wei) 
and return multiplication result fee (wei) in result.

```func CalculateFee(gasUsed interface{}, gasPrice interface{}) *big.Int```

Formula calculation fee in Ethereum Blockchain is very simple

```fee = gasUsed * gasPrice```

Example
```
gasUsed := 0x1
gasPrice := 0x1
fee := CalculateFee(gasUsed, gasPrice) // equal 2 wei
```

## Convert units
This library allows to easy convert different units (wei to eth, eth to gwei, etc...)

Function Eth2, Gwei2, Wei2 accepts params type unit and value as a result returns a converted value

```
func Eth2(unit Unit, value interface{}) *big.Float
func Gwei2(unit Unit, value interface{}) *big.Float
func Wei2(unit Unit, value interface{}) *big.Float
```  

P.S. this beta version and some features may not work correctly  