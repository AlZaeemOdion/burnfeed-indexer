// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// BurnFeedProtocolClientMetaData contains all meta data concerning the BurnFeedProtocolClient contract.
var BurnFeedProtocolClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"}],\"name\":\"Actions\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"PubKey\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"burn\",\"type\":\"uint256\"}],\"name\":\"publishActions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"registerPubKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userPubkeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkeys\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BurnFeedProtocolClientABI is the input ABI used to generate the binding from.
// Deprecated: Use BurnFeedProtocolClientMetaData.ABI instead.
var BurnFeedProtocolClientABI = BurnFeedProtocolClientMetaData.ABI

// BurnFeedProtocolClient is an auto generated Go binding around an Ethereum contract.
type BurnFeedProtocolClient struct {
	BurnFeedProtocolClientCaller     // Read-only binding to the contract
	BurnFeedProtocolClientTransactor // Write-only binding to the contract
	BurnFeedProtocolClientFilterer   // Log filterer for contract events
}

// BurnFeedProtocolClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnFeedProtocolClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnFeedProtocolClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnFeedProtocolClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnFeedProtocolClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnFeedProtocolClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnFeedProtocolClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnFeedProtocolClientSession struct {
	Contract     *BurnFeedProtocolClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BurnFeedProtocolClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnFeedProtocolClientCallerSession struct {
	Contract *BurnFeedProtocolClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// BurnFeedProtocolClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnFeedProtocolClientTransactorSession struct {
	Contract     *BurnFeedProtocolClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// BurnFeedProtocolClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnFeedProtocolClientRaw struct {
	Contract *BurnFeedProtocolClient // Generic contract binding to access the raw methods on
}

// BurnFeedProtocolClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnFeedProtocolClientCallerRaw struct {
	Contract *BurnFeedProtocolClientCaller // Generic read-only contract binding to access the raw methods on
}

// BurnFeedProtocolClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnFeedProtocolClientTransactorRaw struct {
	Contract *BurnFeedProtocolClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnFeedProtocolClient creates a new instance of BurnFeedProtocolClient, bound to a specific deployed contract.
func NewBurnFeedProtocolClient(address common.Address, backend bind.ContractBackend) (*BurnFeedProtocolClient, error) {
	contract, err := bindBurnFeedProtocolClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnFeedProtocolClient{BurnFeedProtocolClientCaller: BurnFeedProtocolClientCaller{contract: contract}, BurnFeedProtocolClientTransactor: BurnFeedProtocolClientTransactor{contract: contract}, BurnFeedProtocolClientFilterer: BurnFeedProtocolClientFilterer{contract: contract}}, nil
}

// NewBurnFeedProtocolClientCaller creates a new read-only instance of BurnFeedProtocolClient, bound to a specific deployed contract.
func NewBurnFeedProtocolClientCaller(address common.Address, caller bind.ContractCaller) (*BurnFeedProtocolClientCaller, error) {
	contract, err := bindBurnFeedProtocolClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnFeedProtocolClientCaller{contract: contract}, nil
}

// NewBurnFeedProtocolClientTransactor creates a new write-only instance of BurnFeedProtocolClient, bound to a specific deployed contract.
func NewBurnFeedProtocolClientTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnFeedProtocolClientTransactor, error) {
	contract, err := bindBurnFeedProtocolClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnFeedProtocolClientTransactor{contract: contract}, nil
}

// NewBurnFeedProtocolClientFilterer creates a new log filterer instance of BurnFeedProtocolClient, bound to a specific deployed contract.
func NewBurnFeedProtocolClientFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnFeedProtocolClientFilterer, error) {
	contract, err := bindBurnFeedProtocolClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnFeedProtocolClientFilterer{contract: contract}, nil
}

// bindBurnFeedProtocolClient binds a generic wrapper to an already deployed contract.
func bindBurnFeedProtocolClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnFeedProtocolClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnFeedProtocolClient *BurnFeedProtocolClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnFeedProtocolClient.Contract.BurnFeedProtocolClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnFeedProtocolClient *BurnFeedProtocolClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnFeedProtocolClient.Contract.BurnFeedProtocolClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnFeedProtocolClient *BurnFeedProtocolClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnFeedProtocolClient.Contract.BurnFeedProtocolClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnFeedProtocolClient *BurnFeedProtocolClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnFeedProtocolClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnFeedProtocolClient *BurnFeedProtocolClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnFeedProtocolClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnFeedProtocolClient *BurnFeedProtocolClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnFeedProtocolClient.Contract.contract.Transact(opts, method, params...)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnFeedProtocolClient.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientSession) Token() (common.Address, error) {
	return _BurnFeedProtocolClient.Contract.Token(&_BurnFeedProtocolClient.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientCallerSession) Token() (common.Address, error) {
	return _BurnFeedProtocolClient.Contract.Token(&_BurnFeedProtocolClient.CallOpts)
}

// UserPubkeys is a free data retrieval call binding the contract method 0x4f9475ff.
//
// Solidity: function userPubkeys(address user, uint256 ) view returns(bytes pubkeys)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientCaller) UserPubkeys(opts *bind.CallOpts, user common.Address, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _BurnFeedProtocolClient.contract.Call(opts, &out, "userPubkeys", user, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// UserPubkeys is a free data retrieval call binding the contract method 0x4f9475ff.
//
// Solidity: function userPubkeys(address user, uint256 ) view returns(bytes pubkeys)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientSession) UserPubkeys(user common.Address, arg1 *big.Int) ([]byte, error) {
	return _BurnFeedProtocolClient.Contract.UserPubkeys(&_BurnFeedProtocolClient.CallOpts, user, arg1)
}

// UserPubkeys is a free data retrieval call binding the contract method 0x4f9475ff.
//
// Solidity: function userPubkeys(address user, uint256 ) view returns(bytes pubkeys)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientCallerSession) UserPubkeys(user common.Address, arg1 *big.Int) ([]byte, error) {
	return _BurnFeedProtocolClient.Contract.UserPubkeys(&_BurnFeedProtocolClient.CallOpts, user, arg1)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnFeedProtocolClient.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientSession) Vault() (common.Address, error) {
	return _BurnFeedProtocolClient.Contract.Vault(&_BurnFeedProtocolClient.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientCallerSession) Vault() (common.Address, error) {
	return _BurnFeedProtocolClient.Contract.Vault(&_BurnFeedProtocolClient.CallOpts)
}

// PublishActions is a paid mutator transaction binding the contract method 0x54b9c3b9.
//
// Solidity: function publishActions(string uri, uint256 burn) returns()
func (_BurnFeedProtocolClient *BurnFeedProtocolClientTransactor) PublishActions(opts *bind.TransactOpts, uri string, burn *big.Int) (*types.Transaction, error) {
	return _BurnFeedProtocolClient.contract.Transact(opts, "publishActions", uri, burn)
}

// PublishActions is a paid mutator transaction binding the contract method 0x54b9c3b9.
//
// Solidity: function publishActions(string uri, uint256 burn) returns()
func (_BurnFeedProtocolClient *BurnFeedProtocolClientSession) PublishActions(uri string, burn *big.Int) (*types.Transaction, error) {
	return _BurnFeedProtocolClient.Contract.PublishActions(&_BurnFeedProtocolClient.TransactOpts, uri, burn)
}

// PublishActions is a paid mutator transaction binding the contract method 0x54b9c3b9.
//
// Solidity: function publishActions(string uri, uint256 burn) returns()
func (_BurnFeedProtocolClient *BurnFeedProtocolClientTransactorSession) PublishActions(uri string, burn *big.Int) (*types.Transaction, error) {
	return _BurnFeedProtocolClient.Contract.PublishActions(&_BurnFeedProtocolClient.TransactOpts, uri, burn)
}

// RegisterPubKey is a paid mutator transaction binding the contract method 0x3a7bf995.
//
// Solidity: function registerPubKey(bytes pubkey) returns()
func (_BurnFeedProtocolClient *BurnFeedProtocolClientTransactor) RegisterPubKey(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
	return _BurnFeedProtocolClient.contract.Transact(opts, "registerPubKey", pubkey)
}

// RegisterPubKey is a paid mutator transaction binding the contract method 0x3a7bf995.
//
// Solidity: function registerPubKey(bytes pubkey) returns()
func (_BurnFeedProtocolClient *BurnFeedProtocolClientSession) RegisterPubKey(pubkey []byte) (*types.Transaction, error) {
	return _BurnFeedProtocolClient.Contract.RegisterPubKey(&_BurnFeedProtocolClient.TransactOpts, pubkey)
}

// RegisterPubKey is a paid mutator transaction binding the contract method 0x3a7bf995.
//
// Solidity: function registerPubKey(bytes pubkey) returns()
func (_BurnFeedProtocolClient *BurnFeedProtocolClientTransactorSession) RegisterPubKey(pubkey []byte) (*types.Transaction, error) {
	return _BurnFeedProtocolClient.Contract.RegisterPubKey(&_BurnFeedProtocolClient.TransactOpts, pubkey)
}

// BurnFeedProtocolClientActionsIterator is returned from FilterActions and is used to iterate over the raw logs and unpacked data for Actions events raised by the BurnFeedProtocolClient contract.
type BurnFeedProtocolClientActionsIterator struct {
	Event *BurnFeedProtocolClientActions // Event containing the contract specifics and raw log

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
func (it *BurnFeedProtocolClientActionsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFeedProtocolClientActions)
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
		it.Event = new(BurnFeedProtocolClientActions)
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
func (it *BurnFeedProtocolClientActionsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnFeedProtocolClientActionsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnFeedProtocolClientActions represents a Actions event raised by the BurnFeedProtocolClient contract.
type BurnFeedProtocolClientActions struct {
	User common.Address
	Uri  string
	Burn *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterActions is a free log retrieval operation binding the contract event 0x094301d9f9d11fc5e8c248a58b65af5831c7efcd896c53522ccbe3be619b21fd.
//
// Solidity: event Actions(address indexed user, string uri, uint256 burn)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientFilterer) FilterActions(opts *bind.FilterOpts, user []common.Address) (*BurnFeedProtocolClientActionsIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BurnFeedProtocolClient.contract.FilterLogs(opts, "Actions", userRule)
	if err != nil {
		return nil, err
	}
	return &BurnFeedProtocolClientActionsIterator{contract: _BurnFeedProtocolClient.contract, event: "Actions", logs: logs, sub: sub}, nil
}

// WatchActions is a free log subscription operation binding the contract event 0x094301d9f9d11fc5e8c248a58b65af5831c7efcd896c53522ccbe3be619b21fd.
//
// Solidity: event Actions(address indexed user, string uri, uint256 burn)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientFilterer) WatchActions(opts *bind.WatchOpts, sink chan<- *BurnFeedProtocolClientActions, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BurnFeedProtocolClient.contract.WatchLogs(opts, "Actions", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnFeedProtocolClientActions)
				if err := _BurnFeedProtocolClient.contract.UnpackLog(event, "Actions", log); err != nil {
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

// ParseActions is a log parse operation binding the contract event 0x094301d9f9d11fc5e8c248a58b65af5831c7efcd896c53522ccbe3be619b21fd.
//
// Solidity: event Actions(address indexed user, string uri, uint256 burn)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientFilterer) ParseActions(log types.Log) (*BurnFeedProtocolClientActions, error) {
	event := new(BurnFeedProtocolClientActions)
	if err := _BurnFeedProtocolClient.contract.UnpackLog(event, "Actions", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnFeedProtocolClientPubKeyIterator is returned from FilterPubKey and is used to iterate over the raw logs and unpacked data for PubKey events raised by the BurnFeedProtocolClient contract.
type BurnFeedProtocolClientPubKeyIterator struct {
	Event *BurnFeedProtocolClientPubKey // Event containing the contract specifics and raw log

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
func (it *BurnFeedProtocolClientPubKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFeedProtocolClientPubKey)
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
		it.Event = new(BurnFeedProtocolClientPubKey)
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
func (it *BurnFeedProtocolClientPubKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnFeedProtocolClientPubKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnFeedProtocolClientPubKey represents a PubKey event raised by the BurnFeedProtocolClient contract.
type BurnFeedProtocolClientPubKey struct {
	User   common.Address
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPubKey is a free log retrieval operation binding the contract event 0x0ceb751f31e47453ca6cc79585f90f5068e6b6129b4a19aa7a83610acaab4a5d.
//
// Solidity: event PubKey(address indexed user, bytes pubkey)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientFilterer) FilterPubKey(opts *bind.FilterOpts, user []common.Address) (*BurnFeedProtocolClientPubKeyIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BurnFeedProtocolClient.contract.FilterLogs(opts, "PubKey", userRule)
	if err != nil {
		return nil, err
	}
	return &BurnFeedProtocolClientPubKeyIterator{contract: _BurnFeedProtocolClient.contract, event: "PubKey", logs: logs, sub: sub}, nil
}

// WatchPubKey is a free log subscription operation binding the contract event 0x0ceb751f31e47453ca6cc79585f90f5068e6b6129b4a19aa7a83610acaab4a5d.
//
// Solidity: event PubKey(address indexed user, bytes pubkey)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientFilterer) WatchPubKey(opts *bind.WatchOpts, sink chan<- *BurnFeedProtocolClientPubKey, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BurnFeedProtocolClient.contract.WatchLogs(opts, "PubKey", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnFeedProtocolClientPubKey)
				if err := _BurnFeedProtocolClient.contract.UnpackLog(event, "PubKey", log); err != nil {
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

// ParsePubKey is a log parse operation binding the contract event 0x0ceb751f31e47453ca6cc79585f90f5068e6b6129b4a19aa7a83610acaab4a5d.
//
// Solidity: event PubKey(address indexed user, bytes pubkey)
func (_BurnFeedProtocolClient *BurnFeedProtocolClientFilterer) ParsePubKey(log types.Log) (*BurnFeedProtocolClientPubKey, error) {
	event := new(BurnFeedProtocolClientPubKey)
	if err := _BurnFeedProtocolClient.contract.UnpackLog(event, "PubKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
