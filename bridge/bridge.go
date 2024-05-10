// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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
)

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"RequestBridge\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"etherToll\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inKindFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"requestBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newChargeTime\",\"type\":\"bool\"}],\"name\":\"setChargeTime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"inKind\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"flat\",\"type\":\"bool\"}],\"name\":\"setFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToChainToEtherTollFlat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToChainToEtherTollRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToChainToInKindFeeFlat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToChainToInKindFeeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506118c98061005b5f395ff3fe60806040526004361061009f575f3560e01c80637c7f3c89116100635780637c7f3c89146101d55780638f28397014610211578063a1db97821461024d578063c7940bf314610289578063e031c552146102c5578063f4d8a65514610301576100a6565b80631a2873d6146100a85780633bed33ce146100e45780633ee6024b146101205780635546299e1461015c57806372f3af0614610199576100a6565b366100a657005b005b3480156100b3575f80fd5b506100ce60048036038101906100c99190611152565b610331565b6040516100db919061119f565b60405180910390f35b3480156100ef575f80fd5b5061010a600480360381019061010591906111b8565b610351565b60405161011791906111fd565b60405180910390f35b34801561012b575f80fd5b5061014660048036038101906101419190611152565b610495565b604051610153919061119f565b60405180910390f35b348015610167575f80fd5b50610182600480360381019061017d9190611216565b6104b5565b604051610190929190611266565b60405180910390f35b3480156101a4575f80fd5b506101bf60048036038101906101ba9190611152565b61064d565b6040516101cc919061119f565b60405180910390f35b3480156101e0575f80fd5b506101fb60048036038101906101f69190611152565b61066d565b604051610208919061119f565b60405180910390f35b34801561021c575f80fd5b506102376004803603810190610232919061128d565b61068d565b60405161024491906111fd565b60405180910390f35b348015610258575f80fd5b50610273600480360381019061026e9190611152565b610764565b60405161028091906111fd565b60405180910390f35b348015610294575f80fd5b506102af60048036038101906102aa91906112e2565b61092e565b6040516102bc91906111fd565b60405180910390f35b3480156102d0575f80fd5b506102eb60048036038101906102e6919061130d565b6109df565b6040516102f891906111fd565b60405180910390f35b61031b600480360381019061031691906113e5565b610bdf565b60405161032891906111fd565b60405180910390f35b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d7906114d5565b60405180910390fd5b5f803373ffffffffffffffffffffffffffffffffffffffff168460405161040690611520565b5f6040518083038185875af1925050503d805f8114610440576040519150601f19603f3d011682016040523d82523d5f602084013e610445565b606091505b50915091508161048a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104819061157e565b60405180910390fd5b600192505050919050565b6004602052815f5260405f20602052805f5260405f205f91509150505481565b5f8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8581526020019081526020015f20549150670de0b6b3a764000060025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8681526020019081526020015f20548661056791906115c9565b6105719190611637565b8261057c9190611667565b915060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8581526020019081526020015f20549050670de0b6b3a764000060045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8681526020019081526020015f20548661062e91906115c9565b6106389190611637565b816106439190611667565b9050935093915050565b6003602052815f5260405f20602052805f5260405f205f91509150505481565b6001602052815f5260405f20602052805f5260405f205f91509150505481565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461071c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610713906114d5565b60405180910390fd5b815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019050919050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ea906114d5565b60405180910390fd5b73dac17f958d2ee523a2206206994597c13d831ec773ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108a7578273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b81526004016108759291906116a9565b5f604051808303815f87803b15801561088c575f80fd5b505af115801561089e573d5f803e3d5ffd5b50505050610924565b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b81526004016108e29291906116a9565b6020604051808303815f875af11580156108fe573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061092291906116e4565b505b6001905092915050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b4906114d5565b60405180910390fd5b815f60146101000a81548160ff02191690831515021790555060019050919050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a6e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a65906114d5565b60405180910390fd5b8215610b27578115610ad0578460035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f2081905550610b22565b8460045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f20819055505b610bd6565b8115610b83578460015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f2081905550610bd5565b8460025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f20819055505b5b95945050505050565b5f8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f20549050670de0b6b3a764000060025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8981526020019081526020015f205489610c9191906115c9565b610c9b9190611637565b81610ca69190611667565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610d30578781610ce79190611667565b9050803411610d2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2290611759565b60405180910390fd5b610d73565b803411610d72576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d6990611759565b60405180910390fd5b5b5f60149054906101000a900460ff1615610f1f575f60035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8981526020019081526020015f20549050670de0b6b3a764000060045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8a81526020019081526020015f20548a610e3891906115c9565b610e429190611637565b81610e4d9190611667565b90508373ffffffffffffffffffffffffffffffffffffffff166323b872dd33308c85610e799190611667565b6040518463ffffffff1660e01b8152600401610e9793929190611777565b6020604051808303815f875af1158015610eb3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ed791906116e4565b507f7b79e3f7a30b61eee6994dca00804d4dc8b92ec69c59ebe012e4ae274e58193e89468a8a8a8a604051610f1196959493929190611806565b60405180910390a1506110ae565b8273ffffffffffffffffffffffffffffffffffffffff166323b872dd33308b6040518463ffffffff1660e01b8152600401610f5c93929190611777565b6020604051808303815f875af1158015610f78573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f9c91906116e4565b50670de0b6b3a764000060045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8981526020019081526020015f205489610ffe91906115c9565b6110089190611637565b886110139190611860565b97507f7b79e3f7a30b61eee6994dca00804d4dc8b92ec69c59ebe012e4ae274e58193e60035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8981526020019081526020015f20548961108e9190611860565b46898989896040516110a596959493929190611806565b60405180910390a15b60019150509695505050505050565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6110ee826110c5565b9050919050565b6110fe816110e4565b8114611108575f80fd5b50565b5f81359050611119816110f5565b92915050565b5f819050919050565b6111318161111f565b811461113b575f80fd5b50565b5f8135905061114c81611128565b92915050565b5f8060408385031215611168576111676110bd565b5b5f6111758582860161110b565b92505060206111868582860161113e565b9150509250929050565b6111998161111f565b82525050565b5f6020820190506111b25f830184611190565b92915050565b5f602082840312156111cd576111cc6110bd565b5b5f6111da8482850161113e565b91505092915050565b5f8115159050919050565b6111f7816111e3565b82525050565b5f6020820190506112105f8301846111ee565b92915050565b5f805f6060848603121561122d5761122c6110bd565b5b5f61123a8682870161113e565b935050602061124b8682870161113e565b925050604061125c8682870161110b565b9150509250925092565b5f6040820190506112795f830185611190565b6112866020830184611190565b9392505050565b5f602082840312156112a2576112a16110bd565b5b5f6112af8482850161110b565b91505092915050565b6112c1816111e3565b81146112cb575f80fd5b50565b5f813590506112dc816112b8565b92915050565b5f602082840312156112f7576112f66110bd565b5b5f611304848285016112ce565b91505092915050565b5f805f805f60a08688031215611326576113256110bd565b5b5f6113338882890161113e565b95505060206113448882890161113e565b94505060406113558882890161110b565b9350506060611366888289016112ce565b9250506080611377888289016112ce565b9150509295509295909350565b5f80fd5b5f80fd5b5f80fd5b5f8083601f8401126113a5576113a4611384565b5b8235905067ffffffffffffffff8111156113c2576113c1611388565b5b6020830191508360018202830111156113de576113dd61138c565b5b9250929050565b5f805f805f8060a087890312156113ff576113fe6110bd565b5b5f61140c89828a0161113e565b965050602061141d89828a0161113e565b955050604061142e89828a0161110b565b945050606087013567ffffffffffffffff81111561144f5761144e6110c1565b5b61145b89828a01611390565b9350935050608061146e89828a0161110b565b9150509295509295509295565b5f82825260208201905092915050565b7f43616e206f6e6c792062652063616c6c65642062792061646d696e00000000005f82015250565b5f6114bf601b8361147b565b91506114ca8261148b565b602082019050919050565b5f6020820190508181035f8301526114ec816114b3565b9050919050565b5f81905092915050565b50565b5f61150b5f836114f3565b9150611516826114fd565b5f82019050919050565b5f61152a82611500565b9150819050919050565b7f4661696c656420746f2073656e642045746865720000000000000000000000005f82015250565b5f61156860148361147b565b915061157382611534565b602082019050919050565b5f6020820190508181035f8301526115958161155c565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6115d38261111f565b91506115de8361111f565b92508282026115ec8161111f565b915082820484148315176116035761160261159c565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6116418261111f565b915061164c8361111f565b92508261165c5761165b61160a565b5b828204905092915050565b5f6116718261111f565b915061167c8361111f565b92508282019050808211156116945761169361159c565b5b92915050565b6116a3816110e4565b82525050565b5f6040820190506116bc5f83018561169a565b6116c96020830184611190565b9392505050565b5f815190506116de816112b8565b92915050565b5f602082840312156116f9576116f86110bd565b5b5f611706848285016116d0565b91505092915050565b7f546f6c6c207061696420697320696e73756666696369656e74000000000000005f82015250565b5f61174360198361147b565b915061174e8261170f565b602082019050919050565b5f6020820190508181035f83015261177081611737565b9050919050565b5f60608201905061178a5f83018661169a565b611797602083018561169a565b6117a46040830184611190565b949350505050565b5f82825260208201905092915050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f6117e583856117ac565b93506117f28385846117bc565b6117fb836117ca565b840190509392505050565b5f60a0820190506118195f830189611190565b6118266020830188611190565b6118336040830187611190565b611840606083018661169a565b81810360808301526118538184866117da565b9050979650505050505050565b5f61186a8261111f565b91506118758361111f565b925082820390508181111561188d5761188c61159c565b5b9291505056fea2646970667358221220c286d8c1c02c0f0c2eeb5e91eb77a7d2831544560aa27a34be31727a23644a6464736f6c63430008190033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// EstimateFee is a free data retrieval call binding the contract method 0x5546299e.
//
// Solidity: function estimateFee(uint256 amount, uint256 destinationChainID, address token) view returns(uint256 etherToll, uint256 inKindFee)
func (_Bridge *BridgeCaller) EstimateFee(opts *bind.CallOpts, amount *big.Int, destinationChainID *big.Int, token common.Address) (struct {
	EtherToll *big.Int
	InKindFee *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "estimateFee", amount, destinationChainID, token)

	outstruct := new(struct {
		EtherToll *big.Int
		InKindFee *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EtherToll = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.InKindFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateFee is a free data retrieval call binding the contract method 0x5546299e.
//
// Solidity: function estimateFee(uint256 amount, uint256 destinationChainID, address token) view returns(uint256 etherToll, uint256 inKindFee)
func (_Bridge *BridgeSession) EstimateFee(amount *big.Int, destinationChainID *big.Int, token common.Address) (struct {
	EtherToll *big.Int
	InKindFee *big.Int
}, error) {
	return _Bridge.Contract.EstimateFee(&_Bridge.CallOpts, amount, destinationChainID, token)
}

// EstimateFee is a free data retrieval call binding the contract method 0x5546299e.
//
// Solidity: function estimateFee(uint256 amount, uint256 destinationChainID, address token) view returns(uint256 etherToll, uint256 inKindFee)
func (_Bridge *BridgeCallerSession) EstimateFee(amount *big.Int, destinationChainID *big.Int, token common.Address) (struct {
	EtherToll *big.Int
	InKindFee *big.Int
}, error) {
	return _Bridge.Contract.EstimateFee(&_Bridge.CallOpts, amount, destinationChainID, token)
}

// TokenToChainToEtherTollFlat is a free data retrieval call binding the contract method 0x7c7f3c89.
//
// Solidity: function tokenToChainToEtherTollFlat(address , uint256 ) view returns(uint256)
func (_Bridge *BridgeCaller) TokenToChainToEtherTollFlat(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tokenToChainToEtherTollFlat", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToChainToEtherTollFlat is a free data retrieval call binding the contract method 0x7c7f3c89.
//
// Solidity: function tokenToChainToEtherTollFlat(address , uint256 ) view returns(uint256)
func (_Bridge *BridgeSession) TokenToChainToEtherTollFlat(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.TokenToChainToEtherTollFlat(&_Bridge.CallOpts, arg0, arg1)
}

// TokenToChainToEtherTollFlat is a free data retrieval call binding the contract method 0x7c7f3c89.
//
// Solidity: function tokenToChainToEtherTollFlat(address , uint256 ) view returns(uint256)
func (_Bridge *BridgeCallerSession) TokenToChainToEtherTollFlat(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.TokenToChainToEtherTollFlat(&_Bridge.CallOpts, arg0, arg1)
}

// TokenToChainToEtherTollRatio is a free data retrieval call binding the contract method 0x1a2873d6.
//
// Solidity: function tokenToChainToEtherTollRatio(address , uint256 ) view returns(uint256)
func (_Bridge *BridgeCaller) TokenToChainToEtherTollRatio(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tokenToChainToEtherTollRatio", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToChainToEtherTollRatio is a free data retrieval call binding the contract method 0x1a2873d6.
//
// Solidity: function tokenToChainToEtherTollRatio(address , uint256 ) view returns(uint256)
func (_Bridge *BridgeSession) TokenToChainToEtherTollRatio(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.TokenToChainToEtherTollRatio(&_Bridge.CallOpts, arg0, arg1)
}

// TokenToChainToEtherTollRatio is a free data retrieval call binding the contract method 0x1a2873d6.
//
// Solidity: function tokenToChainToEtherTollRatio(address , uint256 ) view returns(uint256)
func (_Bridge *BridgeCallerSession) TokenToChainToEtherTollRatio(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.TokenToChainToEtherTollRatio(&_Bridge.CallOpts, arg0, arg1)
}

// TokenToChainToInKindFeeFlat is a free data retrieval call binding the contract method 0x72f3af06.
//
// Solidity: function tokenToChainToInKindFeeFlat(address , uint256 ) view returns(uint256)
func (_Bridge *BridgeCaller) TokenToChainToInKindFeeFlat(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tokenToChainToInKindFeeFlat", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToChainToInKindFeeFlat is a free data retrieval call binding the contract method 0x72f3af06.
//
// Solidity: function tokenToChainToInKindFeeFlat(address , uint256 ) view returns(uint256)
func (_Bridge *BridgeSession) TokenToChainToInKindFeeFlat(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.TokenToChainToInKindFeeFlat(&_Bridge.CallOpts, arg0, arg1)
}

// TokenToChainToInKindFeeFlat is a free data retrieval call binding the contract method 0x72f3af06.
//
// Solidity: function tokenToChainToInKindFeeFlat(address , uint256 ) view returns(uint256)
func (_Bridge *BridgeCallerSession) TokenToChainToInKindFeeFlat(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.TokenToChainToInKindFeeFlat(&_Bridge.CallOpts, arg0, arg1)
}

// TokenToChainToInKindFeeRatio is a free data retrieval call binding the contract method 0x3ee6024b.
//
// Solidity: function tokenToChainToInKindFeeRatio(address , uint256 ) view returns(uint256)
func (_Bridge *BridgeCaller) TokenToChainToInKindFeeRatio(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tokenToChainToInKindFeeRatio", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToChainToInKindFeeRatio is a free data retrieval call binding the contract method 0x3ee6024b.
//
// Solidity: function tokenToChainToInKindFeeRatio(address , uint256 ) view returns(uint256)
func (_Bridge *BridgeSession) TokenToChainToInKindFeeRatio(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.TokenToChainToInKindFeeRatio(&_Bridge.CallOpts, arg0, arg1)
}

// TokenToChainToInKindFeeRatio is a free data retrieval call binding the contract method 0x3ee6024b.
//
// Solidity: function tokenToChainToInKindFeeRatio(address , uint256 ) view returns(uint256)
func (_Bridge *BridgeCallerSession) TokenToChainToInKindFeeRatio(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.TokenToChainToInKindFeeRatio(&_Bridge.CallOpts, arg0, arg1)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns(bool)
func (_Bridge *BridgeTransactor) ChangeAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "changeAdmin", _admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns(bool)
func (_Bridge *BridgeSession) ChangeAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeAdmin(&_Bridge.TransactOpts, _admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns(bool)
func (_Bridge *BridgeTransactorSession) ChangeAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeAdmin(&_Bridge.TransactOpts, _admin)
}

// RequestBridge is a paid mutator transaction binding the contract method 0xf4d8a655.
//
// Solidity: function requestBridge(uint256 amount, uint256 destinationChainID, address destinationAddress, bytes extraData, address token) payable returns(bool)
func (_Bridge *BridgeTransactor) RequestBridge(opts *bind.TransactOpts, amount *big.Int, destinationChainID *big.Int, destinationAddress common.Address, extraData []byte, token common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "requestBridge", amount, destinationChainID, destinationAddress, extraData, token)
}

// RequestBridge is a paid mutator transaction binding the contract method 0xf4d8a655.
//
// Solidity: function requestBridge(uint256 amount, uint256 destinationChainID, address destinationAddress, bytes extraData, address token) payable returns(bool)
func (_Bridge *BridgeSession) RequestBridge(amount *big.Int, destinationChainID *big.Int, destinationAddress common.Address, extraData []byte, token common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RequestBridge(&_Bridge.TransactOpts, amount, destinationChainID, destinationAddress, extraData, token)
}

// RequestBridge is a paid mutator transaction binding the contract method 0xf4d8a655.
//
// Solidity: function requestBridge(uint256 amount, uint256 destinationChainID, address destinationAddress, bytes extraData, address token) payable returns(bool)
func (_Bridge *BridgeTransactorSession) RequestBridge(amount *big.Int, destinationChainID *big.Int, destinationAddress common.Address, extraData []byte, token common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RequestBridge(&_Bridge.TransactOpts, amount, destinationChainID, destinationAddress, extraData, token)
}

// SetChargeTime is a paid mutator transaction binding the contract method 0xc7940bf3.
//
// Solidity: function setChargeTime(bool newChargeTime) returns(bool)
func (_Bridge *BridgeTransactor) SetChargeTime(opts *bind.TransactOpts, newChargeTime bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setChargeTime", newChargeTime)
}

// SetChargeTime is a paid mutator transaction binding the contract method 0xc7940bf3.
//
// Solidity: function setChargeTime(bool newChargeTime) returns(bool)
func (_Bridge *BridgeSession) SetChargeTime(newChargeTime bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetChargeTime(&_Bridge.TransactOpts, newChargeTime)
}

// SetChargeTime is a paid mutator transaction binding the contract method 0xc7940bf3.
//
// Solidity: function setChargeTime(bool newChargeTime) returns(bool)
func (_Bridge *BridgeTransactorSession) SetChargeTime(newChargeTime bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetChargeTime(&_Bridge.TransactOpts, newChargeTime)
}

// SetFees is a paid mutator transaction binding the contract method 0xe031c552.
//
// Solidity: function setFees(uint256 destinationChain, uint256 fee, address token, bool inKind, bool flat) returns(bool)
func (_Bridge *BridgeTransactor) SetFees(opts *bind.TransactOpts, destinationChain *big.Int, fee *big.Int, token common.Address, inKind bool, flat bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setFees", destinationChain, fee, token, inKind, flat)
}

// SetFees is a paid mutator transaction binding the contract method 0xe031c552.
//
// Solidity: function setFees(uint256 destinationChain, uint256 fee, address token, bool inKind, bool flat) returns(bool)
func (_Bridge *BridgeSession) SetFees(destinationChain *big.Int, fee *big.Int, token common.Address, inKind bool, flat bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetFees(&_Bridge.TransactOpts, destinationChain, fee, token, inKind, flat)
}

// SetFees is a paid mutator transaction binding the contract method 0xe031c552.
//
// Solidity: function setFees(uint256 destinationChain, uint256 fee, address token, bool inKind, bool flat) returns(bool)
func (_Bridge *BridgeTransactorSession) SetFees(destinationChain *big.Int, fee *big.Int, token common.Address, inKind bool, flat bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetFees(&_Bridge.TransactOpts, destinationChain, fee, token, inKind, flat)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address token, uint256 amountToWithdraw) returns(bool)
func (_Bridge *BridgeTransactor) WithdrawERC20(opts *bind.TransactOpts, token common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC20", token, amountToWithdraw)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address token, uint256 amountToWithdraw) returns(bool)
func (_Bridge *BridgeSession) WithdrawERC20(token common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20(&_Bridge.TransactOpts, token, amountToWithdraw)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address token, uint256 amountToWithdraw) returns(bool)
func (_Bridge *BridgeTransactorSession) WithdrawERC20(token common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20(&_Bridge.TransactOpts, token, amountToWithdraw)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amountToWithdraw) returns(bool)
func (_Bridge *BridgeTransactor) WithdrawEther(opts *bind.TransactOpts, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawEther", amountToWithdraw)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amountToWithdraw) returns(bool)
func (_Bridge *BridgeSession) WithdrawEther(amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawEther(&_Bridge.TransactOpts, amountToWithdraw)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 amountToWithdraw) returns(bool)
func (_Bridge *BridgeTransactorSession) WithdrawEther(amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawEther(&_Bridge.TransactOpts, amountToWithdraw)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Bridge *BridgeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Bridge *BridgeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Fallback(&_Bridge.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Bridge *BridgeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Fallback(&_Bridge.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// BridgeRequestBridgeIterator is returned from FilterRequestBridge and is used to iterate over the raw logs and unpacked data for RequestBridge events raised by the Bridge contract.
type BridgeRequestBridgeIterator struct {
	Event *BridgeRequestBridge // Event containing the contract specifics and raw log

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
func (it *BridgeRequestBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRequestBridge)
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
		it.Event = new(BridgeRequestBridge)
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
func (it *BridgeRequestBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRequestBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRequestBridge represents a RequestBridge event raised by the Bridge contract.
type BridgeRequestBridge struct {
	Amount             *big.Int
	OriginChainID      *big.Int
	DestinationChainID *big.Int
	Recipient          common.Address
	ExtraData          []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRequestBridge is a free log retrieval operation binding the contract event 0x7b79e3f7a30b61eee6994dca00804d4dc8b92ec69c59ebe012e4ae274e58193e.
//
// Solidity: event RequestBridge(uint256 amount, uint256 originChainID, uint256 destinationChainID, address recipient, bytes extraData)
func (_Bridge *BridgeFilterer) FilterRequestBridge(opts *bind.FilterOpts) (*BridgeRequestBridgeIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RequestBridge")
	if err != nil {
		return nil, err
	}
	return &BridgeRequestBridgeIterator{contract: _Bridge.contract, event: "RequestBridge", logs: logs, sub: sub}, nil
}

// WatchRequestBridge is a free log subscription operation binding the contract event 0x7b79e3f7a30b61eee6994dca00804d4dc8b92ec69c59ebe012e4ae274e58193e.
//
// Solidity: event RequestBridge(uint256 amount, uint256 originChainID, uint256 destinationChainID, address recipient, bytes extraData)
func (_Bridge *BridgeFilterer) WatchRequestBridge(opts *bind.WatchOpts, sink chan<- *BridgeRequestBridge) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RequestBridge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRequestBridge)
				if err := _Bridge.contract.UnpackLog(event, "RequestBridge", log); err != nil {
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

// ParseRequestBridge is a log parse operation binding the contract event 0x7b79e3f7a30b61eee6994dca00804d4dc8b92ec69c59ebe012e4ae274e58193e.
//
// Solidity: event RequestBridge(uint256 amount, uint256 originChainID, uint256 destinationChainID, address recipient, bytes extraData)
func (_Bridge *BridgeFilterer) ParseRequestBridge(log types.Log) (*BridgeRequestBridge, error) {
	event := new(BridgeRequestBridge)
	if err := _Bridge.contract.UnpackLog(event, "RequestBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
