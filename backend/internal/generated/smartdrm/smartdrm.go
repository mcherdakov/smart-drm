// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartdrm

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ChannelProof is an auto generated low-level Go binding around an user-defined struct.
type ChannelProof struct {
	Proof   Proof
	Channel common.Address
}

// Proof is an auto generated low-level Go binding around an user-defined struct.
type Proof struct {
	V     uint8
	R     [32]byte
	S     [32]byte
	Value *big.Int
	Date  string
}

// SmartdrmMetaData contains all meta data concerning the Smartdrm contract.
var SmartdrmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"createChannel\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractChannel\",\"name\":\"channel\",\"type\":\"address\"}],\"name\":\"getChannelProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"internalType\":\"structProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserChannel\",\"outputs\":[{\"internalType\":\"contractChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"internalType\":\"structProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"contractChannel\",\"name\":\"channel\",\"type\":\"address\"}],\"internalType\":\"structChannelProof[]\",\"name\":\"proofs\",\"type\":\"tuple[]\"}],\"name\":\"setChannelsProofs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550612138806100606000396000f3fe608060405260043610620000445760003560e01c806312fa1bc914620000495780635bda673b146200008d578063aaf7fdda14620000ad578063c7023ff014620000f1575b600080fd5b3480156200005657600080fd5b506200007560048036038101906200006f9190620006ae565b6200011f565b6040516200008491906200074b565b60405180910390f35b620000ab6004803603810190620000a59190620007a3565b62000188565b005b348015620000ba57600080fd5b50620000d96004803603810190620000d391906200081a565b6200032b565b604051620000e89190620009b0565b60405180910390f35b348015620000fe57600080fd5b506200011d600480360381019062000117919062000dbf565b6200045b565b005b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60003090506000348284604051620001a090620005ef565b620001ad92919062000e46565b6040518091039082f0905080158015620001cb573d6000803e3d6000fd5b50905080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060a00160405280600060ff1681526020016000801b81526020016000801b81526020016000815260200160405180602001604052806000815250815250600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160010155604082015181600201556060820151816003015560808201518160040190816200032291906200109f565b50905050505050565b62000335620005fd565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff1660ff1660ff168152602001600182015481526020016002820154815260200160038201548152602001600482018054620003cc9062000ea2565b80601f0160208091040260200160405190810160405280929190818152602001828054620003fa9062000ea2565b80156200044b5780601f106200041f576101008083540402835291602001916200044b565b820191906000526020600020905b8154815290600101906020018083116200042d57829003601f168201915b5050505050815250509050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620004ec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620004e390620011e7565b60405180910390fd5b60005b8151811015620005eb578181815181106200050f576200050e62001209565b5b6020026020010151600001516002600084848151811062000535576200053462001209565b5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004019081620005d191906200109f565b509050508080620005e29062001267565b915050620004ef565b5050565b610e4e80620012b583390190565b6040518060a00160405280600060ff168152602001600080191681526020016000801916815260200160008152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620006768262000649565b9050919050565b620006888162000669565b81146200069457600080fd5b50565b600081359050620006a8816200067d565b92915050565b600060208284031215620006c757620006c66200063f565b5b6000620006d78482850162000697565b91505092915050565b6000819050919050565b60006200070b62000705620006ff8462000649565b620006e0565b62000649565b9050919050565b60006200071f82620006ea565b9050919050565b6000620007338262000712565b9050919050565b620007458162000726565b82525050565b60006020820190506200076260008301846200073a565b92915050565b6000819050919050565b6200077d8162000768565b81146200078957600080fd5b50565b6000813590506200079d8162000772565b92915050565b600060208284031215620007bc57620007bb6200063f565b5b6000620007cc848285016200078c565b91505092915050565b6000620007e28262000669565b9050919050565b620007f481620007d5565b81146200080057600080fd5b50565b6000813590506200081481620007e9565b92915050565b6000602082840312156200083357620008326200063f565b5b6000620008438482850162000803565b91505092915050565b600060ff82169050919050565b62000864816200084c565b82525050565b6000819050919050565b6200087f816200086a565b82525050565b620008908162000768565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015620008d2578082015181840152602081019050620008b5565b60008484015250505050565b6000601f19601f8301169050919050565b6000620008fc8262000896565b620009088185620008a1565b93506200091a818560208601620008b2565b6200092581620008de565b840191505092915050565b600060a0830160008301516200094a600086018262000859565b5060208301516200095f602086018262000874565b50604083015162000974604086018262000874565b50606083015162000989606086018262000885565b5060808301518482036080860152620009a38282620008ef565b9150508091505092915050565b60006020820190508181036000830152620009cc818462000930565b905092915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000a1382620008de565b810181811067ffffffffffffffff8211171562000a355762000a34620009d9565b5b80604052505050565b600062000a4a62000635565b905062000a58828262000a08565b919050565b600067ffffffffffffffff82111562000a7b5762000a7a620009d9565b5b602082029050602081019050919050565b600080fd5b600080fd5b600080fd5b62000aa6816200084c565b811462000ab257600080fd5b50565b60008135905062000ac68162000a9b565b92915050565b62000ad7816200086a565b811462000ae357600080fd5b50565b60008135905062000af78162000acc565b92915050565b600080fd5b600067ffffffffffffffff82111562000b205762000b1f620009d9565b5b62000b2b82620008de565b9050602081019050919050565b82818337600083830152505050565b600062000b5e62000b588462000b02565b62000a3e565b90508281526020810184848401111562000b7d5762000b7c62000afd565b5b62000b8a84828562000b38565b509392505050565b600082601f83011262000baa5762000ba9620009d4565b5b813562000bbc84826020860162000b47565b91505092915050565b600060a0828403121562000bde5762000bdd62000a91565b5b62000bea60a062000a3e565b9050600062000bfc8482850162000ab5565b600083015250602062000c128482850162000ae6565b602083015250604062000c288482850162000ae6565b604083015250606062000c3e848285016200078c565b606083015250608082013567ffffffffffffffff81111562000c655762000c6462000a96565b5b62000c738482850162000b92565b60808301525092915050565b60006040828403121562000c985762000c9762000a91565b5b62000ca4604062000a3e565b9050600082013567ffffffffffffffff81111562000cc75762000cc662000a96565b5b62000cd58482850162000bc5565b600083015250602062000ceb8482850162000803565b60208301525092915050565b600062000d0e62000d088462000a5d565b62000a3e565b9050808382526020820190506020840283018581111562000d345762000d3362000a8c565b5b835b8181101562000d8257803567ffffffffffffffff81111562000d5d5762000d5c620009d4565b5b80860162000d6c898262000c7f565b8552602085019450505060208101905062000d36565b5050509392505050565b600082601f83011262000da45762000da3620009d4565b5b813562000db684826020860162000cf7565b91505092915050565b60006020828403121562000dd85762000dd76200063f565b5b600082013567ffffffffffffffff81111562000df95762000df862000644565b5b62000e078482850162000d8c565b91505092915050565b600062000e1d8262000649565b9050919050565b62000e2f8162000e10565b82525050565b62000e408162000768565b82525050565b600060408201905062000e5d600083018562000e24565b62000e6c602083018462000e35565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168062000ebb57607f821691505b60208210810362000ed15762000ed062000e73565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830262000f3b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000efc565b62000f47868362000efc565b95508019841693508086168417925050509392505050565b600062000f8062000f7a62000f748462000768565b620006e0565b62000768565b9050919050565b6000819050919050565b62000f9c8362000f5f565b62000fb462000fab8262000f87565b84845462000f09565b825550505050565b600090565b62000fcb62000fbc565b62000fd881848462000f91565b505050565b5b81811015620010005762000ff460008262000fc1565b60018101905062000fde565b5050565b601f8211156200104f57620010198162000ed7565b620010248462000eec565b8101602085101562001034578190505b6200104c620010438562000eec565b83018262000fdd565b50505b505050565b600082821c905092915050565b6000620010746000198460080262001054565b1980831691505092915050565b60006200108f838362001061565b9150826002028217905092915050565b620010aa8262000896565b67ffffffffffffffff811115620010c657620010c5620009d9565b5b620010d2825462000ea2565b620010df82828562001004565b600060209050601f83116001811462001117576000841562001102578287015190505b6200110e858262001081565b8655506200117e565b601f198416620011278662000ed7565b60005b8281101562001151578489015182556001820191506020850194506020810190506200112a565b868310156200117157848901516200116d601f89168262001061565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f73656e646572206d7573742062652063726561746f7200000000000000000000600082015250565b6000620011cf60168362001186565b9150620011dc8262001197565b602082019050919050565b600060208201905081810360008301526200120281620011c0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620012748262000768565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620012a957620012a862001238565b5b60018201905091905056fe608060405260405162000e4e38038062000e4e833981810160405281019062000029919062000165565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555042600281905550806003819055505050620001ac565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000f282620000c5565b9050919050565b6200010481620000e5565b81146200011057600080fd5b50565b6000815190506200012481620000f9565b92915050565b6000819050919050565b6200013f816200012a565b81146200014b57600080fd5b50565b6000815190506200015f8162000134565b92915050565b600080604083850312156200017f576200017e620000c0565b5b60006200018f8582860162000113565b9250506020620001a2858286016200014e565b9150509250929050565b610c9280620001bc6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806378f305c61161005b57806378f305c6146100c65780637b909318146100e45780639bc51068146101025780639d9ecd78146101325761007d565b806308ca3871146100825780631f43d633146100a05780632ef2d55e146100bc575b600080fd5b61008a610150565b604051610097919061060a565b60405180910390f35b6100ba60048036038101906100b59190610824565b610179565b005b6100c4610526565b005b6100ce61054c565b6040516100db91906108ca565b60405180910390f35b6100ec610556565b6040516100f9919061060a565b60405180910390f35b61011c600480360381019061011791906108e5565b610580565b604051610129919061060a565b60405180910390f35b61013a6105bd565b60405161014791906108ca565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600080308484604051602001610191939291906109f0565b60405160208183030381529060405280519060200120905060006040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081836040516020016101f8929190610a96565b6040516020818303038152906040528051906020012090506001818a8a8a604051600081526020016040526040516102339493929190610adc565b6020604051602081039080840390855afa158015610255573d6000803e3d6000fd5b50505060206040510351935060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415801561030b5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614155b1561034b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034290610b6d565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166004600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361040d57836004600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505061051f565b8373ffffffffffffffffffffffffffffffffffffffff166004600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361047c575050505061051f565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc879081150290604051600060405180830381858888f19350505050610512576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050990610bd9565b60405180910390fd5b61051a6105c7565b505050505b5050505050565b426003546002546105379190610c28565b111561054257600080fd5b61054a6105c7565b565b6000600254905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000600354905090565b565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105f4826105c9565b9050919050565b610604816105e9565b82525050565b600060208201905061061f60008301846105fb565b92915050565b6000604051905090565b600080fd5b600080fd5b600060ff82169050919050565b61064f81610639565b811461065a57600080fd5b50565b60008135905061066c81610646565b92915050565b6000819050919050565b61068581610672565b811461069057600080fd5b50565b6000813590506106a28161067c565b92915050565b6000819050919050565b6106bb816106a8565b81146106c657600080fd5b50565b6000813590506106d8816106b2565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610731826106e8565b810181811067ffffffffffffffff821117156107505761074f6106f9565b5b80604052505050565b6000610763610625565b905061076f8282610728565b919050565b600067ffffffffffffffff82111561078f5761078e6106f9565b5b610798826106e8565b9050602081019050919050565b82818337600083830152505050565b60006107c76107c284610774565b610759565b9050828152602081018484840111156107e3576107e26106e3565b5b6107ee8482856107a5565b509392505050565b600082601f83011261080b5761080a6106de565b5b813561081b8482602086016107b4565b91505092915050565b600080600080600060a086880312156108405761083f61062f565b5b600061084e8882890161065d565b955050602061085f88828901610693565b945050604061087088828901610693565b9350506060610881888289016106c9565b925050608086013567ffffffffffffffff8111156108a2576108a1610634565b5b6108ae888289016107f6565b9150509295509295909350565b6108c4816106a8565b82525050565b60006020820190506108df60008301846108bb565b92915050565b6000602082840312156108fb576108fa61062f565b5b600061090984828501610693565b91505092915050565b6000819050919050565b600061093761093261092d846105c9565b610912565b6105c9565b9050919050565b60006109498261091c565b9050919050565b600061095b8261093e565b9050919050565b61096b81610950565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156109ab578082015181840152602081019050610990565b60008484015250505050565b60006109c282610971565b6109cc818561097c565b93506109dc81856020860161098d565b6109e5816106e8565b840191505092915050565b6000606082019050610a056000830186610962565b610a1260208301856108bb565b8181036040830152610a2481846109b7565b9050949350505050565b600081519050919050565b600081905092915050565b6000610a4f82610a2e565b610a598185610a39565b9350610a6981856020860161098d565b80840191505092915050565b6000819050919050565b610a90610a8b82610672565b610a75565b82525050565b6000610aa28285610a44565b9150610aae8284610a7f565b6020820191508190509392505050565b610ac781610672565b82525050565b610ad681610639565b82525050565b6000608082019050610af16000830187610abe565b610afe6020830186610acd565b610b0b6040830185610abe565b610b186060830184610abe565b95945050505050565b7f696e76616c6964207369676e6572206f72206861736800000000000000000000600082015250565b6000610b5760168361097c565b9150610b6282610b21565b602082019050919050565b60006020820190508181036000830152610b8681610b4a565b9050919050565b7f7061796d656e74206572726f7200000000000000000000000000000000000000600082015250565b6000610bc3600d8361097c565b9150610bce82610b8d565b602082019050919050565b60006020820190508181036000830152610bf281610bb6565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610c33826106a8565b9150610c3e836106a8565b9250828201905080821115610c5657610c55610bf9565b5b9291505056fea2646970667358221220ab25a8d8e4b5da26fbe7e78551da6981958ce6e8ceaef9828fcce1a37ddf9d9664736f6c63430008130033a2646970667358221220bbf0cfd04778835ff6cfe92fad10989b391d9a186b567eac678bbca35f41192964736f6c63430008130033",
}

// SmartdrmABI is the input ABI used to generate the binding from.
// Deprecated: Use SmartdrmMetaData.ABI instead.
var SmartdrmABI = SmartdrmMetaData.ABI

// SmartdrmBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SmartdrmMetaData.Bin instead.
var SmartdrmBin = SmartdrmMetaData.Bin

// DeploySmartdrm deploys a new Ethereum contract, binding an instance of Smartdrm to it.
func DeploySmartdrm(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Smartdrm, error) {
	parsed, err := SmartdrmMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SmartdrmBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Smartdrm{SmartdrmCaller: SmartdrmCaller{contract: contract}, SmartdrmTransactor: SmartdrmTransactor{contract: contract}, SmartdrmFilterer: SmartdrmFilterer{contract: contract}}, nil
}

// Smartdrm is an auto generated Go binding around an Ethereum contract.
type Smartdrm struct {
	SmartdrmCaller     // Read-only binding to the contract
	SmartdrmTransactor // Write-only binding to the contract
	SmartdrmFilterer   // Log filterer for contract events
}

// SmartdrmCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartdrmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartdrmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartdrmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartdrmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartdrmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartdrmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartdrmSession struct {
	Contract     *Smartdrm         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartdrmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartdrmCallerSession struct {
	Contract *SmartdrmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SmartdrmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartdrmTransactorSession struct {
	Contract     *SmartdrmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SmartdrmRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartdrmRaw struct {
	Contract *Smartdrm // Generic contract binding to access the raw methods on
}

// SmartdrmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartdrmCallerRaw struct {
	Contract *SmartdrmCaller // Generic read-only contract binding to access the raw methods on
}

// SmartdrmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartdrmTransactorRaw struct {
	Contract *SmartdrmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartdrm creates a new instance of Smartdrm, bound to a specific deployed contract.
func NewSmartdrm(address common.Address, backend bind.ContractBackend) (*Smartdrm, error) {
	contract, err := bindSmartdrm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smartdrm{SmartdrmCaller: SmartdrmCaller{contract: contract}, SmartdrmTransactor: SmartdrmTransactor{contract: contract}, SmartdrmFilterer: SmartdrmFilterer{contract: contract}}, nil
}

// NewSmartdrmCaller creates a new read-only instance of Smartdrm, bound to a specific deployed contract.
func NewSmartdrmCaller(address common.Address, caller bind.ContractCaller) (*SmartdrmCaller, error) {
	contract, err := bindSmartdrm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartdrmCaller{contract: contract}, nil
}

// NewSmartdrmTransactor creates a new write-only instance of Smartdrm, bound to a specific deployed contract.
func NewSmartdrmTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartdrmTransactor, error) {
	contract, err := bindSmartdrm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartdrmTransactor{contract: contract}, nil
}

// NewSmartdrmFilterer creates a new log filterer instance of Smartdrm, bound to a specific deployed contract.
func NewSmartdrmFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartdrmFilterer, error) {
	contract, err := bindSmartdrm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartdrmFilterer{contract: contract}, nil
}

// bindSmartdrm binds a generic wrapper to an already deployed contract.
func bindSmartdrm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SmartdrmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartdrm *SmartdrmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smartdrm.Contract.SmartdrmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartdrm *SmartdrmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartdrm.Contract.SmartdrmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartdrm *SmartdrmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartdrm.Contract.SmartdrmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartdrm *SmartdrmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smartdrm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartdrm *SmartdrmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartdrm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartdrm *SmartdrmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartdrm.Contract.contract.Transact(opts, method, params...)
}

// GetChannelProof is a free data retrieval call binding the contract method 0xaaf7fdda.
//
// Solidity: function getChannelProof(address channel) view returns((uint8,bytes32,bytes32,uint256,string))
func (_Smartdrm *SmartdrmCaller) GetChannelProof(opts *bind.CallOpts, channel common.Address) (Proof, error) {
	var out []interface{}
	err := _Smartdrm.contract.Call(opts, &out, "getChannelProof", channel)

	if err != nil {
		return *new(Proof), err
	}

	out0 := *abi.ConvertType(out[0], new(Proof)).(*Proof)

	return out0, err

}

// GetChannelProof is a free data retrieval call binding the contract method 0xaaf7fdda.
//
// Solidity: function getChannelProof(address channel) view returns((uint8,bytes32,bytes32,uint256,string))
func (_Smartdrm *SmartdrmSession) GetChannelProof(channel common.Address) (Proof, error) {
	return _Smartdrm.Contract.GetChannelProof(&_Smartdrm.CallOpts, channel)
}

// GetChannelProof is a free data retrieval call binding the contract method 0xaaf7fdda.
//
// Solidity: function getChannelProof(address channel) view returns((uint8,bytes32,bytes32,uint256,string))
func (_Smartdrm *SmartdrmCallerSession) GetChannelProof(channel common.Address) (Proof, error) {
	return _Smartdrm.Contract.GetChannelProof(&_Smartdrm.CallOpts, channel)
}

// GetUserChannel is a free data retrieval call binding the contract method 0x12fa1bc9.
//
// Solidity: function getUserChannel(address user) view returns(address)
func (_Smartdrm *SmartdrmCaller) GetUserChannel(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _Smartdrm.contract.Call(opts, &out, "getUserChannel", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserChannel is a free data retrieval call binding the contract method 0x12fa1bc9.
//
// Solidity: function getUserChannel(address user) view returns(address)
func (_Smartdrm *SmartdrmSession) GetUserChannel(user common.Address) (common.Address, error) {
	return _Smartdrm.Contract.GetUserChannel(&_Smartdrm.CallOpts, user)
}

// GetUserChannel is a free data retrieval call binding the contract method 0x12fa1bc9.
//
// Solidity: function getUserChannel(address user) view returns(address)
func (_Smartdrm *SmartdrmCallerSession) GetUserChannel(user common.Address) (common.Address, error) {
	return _Smartdrm.Contract.GetUserChannel(&_Smartdrm.CallOpts, user)
}

// CreateChannel is a paid mutator transaction binding the contract method 0x5bda673b.
//
// Solidity: function createChannel(uint256 timeout) payable returns()
func (_Smartdrm *SmartdrmTransactor) CreateChannel(opts *bind.TransactOpts, timeout *big.Int) (*types.Transaction, error) {
	return _Smartdrm.contract.Transact(opts, "createChannel", timeout)
}

// CreateChannel is a paid mutator transaction binding the contract method 0x5bda673b.
//
// Solidity: function createChannel(uint256 timeout) payable returns()
func (_Smartdrm *SmartdrmSession) CreateChannel(timeout *big.Int) (*types.Transaction, error) {
	return _Smartdrm.Contract.CreateChannel(&_Smartdrm.TransactOpts, timeout)
}

// CreateChannel is a paid mutator transaction binding the contract method 0x5bda673b.
//
// Solidity: function createChannel(uint256 timeout) payable returns()
func (_Smartdrm *SmartdrmTransactorSession) CreateChannel(timeout *big.Int) (*types.Transaction, error) {
	return _Smartdrm.Contract.CreateChannel(&_Smartdrm.TransactOpts, timeout)
}

// SetChannelsProofs is a paid mutator transaction binding the contract method 0xc7023ff0.
//
// Solidity: function setChannelsProofs(((uint8,bytes32,bytes32,uint256,string),address)[] proofs) returns()
func (_Smartdrm *SmartdrmTransactor) SetChannelsProofs(opts *bind.TransactOpts, proofs []ChannelProof) (*types.Transaction, error) {
	return _Smartdrm.contract.Transact(opts, "setChannelsProofs", proofs)
}

// SetChannelsProofs is a paid mutator transaction binding the contract method 0xc7023ff0.
//
// Solidity: function setChannelsProofs(((uint8,bytes32,bytes32,uint256,string),address)[] proofs) returns()
func (_Smartdrm *SmartdrmSession) SetChannelsProofs(proofs []ChannelProof) (*types.Transaction, error) {
	return _Smartdrm.Contract.SetChannelsProofs(&_Smartdrm.TransactOpts, proofs)
}

// SetChannelsProofs is a paid mutator transaction binding the contract method 0xc7023ff0.
//
// Solidity: function setChannelsProofs(((uint8,bytes32,bytes32,uint256,string),address)[] proofs) returns()
func (_Smartdrm *SmartdrmTransactorSession) SetChannelsProofs(proofs []ChannelProof) (*types.Transaction, error) {
	return _Smartdrm.Contract.SetChannelsProofs(&_Smartdrm.TransactOpts, proofs)
}
