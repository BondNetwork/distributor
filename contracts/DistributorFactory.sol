// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Types} from "contracts/libraries/Types.sol";
import {Events} from "contracts/libraries/Events.sol";
import {Errors} from "contracts/libraries/Errors.sol";
import {BondUpgradeable} from "contracts/upgradeability/BondUpgradeable.sol";
import {MerkleDistributor} from "contracts/MerkleDistributor.sol";
import {IMerkleDistributor} from "contracts/interfaces/IMerkleDistributor.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IWETH} from "contracts/interfaces/IWETH.sol";
import "hardhat/console.sol";

contract DistributorFactory is BondUpgradeable {
    using SafeERC20 for IERC20;

    uint32 internal constant REVISION = 1;
    IWETH internal _IWETH;

    mapping(bytes32 taskId => Types.DistributorData data) internal _taskItems;

    function initialize(address wethAddress) external initializer {
        __BondUpgradeable_init();
        _IWETH = IWETH(wethAddress);
    }

    function getVersion() external pure returns (uint32) {
        return REVISION;
    }

    /**
     * @dev Get WETH address
     */
    function getWETHAddress() external view returns (address) {
        return address(_IWETH);
    }

    function createDistributorByToken(
        Types.CreateDistributorParams calldata params
    ) external returns (address) {
        console.log("token ", params.token);

        (bytes32 taskId, address distributorAddress) = _createDistributor(
            params
        );
        if (params.token != address(0)) {
            console.log(
                "account: %s %s",
                msg.sender,
                IERC20(params.token).balanceOf(msg.sender)
            );
            console.log(
                "allowance: %s %s ",
                IERC20(params.token).allowance(msg.sender, address(this)),
                params.amount
            );
            IERC20(params.token).safeTransferFrom(
                msg.sender,
                distributorAddress,
                params.amount
            );
        }
        _taskItems[taskId].distributor = distributorAddress;
        _taskItems[taskId].owner = msg.sender;
        emit Events.DistributorCreated(
            params.projectId,
            params.taskId,
            msg.sender,
            distributorAddress,
            params.token,
            params.amount,
            block.timestamp
        );
        return distributorAddress;
    }

    function createDistributorByEth(
        Types.CreateDistributorParams calldata params
    ) external payable returns (address) {
        console.log("msg.value %s,%s ", msg.value, params.amount);

        require(msg.value >= 0, "Not enough ETH sent");
        Types.CreateDistributorParams memory tmpParams;
        tmpParams.aggregatorAddress = params.aggregatorAddress;
        tmpParams.token = address(_IWETH);
        tmpParams.projectId = params.projectId;
        tmpParams.taskId = params.taskId;
        tmpParams.amount = msg.value;
        tmpParams.startTimestamp = params.startTimestamp;
        tmpParams.endTimestamp = params.endTimestamp;
        tmpParams.rewardPerBatch = params.rewardPerBatch;
        //
        (bytes32 taskId, address distributorAddress) = _createDistributor(
            tmpParams
        );

        _IWETH.deposit{value: tmpParams.amount}();

        _IWETH.transfer(distributorAddress, tmpParams.amount);

        console.log("IWETH transfer %s ", tmpParams.amount);

        _taskItems[taskId].distributor = distributorAddress;
        _taskItems[taskId].owner = msg.sender;

        emit Events.DistributorCreated(
            tmpParams.projectId,
            tmpParams.taskId,
            msg.sender,
            distributorAddress,
            tmpParams.token,
            tmpParams.amount,
            block.timestamp
        );
        return distributorAddress;
    }

    function updateDistributor(
        Types.CreateDistributorParams calldata params
    ) external {
        bytes32 taskId = keccak256(bytes(params.taskId));
        address oldAddress = _taskItems[taskId].distributor;
        if (oldAddress == address(0)) {
            revert Errors.TaskItemNoExist();
        }
        if (msg.sender != _taskItems[taskId].owner) {
            revert Errors.NotOwner();
        }
        address distributorAddress = address(
            new MerkleDistributor(address(this), params)
        );
        console.log("update distributor: %s ", distributorAddress);
        _taskItems[taskId].distributor = distributorAddress;
        emit Events.DistributorUpdate(
            params.projectId,
            params.taskId,
            msg.sender,
            oldAddress,
            distributorAddress,
            block.timestamp
        );
    }

    function getDistributorAddress(
        string memory taskId
    ) external view returns (address) {
        bytes32 id = keccak256(bytes(taskId));
        return _taskItems[id].distributor;
    }

    function depositETH(
        address distributorAddress
    ) external payable{
        _IWETH.deposit{value:  msg.value}();
        _IWETH.transfer(distributorAddress,  msg.value);
        console.log("depositETH %s ",  msg.value);
    }

    function withdrawETH(
        address distributorAddress,
        uint256 amount,
        address to
    ) external {
        uint256 userBalance = IMerkleDistributor(distributorAddress)
            .getTokenBalance();

        console.log(
            "distributorAddress %s userBalance %s ",
            distributorAddress,
            userBalance
        );

        uint256 amountToWithdraw = amount;
        // if amount is equal to uint(-1), the user wants to redeem everything
        if (amount == type(uint256).max) {
            amountToWithdraw = userBalance;
        }

        console.log("amountToWithdraw %s ", amountToWithdraw);

        IMerkleDistributor(distributorAddress).withdraw(amountToWithdraw);

        _IWETH.withdraw(amountToWithdraw);

        _safeTransferETH(to, amountToWithdraw);
    }

    function _createDistributor(
        Types.CreateDistributorParams memory params
    ) private returns (bytes32, address) {
        bytes32 taskId = keccak256(bytes(params.taskId));
        if (_taskItems[taskId].distributor != address(0)) {
            revert Errors.TaskItemAlreadyExist();
        }
        address distributorAddress = address(
            new MerkleDistributor(address(this), params)
        );
        console.log("new distributor: %s ", distributorAddress);
        return (taskId, distributorAddress);
    }

    function _safeTransferETH(address to, uint256 amount) internal {
        if (amount > 0) {
            bool success;
            assembly {
                success := call(gas(), to, amount, 0, 0, 0, 0)
            }
            if (!success) {
                revert Errors.ETHTransferFailed();
            }
        }
    }

    function getTokenBalance() internal view returns (uint256) {
        return _IWETH.balanceOf(address(this));
    }

    /**
     * @dev Only WETH contract is allowed to transfer ETH here. Prevent other addresses to send Ether to this contract.
     */
    receive() external payable {
        require(msg.sender == address(_IWETH), "Receive not allowed");
    }

    /**
     * @dev Revert fallback calls
     */
    fallback() external payable {
        revert("Fallback not allowed");
    }
}
