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

    uint32 internal constant REVISION = 5;
    IWETH internal _IWETH;

    mapping(bytes32 taskId => address distributor) internal _taskItems;

    function initialize(address wethAddress) external initializer {
        __BondUpgradeable_init();
        _IWETH = IWETH(wethAddress);
    }

    function getVersion() external pure returns (uint32) {
        return REVISION;
    }

    function createDistributorByToken(
        Types.CreateDistributorParams calldata params
    ) external returns (address) {
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
        _taskItems[taskId] = distributorAddress;
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
        console.log("msg.value %s,%s ", msg.value,params.amount);

        require(msg.value >= params.amount, "Not enough ETH sent");
        //
        (bytes32 taskId, address distributorAddress) = _createDistributor(
            params
        );
       
        _IWETH.deposit{value: msg.value}();

        console.log("IWETH deposit %s ", msg.value);

         _IWETH.transfer(distributorAddress, params.amount);
       //_IWETH.transferFrom(msg.sender, distributorAddress, params.amount);
        
        console.log("IWETH transfer %s ", params.amount);

        _taskItems[taskId] = distributorAddress;

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

    function updateDistributor(
        Types.CreateDistributorParams calldata params
    ) external {
        (bytes32 taskId, address distributorAddress) = _createDistributor(
            params
        );
        address oldAddress = _taskItems[taskId];
        _taskItems[taskId] = distributorAddress;

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
        return _taskItems[id];
    }

    function withdrawETH(
        address distributorAddress,
        uint256 amount,
        address to
    ) external {
        console.log("distributorAddress %s ", distributorAddress);

        uint256 userBalance = IMerkleDistributor(distributorAddress).getTokenBalance();

        console.log("userBalance %s ", userBalance);

        uint256 amountToWithdraw = amount;
        if (amount > userBalance ) {
            amountToWithdraw = userBalance;
        }

        console.log("amountToWithdraw %s ", amountToWithdraw);

        IMerkleDistributor(distributorAddress).withdraw(amountToWithdraw);
 
        console.log("Balance %s %s ",getTokenBalance(),amountToWithdraw);
        
        _IWETH.withdraw(amountToWithdraw);

        console.log("safeTransferETH %s %s ",to,amountToWithdraw); 
        _safeTransferETH(to, amountToWithdraw);
    }

    function _createDistributor(
        Types.CreateDistributorParams calldata params
    ) private returns (bytes32, address) {
        bytes32 taskId = keccak256(bytes(params.taskId));
        if (_taskItems[taskId] != address(0)) {
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
}
