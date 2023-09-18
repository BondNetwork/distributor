// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
import "./MerkleDistributor.sol";

contract DistributorFactory {
    event DistributorCreated(
        address deployer,
        string  pId,
        uint256 tId,
        address token,
        uint256 timestamp
    );

    function createDistributor(
        address aggregatorProxy,
        address token,
        string memory pId,
        uint256 tId,
        address owner
    ) public returns (address) {
        MerkleDistributor distributorObj = new MerkleDistributor(aggregatorProxy, token, pId, tId);
        emit DistributorCreated(owner, pId, tId, address(distributorObj), block.timestamp);
        return address(distributorObj);
    }
}
