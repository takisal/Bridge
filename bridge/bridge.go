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

// BridgeRequestBridge is an auto generated low-level Go binding around an user-defined struct.
type BridgeRequestBridge struct {
	Amount             *big.Int
	OriginChainID      uint64
	DestinationChainID uint64
	Recipient          common.Address
	ExtraData          []byte
	BlockNumber        *big.Int
	Token              common.Address
}

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"checkForValidRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"originChainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destinationChainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridge.RequestBridge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"etherToll\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inKindFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destinationChainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"requestBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"originChainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destinationChainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newChargeTime\",\"type\":\"bool\"}],\"name\":\"setChargeTime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"inKind\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"flat\",\"type\":\"bool\"}],\"name\":\"setFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToChainToEtherTollFlat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToChainToEtherTollRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToChainToInKindFeeFlat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToChainToInKindFeeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506126ec8061005b5f395ff3fe6080604052600436106100c5575f3560e01c806372f3af061161007e5780638f283970116100585780638f283970146102e5578063a1db978214610321578063c7940bf31461035d578063e031c55214610399576100cc565b806372f3af061461022b5780637c7f3c891461026757806381d12c58146102a3576100cc565b80630c1fe6b2146100ce5780631a2873d61461010a57806324ca571f146101465780633bed33ce146101765780633ee6024b146101b25780635546299e146101ee576100cc565b366100cc57005b005b3480156100d9575f80fd5b506100f460048036038101906100ef9190611964565b6103d5565b6040516101019190611b08565b60405180910390f35b348015610115575f80fd5b50610130600480360381019061012b9190611b52565b610667565b60405161013d9190611b9f565b60405180910390f35b610160600480360381019061015b9190611c43565b610687565b60405161016d9190611cf3565b60405180910390f35b348015610181575f80fd5b5061019c60048036038101906101979190611964565b610efb565b6040516101a99190611cf3565b60405180910390f35b3480156101bd575f80fd5b506101d860048036038101906101d39190611b52565b61103f565b6040516101e59190611b9f565b60405180910390f35b3480156101f9575f80fd5b50610214600480360381019061020f9190611d0c565b61105f565b604051610222929190611d5c565b60405180910390f35b348015610236575f80fd5b50610251600480360381019061024c9190611b52565b6111f7565b60405161025e9190611b9f565b60405180910390f35b348015610272575f80fd5b5061028d60048036038101906102889190611b52565b611217565b60405161029a9190611b9f565b60405180910390f35b3480156102ae575f80fd5b506102c960048036038101906102c49190611964565b611237565b6040516102dc9796959493929190611de9565b60405180910390f35b3480156102f0575f80fd5b5061030b60048036038101906103069190611e5d565b611360565b6040516103189190611cf3565b60405180910390f35b34801561032c575f80fd5b5061034760048036038101906103429190611b52565b611437565b6040516103549190611cf3565b60405180910390f35b348015610368575f80fd5b50610383600480360381019061037e9190611eb2565b611601565b6040516103909190611cf3565b60405180910390f35b3480156103a4575f80fd5b506103bf60048036038101906103ba9190611edd565b6116b2565b6040516103cc9190611cf3565b60405180910390f35b6103dd6118b2565b5f60065f8481526020019081526020015f205f015411610432576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042990611fae565b60405180910390fd5b60054361043f9190611ff9565b60065f8481526020019081526020015f206004015410610494576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048b90612076565b60405180910390fd5b60065f8381526020019081526020015f206040518060e00160405290815f8201548152602001600182015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054610581906120c1565b80601f01602080910402602001604051908101604052809291908181526020018280546105ad906120c1565b80156105f85780601f106105cf576101008083540402835291602001916105f8565b820191905f5260205f20905b8154815290600101906020018083116105db57829003601f168201915b5050505050815260200160048201548152602001600582015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b6003602052815f5260405f20602052805f5260405f205f91509150505481565b5f8060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1681526020019081526020015f20549050670de0b6b3a764000060035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8967ffffffffffffffff1681526020019081526020015f20548961074d91906120f1565b610757919061215f565b81610762919061218f565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036107ec5787816107a3919061218f565b90508034116107e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107de9061220c565b60405180910390fd5b61082f565b80341161082e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108259061220c565b60405180910390fd5b5b5f881161083a575f80fd5b5f60149054906101000a900460ff1615610b96575f60045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8967ffffffffffffffff1681526020019081526020015f20549050670de0b6b3a764000060055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8a67ffffffffffffffff1681526020019081526020015f20548a61091391906120f1565b61091d919061215f565b81610928919061218f565b90508373ffffffffffffffffffffffffffffffffffffffff166323b872dd33308c85610954919061218f565b6040518463ffffffff1660e01b81526004016109729392919061222a565b6020604051808303815f875af115801561098e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109b29190612273565b506040518060e001604052808a81526020014667ffffffffffffffff1681526020018967ffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014381526020018573ffffffffffffffffffffffffffffffffffffffff1681525060065f60015481526020019081526020015f205f820151815f01556020820151816001015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816003019081610b3c9190612468565b5060a0820151816004015560c0820151816005015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505050610ed5565b8273ffffffffffffffffffffffffffffffffffffffff166323b872dd33308b6040518463ffffffff1660e01b8152600401610bd39392919061222a565b6020604051808303815f875af1158015610bef573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c139190612273565b50670de0b6b3a764000060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8967ffffffffffffffff1681526020019081526020015f205489610c7f91906120f1565b610c89919061215f565b88610c949190611ff9565b97506040518060e0016040528060045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8a67ffffffffffffffff1681526020019081526020015f20548a610d039190611ff9565b81526020014667ffffffffffffffff1681526020018867ffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff16815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014381526020018473ffffffffffffffffffffffffffffffffffffffff1681525060065f60015481526020019081526020015f205f820151815f01556020820151816001015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816003019081610e809190612468565b5060a0820151816004015560c0820151816005015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050505b60015f815480929190610ee790612537565b919050555060019150509695505050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f81906125c8565b60405180910390fd5b5f803373ffffffffffffffffffffffffffffffffffffffff1684604051610fb090612613565b5f6040518083038185875af1925050503d805f8114610fea576040519150601f19603f3d011682016040523d82523d5f602084013e610fef565b606091505b509150915081611034576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102b90612671565b60405180910390fd5b600192505050919050565b6005602052815f5260405f20602052805f5260405f205f91509150505481565b5f8060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8581526020019081526020015f20549150670de0b6b3a764000060035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8681526020019081526020015f20548661111191906120f1565b61111b919061215f565b82611126919061218f565b915060045f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8581526020019081526020015f20549050670de0b6b3a764000060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8681526020019081526020015f2054866111d891906120f1565b6111e2919061215f565b816111ed919061218f565b9050935093915050565b6004602052815f5260405f20602052805f5260405f205f91509150505481565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b6006602052805f5260405f205f91509050805f015490806001015f9054906101000a900467ffffffffffffffff16908060010160089054906101000a900467ffffffffffffffff1690806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030180546112b4906120c1565b80601f01602080910402602001604051908101604052809291908181526020018280546112e0906120c1565b801561132b5780601f106113025761010080835404028352916020019161132b565b820191905f5260205f20905b81548152906001019060200180831161130e57829003601f168201915b505050505090806004015490806005015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905087565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146113ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113e6906125c8565b60405180910390fd5b815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019050919050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146114c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114bd906125c8565b60405180910390fd5b73dac17f958d2ee523a2206206994597c13d831ec773ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361157a578273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b815260040161154892919061268f565b5f604051808303815f87803b15801561155f575f80fd5b505af1158015611571573d5f803e3d5ffd5b505050506115f7565b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b81526004016115b592919061268f565b6020604051808303815f875af11580156115d1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115f59190612273565b505b6001905092915050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611690576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611687906125c8565b60405180910390fd5b815f60146101000a81548160ff02191690831515021790555060019050919050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611741576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611738906125c8565b60405180910390fd5b82156117fa5781156117a3578460045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f20819055506117f5565b8460055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f20819055505b6118a9565b8115611856578460025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f20819055506118a8565b8460035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f20819055505b5b95945050505050565b6040518060e001604052805f81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff168152602001606081526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b5f80fd5b5f80fd5b5f819050919050565b61194381611931565b811461194d575f80fd5b50565b5f8135905061195e8161193a565b92915050565b5f6020828403121561197957611978611929565b5b5f61198684828501611950565b91505092915050565b61199881611931565b82525050565b5f67ffffffffffffffff82169050919050565b6119ba8161199e565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6119e9826119c0565b9050919050565b6119f9816119df565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611a41826119ff565b611a4b8185611a09565b9350611a5b818560208601611a19565b611a6481611a27565b840191505092915050565b5f60e083015f830151611a845f86018261198f565b506020830151611a9760208601826119b1565b506040830151611aaa60408601826119b1565b506060830151611abd60608601826119f0565b5060808301518482036080860152611ad58282611a37565b91505060a0830151611aea60a086018261198f565b5060c0830151611afd60c08601826119f0565b508091505092915050565b5f6020820190508181035f830152611b208184611a6f565b905092915050565b611b31816119df565b8114611b3b575f80fd5b50565b5f81359050611b4c81611b28565b92915050565b5f8060408385031215611b6857611b67611929565b5b5f611b7585828601611b3e565b9250506020611b8685828601611950565b9150509250929050565b611b9981611931565b82525050565b5f602082019050611bb25f830184611b90565b92915050565b611bc18161199e565b8114611bcb575f80fd5b50565b5f81359050611bdc81611bb8565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112611c0357611c02611be2565b5b8235905067ffffffffffffffff811115611c2057611c1f611be6565b5b602083019150836001820283011115611c3c57611c3b611bea565b5b9250929050565b5f805f805f8060a08789031215611c5d57611c5c611929565b5b5f611c6a89828a01611950565b9650506020611c7b89828a01611bce565b9550506040611c8c89828a01611b3e565b945050606087013567ffffffffffffffff811115611cad57611cac61192d565b5b611cb989828a01611bee565b93509350506080611ccc89828a01611b3e565b9150509295509295509295565b5f8115159050919050565b611ced81611cd9565b82525050565b5f602082019050611d065f830184611ce4565b92915050565b5f805f60608486031215611d2357611d22611929565b5b5f611d3086828701611950565b9350506020611d4186828701611950565b9250506040611d5286828701611b3e565b9150509250925092565b5f604082019050611d6f5f830185611b90565b611d7c6020830184611b90565b9392505050565b611d8c8161199e565b82525050565b611d9b816119df565b82525050565b5f82825260208201905092915050565b5f611dbb826119ff565b611dc58185611da1565b9350611dd5818560208601611a19565b611dde81611a27565b840191505092915050565b5f60e082019050611dfc5f83018a611b90565b611e096020830189611d83565b611e166040830188611d83565b611e236060830187611d92565b8181036080830152611e358186611db1565b9050611e4460a0830185611b90565b611e5160c0830184611d92565b98975050505050505050565b5f60208284031215611e7257611e71611929565b5b5f611e7f84828501611b3e565b91505092915050565b611e9181611cd9565b8114611e9b575f80fd5b50565b5f81359050611eac81611e88565b92915050565b5f60208284031215611ec757611ec6611929565b5b5f611ed484828501611e9e565b91505092915050565b5f805f805f60a08688031215611ef657611ef5611929565b5b5f611f0388828901611950565b9550506020611f1488828901611950565b9450506040611f2588828901611b3e565b9350506060611f3688828901611e9e565b9250506080611f4788828901611e9e565b9150509295509295909350565b5f82825260208201905092915050565b7f446f6573206e6f742065786973740000000000000000000000000000000000005f82015250565b5f611f98600e83611f54565b9150611fa382611f64565b602082019050919050565b5f6020820190508181035f830152611fc581611f8c565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61200382611931565b915061200e83611931565b925082820390508181111561202657612025611fcc565b5b92915050565b7f546f6f20736f6f6e0000000000000000000000000000000000000000000000005f82015250565b5f612060600883611f54565b915061206b8261202c565b602082019050919050565b5f6020820190508181035f83015261208d81612054565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806120d857607f821691505b6020821081036120eb576120ea612094565b5b50919050565b5f6120fb82611931565b915061210683611931565b925082820261211481611931565b9150828204841483151761212b5761212a611fcc565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61216982611931565b915061217483611931565b92508261218457612183612132565b5b828204905092915050565b5f61219982611931565b91506121a483611931565b92508282019050808211156121bc576121bb611fcc565b5b92915050565b7f546f6c6c207061696420697320696e73756666696369656e74000000000000005f82015250565b5f6121f6601983611f54565b9150612201826121c2565b602082019050919050565b5f6020820190508181035f830152612223816121ea565b9050919050565b5f60608201905061223d5f830186611d92565b61224a6020830185611d92565b6122576040830184611b90565b949350505050565b5f8151905061226d81611e88565b92915050565b5f6020828403121561228857612287611929565b5b5f6122958482850161225f565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026123277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826122ec565b61233186836122ec565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61236c61236761236284611931565b612349565b611931565b9050919050565b5f819050919050565b61238583612352565b61239961239182612373565b8484546122f8565b825550505050565b5f90565b6123ad6123a1565b6123b881848461237c565b505050565b5b818110156123db576123d05f826123a5565b6001810190506123be565b5050565b601f821115612420576123f1816122cb565b6123fa846122dd565b81016020851015612409578190505b61241d612415856122dd565b8301826123bd565b50505b505050565b5f82821c905092915050565b5f6124405f1984600802612425565b1980831691505092915050565b5f6124588383612431565b9150826002028217905092915050565b612471826119ff565b67ffffffffffffffff81111561248a5761248961229e565b5b61249482546120c1565b61249f8282856123df565b5f60209050601f8311600181146124d0575f84156124be578287015190505b6124c8858261244d565b86555061252f565b601f1984166124de866122cb565b5f5b82811015612505578489015182556001820191506020850194506020810190506124e0565b86831015612522578489015161251e601f891682612431565b8355505b6001600288020188555050505b505050505050565b5f61254182611931565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361257357612572611fcc565b5b600182019050919050565b7f43616e206f6e6c792062652063616c6c65642062792061646d696e00000000005f82015250565b5f6125b2601b83611f54565b91506125bd8261257e565b602082019050919050565b5f6020820190508181035f8301526125df816125a6565b9050919050565b5f81905092915050565b50565b5f6125fe5f836125e6565b9150612609826125f0565b5f82019050919050565b5f61261d826125f3565b9150819050919050565b7f4661696c656420746f2073656e642045746865720000000000000000000000005f82015250565b5f61265b601483611f54565b915061266682612627565b602082019050919050565b5f6020820190508181035f8301526126888161264f565b9050919050565b5f6040820190506126a25f830185611d92565b6126af6020830184611b90565b939250505056fea2646970667358221220f2b1eab9197e8f96b3b42d6ec87c838e5241f2922c928ba4001bed208c6a40ee64736f6c63430008190033",
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

// CheckForValidRequest is a free data retrieval call binding the contract method 0x0c1fe6b2.
//
// Solidity: function checkForValidRequest(uint256 num) view returns((uint256,uint64,uint64,address,bytes,uint256,address))
func (_Bridge *BridgeCaller) CheckForValidRequest(opts *bind.CallOpts, num *big.Int) (BridgeRequestBridge, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "checkForValidRequest", num)

	if err != nil {
		return *new(BridgeRequestBridge), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeRequestBridge)).(*BridgeRequestBridge)

	return out0, err

}

// CheckForValidRequest is a free data retrieval call binding the contract method 0x0c1fe6b2.
//
// Solidity: function checkForValidRequest(uint256 num) view returns((uint256,uint64,uint64,address,bytes,uint256,address))
func (_Bridge *BridgeSession) CheckForValidRequest(num *big.Int) (BridgeRequestBridge, error) {
	return _Bridge.Contract.CheckForValidRequest(&_Bridge.CallOpts, num)
}

// CheckForValidRequest is a free data retrieval call binding the contract method 0x0c1fe6b2.
//
// Solidity: function checkForValidRequest(uint256 num) view returns((uint256,uint64,uint64,address,bytes,uint256,address))
func (_Bridge *BridgeCallerSession) CheckForValidRequest(num *big.Int) (BridgeRequestBridge, error) {
	return _Bridge.Contract.CheckForValidRequest(&_Bridge.CallOpts, num)
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

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 amount, uint64 originChainID, uint64 destinationChainID, address recipient, bytes extraData, uint256 blockNumber, address token)
func (_Bridge *BridgeCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount             *big.Int
	OriginChainID      uint64
	DestinationChainID uint64
	Recipient          common.Address
	ExtraData          []byte
	BlockNumber        *big.Int
	Token              common.Address
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Amount             *big.Int
		OriginChainID      uint64
		DestinationChainID uint64
		Recipient          common.Address
		ExtraData          []byte
		BlockNumber        *big.Int
		Token              common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OriginChainID = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.DestinationChainID = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Recipient = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ExtraData = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.BlockNumber = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 amount, uint64 originChainID, uint64 destinationChainID, address recipient, bytes extraData, uint256 blockNumber, address token)
func (_Bridge *BridgeSession) Requests(arg0 *big.Int) (struct {
	Amount             *big.Int
	OriginChainID      uint64
	DestinationChainID uint64
	Recipient          common.Address
	ExtraData          []byte
	BlockNumber        *big.Int
	Token              common.Address
}, error) {
	return _Bridge.Contract.Requests(&_Bridge.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 amount, uint64 originChainID, uint64 destinationChainID, address recipient, bytes extraData, uint256 blockNumber, address token)
func (_Bridge *BridgeCallerSession) Requests(arg0 *big.Int) (struct {
	Amount             *big.Int
	OriginChainID      uint64
	DestinationChainID uint64
	Recipient          common.Address
	ExtraData          []byte
	BlockNumber        *big.Int
	Token              common.Address
}, error) {
	return _Bridge.Contract.Requests(&_Bridge.CallOpts, arg0)
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

// RequestBridge is a paid mutator transaction binding the contract method 0x24ca571f.
//
// Solidity: function requestBridge(uint256 amount, uint64 destinationChainID, address destinationAddress, bytes extraData, address token) payable returns(bool)
func (_Bridge *BridgeTransactor) RequestBridge(opts *bind.TransactOpts, amount *big.Int, destinationChainID uint64, destinationAddress common.Address, extraData []byte, token common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "requestBridge", amount, destinationChainID, destinationAddress, extraData, token)
}

// RequestBridge is a paid mutator transaction binding the contract method 0x24ca571f.
//
// Solidity: function requestBridge(uint256 amount, uint64 destinationChainID, address destinationAddress, bytes extraData, address token) payable returns(bool)
func (_Bridge *BridgeSession) RequestBridge(amount *big.Int, destinationChainID uint64, destinationAddress common.Address, extraData []byte, token common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RequestBridge(&_Bridge.TransactOpts, amount, destinationChainID, destinationAddress, extraData, token)
}

// RequestBridge is a paid mutator transaction binding the contract method 0x24ca571f.
//
// Solidity: function requestBridge(uint256 amount, uint64 destinationChainID, address destinationAddress, bytes extraData, address token) payable returns(bool)
func (_Bridge *BridgeTransactorSession) RequestBridge(amount *big.Int, destinationChainID uint64, destinationAddress common.Address, extraData []byte, token common.Address) (*types.Transaction, error) {
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
