// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity >=0.8.0 <0.9.0;
import {IMerkleDistributor} from "./interfaces/IMerkleDistributor.sol";
import "./interfaces/AggregatorMerkleInterface.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";


contract MerkleDistributor is IMerkleDistributor {
    using SafeERC20 for IERC20;
    AggregatorMerkleInterface internal merkleAggregator;
    address public   token;
    uint256 public   projectId;
    bytes32 public   merkleRoot;
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

    function updateRoot() public  returns (bool){
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
    public  virtual override returns (bool) {
        if(!updateRoot()) return false;
        if(batch != curBatch) return false;
        if (isClaimed(curBatch,index)) return false;

        // Verify the merkle proof.
        bytes32 node = keccak256(abi.encodePacked(index, account, amount));
        if (!MerkleProof.verify(merkleProof, merkleRoot, node)) return false;

        // Mark it claimed and send the token.
        _setClaimed(index);

        IERC20(token).safeTransfer(account, amount);
        emit Claimed(batch, index, account, amount);
        return true;
    }
}
