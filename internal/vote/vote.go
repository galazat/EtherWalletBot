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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"CandidatNames\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Candidats\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VoteCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Candidat\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"voted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winnerName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"winnerName_\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winninCandidat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"winninCandidat_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161074d38038061074d83398101604081905261002f916100e2565b600080546001600160a01b03191633908117825581526001602081905260408220555b81518110156100c5576002604051806040016040528084848151811061007a5761007a61019f565b602090810291909101810151825260009181018290528354600181810186559483529181902083516002909302019182559190910151910155806100bd816101b5565b915050610052565b50506101dc565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156100f557600080fd5b82516001600160401b038082111561010c57600080fd5b818501915085601f83011261012057600080fd5b815181811115610132576101326100cc565b8060051b604051601f19603f83011681018181108582111715610157576101576100cc565b60405291825284820192508381018501918883111561017557600080fd5b938501935b828510156101935784518452938501939285019261017a565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600182016101d557634e487b7160e01b600052601160045260246000fd5b5060010190565b610562806101eb6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063a3ec138d1161005b578063a3ec138d146100d7578063ae0222e614610126578063bfe875d614610151578063e2ba53f01461016757600080fd5b80630121b93f146100825780634b2a53b8146100975780639e7b8d61146100c4575b600080fd5b61009561009036600461045f565b61016f565b005b6100aa6100a536600461045f565b610269565b604080519283526020830191909152015b60405180910390f35b6100956100d2366004610478565b610297565b61010b6100e5366004610478565b600160208190526000918252604090912080549181015460029091015460ff9091169083565b604080519384529115156020840152908201526060016100bb565b600054610139906001600160a01b031681565b6040516001600160a01b0390911681526020016100bb565b6101596103af565b6040519081526020016100bb565b61015961042c565b33600090815260016020526040812080549091036101cb5760405162461bcd60e51b8152602060048201526014602482015273486173206e6f20726967687420746f20766f746560601b60448201526064015b60405180910390fd5b600181015460ff16156102115760405162461bcd60e51b815260206004820152600e60248201526d20b63932b0b23c903b37ba32b21760911b60448201526064016101c2565b6001818101805460ff1916909117905560028082018390558154815490919084908110610240576102406104a8565b9060005260206000209060020201600101600082825461026091906104d4565b90915550505050565b6002818154811061027957600080fd5b60009182526020909120600290910201805460019091015490915082565b6000546001600160a01b031633146103025760405162461bcd60e51b815260206004820152602860248201527f4f6e6c7920566f746543726561746f722063616e2067697665207269676874206044820152673a37903b37ba329760c11b60648201526084016101c2565b6001600160a01b0381166000908152600160208190526040909120015460ff161561036f5760405162461bcd60e51b815260206004820152601860248201527f54686520766f74657220616c726561647920766f7465642e000000000000000060448201526064016101c2565b6001600160a01b0381166000908152600160205260409020541561039257600080fd5b6001600160a01b0316600090815260016020819052604090912055565b600080805b6002548110156104275781600282815481106103d2576103d26104a8565b906000526020600020906002020160010154111561041557600281815481106103fd576103fd6104a8565b90600052602060002090600202016001015491508092505b8061041f816104ed565b9150506103b4565b505090565b600060026104386103af565b81548110610448576104486104a8565b906000526020600020906002020160000154905090565b60006020828403121561047157600080fd5b5035919050565b60006020828403121561048a57600080fd5b81356001600160a01b03811681146104a157600080fd5b9392505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b808201808211156104e7576104e76104be565b92915050565b6000600182016104ff576104ff6104be565b506001019056fea26469706673582212201d63478558738337d2dfa534377ac729ec9350fcb8805e73779184c616da8fa264736f6c637828302e382e32302d646576656c6f702e323032332e322e32352b636f6d6d69742e39383334303737360059",
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
