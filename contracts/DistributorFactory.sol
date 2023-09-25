// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
import "./MerkleDistributor.sol";

contract DistributorFactory {
    event DistributorCreated(
        address deployer,
        string  pId,
        string  tId,
        address distributorAddress,
        uint256 timestamp
    );

    struct TaskItem
    {
        bool isUsed;
        address distributorAddress;
    }

    mapping(string => TaskItem) private taskIds;

    function createDistributor(
        address aggregatorProxy,
        address token,
        string memory pId,
        string memory tId,
        uint256 amount,
        uint256 startTimestamp,
        uint256 endTimestamp
    ) public returns (address) {
        require(taskIds[tId].isUsed == false, "taskid was created");
        MerkleDistributor distributorObj = new MerkleDistributor(aggregatorProxy, token, pId, tId, amount, startTimestamp, endTimestamp);
        taskIds[tId].isUsed = true;
        taskIds[tId].distributorAddress = address(distributorObj);
        emit DistributorCreated(msg.sender, pId, tId, address(distributorObj), block.timestamp);
        return address(distributorObj);
    }
}
