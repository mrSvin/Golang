1. go get all
2. solcjs --optimize --abi --bin ./SimpleStorage.sol -o build
3. abigen --abi=./build/SimpleStorage_sol_SimpleStorage.abi --bin=./build/SimpleStorage_sol_SimpleStorage.bin --pkg=api --out=./api/SimpleStorage.go
