run:
	go run ./cmd/service/

gen-abi:
	solc --abi ../contract/contracts/SmartDRM.sol -o abi --overwrite && \
	solc --abi ../contract/contracts/Channel.sol -o abi --overwrite

gen: gen-abi
	abigen --abi=abi/SmartDRM.abi --pkg=smartdrm --out=internal/generated/smartdrm/smartdrm.go && \
	abigen --abi=abi/Channel.abi --pkg=channel --out=internal/generated/channel/channel.go
