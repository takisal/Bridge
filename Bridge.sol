// SPDX-License-Identifier: MIT

pragma solidity ^0.8.12;

import "IERC20.sol";

interface IERC20USDT {
    function approve(address _spender, uint _value) external;

    function transfer(address _to, uint _value) external;
}

contract Bridge {
    address admin;
    bool preCharge;
    uint sequenceNumber;
    struct RequestBridge {
        uint amount;
        uint64 originChainID;
        uint64 destinationChainID;
        address recipient;
        bytes extraData;
        uint blockNumber;
        address token;
    }
    mapping(address => mapping(uint => uint))
        public tokenToChainToEtherTollFlat;
    mapping(address => mapping(uint => uint))
        public tokenToChainToEtherTollRatio;

    mapping(address => mapping(uint => uint))
        public tokenToChainToInKindFeeFlat;
    mapping(address => mapping(uint => uint))
        public tokenToChainToInKindFeeRatio;
    mapping(uint => RequestBridge) public requests;
    uint constant ratioCap = 10 ** 18;

    constructor() {
        admin = msg.sender;
    }

    modifier onlyAdmin() {
        require(msg.sender == admin, "Can only be called by admin");
        _;
    }

    function estimateFee(
        uint amount,
        uint destinationChainID,
        address token
    ) public view returns (uint etherToll, uint inKindFee) {
        etherToll = tokenToChainToEtherTollFlat[token][destinationChainID];
        etherToll =
            etherToll +
            ((amount *
                tokenToChainToEtherTollRatio[token][destinationChainID]) /
                ratioCap);
        inKindFee = tokenToChainToInKindFeeFlat[token][destinationChainID];
        inKindFee =
            inKindFee +
            ((amount *
                tokenToChainToInKindFeeRatio[token][destinationChainID]) /
                ratioCap);
    }

    function requestBridge(
        uint amount,
        uint64 destinationChainID,
        address destinationAddress,
        bytes calldata extraData,
        address token
    ) public payable returns (bool) {
        uint etherDue = tokenToChainToEtherTollFlat[token][destinationChainID];
        etherDue =
            etherDue +
            ((amount *
                tokenToChainToEtherTollRatio[token][destinationChainID]) /
                ratioCap);
        if (token == address(0)) {
            etherDue += amount;
            require(msg.value > etherDue, "Toll paid is insufficient");
        } else {
            require(msg.value > etherDue, "Toll paid is insufficient");
        }
        require(amount > 0);
        if (preCharge) {
            uint inKindFee = tokenToChainToInKindFeeFlat[token][
                destinationChainID
            ];
            inKindFee =
                inKindFee +
                ((amount *
                    tokenToChainToInKindFeeRatio[token][destinationChainID]) /
                    ratioCap);
            IERC20(token).transferFrom(msg.sender, admin, inKindFee + amount);
            requests[sequenceNumber] = RequestBridge(
                amount,
                uint64(block.chainid),
                destinationChainID,
                destinationAddress,
                extraData,
                block.number,
                token
            );
        } else {
            IERC20(token).transferFrom(msg.sender, admin, amount);
            amount =
                amount -
                ((amount *
                    tokenToChainToInKindFeeRatio[token][destinationChainID]) /
                    ratioCap);
            requests[sequenceNumber] = RequestBridge(
                amount - tokenToChainToInKindFeeFlat[token][destinationChainID],
                uint64(block.chainid),
                destinationChainID,
                destinationAddress,
                extraData,
                block.number,
                token
            );
        }
        payable(admin).transfer(msg.value);
        sequenceNumber++;
        return true;
    }

    function changeAdmin(address _admin) public onlyAdmin returns (bool) {
        admin = _admin;
        return true;
    }

    function setChargeTime(bool newChargeTime) public onlyAdmin returns (bool) {
        preCharge = newChargeTime;
        return true;
    }

    function setFees(
        uint destinationChain,
        uint fee,
        address token,
        bool inKind,
        bool flat
    ) public onlyAdmin returns (bool) {
        if (inKind) {
            if (flat) {
                tokenToChainToInKindFeeFlat[token][destinationChain] = fee;
            } else {
                tokenToChainToInKindFeeRatio[token][destinationChain] = fee;
            }
        } else {
            if (flat) {
                tokenToChainToEtherTollFlat[token][destinationChain] = fee;
            } else {
                tokenToChainToEtherTollRatio[token][destinationChain] = fee;
            }
        }
    }

    function checkForValidRequest(
        uint num
    ) public view returns (RequestBridge memory) {
        require(requests[num].amount > 0, "Does not exist");
        require(requests[num].blockNumber < block.number - 5, "Too soon");
        return requests[num];
    }

    function withdrawERC20(
        address token,
        uint256 amountToWithdraw
    ) public onlyAdmin returns (bool) {
        if (token == address(0xdAC17F958D2ee523a2206206994597C13D831ec7)) {
            IERC20USDT(token).transfer(msg.sender, amountToWithdraw);
        } else {
            IERC20(token).transfer(msg.sender, amountToWithdraw);
        }
        return true;
    }

    function withdrawEther(
        uint256 amountToWithdraw
    ) public onlyAdmin returns (bool) {
        (bool sent, bytes memory data) = payable(msg.sender).call{
            value: amountToWithdraw
        }("");
        require(sent, "Failed to send Ether");
        return true;
    }

    receive() external payable {}

    fallback() external payable {}
}
