run:
	go run ./cmd/service/

gen-abi:
	solc --abi ../contract/contracts/SmartDRM.sol -o abi --overwrite

gen: gen-abi
	abigen --abi=abi/SmartDRM.abi --pkg=smartdrm --out=internal/generated/smartdrm/smartdrm.go
