// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;
import "hardhat/console.sol";
import {IMerkleDistributor} from "./interfaces/IMerkleDistributor.sol";
import "./interfaces/AggregatorMerkleInterface.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract MerkleDistributor is IMerkleDistributor, ReentrancyGuard {
    using SafeERC20 for IERC20;
    AggregatorMerkleInterface internal merkleAggregator;
    address override public   token;
    uint256 public   projectId;
    bytes32 override public   merkleRoot;
    uint256 public   curBatch;
    uint80  public   roundId;
    uint256 public   startedAt;
    uint256 public   updatedAt;

    struct IndexValue {uint256 keyIndex; bool value;}
    struct KeyFlag { uint256 key; bool deleted; }

    mapping(uint256 => IndexValue) private claimedBitMap;
    KeyFlag[] private claimedkeys;
    
    constructor(address aggregatorProxy_, address token_, uint256 projectId_) {
        merkleAggregator = AggregatorMerkleInterface(aggregatorProxy_);
        token = token_;
        projectId = projectId_;
    }

    function updateRoot() private  returns (bool){
        if(merkleAggregator.isLocked())
        {
            return false;
        }
        uint256  curBatchNew;
        bytes32  merkleRootNew;
        uint80   roundIdNew;
        uint256  startedAtNew;
        uint256  updatedAtNew;
        (
            roundIdNew,
            curBatchNew,
            merkleRootNew,
            startedAtNew,
            updatedAtNew
        ) = merkleAggregator.latestMerkleRoundData();

        if(curBatchNew != curBatch)
        {    
            uint256 arrayLength = claimedkeys.length;
            for (uint i = 0; i < arrayLength; i++) {
                delete claimedBitMap[claimedkeys[i].key];      
            }
            while(claimedkeys.length > 0){
                claimedkeys.pop();
            }
            curBatch = curBatchNew;
            merkleRoot = merkleRootNew;
            roundId = roundIdNew;
            startedAt = startedAtNew;
            updatedAt = updatedAtNew;
            //console.log("updateRoot curBatch:%d roundId:%d\n", curBatch,roundId);
            //console.logBytes32(merkleRoot);
        }   
        return true;
    }

    function isClaimed(uint256 batch, uint256 index)
    public view override returns (bool) {
        if(curBatch != batch)
            return true;
        return claimedBitMap[index].value;
    }

    function _setClaimed(uint256 index) private {
        IndexValue memory claimedIndexValue = claimedBitMap[index];
        if (claimedIndexValue.value == false)
        {
            uint256 keyIndex = claimedkeys.length;
            claimedBitMap[index].value = true;
            claimedBitMap[index].keyIndex = keyIndex;
            claimedkeys.push(KeyFlag({key:index, deleted:false}));
        }
    }

    function claim(uint256 batch, uint256 index, address account, uint256 amount, bytes32[] calldata merkleProof)
    public  nonReentrant virtual override returns (bool) {
        if(msg.sender != account) revert("account is error");
        if(!updateRoot()) revert("updateRoot is error");
        if(batch != curBatch) revert("batch is error");
        if(isClaimed(curBatch,index)) revert("isClaimed");
        // Verify the merkle proof.
        bytes32 node = keccak256(abi.encodePacked(index, account, amount));
        if (!MerkleProof.verify(merkleProof, merkleRoot, node)) revert("proof verify error");

        // Mark it claimed and send the token.
        _setClaimed(index);
        console.log("claim 02 toarr:%s, amount:%d, thisaddr:%s\n", account, amount, address(this));
        IERC20(token).transfer(account, amount);
        emit Claimed(batch, index, account, amount);
        return true;
    }
}
