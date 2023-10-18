// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

library Types {

    struct CreateDistributorParams {
        address aggregatorAddress;
        address token;
        address account;
        uint256 ethValue;
        string  projectId;
        string  taskId;
        uint256 amount;
        uint256 startTimestamp;
        uint256 endTimestamp;
        uint256 rewardPerBatch;
    }

    struct IndexValue {
        uint256 keyIndex;
        bool value;
    }

    struct KeyFlag {
        uint256 key;
        bool deleted;
    }
}