// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity >=0.8.0 <0.9.0;
import {IMerkleDistributor} from "./interfaces/IMerkleDistributor.sol";
import "./interfaces/AggregatorMerkleInterface.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";


contract MerkleDistributor is IMerkleDistributor {
    using SafeERC20 for IERC20;

    struct ClaimItem{
        uint256 _batchId;
        uint256 _index;
        address _account;
        bool _isClaimed;
    }

    AggregatorMerkleInterface internal merkleAggregator;
    address public   token;
    uint256 public   projectId;
    bytes32 public   merkleRoot;
    uint256 public   curBatch;
    uint80 public    roundId;
    uint256 public   startedAt;
    uint256 public   updatedAt;

    // This is a packed array of booleans.
    mapping(uint256 => ClaimItem) private claimedBitMap;

    constructor(address aggregatorProxy_, address token_, uint256 projectId_) {
        merkleAggregator = AggregatorMerkleInterface(aggregatorProxy_);
        token = token_;
        projectId = projectId_;
    }

    function isLocked() public returns (bool){
        return merkleAggregator.isLocked();
    }

    function getRoot() public returns (bool){
        if(isLocked())
            return false;
        (
            roundId,
            curBatch,
            merkleRoot,
            startedAt,
            updatedAt
        ) = merkleAggregator.latestMerkleRoundData();
        return true;
    }

    function isClaimed(uint256 batch, uint256 index)
    public view override returns (bool) {
        if(curBatch != batch)
            return true;
        ClaimItem memory claim  =  claimedBitMap[batch];
        if(claim._index == index && claim._isClaimed == true)
            return true;
        return false;
    }

    function _setClaimed(uint256 index, address account) private {
        claimedBitMap[index] = ClaimItem({
        _batchId:curBatch,
        _index:index,
        _account:account,
        _isClaimed:true
        });
    }

    function claim(uint256 batch, uint256 index, address account, uint256 amount, bytes32[] calldata merkleProof)
    public  virtual override returns (bool) {
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
