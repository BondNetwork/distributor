// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

library Events {
    event DistributorCreated(
        string  projectId,
        string  taskId,
        address deployer,
        address distributorAddress,
        address tokenAddress,
        uint256 amount,
        uint256 timestamp
    );

    event DistributorUpdate(
        string  projectId,
        string  taskId,
        address deployer,
        address oldDistributorAddress,
        address newDistributorAddress,
        uint256 timestamp
    );

    event DistributorDeposit(
        string  projectId,
        string  taskId,
        address user,
        uint256 amount,
        uint256 timestamp
    );

    event DistributorWithdraw(
        string  projectId,
        string  taskId,
        address user,
        uint256 amount,
        uint256 timestamp
    );
}
