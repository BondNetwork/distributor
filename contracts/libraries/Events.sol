// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

library Events {
    event DistributorCreated(
        string indexed taskId,
        string projectId,
        address deployer,
        address distributorAddress,
        string  tokenType,
        uint256 amount,
        uint256 timestamp
    );

    event DistributorUpdate(
        string indexed taskId,
        string projectId,
        address deployer,
        address oldDistributorAddress,
        address newDistributorAddress,
        uint256 timestamp
    );

    event DistributorDeposit(
        string indexed taskId,
        address indexed user,
        uint256 amount
    );

    event DistributorWithdraw(
        string indexed taskId,
        address indexed user,
        uint256 amount
    );
}
