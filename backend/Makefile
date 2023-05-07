run:
	go run ./cmd/service/

gen-abi:
	solc --abi ../contract/contracts/SmartDRM.sol -o abi --overwrite && \
	solc --abi ../contract/contracts/Channel.sol -o abi --overwrite

gen-bytecode:
	solc --bin ../contract/contracts/Channel.sol -o bin --overwrite
	solc --bin ../contract/contracts/SmartDRM.sol -o bin --overwrite

gen: gen-abi
	abigen --abi=abi/SmartDRM.abi --pkg=smartdrm --out=internal/generated/smartdrm/smartdrm.go --bin bin/SmartDRM.bin && \
	abigen --abi=abi/Channel.abi --pkg=channel --out=internal/generated/channel/channel.go
