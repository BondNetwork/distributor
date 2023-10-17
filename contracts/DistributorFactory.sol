// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "./MerkleDistributor.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract DistributorFactory is Initializable{
    event DistributorCreated(
        address deployer,
        string  pId,
        string  tId,
        address distributorAddress,
        uint256 timestamp
    );

    event DistributorUpdate(
        address deployer,
        string  pId,
        string  tId,
        address oldDistributorAddress,
        address newDistributorAddress,
        uint256 timestamp
    );

    struct TaskItem
    {
        bool isUsed;
        address distributorAddress;
    }

    mapping(string => TaskItem) private taskIds;

    function initialize (
    ) external initializer{
    }

    function createDistributor(
        address aggregatorAddress,
        address token,
        string memory pId,
        string memory tId,
        uint256 amount,
        uint256 startTimestamp,
        uint256 endTimestamp
    ) external{
        require(taskIds[tId].isUsed == false, "taskid was created");
        MerkleDistributor distributorObj = new MerkleDistributor(aggregatorAddress, token, pId, tId, amount, startTimestamp, endTimestamp);
        taskIds[tId].isUsed = true;
        taskIds[tId].distributorAddress = address(distributorObj);
        emit DistributorCreated(msg.sender, pId, tId, address(distributorObj), block.timestamp);
    }

    function updateDistributor(address aggregatorAddress,
        address token,
        string memory pId,
        string memory tId,
        uint256 amount,
        uint256 startTimestamp,
        uint256 endTimestamp)
    external{
        require(taskIds[tId].isUsed == true, "taskid is not exist");
        MerkleDistributor distributorObj = new MerkleDistributor(aggregatorAddress, token, pId, tId, amount, startTimestamp, endTimestamp);
        taskIds[tId].distributorAddress = address(distributorObj);
        emit DistributorUpdate(msg.sender, pId, tId, taskIds[tId].distributorAddress, address(distributorObj), block.timestamp);
    }

    function distributorFactoryVersion()
    external
    pure
    returns (string memory){
        return "1.0.0";
    }

    function getDistributorAddress(string memory tId)
    external
    view
    returns (address){
        return taskIds[tId].distributorAddress;
    }
}
