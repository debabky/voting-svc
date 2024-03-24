// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// VerifierHelperProofPoints is an auto generated low-level Go binding around an user-defined struct.
type VerifierHelperProofPoints struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// VotingVotingParams is an auto generated low-level Go binding around an user-defined struct.
type VotingVotingParams struct {
	StartTimestamp *big.Int
	Duration       *big.Int
	Candidates     []*big.Int
}

// VotingVotingPublicConfig is an auto generated low-level Go binding around an user-defined struct.
type VotingVotingPublicConfig struct {
	StartTimestamp     *big.Int
	EndTimestamp       *big.Int
	Status             uint8
	Candidates         []*big.Int
	VotesPerCandidates []*big.Int
}

// VotingMetaData contains all meta data concerning the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationMerkleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"verifier_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"candidates\",\"type\":\"uint256[]\"}],\"internalType\":\"structVoting.VotingParams\",\"name\":\"config_\",\"type\":\"tuple\"}],\"name\":\"__Voting_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blindedNullifiers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotingInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumVoting.VotingStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"candidates\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"votesPerCandidates\",\"type\":\"uint256[]\"}],\"internalType\":\"structVoting.VotingPublicConfig\",\"name\":\"info_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrationMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrationTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"candidates_\",\"type\":\"uint256[5]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"nullifierHash_\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"votesForCandidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VotingABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMetaData.ABI instead.
var VotingABI = VotingMetaData.ABI

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// BlindedNullifiers is a free data retrieval call binding the contract method 0x24526b1b.
//
// Solidity: function blindedNullifiers(uint256 ) view returns(bool)
func (_Voting *VotingCaller) BlindedNullifiers(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "blindedNullifiers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlindedNullifiers is a free data retrieval call binding the contract method 0x24526b1b.
//
// Solidity: function blindedNullifiers(uint256 ) view returns(bool)
func (_Voting *VotingSession) BlindedNullifiers(arg0 *big.Int) (bool, error) {
	return _Voting.Contract.BlindedNullifiers(&_Voting.CallOpts, arg0)
}

// BlindedNullifiers is a free data retrieval call binding the contract method 0x24526b1b.
//
// Solidity: function blindedNullifiers(uint256 ) view returns(bool)
func (_Voting *VotingCallerSession) BlindedNullifiers(arg0 *big.Int) (bool, error) {
	return _Voting.Contract.BlindedNullifiers(&_Voting.CallOpts, arg0)
}

// GetVotingInfo is a free data retrieval call binding the contract method 0x359af4b2.
//
// Solidity: function getVotingInfo() view returns((uint256,uint256,uint8,uint256[],uint256[]) info_)
func (_Voting *VotingCaller) GetVotingInfo(opts *bind.CallOpts) (VotingVotingPublicConfig, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getVotingInfo")

	if err != nil {
		return *new(VotingVotingPublicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(VotingVotingPublicConfig)).(*VotingVotingPublicConfig)

	return out0, err

}

// GetVotingInfo is a free data retrieval call binding the contract method 0x359af4b2.
//
// Solidity: function getVotingInfo() view returns((uint256,uint256,uint8,uint256[],uint256[]) info_)
func (_Voting *VotingSession) GetVotingInfo() (VotingVotingPublicConfig, error) {
	return _Voting.Contract.GetVotingInfo(&_Voting.CallOpts)
}

// GetVotingInfo is a free data retrieval call binding the contract method 0x359af4b2.
//
// Solidity: function getVotingInfo() view returns((uint256,uint256,uint8,uint256[],uint256[]) info_)
func (_Voting *VotingCallerSession) GetVotingInfo() (VotingVotingPublicConfig, error) {
	return _Voting.Contract.GetVotingInfo(&_Voting.CallOpts)
}

// RegistrationMerkleRoot is a free data retrieval call binding the contract method 0xfa1b3180.
//
// Solidity: function registrationMerkleRoot() view returns(bytes32)
func (_Voting *VotingCaller) RegistrationMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "registrationMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RegistrationMerkleRoot is a free data retrieval call binding the contract method 0xfa1b3180.
//
// Solidity: function registrationMerkleRoot() view returns(bytes32)
func (_Voting *VotingSession) RegistrationMerkleRoot() ([32]byte, error) {
	return _Voting.Contract.RegistrationMerkleRoot(&_Voting.CallOpts)
}

// RegistrationMerkleRoot is a free data retrieval call binding the contract method 0xfa1b3180.
//
// Solidity: function registrationMerkleRoot() view returns(bytes32)
func (_Voting *VotingCallerSession) RegistrationMerkleRoot() ([32]byte, error) {
	return _Voting.Contract.RegistrationMerkleRoot(&_Voting.CallOpts)
}

// RegistrationTimestamp is a free data retrieval call binding the contract method 0x1e801262.
//
// Solidity: function registrationTimestamp() view returns(uint256)
func (_Voting *VotingCaller) RegistrationTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "registrationTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistrationTimestamp is a free data retrieval call binding the contract method 0x1e801262.
//
// Solidity: function registrationTimestamp() view returns(uint256)
func (_Voting *VotingSession) RegistrationTimestamp() (*big.Int, error) {
	return _Voting.Contract.RegistrationTimestamp(&_Voting.CallOpts)
}

// RegistrationTimestamp is a free data retrieval call binding the contract method 0x1e801262.
//
// Solidity: function registrationTimestamp() view returns(uint256)
func (_Voting *VotingCallerSession) RegistrationTimestamp() (*big.Int, error) {
	return _Voting.Contract.RegistrationTimestamp(&_Voting.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Voting *VotingCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Voting *VotingSession) Verifier() (common.Address, error) {
	return _Voting.Contract.Verifier(&_Voting.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Voting *VotingCallerSession) Verifier() (common.Address, error) {
	return _Voting.Contract.Verifier(&_Voting.CallOpts)
}

// VotesForCandidates is a free data retrieval call binding the contract method 0x055ef41c.
//
// Solidity: function votesForCandidates(uint256 ) view returns(uint256)
func (_Voting *VotingCaller) VotesForCandidates(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "votesForCandidates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotesForCandidates is a free data retrieval call binding the contract method 0x055ef41c.
//
// Solidity: function votesForCandidates(uint256 ) view returns(uint256)
func (_Voting *VotingSession) VotesForCandidates(arg0 *big.Int) (*big.Int, error) {
	return _Voting.Contract.VotesForCandidates(&_Voting.CallOpts, arg0)
}

// VotesForCandidates is a free data retrieval call binding the contract method 0x055ef41c.
//
// Solidity: function votesForCandidates(uint256 ) view returns(uint256)
func (_Voting *VotingCallerSession) VotesForCandidates(arg0 *big.Int) (*big.Int, error) {
	return _Voting.Contract.VotesForCandidates(&_Voting.CallOpts, arg0)
}

// VotingInit is a paid mutator transaction binding the contract method 0x36cd7f05.
//
// Solidity: function __Voting_init(bytes32 registrationMerkleRoot_, address verifier_, (uint256,uint256,uint256[]) config_) returns()
func (_Voting *VotingTransactor) VotingInit(opts *bind.TransactOpts, registrationMerkleRoot_ [32]byte, verifier_ common.Address, config_ VotingVotingParams) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "__Voting_init", registrationMerkleRoot_, verifier_, config_)
}

// VotingInit is a paid mutator transaction binding the contract method 0x36cd7f05.
//
// Solidity: function __Voting_init(bytes32 registrationMerkleRoot_, address verifier_, (uint256,uint256,uint256[]) config_) returns()
func (_Voting *VotingSession) VotingInit(registrationMerkleRoot_ [32]byte, verifier_ common.Address, config_ VotingVotingParams) (*types.Transaction, error) {
	return _Voting.Contract.VotingInit(&_Voting.TransactOpts, registrationMerkleRoot_, verifier_, config_)
}

// VotingInit is a paid mutator transaction binding the contract method 0x36cd7f05.
//
// Solidity: function __Voting_init(bytes32 registrationMerkleRoot_, address verifier_, (uint256,uint256,uint256[]) config_) returns()
func (_Voting *VotingTransactorSession) VotingInit(registrationMerkleRoot_ [32]byte, verifier_ common.Address, config_ VotingVotingParams) (*types.Transaction, error) {
	return _Voting.Contract.VotingInit(&_Voting.TransactOpts, registrationMerkleRoot_, verifier_, config_)
}

// Vote is a paid mutator transaction binding the contract method 0x68a866f0.
//
// Solidity: function vote(uint256[5] candidates_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_, uint256 nullifierHash_) returns()
func (_Voting *VotingTransactor) Vote(opts *bind.TransactOpts, candidates_ [5]*big.Int, zkPoints_ VerifierHelperProofPoints, nullifierHash_ *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "vote", candidates_, zkPoints_, nullifierHash_)
}

// Vote is a paid mutator transaction binding the contract method 0x68a866f0.
//
// Solidity: function vote(uint256[5] candidates_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_, uint256 nullifierHash_) returns()
func (_Voting *VotingSession) Vote(candidates_ [5]*big.Int, zkPoints_ VerifierHelperProofPoints, nullifierHash_ *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, candidates_, zkPoints_, nullifierHash_)
}

// Vote is a paid mutator transaction binding the contract method 0x68a866f0.
//
// Solidity: function vote(uint256[5] candidates_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_, uint256 nullifierHash_) returns()
func (_Voting *VotingTransactorSession) Vote(candidates_ [5]*big.Int, zkPoints_ VerifierHelperProofPoints, nullifierHash_ *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, candidates_, zkPoints_, nullifierHash_)
}

// VotingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Voting contract.
type VotingInitializedIterator struct {
	Event *VotingInitialized // Event containing the contract specifics and raw log

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
func (it *VotingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingInitialized)
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
		it.Event = new(VotingInitialized)
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
func (it *VotingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingInitialized represents a Initialized event raised by the Voting contract.
type VotingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voting *VotingFilterer) FilterInitialized(opts *bind.FilterOpts) (*VotingInitializedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VotingInitializedIterator{contract: _Voting.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voting *VotingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VotingInitialized) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingInitialized)
				if err := _Voting.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voting *VotingFilterer) ParseInitialized(log types.Log) (*VotingInitialized, error) {
	event := new(VotingInitialized)
	if err := _Voting.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
