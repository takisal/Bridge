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
}

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"checkForValidRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"originChainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destinationChainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structBridge.RequestBridge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"etherToll\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inKindFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destinationChainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"requestBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"originChainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destinationChainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newChargeTime\",\"type\":\"bool\"}],\"name\":\"setChargeTime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"inKind\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"flat\",\"type\":\"bool\"}],\"name\":\"setFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToChainToEtherTollFlat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToChainToEtherTollRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToChainToInKindFeeFlat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToChainToInKindFeeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506125708061005b5f395ff3fe6080604052600436106100c5575f3560e01c806372f3af061161007e5780638f283970116100585780638f283970146102e4578063a1db978214610320578063c7940bf31461035c578063e031c55214610398576100cc565b806372f3af061461022b5780637c7f3c891461026757806381d12c58146102a3576100cc565b80630c1fe6b2146100ce5780631a2873d61461010a57806324ca571f146101465780633bed33ce146101765780633ee6024b146101b25780635546299e146101ee576100cc565b366100cc57005b005b3480156100d9575f80fd5b506100f460048036038101906100ef9190611809565b6103d4565b604051610101919061199a565b60405180910390f35b348015610115575f80fd5b50610130600480360381019061012b91906119e4565b610611565b60405161013d9190611a31565b60405180910390f35b610160600480360381019061015b9190611ad5565b610631565b60405161016d9190611b85565b60405180910390f35b348015610181575f80fd5b5061019c60048036038101906101979190611809565b610de1565b6040516101a99190611b85565b60405180910390f35b3480156101bd575f80fd5b506101d860048036038101906101d391906119e4565b610f25565b6040516101e59190611a31565b60405180910390f35b3480156101f9575f80fd5b50610214600480360381019061020f9190611b9e565b610f45565b604051610222929190611bee565b60405180910390f35b348015610236575f80fd5b50610251600480360381019061024c91906119e4565b6110dd565b60405161025e9190611a31565b60405180910390f35b348015610272575f80fd5b5061028d600480360381019061028891906119e4565b6110fd565b60405161029a9190611a31565b60405180910390f35b3480156102ae575f80fd5b506102c960048036038101906102c49190611809565b61111d565b6040516102db96959493929190611c7b565b60405180910390f35b3480156102ef575f80fd5b5061030a60048036038101906103059190611ce1565b611221565b6040516103179190611b85565b60405180910390f35b34801561032b575f80fd5b50610346600480360381019061034191906119e4565b6112f8565b6040516103539190611b85565b60405180910390f35b348015610367575f80fd5b50610382600480360381019061037d9190611d36565b6114c2565b60405161038f9190611b85565b60405180910390f35b3480156103a3575f80fd5b506103be60048036038101906103b99190611d61565b611573565b6040516103cb9190611b85565b60405180910390f35b6103dc611773565b5f60065f8481526020019081526020015f205f015411610431576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042890611e32565b60405180910390fd5b60054361043e9190611e7d565b60065f8481526020019081526020015f206004015410610493576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048a90611efa565b60405180910390fd5b60065f8381526020019081526020015f206040518060c00160405290815f8201548152602001600182015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805461058090611f45565b80601f01602080910402602001604051908101604052809291908181526020018280546105ac90611f45565b80156105f75780601f106105ce576101008083540402835291602001916105f7565b820191905f5260205f20905b8154815290600101906020018083116105da57829003601f168201915b505050505081526020016004820154815250509050919050565b6003602052815f5260405f20602052805f5260405f205f91509150505481565b5f8060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1681526020019081526020015f20549050670de0b6b3a764000060035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8967ffffffffffffffff1681526020019081526020015f2054896106f79190611f75565b6107019190611fe3565b8161070c9190612013565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361079657878161074d9190612013565b9050803411610791576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078890612090565b60405180910390fd5b6107d9565b8034116107d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107cf90612090565b60405180910390fd5b5b5f88116107e4575f80fd5b5f60149054906101000a900460ff1615610ade575f60045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8967ffffffffffffffff1681526020019081526020015f20549050670de0b6b3a764000060055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8a67ffffffffffffffff1681526020019081526020015f20548a6108bd9190611f75565b6108c79190611fe3565b816108d29190612013565b90508373ffffffffffffffffffffffffffffffffffffffff166323b872dd33308c856108fe9190612013565b6040518463ffffffff1660e01b815260040161091c939291906120ae565b6020604051808303815f875af1158015610938573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061095c91906120f7565b506040518060c001604052808a81526020014667ffffffffffffffff1681526020018967ffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014381525060065f60015481526020019081526020015f205f820151815f01556020820151816001015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816003019081610aca91906122ec565b5060a0820151816004015590505050610dbb565b8273ffffffffffffffffffffffffffffffffffffffff166323b872dd33308b6040518463ffffffff1660e01b8152600401610b1b939291906120ae565b6020604051808303815f875af1158015610b37573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b5b91906120f7565b50670de0b6b3a764000060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8967ffffffffffffffff1681526020019081526020015f205489610bc79190611f75565b610bd19190611fe3565b88610bdc9190611e7d565b97506040518060c0016040528060045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8a67ffffffffffffffff1681526020019081526020015f20548a610c4b9190611e7d565b81526020014667ffffffffffffffff1681526020018867ffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff16815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014381525060065f60015481526020019081526020015f205f820151815f01556020820151816001015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816003019081610dac91906122ec565b5060a082015181600401559050505b60015f815480929190610dcd906123bb565b919050555060019150509695505050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e679061244c565b60405180910390fd5b5f803373ffffffffffffffffffffffffffffffffffffffff1684604051610e9690612497565b5f6040518083038185875af1925050503d805f8114610ed0576040519150601f19603f3d011682016040523d82523d5f602084013e610ed5565b606091505b509150915081610f1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f11906124f5565b60405180910390fd5b600192505050919050565b6005602052815f5260405f20602052805f5260405f205f91509150505481565b5f8060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8581526020019081526020015f20549150670de0b6b3a764000060035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8681526020019081526020015f205486610ff79190611f75565b6110019190611fe3565b8261100c9190612013565b915060045f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8581526020019081526020015f20549050670de0b6b3a764000060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8681526020019081526020015f2054866110be9190611f75565b6110c89190611fe3565b816110d39190612013565b9050935093915050565b6004602052815f5260405f20602052805f5260405f205f91509150505481565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b6006602052805f5260405f205f91509050805f015490806001015f9054906101000a900467ffffffffffffffff16908060010160089054906101000a900467ffffffffffffffff1690806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600301805461119a90611f45565b80601f01602080910402602001604051908101604052809291908181526020018280546111c690611f45565b80156112115780601f106111e857610100808354040283529160200191611211565b820191905f5260205f20905b8154815290600101906020018083116111f457829003601f168201915b5050505050908060040154905086565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146112b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112a79061244c565b60405180910390fd5b815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019050919050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611387576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161137e9061244c565b60405180910390fd5b73dac17f958d2ee523a2206206994597c13d831ec773ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361143b578273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b8152600401611409929190612513565b5f604051808303815f87803b158015611420575f80fd5b505af1158015611432573d5f803e3d5ffd5b505050506114b8565b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b8152600401611476929190612513565b6020604051808303815f875af1158015611492573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114b691906120f7565b505b6001905092915050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611551576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115489061244c565b60405180910390fd5b815f60146101000a81548160ff02191690831515021790555060019050919050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611602576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115f99061244c565b60405180910390fd5b82156116bb578115611664578460045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f20819055506116b6565b8460055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f20819055505b61176a565b8115611717578460025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f2081905550611769565b8460035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8881526020019081526020015f20819055505b5b95945050505050565b6040518060c001604052805f81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff168152602001606081526020015f81525090565b5f80fd5b5f80fd5b5f819050919050565b6117e8816117d6565b81146117f2575f80fd5b50565b5f81359050611803816117df565b92915050565b5f6020828403121561181e5761181d6117ce565b5b5f61182b848285016117f5565b91505092915050565b61183d816117d6565b82525050565b5f67ffffffffffffffff82169050919050565b61185f81611843565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61188e82611865565b9050919050565b61189e81611884565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6118e6826118a4565b6118f081856118ae565b93506119008185602086016118be565b611909816118cc565b840191505092915050565b5f60c083015f8301516119295f860182611834565b50602083015161193c6020860182611856565b50604083015161194f6040860182611856565b5060608301516119626060860182611895565b506080830151848203608086015261197a82826118dc565b91505060a083015161198f60a0860182611834565b508091505092915050565b5f6020820190508181035f8301526119b28184611914565b905092915050565b6119c381611884565b81146119cd575f80fd5b50565b5f813590506119de816119ba565b92915050565b5f80604083850312156119fa576119f96117ce565b5b5f611a07858286016119d0565b9250506020611a18858286016117f5565b9150509250929050565b611a2b816117d6565b82525050565b5f602082019050611a445f830184611a22565b92915050565b611a5381611843565b8114611a5d575f80fd5b50565b5f81359050611a6e81611a4a565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112611a9557611a94611a74565b5b8235905067ffffffffffffffff811115611ab257611ab1611a78565b5b602083019150836001820283011115611ace57611acd611a7c565b5b9250929050565b5f805f805f8060a08789031215611aef57611aee6117ce565b5b5f611afc89828a016117f5565b9650506020611b0d89828a01611a60565b9550506040611b1e89828a016119d0565b945050606087013567ffffffffffffffff811115611b3f57611b3e6117d2565b5b611b4b89828a01611a80565b93509350506080611b5e89828a016119d0565b9150509295509295509295565b5f8115159050919050565b611b7f81611b6b565b82525050565b5f602082019050611b985f830184611b76565b92915050565b5f805f60608486031215611bb557611bb46117ce565b5b5f611bc2868287016117f5565b9350506020611bd3868287016117f5565b9250506040611be4868287016119d0565b9150509250925092565b5f604082019050611c015f830185611a22565b611c0e6020830184611a22565b9392505050565b611c1e81611843565b82525050565b611c2d81611884565b82525050565b5f82825260208201905092915050565b5f611c4d826118a4565b611c578185611c33565b9350611c678185602086016118be565b611c70816118cc565b840191505092915050565b5f60c082019050611c8e5f830189611a22565b611c9b6020830188611c15565b611ca86040830187611c15565b611cb56060830186611c24565b8181036080830152611cc78185611c43565b9050611cd660a0830184611a22565b979650505050505050565b5f60208284031215611cf657611cf56117ce565b5b5f611d03848285016119d0565b91505092915050565b611d1581611b6b565b8114611d1f575f80fd5b50565b5f81359050611d3081611d0c565b92915050565b5f60208284031215611d4b57611d4a6117ce565b5b5f611d5884828501611d22565b91505092915050565b5f805f805f60a08688031215611d7a57611d796117ce565b5b5f611d87888289016117f5565b9550506020611d98888289016117f5565b9450506040611da9888289016119d0565b9350506060611dba88828901611d22565b9250506080611dcb88828901611d22565b9150509295509295909350565b5f82825260208201905092915050565b7f446f6573206e6f742065786973740000000000000000000000000000000000005f82015250565b5f611e1c600e83611dd8565b9150611e2782611de8565b602082019050919050565b5f6020820190508181035f830152611e4981611e10565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611e87826117d6565b9150611e92836117d6565b9250828203905081811115611eaa57611ea9611e50565b5b92915050565b7f546f6f20736f6f6e0000000000000000000000000000000000000000000000005f82015250565b5f611ee4600883611dd8565b9150611eef82611eb0565b602082019050919050565b5f6020820190508181035f830152611f1181611ed8565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611f5c57607f821691505b602082108103611f6f57611f6e611f18565b5b50919050565b5f611f7f826117d6565b9150611f8a836117d6565b9250828202611f98816117d6565b91508282048414831517611faf57611fae611e50565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f611fed826117d6565b9150611ff8836117d6565b92508261200857612007611fb6565b5b828204905092915050565b5f61201d826117d6565b9150612028836117d6565b92508282019050808211156120405761203f611e50565b5b92915050565b7f546f6c6c207061696420697320696e73756666696369656e74000000000000005f82015250565b5f61207a601983611dd8565b915061208582612046565b602082019050919050565b5f6020820190508181035f8301526120a78161206e565b9050919050565b5f6060820190506120c15f830186611c24565b6120ce6020830185611c24565b6120db6040830184611a22565b949350505050565b5f815190506120f181611d0c565b92915050565b5f6020828403121561210c5761210b6117ce565b5b5f612119848285016120e3565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026121ab7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612170565b6121b58683612170565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6121f06121eb6121e6846117d6565b6121cd565b6117d6565b9050919050565b5f819050919050565b612209836121d6565b61221d612215826121f7565b84845461217c565b825550505050565b5f90565b612231612225565b61223c818484612200565b505050565b5b8181101561225f576122545f82612229565b600181019050612242565b5050565b601f8211156122a4576122758161214f565b61227e84612161565b8101602085101561228d578190505b6122a161229985612161565b830182612241565b50505b505050565b5f82821c905092915050565b5f6122c45f19846008026122a9565b1980831691505092915050565b5f6122dc83836122b5565b9150826002028217905092915050565b6122f5826118a4565b67ffffffffffffffff81111561230e5761230d612122565b5b6123188254611f45565b612323828285612263565b5f60209050601f831160018114612354575f8415612342578287015190505b61234c85826122d1565b8655506123b3565b601f1984166123628661214f565b5f5b8281101561238957848901518255600182019150602085019450602081019050612364565b868310156123a657848901516123a2601f8916826122b5565b8355505b6001600288020188555050505b505050505050565b5f6123c5826117d6565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036123f7576123f6611e50565b5b600182019050919050565b7f43616e206f6e6c792062652063616c6c65642062792061646d696e00000000005f82015250565b5f612436601b83611dd8565b915061244182612402565b602082019050919050565b5f6020820190508181035f8301526124638161242a565b9050919050565b5f81905092915050565b50565b5f6124825f8361246a565b915061248d82612474565b5f82019050919050565b5f6124a182612477565b9150819050919050565b7f4661696c656420746f2073656e642045746865720000000000000000000000005f82015250565b5f6124df601483611dd8565b91506124ea826124ab565b602082019050919050565b5f6020820190508181035f83015261250c816124d3565b9050919050565b5f6040820190506125265f830185611c24565b6125336020830184611a22565b939250505056fea264697066735822122019305a0afe196ae4f9a6532248fbd90d373aa04b9ba84d841da1073ed518064864736f6c63430008190033",
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
// Solidity: function checkForValidRequest(uint256 num) view returns((uint256,uint64,uint64,address,bytes,uint256))
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
// Solidity: function checkForValidRequest(uint256 num) view returns((uint256,uint64,uint64,address,bytes,uint256))
func (_Bridge *BridgeSession) CheckForValidRequest(num *big.Int) (BridgeRequestBridge, error) {
	return _Bridge.Contract.CheckForValidRequest(&_Bridge.CallOpts, num)
}

// CheckForValidRequest is a free data retrieval call binding the contract method 0x0c1fe6b2.
//
// Solidity: function checkForValidRequest(uint256 num) view returns((uint256,uint64,uint64,address,bytes,uint256))
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
// Solidity: function requests(uint256 ) view returns(uint256 amount, uint64 originChainID, uint64 destinationChainID, address recipient, bytes extraData, uint256 blockNumber)
func (_Bridge *BridgeCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount             *big.Int
	OriginChainID      uint64
	DestinationChainID uint64
	Recipient          common.Address
	ExtraData          []byte
	BlockNumber        *big.Int
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

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 amount, uint64 originChainID, uint64 destinationChainID, address recipient, bytes extraData, uint256 blockNumber)
func (_Bridge *BridgeSession) Requests(arg0 *big.Int) (struct {
	Amount             *big.Int
	OriginChainID      uint64
	DestinationChainID uint64
	Recipient          common.Address
	ExtraData          []byte
	BlockNumber        *big.Int
}, error) {
	return _Bridge.Contract.Requests(&_Bridge.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 amount, uint64 originChainID, uint64 destinationChainID, address recipient, bytes extraData, uint256 blockNumber)
func (_Bridge *BridgeCallerSession) Requests(arg0 *big.Int) (struct {
	Amount             *big.Int
	OriginChainID      uint64
	DestinationChainID uint64
	Recipient          common.Address
	ExtraData          []byte
	BlockNumber        *big.Int
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
