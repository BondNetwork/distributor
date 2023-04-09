// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity >=0.8.1 <0.9.0;

import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import {IMerkleDistributor} from "./IMerkleDistributor.sol";

import "./interfaces/AggregatorMerkleInterface.sol";

contract MerkleDistributor is IMerkleDistributor {
    using SafeERC20 for IERC20;

    struct ClaimedItem{
        uint256 _index;
        address _account;
        bool _isClaimed;
    }

    AggregatorMerkleInterface internal merkleAggregator;
    address public immutable override token;
    uint256 public immutable override projectId;
    bytes32 public immutable override merkleRoot;
    uint256 public immutable override curBatch;
    

    // This is a packed array of booleans.
    mapping(uint256 => ClaimItem) private claimedBitMap;

    constructor(address aggregatorProxy_, address token_, uint256 projectId_) {
        merkleAggregator = AggregatorMerkleInterface(aggregatorProxy_);
        token = token_;
        projectId_ = projectId;
    }

    function isLocked() public returns (bool){
        return merkleAggregator.isLocked();
    }

    function getRoot() public returns (bool){
        if(isLocked())
            return false;
        (
            /*uint80 roundID*/,
            batch,
            merkleRoot,
            /*uint startedAt*/,
            /*uint timeStamp*/,
        ) = merkleAggregator.lastMerkleRoot();
        return true;
    }

    function isClaimed(uint256 batch, uint256 index)
    public
    virtual
    returns (bool) {
        if(curBatch != batch)
            return true;
        ClaimItem memory claim  =  claimedBitMap[batch];
        if(claim._index == index && claim._isClaimed == true)
            return true;
        return false;
    }

    function _setClaimed(uint256 index) private {
        claimedBitMap[index] = ClaimItem({
        _batch:curBatch,
        _index:index,
        _isClaimed:true
        });
    }

    function claim(uint256 batch, uint256 index, address account, uint256 amount, bytes32[] calldata merkleProof)
    public
    virtual
    returns (bool)
    {
        if(isLocked()) return false;
        if(batch != curBatch) return false;
        if (isClaimed(curBatch,index)) return false;

        // Verify the merkle proof.
        bytes32 node = keccak256(abi.encodePacked(index, account, amount));
        if (!MerkleProof.verify(merkleProof, merkleRoot, node)) return false;

        // Mark it claimed and send the token.
        _setClaimed(index, account);


        IERC20(token).safeTransfer(account, amount);

        emit Claimed(batch, index, account, amount);
        return true;
    }
}
