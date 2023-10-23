// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

// Allows anyone to claim a token if they exist in a merkle root.
interface IMerkleDistributor {
    // Returns the address of the token distributed by this contract.
    function token() external view returns (address);

    // Returns the merkle root of the merkle tree containing account balances available to claim.
    function merkleRoot() external view returns (bytes32);

    // Returns true if the index has been marked claimed.
    function isClaimed(
        uint256 batch,
        uint256 index
    ) external view returns (bool);

    // Claim the given amount of the token to the given address. Reverts if the inputs are invalid.
    function claim(
        uint256 batch,
        uint256 index,
        address account,
        uint256 amount,
        bytes32[] calldata merkleProof
    ) external returns (bool);

    // This event is triggered whenever a call to #claim succeeds.
    event Claimed(
        string  indexed projectId,
        string  taskId,
        uint256 indexed batch,
        uint256 indexed index,
        address account,
        uint256 amount
    );

    function withdraw(uint256 amount) external;

    function getTokenBalance() external view returns (uint256);
}
