// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vote

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

// VoteMetaData contains all meta data concerning the Vote contract.
var VoteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"CandidatNames\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Candidats\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TotalCandidats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VoteCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Candidat\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"voted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winnerName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"winnerName_\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winninCandidat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"winninCandidat_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161076738038061076783398101604081905261002f916100e8565b600080546001600160a01b03191633908117825581526001602081905260408220555b81518110156100c5576002604051806040016040528084848151811061007a5761007a6101a5565b602090810291909101810151825260009181018290528354600181810186559483529181902083516002909302019182559190910151910155806100bd816101bb565b915050610052565b50506002546003556101e2565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156100fb57600080fd5b82516001600160401b038082111561011257600080fd5b818501915085601f83011261012657600080fd5b815181811115610138576101386100d2565b8060051b604051601f19603f8301168101818110858211171561015d5761015d6100d2565b60405291825284820192508381018501918883111561017b57600080fd5b938501935b8285101561019957845184529385019392850192610180565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600182016101db57634e487b7160e01b600052601160045260246000fd5b5060010190565b610576806101f16000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063ae0222e61161005b578063ae0222e614610131578063bfe875d61461015c578063e2ba53f014610172578063f2a2f9721461017a57600080fd5b80630121b93f1461008d5780634b2a53b8146100a25780639e7b8d61146100cf578063a3ec138d146100e2575b600080fd5b6100a061009b366004610473565b610183565b005b6100b56100b0366004610473565b61027d565b604080519283526020830191909152015b60405180910390f35b6100a06100dd36600461048c565b6102ab565b6101166100f036600461048c565b600160208190526000918252604090912080549181015460029091015460ff9091169083565b604080519384529115156020840152908201526060016100c6565b600054610144906001600160a01b031681565b6040516001600160a01b0390911681526020016100c6565b6101646103c3565b6040519081526020016100c6565b610164610440565b61016460035481565b33600090815260016020526040812080549091036101df5760405162461bcd60e51b8152602060048201526014602482015273486173206e6f20726967687420746f20766f746560601b60448201526064015b60405180910390fd5b600181015460ff16156102255760405162461bcd60e51b815260206004820152600e60248201526d20b63932b0b23c903b37ba32b21760911b60448201526064016101d6565b6001818101805460ff1916909117905560028082018390558154815490919084908110610254576102546104bc565b9060005260206000209060020201600101600082825461027491906104e8565b90915550505050565b6002818154811061028d57600080fd5b60009182526020909120600290910201805460019091015490915082565b6000546001600160a01b031633146103165760405162461bcd60e51b815260206004820152602860248201527f4f6e6c7920566f746543726561746f722063616e2067697665207269676874206044820152673a37903b37ba329760c11b60648201526084016101d6565b6001600160a01b0381166000908152600160208190526040909120015460ff16156103835760405162461bcd60e51b815260206004820152601860248201527f54686520766f74657220616c726561647920766f7465642e000000000000000060448201526064016101d6565b6001600160a01b038116600090815260016020526040902054156103a657600080fd5b6001600160a01b0316600090815260016020819052604090912055565b600080805b60025481101561043b5781600282815481106103e6576103e66104bc565b90600052602060002090600202016001015411156104295760028181548110610411576104116104bc565b90600052602060002090600202016001015491508092505b8061043381610501565b9150506103c8565b505090565b6000600261044c6103c3565b8154811061045c5761045c6104bc565b906000526020600020906002020160000154905090565b60006020828403121561048557600080fd5b5035919050565b60006020828403121561049e57600080fd5b81356001600160a01b03811681146104b557600080fd5b9392505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b808201808211156104fb576104fb6104d2565b92915050565b600060018201610513576105136104d2565b506001019056fea2646970667358221220050fac2c6531c18409f300a746760a3a2fb5b7f3a22540c7420b7446fc45dabb64736f6c637828302e382e32302d646576656c6f702e323032332e322e32352b636f6d6d69742e39383334303737360059",
}

// VoteABI is the input ABI used to generate the binding from.
// Deprecated: Use VoteMetaData.ABI instead.
var VoteABI = VoteMetaData.ABI

// VoteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VoteMetaData.Bin instead.
var VoteBin = VoteMetaData.Bin

// DeployVote deploys a new Ethereum contract, binding an instance of Vote to it.
func DeployVote(auth *bind.TransactOpts, backend bind.ContractBackend, CandidatNames [][32]byte) (common.Address, *types.Transaction, *Vote, error) {
	parsed, err := VoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VoteBin), backend, CandidatNames)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}

// Candidats is a free data retrieval call binding the contract method 0x4b2a53b8.
//
// Solidity: function Candidats(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Vote *VoteCaller) Candidats(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "Candidats", arg0)

	outstruct := new(struct {
		Name      [32]byte
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Candidats is a free data retrieval call binding the contract method 0x4b2a53b8.
//
// Solidity: function Candidats(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Vote *VoteSession) Candidats(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Vote.Contract.Candidats(&_Vote.CallOpts, arg0)
}

// Candidats is a free data retrieval call binding the contract method 0x4b2a53b8.
//
// Solidity: function Candidats(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Vote *VoteCallerSession) Candidats(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Vote.Contract.Candidats(&_Vote.CallOpts, arg0)
}

// TotalCandidats is a free data retrieval call binding the contract method 0xf2a2f972.
//
// Solidity: function TotalCandidats() view returns(uint256)
func (_Vote *VoteCaller) TotalCandidats(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "TotalCandidats")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCandidats is a free data retrieval call binding the contract method 0xf2a2f972.
//
// Solidity: function TotalCandidats() view returns(uint256)
func (_Vote *VoteSession) TotalCandidats() (*big.Int, error) {
	return _Vote.Contract.TotalCandidats(&_Vote.CallOpts)
}

// TotalCandidats is a free data retrieval call binding the contract method 0xf2a2f972.
//
// Solidity: function TotalCandidats() view returns(uint256)
func (_Vote *VoteCallerSession) TotalCandidats() (*big.Int, error) {
	return _Vote.Contract.TotalCandidats(&_Vote.CallOpts)
}

// VoteCreator is a free data retrieval call binding the contract method 0xae0222e6.
//
// Solidity: function VoteCreator() view returns(address)
func (_Vote *VoteCaller) VoteCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "VoteCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteCreator is a free data retrieval call binding the contract method 0xae0222e6.
//
// Solidity: function VoteCreator() view returns(address)
func (_Vote *VoteSession) VoteCreator() (common.Address, error) {
	return _Vote.Contract.VoteCreator(&_Vote.CallOpts)
}

// VoteCreator is a free data retrieval call binding the contract method 0xae0222e6.
//
// Solidity: function VoteCreator() view returns(address)
func (_Vote *VoteCallerSession) VoteCreator() (common.Address, error) {
	return _Vote.Contract.VoteCreator(&_Vote.CallOpts)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, uint256 vote)
func (_Vote *VoteCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	Weight *big.Int
	Voted  bool
	Vote   *big.Int
}, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		Weight *big.Int
		Voted  bool
		Vote   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Weight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Voted = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Vote = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, uint256 vote)
func (_Vote *VoteSession) Voters(arg0 common.Address) (struct {
	Weight *big.Int
	Voted  bool
	Vote   *big.Int
}, error) {
	return _Vote.Contract.Voters(&_Vote.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, uint256 vote)
func (_Vote *VoteCallerSession) Voters(arg0 common.Address) (struct {
	Weight *big.Int
	Voted  bool
	Vote   *big.Int
}, error) {
	return _Vote.Contract.Voters(&_Vote.CallOpts, arg0)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Vote *VoteCaller) WinnerName(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "winnerName")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Vote *VoteSession) WinnerName() ([32]byte, error) {
	return _Vote.Contract.WinnerName(&_Vote.CallOpts)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Vote *VoteCallerSession) WinnerName() ([32]byte, error) {
	return _Vote.Contract.WinnerName(&_Vote.CallOpts)
}

// WinninCandidat is a free data retrieval call binding the contract method 0xbfe875d6.
//
// Solidity: function winninCandidat() view returns(uint256 winninCandidat_)
func (_Vote *VoteCaller) WinninCandidat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "winninCandidat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinninCandidat is a free data retrieval call binding the contract method 0xbfe875d6.
//
// Solidity: function winninCandidat() view returns(uint256 winninCandidat_)
func (_Vote *VoteSession) WinninCandidat() (*big.Int, error) {
	return _Vote.Contract.WinninCandidat(&_Vote.CallOpts)
}

// WinninCandidat is a free data retrieval call binding the contract method 0xbfe875d6.
//
// Solidity: function winninCandidat() view returns(uint256 winninCandidat_)
func (_Vote *VoteCallerSession) WinninCandidat() (*big.Int, error) {
	return _Vote.Contract.WinninCandidat(&_Vote.CallOpts)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Vote *VoteTransactor) GiveRightToVote(opts *bind.TransactOpts, voter common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "giveRightToVote", voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Vote *VoteSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Vote.Contract.GiveRightToVote(&_Vote.TransactOpts, voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Vote *VoteTransactorSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Vote.Contract.GiveRightToVote(&_Vote.TransactOpts, voter)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 Candidat) returns()
func (_Vote *VoteTransactor) Vote(opts *bind.TransactOpts, Candidat *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "vote", Candidat)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 Candidat) returns()
func (_Vote *VoteSession) Vote(Candidat *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Vote(&_Vote.TransactOpts, Candidat)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 Candidat) returns()
func (_Vote *VoteTransactorSession) Vote(Candidat *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Vote(&_Vote.TransactOpts, Candidat)
}
