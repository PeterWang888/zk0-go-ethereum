// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake

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

// ValidatorsetMetaData contains all meta data concerning the Validatorset contract.
var ValidatorsetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_initial\",\"type\":\"address[]\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_newSet\",\"type\":\"address[]\"}],\"name\":\"InitiateChange\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"addValidator2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakerAmounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorAmounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"removeValidator2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerChange\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405260405162001f2938038062001f298339818101604052810190620000299190620003fe565b60005b815181101562000128576001600260008484815181106200005257620000516200044f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff0219169083151502179055508060026000848481518110620000c957620000c86200044f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555080806200011f90620004b7565b9150506200002c565b5080600090805190602001906200014192919062000162565b5080600190805190602001906200015a92919062000162565b505062000504565b828054828255906000526020600020908101928215620001de579160200282015b82811115620001dd5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019062000183565b5b509050620001ed9190620001f1565b5090565b5b808211156200020c576000816000905550600101620001f2565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620002748262000229565b810181811067ffffffffffffffff821117156200029657620002956200023a565b5b80604052505050565b6000620002ab62000210565b9050620002b9828262000269565b919050565b600067ffffffffffffffff821115620002dc57620002db6200023a565b5b602082029050602081019050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200031f82620002f2565b9050919050565b620003318162000312565b81146200033d57600080fd5b50565b600081519050620003518162000326565b92915050565b60006200036e6200036884620002be565b6200029f565b90508083825260208201905060208402830185811115620003945762000393620002ed565b5b835b81811015620003c15780620003ac888262000340565b84526020840193505060208101905062000396565b5050509392505050565b600082601f830112620003e357620003e262000224565b5b8151620003f584826020860162000357565b91505092915050565b6000602082840312156200041757620004166200021a565b5b600082015167ffffffffffffffff8111156200043857620004376200021f565b5b6200044684828501620003cb565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b6000620004c482620004ad565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620004f957620004f86200047e565b5b600182019050919050565b611a1580620005146000396000f3fe60806040526004361061008a5760003560e01c80634d238c8e116100595780634d238c8e146101125780634fc2798b1461012e5780638299e3841461015a578063b7ab4db514610164578063f951a18a1461018f57610091565b806312065fe0146100935780631cbe1a9a146100be5780633356a0fd146100da57806340a141ff146100f657610091565b3661009157005b005b34801561009f57600080fd5b506100a86101bb565b6040516100b591906114b9565b60405180910390f35b6100d860048036038101906100d39190611537565b6101c3565b005b6100f460048036038101906100ef9190611537565b610500565b005b610110600480360381019061010b9190611537565b610b97565b005b61012c60048036038101906101279190611537565b610eed565b005b34801561013a57600080fd5b5061014361105e565b6040516101519291906116e0565b60405180910390f35b610162611214565b005b34801561017057600080fd5b5061017961125c565b6040516101869190611717565b60405180910390f35b34801561019b57600080fd5b506101a46112ea565b6040516101b29291906116e0565b60405180910390f35b600047905090565b80600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff161561021e57600080fd5b3334670de0b6b3a76400008111801561024a57506000670de0b6b3a7640000826102489190611768565b145b61025357600080fd5b808273ffffffffffffffffffffffffffffffffffffffff16311161027657600080fd5b60003390503073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050501580156102c1573d6000803e3d6000fd5b5034600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555084600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff021916908315150217905550600080549050600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000859080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506104f9611214565b5050505050565b806000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff1690506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015490508180156105ad575060008054905081105b801561062257508273ffffffffffffffffffffffffffffffffffffffff16600082815481106105df576105de611799565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b61062b57600080fd5b33846000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050670de0b6b3a76400008111801561069b57506000670de0b6b3a7640000826106999190611768565b145b6106a457600080fd5b6000600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461074057600080fd5b60003390506000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156107cf573d6000803e3d6000fd5b506000600260008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549050600180808054905061082991906117f7565b8154811061083a57610839611799565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166001828154811061087957610878611799565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018054806108d3576108d261182b565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009055600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600060016000805490506109c191906117f7565b815481106109d2576109d1611799565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660008281548110610a1157610a10611799565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060026000808481548110610a7157610a70611799565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000805480610aef57610aee61182b565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055600260008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006101000a81549060ff021916905560018201600090555050610b8a611214565b5050505050505050505050565b806000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff1690506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549050818015610c44575060008054905081105b8015610cb957508273ffffffffffffffffffffffffffffffffffffffff1660008281548110610c7657610c75611799565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b610cc257600080fd5b6000600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154905060006001600080549050610d1d91906117f7565b81548110610d2e57610d2d611799565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660008281548110610d6d57610d6c611799565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060026000808481548110610dcd57610dcc611799565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000805480610e4b57610e4a61182b565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006101000a81549060ff021916905560018201600090555050610ee6611214565b5050505050565b80600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff1615610f4857600080fd5b6001600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff021916908315150217905550600080549050600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061105a611214565b5050565b606080600060018054905067ffffffffffffffff8111156110825761108161185a565b5b6040519080825280602002602001820160405280156110b05781602001602082028036833780820191505090505b50905060005b60018054905081101561117f576000600182815481106110d9576110d8611799565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508084848151811061115e5761115d611799565b5b6020026020010181815250505050808061117790611889565b9150506110b6565b506001818180548060200260200160405190810160405280929190818152602001828054801561120457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116111ba575b5050505050915092509250509091565b60014361122191906117f7565b407f55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c89600060405161125291906119bd565b60405180910390a2565b606060008054806020026020016040519081016040528092919081815260200182805480156112e057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611296575b5050505050905090565b606080600060018054905067ffffffffffffffff81111561130e5761130d61185a565b5b60405190808252806020026020018201604052801561133c5781602001602082028036833780820191505090505b50905060005b60018054905081101561140b5760006001828154811061136557611364611799565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050808484815181106113ea576113e9611799565b5b6020026020010181815250505050808061140390611889565b915050611342565b506000818180548060200260200160405190810160405280929190818152602001828054801561149057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611446575b5050505050915092509250509091565b6000819050919050565b6114b3816114a0565b82525050565b60006020820190506114ce60008301846114aa565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611504826114d9565b9050919050565b611514816114f9565b811461151f57600080fd5b50565b6000813590506115318161150b565b92915050565b60006020828403121561154d5761154c6114d4565b5b600061155b84828501611522565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611599816114f9565b82525050565b60006115ab8383611590565b60208301905092915050565b6000602082019050919050565b60006115cf82611564565b6115d9818561156f565b93506115e483611580565b8060005b838110156116155781516115fc888261159f565b9750611607836115b7565b9250506001810190506115e8565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611657816114a0565b82525050565b6000611669838361164e565b60208301905092915050565b6000602082019050919050565b600061168d82611622565b611697818561162d565b93506116a28361163e565b8060005b838110156116d35781516116ba888261165d565b97506116c583611675565b9250506001810190506116a6565b5085935050505092915050565b600060408201905081810360008301526116fa81856115c4565b9050818103602083015261170e8184611682565b90509392505050565b6000602082019050818103600083015261173181846115c4565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611773826114a0565b915061177e836114a0565b92508261178e5761178d611739565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611802826114a0565b915061180d836114a0565b9250828203905081811115611825576118246117c8565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000611894826114a0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036118c6576118c56117c8565b5b600182019050919050565b600081549050919050565b60008190508160005260206000209050919050565b60008160001c9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061193161192c836118f1565b6118fe565b9050919050565b6000611944825461191e565b9050919050565b6000600182019050919050565b6000611963826118d1565b61196d818561156f565b9350611978836118dc565b8060005b838110156119b05761198d82611938565b611997888261159f565b97506119a28361194b565b92505060018101905061197c565b5085935050505092915050565b600060208201905081810360008301526119d78184611958565b90509291505056fea2646970667358221220319716695f1ac6c44da0bb6eabe3593adcd3beb46198af140ba7ab2fbc0edccf64736f6c63430008130033",
}

// ValidatorsetABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorsetMetaData.ABI instead.
var ValidatorsetABI = ValidatorsetMetaData.ABI

// ValidatorsetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorsetMetaData.Bin instead.
var ValidatorsetBin = ValidatorsetMetaData.Bin

// DeployValidatorset deploys a new Ethereum contract, binding an instance of Validatorset to it.
func DeployValidatorset(auth *bind.TransactOpts, backend bind.ContractBackend, _initial []common.Address) (common.Address, *types.Transaction, *Validatorset, error) {
	parsed, err := ValidatorsetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorsetBin), backend, _initial)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validatorset{ValidatorsetCaller: ValidatorsetCaller{contract: contract}, ValidatorsetTransactor: ValidatorsetTransactor{contract: contract}, ValidatorsetFilterer: ValidatorsetFilterer{contract: contract}}, nil
}

// Validatorset is an auto generated Go binding around an Ethereum contract.
type Validatorset struct {
	ValidatorsetCaller     // Read-only binding to the contract
	ValidatorsetTransactor // Write-only binding to the contract
	ValidatorsetFilterer   // Log filterer for contract events
}

// ValidatorsetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorsetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorsetSession struct {
	Contract     *Validatorset     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorsetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorsetCallerSession struct {
	Contract *ValidatorsetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ValidatorsetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorsetTransactorSession struct {
	Contract     *ValidatorsetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorsetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorsetRaw struct {
	Contract *Validatorset // Generic contract binding to access the raw methods on
}

// ValidatorsetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorsetCallerRaw struct {
	Contract *ValidatorsetCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorsetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorsetTransactorRaw struct {
	Contract *ValidatorsetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorset creates a new instance of Validatorset, bound to a specific deployed contract.
func NewValidatorset(address common.Address, backend bind.ContractBackend) (*Validatorset, error) {
	contract, err := bindValidatorset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validatorset{ValidatorsetCaller: ValidatorsetCaller{contract: contract}, ValidatorsetTransactor: ValidatorsetTransactor{contract: contract}, ValidatorsetFilterer: ValidatorsetFilterer{contract: contract}}, nil
}

// NewValidatorsetCaller creates a new read-only instance of Validatorset, bound to a specific deployed contract.
func NewValidatorsetCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsetCaller, error) {
	contract, err := bindValidatorset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetCaller{contract: contract}, nil
}

// NewValidatorsetTransactor creates a new write-only instance of Validatorset, bound to a specific deployed contract.
func NewValidatorsetTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsetTransactor, error) {
	contract, err := bindValidatorset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetTransactor{contract: contract}, nil
}

// NewValidatorsetFilterer creates a new log filterer instance of Validatorset, bound to a specific deployed contract.
func NewValidatorsetFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorsetFilterer, error) {
	contract, err := bindValidatorset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetFilterer{contract: contract}, nil
}

// bindValidatorset binds a generic wrapper to an already deployed contract.
func bindValidatorset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorsetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorset *ValidatorsetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorset.Contract.ValidatorsetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorset *ValidatorsetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.Contract.ValidatorsetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorset *ValidatorsetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorset.Contract.ValidatorsetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorset *ValidatorsetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorset *ValidatorsetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorset *ValidatorsetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorset.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Validatorset *ValidatorsetSession) GetBalance() (*big.Int, error) {
	return _Validatorset.Contract.GetBalance(&_Validatorset.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) GetBalance() (*big.Int, error) {
	return _Validatorset.Contract.GetBalance(&_Validatorset.CallOpts)
}

// GetStakerAmounts is a free data retrieval call binding the contract method 0x4fc2798b.
//
// Solidity: function getStakerAmounts() view returns(address[], uint256[])
func (_Validatorset *ValidatorsetCaller) GetStakerAmounts(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "getStakerAmounts")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetStakerAmounts is a free data retrieval call binding the contract method 0x4fc2798b.
//
// Solidity: function getStakerAmounts() view returns(address[], uint256[])
func (_Validatorset *ValidatorsetSession) GetStakerAmounts() ([]common.Address, []*big.Int, error) {
	return _Validatorset.Contract.GetStakerAmounts(&_Validatorset.CallOpts)
}

// GetStakerAmounts is a free data retrieval call binding the contract method 0x4fc2798b.
//
// Solidity: function getStakerAmounts() view returns(address[], uint256[])
func (_Validatorset *ValidatorsetCallerSession) GetStakerAmounts() ([]common.Address, []*big.Int, error) {
	return _Validatorset.Contract.GetStakerAmounts(&_Validatorset.CallOpts)
}

// GetValidatorAmounts is a free data retrieval call binding the contract method 0xf951a18a.
//
// Solidity: function getValidatorAmounts() view returns(address[], uint256[])
func (_Validatorset *ValidatorsetCaller) GetValidatorAmounts(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "getValidatorAmounts")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetValidatorAmounts is a free data retrieval call binding the contract method 0xf951a18a.
//
// Solidity: function getValidatorAmounts() view returns(address[], uint256[])
func (_Validatorset *ValidatorsetSession) GetValidatorAmounts() ([]common.Address, []*big.Int, error) {
	return _Validatorset.Contract.GetValidatorAmounts(&_Validatorset.CallOpts)
}

// GetValidatorAmounts is a free data retrieval call binding the contract method 0xf951a18a.
//
// Solidity: function getValidatorAmounts() view returns(address[], uint256[])
func (_Validatorset *ValidatorsetCallerSession) GetValidatorAmounts() ([]common.Address, []*big.Int, error) {
	return _Validatorset.Contract.GetValidatorAmounts(&_Validatorset.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Validatorset *ValidatorsetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Validatorset *ValidatorsetSession) GetValidators() ([]common.Address, error) {
	return _Validatorset.Contract.GetValidators(&_Validatorset.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Validatorset *ValidatorsetCallerSession) GetValidators() ([]common.Address, error) {
	return _Validatorset.Contract.GetValidators(&_Validatorset.CallOpts)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) payable returns()
func (_Validatorset *ValidatorsetTransactor) AddValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "addValidator", _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) payable returns()
func (_Validatorset *ValidatorsetSession) AddValidator(_validator common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.AddValidator(&_Validatorset.TransactOpts, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) payable returns()
func (_Validatorset *ValidatorsetTransactorSession) AddValidator(_validator common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.AddValidator(&_Validatorset.TransactOpts, _validator)
}

// AddValidator2 is a paid mutator transaction binding the contract method 0x1cbe1a9a.
//
// Solidity: function addValidator2(address _validator) payable returns()
func (_Validatorset *ValidatorsetTransactor) AddValidator2(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "addValidator2", _validator)
}

// AddValidator2 is a paid mutator transaction binding the contract method 0x1cbe1a9a.
//
// Solidity: function addValidator2(address _validator) payable returns()
func (_Validatorset *ValidatorsetSession) AddValidator2(_validator common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.AddValidator2(&_Validatorset.TransactOpts, _validator)
}

// AddValidator2 is a paid mutator transaction binding the contract method 0x1cbe1a9a.
//
// Solidity: function addValidator2(address _validator) payable returns()
func (_Validatorset *ValidatorsetTransactorSession) AddValidator2(_validator common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.AddValidator2(&_Validatorset.TransactOpts, _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) payable returns()
func (_Validatorset *ValidatorsetTransactor) RemoveValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "removeValidator", _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) payable returns()
func (_Validatorset *ValidatorsetSession) RemoveValidator(_validator common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.RemoveValidator(&_Validatorset.TransactOpts, _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) payable returns()
func (_Validatorset *ValidatorsetTransactorSession) RemoveValidator(_validator common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.RemoveValidator(&_Validatorset.TransactOpts, _validator)
}

// RemoveValidator2 is a paid mutator transaction binding the contract method 0x3356a0fd.
//
// Solidity: function removeValidator2(address _validator) payable returns()
func (_Validatorset *ValidatorsetTransactor) RemoveValidator2(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "removeValidator2", _validator)
}

// RemoveValidator2 is a paid mutator transaction binding the contract method 0x3356a0fd.
//
// Solidity: function removeValidator2(address _validator) payable returns()
func (_Validatorset *ValidatorsetSession) RemoveValidator2(_validator common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.RemoveValidator2(&_Validatorset.TransactOpts, _validator)
}

// RemoveValidator2 is a paid mutator transaction binding the contract method 0x3356a0fd.
//
// Solidity: function removeValidator2(address _validator) payable returns()
func (_Validatorset *ValidatorsetTransactorSession) RemoveValidator2(_validator common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.RemoveValidator2(&_Validatorset.TransactOpts, _validator)
}

// TriggerChange is a paid mutator transaction binding the contract method 0x8299e384.
//
// Solidity: function triggerChange() payable returns()
func (_Validatorset *ValidatorsetTransactor) TriggerChange(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "triggerChange")
}

// TriggerChange is a paid mutator transaction binding the contract method 0x8299e384.
//
// Solidity: function triggerChange() payable returns()
func (_Validatorset *ValidatorsetSession) TriggerChange() (*types.Transaction, error) {
	return _Validatorset.Contract.TriggerChange(&_Validatorset.TransactOpts)
}

// TriggerChange is a paid mutator transaction binding the contract method 0x8299e384.
//
// Solidity: function triggerChange() payable returns()
func (_Validatorset *ValidatorsetTransactorSession) TriggerChange() (*types.Transaction, error) {
	return _Validatorset.Contract.TriggerChange(&_Validatorset.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Validatorset *ValidatorsetTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Validatorset.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Validatorset *ValidatorsetSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Validatorset.Contract.Fallback(&_Validatorset.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Validatorset *ValidatorsetTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Validatorset.Contract.Fallback(&_Validatorset.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validatorset *ValidatorsetTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validatorset *ValidatorsetSession) Receive() (*types.Transaction, error) {
	return _Validatorset.Contract.Receive(&_Validatorset.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validatorset *ValidatorsetTransactorSession) Receive() (*types.Transaction, error) {
	return _Validatorset.Contract.Receive(&_Validatorset.TransactOpts)
}

// ValidatorsetInitiateChangeIterator is returned from FilterInitiateChange and is used to iterate over the raw logs and unpacked data for InitiateChange events raised by the Validatorset contract.
type ValidatorsetInitiateChangeIterator struct {
	Event *ValidatorsetInitiateChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorsetInitiateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsetInitiateChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorsetInitiateChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorsetInitiateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsetInitiateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsetInitiateChange represents a InitiateChange event raised by the Validatorset contract.
type ValidatorsetInitiateChange struct {
	ParentHash [32]byte
	NewSet     []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitiateChange is a free log retrieval operation binding the contract event 0x55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c89.
//
// Solidity: event InitiateChange(bytes32 indexed _parentHash, address[] _newSet)
func (_Validatorset *ValidatorsetFilterer) FilterInitiateChange(opts *bind.FilterOpts, _parentHash [][32]byte) (*ValidatorsetInitiateChangeIterator, error) {

	var _parentHashRule []interface{}
	for _, _parentHashItem := range _parentHash {
		_parentHashRule = append(_parentHashRule, _parentHashItem)
	}

	logs, sub, err := _Validatorset.contract.FilterLogs(opts, "InitiateChange", _parentHashRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetInitiateChangeIterator{contract: _Validatorset.contract, event: "InitiateChange", logs: logs, sub: sub}, nil
}

// WatchInitiateChange is a free log subscription operation binding the contract event 0x55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c89.
//
// Solidity: event InitiateChange(bytes32 indexed _parentHash, address[] _newSet)
func (_Validatorset *ValidatorsetFilterer) WatchInitiateChange(opts *bind.WatchOpts, sink chan<- *ValidatorsetInitiateChange, _parentHash [][32]byte) (event.Subscription, error) {

	var _parentHashRule []interface{}
	for _, _parentHashItem := range _parentHash {
		_parentHashRule = append(_parentHashRule, _parentHashItem)
	}

	logs, sub, err := _Validatorset.contract.WatchLogs(opts, "InitiateChange", _parentHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsetInitiateChange)
				if err := _Validatorset.contract.UnpackLog(event, "InitiateChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitiateChange is a log parse operation binding the contract event 0x55252fa6eee4741b4e24a74a70e9c11fd2c2281df8d6ea13126ff845f7825c89.
//
// Solidity: event InitiateChange(bytes32 indexed _parentHash, address[] _newSet)
func (_Validatorset *ValidatorsetFilterer) ParseInitiateChange(log types.Log) (*ValidatorsetInitiateChange, error) {
	event := new(ValidatorsetInitiateChange)
	if err := _Validatorset.contract.UnpackLog(event, "InitiateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
