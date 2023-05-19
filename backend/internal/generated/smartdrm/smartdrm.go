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

// CreatorClicks is an auto generated low-level Go binding around an user-defined struct.
type CreatorClicks struct {
	Creator common.Address
	Clicks  uint32
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"closeChannel\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"createChannel\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractChannel\",\"name\":\"channel\",\"type\":\"address\"}],\"name\":\"getChannelProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"internalType\":\"structProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"getCreatorClicks\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserChannel\",\"outputs\":[{\"internalType\":\"contractChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"internalType\":\"structProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"contractChannel\",\"name\":\"channel\",\"type\":\"address\"}],\"internalType\":\"structChannelProof[]\",\"name\":\"proofs\",\"type\":\"tuple[]\"}],\"name\":\"setChannelsProofs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"clicks\",\"type\":\"uint32\"}],\"internalType\":\"structCreatorClicks[]\",\"name\":\"clicks\",\"type\":\"tuple[]\"}],\"name\":\"setCreatorsClicks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"splitBalance\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506126c8806100606000396000f3fe608060405260043610620000865760003560e01c8063aaf7fdda1162000055578063aaf7fdda146200014d578063b28aea561462000191578063b782cce214620001ab578063c7023ff014620001d95762000086565b806312fa1bc9146200008b578063449c847c14620000cf5780635bda673b14620000e957806374c9f3811462000109575b600080fd5b3480156200009857600080fd5b50620000b76004803603810190620000b1919062000a4f565b62000207565b604051620000c6919062000aec565b60405180910390f35b348015620000dc57600080fd5b50620000e762000270565b005b62000107600480360381019062000101919062000b44565b62000272565b005b3480156200011657600080fd5b506200013560048036038101906200012f919062000a4f565b62000415565b60405162000144919062000b97565b60405180910390f35b3480156200015a57600080fd5b5062000179600480360381019062000173919062000bf9565b6200046e565b60405162000188919062000d8f565b60405180910390f35b3480156200019e57600080fd5b50620001a96200059e565b005b348015620001b857600080fd5b50620001d76004803603810190620001d1919062000fab565b620005a0565b005b348015620001e657600080fd5b50620002056004803603810190620001ff91906200134f565b620007fc565b005b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b565b600030905060003482846040516200028a9062000990565b62000297929190620013d6565b6040518091039082f0905080158015620002b5573d6000803e3d6000fd5b50905080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060a00160405280600060ff1681526020016000801b81526020016000801b81526020016000815260200160405180602001604052806000815250815250600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160010155604082015181600201556060820151816003015560808201518160040190816200040c91906200162f565b50905050505050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff169050919050565b620004786200099e565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff1660ff1660ff1681526020016001820154815260200160028201548152602001600382015481526020016004820180546200050f9062001432565b80601f01602080910402602001604051908101604052809291908181526020018280546200053d9062001432565b80156200058e5780601f1062000562576101008083540402835291602001916200058e565b820191906000526020600020905b8154815290600101906020018083116200057057829003601f168201915b5050505050815250509050919050565b565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000631576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620006289062001777565b60405180910390fd5b60005b8151811015620007f8576000600460008484815181106200065a576200065962001799565b5b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff1663ffffffff160362000742576003828281518110620006d557620006d462001799565b5b6020026020010151600001519080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b81818151811062000758576200075762001799565b5b602002602001015160200151600460008484815181106200077e576200077d62001799565b5b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055508080620007ef90620017f7565b91505062000634565b5050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146200088d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620008849062001777565b60405180910390fd5b60005b81518110156200098c57818181518110620008b057620008af62001799565b5b60200260200101516000015160026000848481518110620008d657620008d562001799565b5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160010155604082015181600201556060820151816003015560808201518160040190816200097291906200162f565b5090505080806200098390620017f7565b91505062000890565b5050565b610e4e806200184583390190565b6040518060a00160405280600060ff168152602001600080191681526020016000801916815260200160008152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000a1782620009ea565b9050919050565b62000a298162000a0a565b811462000a3557600080fd5b50565b60008135905062000a498162000a1e565b92915050565b60006020828403121562000a685762000a67620009e0565b5b600062000a788482850162000a38565b91505092915050565b6000819050919050565b600062000aac62000aa662000aa084620009ea565b62000a81565b620009ea565b9050919050565b600062000ac08262000a8b565b9050919050565b600062000ad48262000ab3565b9050919050565b62000ae68162000ac7565b82525050565b600060208201905062000b03600083018462000adb565b92915050565b6000819050919050565b62000b1e8162000b09565b811462000b2a57600080fd5b50565b60008135905062000b3e8162000b13565b92915050565b60006020828403121562000b5d5762000b5c620009e0565b5b600062000b6d8482850162000b2d565b91505092915050565b600063ffffffff82169050919050565b62000b918162000b76565b82525050565b600060208201905062000bae600083018462000b86565b92915050565b600062000bc18262000a0a565b9050919050565b62000bd38162000bb4565b811462000bdf57600080fd5b50565b60008135905062000bf38162000bc8565b92915050565b60006020828403121562000c125762000c11620009e0565b5b600062000c228482850162000be2565b91505092915050565b600060ff82169050919050565b62000c438162000c2b565b82525050565b6000819050919050565b62000c5e8162000c49565b82525050565b62000c6f8162000b09565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b8381101562000cb157808201518184015260208101905062000c94565b60008484015250505050565b6000601f19601f8301169050919050565b600062000cdb8262000c75565b62000ce7818562000c80565b935062000cf981856020860162000c91565b62000d048162000cbd565b840191505092915050565b600060a08301600083015162000d29600086018262000c38565b50602083015162000d3e602086018262000c53565b50604083015162000d53604086018262000c53565b50606083015162000d68606086018262000c64565b506080830151848203608086015262000d82828262000cce565b9150508091505092915050565b6000602082019050818103600083015262000dab818462000d0f565b905092915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000df28262000cbd565b810181811067ffffffffffffffff8211171562000e145762000e1362000db8565b5b80604052505050565b600062000e29620009d6565b905062000e37828262000de7565b919050565b600067ffffffffffffffff82111562000e5a5762000e5962000db8565b5b602082029050602081019050919050565b600080fd5b600080fd5b600080fd5b62000e858162000b76565b811462000e9157600080fd5b50565b60008135905062000ea58162000e7a565b92915050565b60006040828403121562000ec45762000ec362000e70565b5b62000ed0604062000e1d565b9050600062000ee28482850162000a38565b600083015250602062000ef88482850162000e94565b60208301525092915050565b600062000f1b62000f158462000e3c565b62000e1d565b9050808382526020820190506040840283018581111562000f415762000f4062000e6b565b5b835b8181101562000f6e578062000f59888262000eab565b84526020840193505060408101905062000f43565b5050509392505050565b600082601f83011262000f905762000f8f62000db3565b5b813562000fa284826020860162000f04565b91505092915050565b60006020828403121562000fc45762000fc3620009e0565b5b600082013567ffffffffffffffff81111562000fe55762000fe4620009e5565b5b62000ff38482850162000f78565b91505092915050565b600067ffffffffffffffff8211156200101a576200101962000db8565b5b602082029050602081019050919050565b620010368162000c2b565b81146200104257600080fd5b50565b60008135905062001056816200102b565b92915050565b620010678162000c49565b81146200107357600080fd5b50565b60008135905062001087816200105c565b92915050565b600080fd5b600067ffffffffffffffff821115620010b057620010af62000db8565b5b620010bb8262000cbd565b9050602081019050919050565b82818337600083830152505050565b6000620010ee620010e88462001092565b62000e1d565b9050828152602081018484840111156200110d576200110c6200108d565b5b6200111a848285620010c8565b509392505050565b600082601f8301126200113a576200113962000db3565b5b81356200114c848260208601620010d7565b91505092915050565b600060a082840312156200116e576200116d62000e70565b5b6200117a60a062000e1d565b905060006200118c8482850162001045565b6000830152506020620011a28482850162001076565b6020830152506040620011b88482850162001076565b6040830152506060620011ce8482850162000b2d565b606083015250608082013567ffffffffffffffff811115620011f557620011f462000e75565b5b620012038482850162001122565b60808301525092915050565b60006040828403121562001228576200122762000e70565b5b62001234604062000e1d565b9050600082013567ffffffffffffffff81111562001257576200125662000e75565b5b620012658482850162001155565b60008301525060206200127b8482850162000be2565b60208301525092915050565b60006200129e620012988462000ffc565b62000e1d565b90508083825260208201905060208402830185811115620012c457620012c362000e6b565b5b835b818110156200131257803567ffffffffffffffff811115620012ed57620012ec62000db3565b5b808601620012fc89826200120f565b85526020850194505050602081019050620012c6565b5050509392505050565b600082601f83011262001334576200133362000db3565b5b81356200134684826020860162001287565b91505092915050565b600060208284031215620013685762001367620009e0565b5b600082013567ffffffffffffffff811115620013895762001388620009e5565b5b62001397848285016200131c565b91505092915050565b6000620013ad82620009ea565b9050919050565b620013bf81620013a0565b82525050565b620013d08162000b09565b82525050565b6000604082019050620013ed6000830185620013b4565b620013fc6020830184620013c5565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200144b57607f821691505b60208210810362001461576200146062001403565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620014cb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200148c565b620014d786836200148c565b95508019841693508086168417925050509392505050565b6000620015106200150a620015048462000b09565b62000a81565b62000b09565b9050919050565b6000819050919050565b6200152c83620014ef565b620015446200153b8262001517565b84845462001499565b825550505050565b600090565b6200155b6200154c565b6200156881848462001521565b505050565b5b8181101562001590576200158460008262001551565b6001810190506200156e565b5050565b601f821115620015df57620015a98162001467565b620015b4846200147c565b81016020851015620015c4578190505b620015dc620015d3856200147c565b8301826200156d565b50505b505050565b600082821c905092915050565b60006200160460001984600802620015e4565b1980831691505092915050565b60006200161f8383620015f1565b9150826002028217905092915050565b6200163a8262000c75565b67ffffffffffffffff81111562001656576200165562000db8565b5b62001662825462001432565b6200166f82828562001594565b600060209050601f831160018114620016a7576000841562001692578287015190505b6200169e858262001611565b8655506200170e565b601f198416620016b78662001467565b60005b82811015620016e157848901518255600182019150602085019450602081019050620016ba565b86831015620017015784890151620016fd601f891682620015f1565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f73656e646572206d757374206265206f776e6572000000000000000000000000600082015250565b60006200175f60148362001716565b91506200176c8262001727565b602082019050919050565b60006020820190508181036000830152620017928162001750565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620018048262000b09565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620018395762001838620017c8565b5b60018201905091905056fe608060405260405162000e4e38038062000e4e833981810160405281019062000029919062000165565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555042600281905550806003819055505050620001ac565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000f282620000c5565b9050919050565b6200010481620000e5565b81146200011057600080fd5b50565b6000815190506200012481620000f9565b92915050565b6000819050919050565b6200013f816200012a565b81146200014b57600080fd5b50565b6000815190506200015f8162000134565b92915050565b600080604083850312156200017f576200017e620000c0565b5b60006200018f8582860162000113565b9250506020620001a2858286016200014e565b9150509250929050565b610c9280620001bc6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806378f305c61161005b57806378f305c6146100c65780637b909318146100e45780639bc51068146101025780639d9ecd78146101325761007d565b806308ca3871146100825780631f43d633146100a05780632ef2d55e146100bc575b600080fd5b61008a610150565b604051610097919061060a565b60405180910390f35b6100ba60048036038101906100b59190610824565b610179565b005b6100c4610526565b005b6100ce61054c565b6040516100db91906108ca565b60405180910390f35b6100ec610556565b6040516100f9919061060a565b60405180910390f35b61011c600480360381019061011791906108e5565b610580565b604051610129919061060a565b60405180910390f35b61013a6105bd565b60405161014791906108ca565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600080308484604051602001610191939291906109f0565b60405160208183030381529060405280519060200120905060006040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081836040516020016101f8929190610a96565b6040516020818303038152906040528051906020012090506001818a8a8a604051600081526020016040526040516102339493929190610adc565b6020604051602081039080840390855afa158015610255573d6000803e3d6000fd5b50505060206040510351935060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415801561030b5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614155b1561034b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034290610b6d565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166004600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361040d57836004600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505061051f565b8373ffffffffffffffffffffffffffffffffffffffff166004600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361047c575050505061051f565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc879081150290604051600060405180830381858888f19350505050610512576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050990610bd9565b60405180910390fd5b61051a6105c7565b505050505b5050505050565b426003546002546105379190610c28565b111561054257600080fd5b61054a6105c7565b565b6000600254905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000600354905090565b565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105f4826105c9565b9050919050565b610604816105e9565b82525050565b600060208201905061061f60008301846105fb565b92915050565b6000604051905090565b600080fd5b600080fd5b600060ff82169050919050565b61064f81610639565b811461065a57600080fd5b50565b60008135905061066c81610646565b92915050565b6000819050919050565b61068581610672565b811461069057600080fd5b50565b6000813590506106a28161067c565b92915050565b6000819050919050565b6106bb816106a8565b81146106c657600080fd5b50565b6000813590506106d8816106b2565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610731826106e8565b810181811067ffffffffffffffff821117156107505761074f6106f9565b5b80604052505050565b6000610763610625565b905061076f8282610728565b919050565b600067ffffffffffffffff82111561078f5761078e6106f9565b5b610798826106e8565b9050602081019050919050565b82818337600083830152505050565b60006107c76107c284610774565b610759565b9050828152602081018484840111156107e3576107e26106e3565b5b6107ee8482856107a5565b509392505050565b600082601f83011261080b5761080a6106de565b5b813561081b8482602086016107b4565b91505092915050565b600080600080600060a086880312156108405761083f61062f565b5b600061084e8882890161065d565b955050602061085f88828901610693565b945050604061087088828901610693565b9350506060610881888289016106c9565b925050608086013567ffffffffffffffff8111156108a2576108a1610634565b5b6108ae888289016107f6565b9150509295509295909350565b6108c4816106a8565b82525050565b60006020820190506108df60008301846108bb565b92915050565b6000602082840312156108fb576108fa61062f565b5b600061090984828501610693565b91505092915050565b6000819050919050565b600061093761093261092d846105c9565b610912565b6105c9565b9050919050565b60006109498261091c565b9050919050565b600061095b8261093e565b9050919050565b61096b81610950565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156109ab578082015181840152602081019050610990565b60008484015250505050565b60006109c282610971565b6109cc818561097c565b93506109dc81856020860161098d565b6109e5816106e8565b840191505092915050565b6000606082019050610a056000830186610962565b610a1260208301856108bb565b8181036040830152610a2481846109b7565b9050949350505050565b600081519050919050565b600081905092915050565b6000610a4f82610a2e565b610a598185610a39565b9350610a6981856020860161098d565b80840191505092915050565b6000819050919050565b610a90610a8b82610672565b610a75565b82525050565b6000610aa28285610a44565b9150610aae8284610a7f565b6020820191508190509392505050565b610ac781610672565b82525050565b610ad681610639565b82525050565b6000608082019050610af16000830187610abe565b610afe6020830186610acd565b610b0b6040830185610abe565b610b186060830184610abe565b95945050505050565b7f696e76616c6964207369676e6572206f72206861736800000000000000000000600082015250565b6000610b5760168361097c565b9150610b6282610b21565b602082019050919050565b60006020820190508181036000830152610b8681610b4a565b9050919050565b7f7061796d656e74206572726f7200000000000000000000000000000000000000600082015250565b6000610bc3600d8361097c565b9150610bce82610b8d565b602082019050919050565b60006020820190508181036000830152610bf281610bb6565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610c33826106a8565b9150610c3e836106a8565b9250828201905080821115610c5657610c55610bf9565b5b9291505056fea2646970667358221220ab25a8d8e4b5da26fbe7e78551da6981958ce6e8ceaef9828fcce1a37ddf9d9664736f6c63430008130033a2646970667358221220a8d8bd3073e783ee05087033dd8d64e8dff29b03aa3bef2ad5ccac893420e57a64736f6c63430008130033",
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

// CloseChannel is a free data retrieval call binding the contract method 0xb28aea56.
//
// Solidity: function closeChannel() view returns()
func (_Smartdrm *SmartdrmCaller) CloseChannel(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Smartdrm.contract.Call(opts, &out, "closeChannel")

	if err != nil {
		return err
	}

	return err

}

// CloseChannel is a free data retrieval call binding the contract method 0xb28aea56.
//
// Solidity: function closeChannel() view returns()
func (_Smartdrm *SmartdrmSession) CloseChannel() error {
	return _Smartdrm.Contract.CloseChannel(&_Smartdrm.CallOpts)
}

// CloseChannel is a free data retrieval call binding the contract method 0xb28aea56.
//
// Solidity: function closeChannel() view returns()
func (_Smartdrm *SmartdrmCallerSession) CloseChannel() error {
	return _Smartdrm.Contract.CloseChannel(&_Smartdrm.CallOpts)
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

// GetCreatorClicks is a free data retrieval call binding the contract method 0x74c9f381.
//
// Solidity: function getCreatorClicks(address creator) view returns(uint32)
func (_Smartdrm *SmartdrmCaller) GetCreatorClicks(opts *bind.CallOpts, creator common.Address) (uint32, error) {
	var out []interface{}
	err := _Smartdrm.contract.Call(opts, &out, "getCreatorClicks", creator)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetCreatorClicks is a free data retrieval call binding the contract method 0x74c9f381.
//
// Solidity: function getCreatorClicks(address creator) view returns(uint32)
func (_Smartdrm *SmartdrmSession) GetCreatorClicks(creator common.Address) (uint32, error) {
	return _Smartdrm.Contract.GetCreatorClicks(&_Smartdrm.CallOpts, creator)
}

// GetCreatorClicks is a free data retrieval call binding the contract method 0x74c9f381.
//
// Solidity: function getCreatorClicks(address creator) view returns(uint32)
func (_Smartdrm *SmartdrmCallerSession) GetCreatorClicks(creator common.Address) (uint32, error) {
	return _Smartdrm.Contract.GetCreatorClicks(&_Smartdrm.CallOpts, creator)
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

// SplitBalance is a free data retrieval call binding the contract method 0x449c847c.
//
// Solidity: function splitBalance() view returns()
func (_Smartdrm *SmartdrmCaller) SplitBalance(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Smartdrm.contract.Call(opts, &out, "splitBalance")

	if err != nil {
		return err
	}

	return err

}

// SplitBalance is a free data retrieval call binding the contract method 0x449c847c.
//
// Solidity: function splitBalance() view returns()
func (_Smartdrm *SmartdrmSession) SplitBalance() error {
	return _Smartdrm.Contract.SplitBalance(&_Smartdrm.CallOpts)
}

// SplitBalance is a free data retrieval call binding the contract method 0x449c847c.
//
// Solidity: function splitBalance() view returns()
func (_Smartdrm *SmartdrmCallerSession) SplitBalance() error {
	return _Smartdrm.Contract.SplitBalance(&_Smartdrm.CallOpts)
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

// SetCreatorsClicks is a paid mutator transaction binding the contract method 0xb782cce2.
//
// Solidity: function setCreatorsClicks((address,uint32)[] clicks) returns()
func (_Smartdrm *SmartdrmTransactor) SetCreatorsClicks(opts *bind.TransactOpts, clicks []CreatorClicks) (*types.Transaction, error) {
	return _Smartdrm.contract.Transact(opts, "setCreatorsClicks", clicks)
}

// SetCreatorsClicks is a paid mutator transaction binding the contract method 0xb782cce2.
//
// Solidity: function setCreatorsClicks((address,uint32)[] clicks) returns()
func (_Smartdrm *SmartdrmSession) SetCreatorsClicks(clicks []CreatorClicks) (*types.Transaction, error) {
	return _Smartdrm.Contract.SetCreatorsClicks(&_Smartdrm.TransactOpts, clicks)
}

// SetCreatorsClicks is a paid mutator transaction binding the contract method 0xb782cce2.
//
// Solidity: function setCreatorsClicks((address,uint32)[] clicks) returns()
func (_Smartdrm *SmartdrmTransactorSession) SetCreatorsClicks(clicks []CreatorClicks) (*types.Transaction, error) {
	return _Smartdrm.Contract.SetCreatorsClicks(&_Smartdrm.TransactOpts, clicks)
}
