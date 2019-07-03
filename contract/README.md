# Balance
## Deploy example
```
MNEMONIC="" truffle migrate --reset --network rinkeby
```
## Test
1. Parse addresses
```
MNEMONIC="" truffle exec scripts/ParseAddresses.js --network mainnet
```
2. Run tests
```
MNEMONIC="" truffle exec scripts/GetBalanceTest.js --network mainnet
```