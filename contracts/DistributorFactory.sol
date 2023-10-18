// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Types} from "contracts/libraries/Types.sol";
import {Events} from "contracts/libraries/Events.sol";
import {Errors} from "contracts/libraries/Errors.sol";
import {BondUpgradeable} from "contracts/upgradeability/BondUpgradeable.sol";
import {MerkleDistributor} from "contracts/MerkleDistributor.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "hardhat/console.sol";

contract DistributorFactory is BondUpgradeable {
    using SafeERC20 for IERC20;

    uint32 internal constant REVISION = 5;

    mapping(bytes32 taskId => address distributor) internal _taskItems;

    function initialize() external initializer {
        __BondUpgradeable_init();
    }

    function getVersion() public pure returns (uint32) {
        return REVISION;
    }

    function createDistributor(
        Types.CreateDistributorParams calldata params
    ) external returns (address) {
        bytes32 taskId = keccak256(bytes(params.taskId));
        if (_taskItems[taskId] != address(0)) {
            revert Errors.TaskItemAlreadyExist();
        }

        
        address distributorAddress = address(new MerkleDistributor(params));
        _taskItems[taskId] = distributorAddress;
        console.log("new distributor: %s ",distributorAddress);
        if (params.token != address(0)) {
           
            console.log("account: %s %s", params.account, IERC20(params.token).balanceOf(params.account));
            console.log("allowance: %s %s ", IERC20(params.token).allowance(params.account,address(this)),params.amount);
            IERC20(params.token).safeTransferFrom(
                params.account,
                distributorAddress,
                params.amount
            );
        }
        emit Events.DistributorCreated(
            params.taskId,
            params.projectId,
            msg.sender,
            distributorAddress,
            block.timestamp
        );
        return distributorAddress;
    }

    function updateDistributor(
        Types.CreateDistributorParams calldata params
    ) external {
        bytes32 taskId = keccak256(bytes(params.taskId));
        if (_taskItems[taskId] == address(0)) {
            revert Errors.TaskItemNoExist();
        }
        address distributorAddress = address(new MerkleDistributor(params));
        address oldAddress = _taskItems[taskId];
        _taskItems[taskId] = distributorAddress;
        emit Events.DistributorUpdate(
            params.taskId,
            params.projectId,
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
        return _taskItems[id];
    }
}
