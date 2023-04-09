pragma solidity ^0.8.0;

interface AggregatorMerkleInterface {
  function version() external view returns (uint256);

  function isLocked()
    external
    view
    returns (
      bool isLocked
    );

  function latestMerkleRoundData()
    external
    view
    returns (
      uint80 roundId,
      uint256 batchId,
      bytes32 merkelAnswer,
      uint256 startedAt,
      uint256 updatedAt
    );
}