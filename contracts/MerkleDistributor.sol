// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;
import "hardhat/console.sol";
import {IMerkleDistributor} from "./interfaces/IMerkleDistributor.sol";
import "./interfaces/AggregatorMerkleInterface.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import {Pausable} from "@openzeppelin/contracts/security/Pausable.sol";

contract MerkleDistributor is IMerkleDistributor, ReentrancyGuard, Ownable, Pausable {
    using SafeERC20 for IERC20;
    AggregatorMerkleInterface internal merkleAggregator;
    address override public   token;
    string public   projectId;
    bytes32 override public   merkleRoot;
    uint256 public   curBatch;
    uint80  public   roundId;
    uint256 public   startedAt;
    uint256 public   updatedAt;
    string  public   taskId;
    uint256 public   taskAmount;
    uint256 public   taskStartTimestamp;
    uint256 public   taskEndTimestamp;

    struct IndexValue {uint256 keyIndex; bool value;}
    struct KeyFlag { uint256 key; bool deleted; }

    mapping(uint256 => IndexValue) private claimedBitMap;
    KeyFlag[] private claimedkeys;
    event Deposit(address indexed user, uint256 amount);
    event Withdraw(address indexed user, uint256 amount);
    constructor(address aggregator_, address token_, string memory projectId_, string memory taskId_, uint256 amount, uint256 startTimestamp, uint256 endTimestamp) {
        merkleAggregator = AggregatorMerkleInterface(aggregator_);
        token = token_;
        projectId = projectId_;
        taskId = taskId_;
        taskAmount = amount;
        taskStartTimestamp = startTimestamp;
        taskEndTimestamp = endTimestamp;
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
        ) = merkleAggregator.latestMerkleRoundData(taskId);

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
    public  whenNotPaused nonReentrant virtual override returns (bool) {
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
    
    function taskBaseInfo()
    public view returns (address, string memory, string memory, uint256, uint256, uint256) {
        return (token, projectId, taskId, taskAmount, taskStartTimestamp, taskEndTimestamp);
    }

    function deposit(uint256 amount) external onlyOwner nonReentrant {
        require(amount > 0, "Deposit: Amount must be > 0");
        IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        emit Deposit(msg.sender, amount);
    }

    function withdraw(uint256 amount) external onlyOwner nonReentrant {
        IERC20(token).safeTransfer(msg.sender, amount);
        emit Withdraw(msg.sender, amount);
    }

    function pauseDistribution() external onlyOwner whenNotPaused {
        _pause();
    }

    function unpauseDistribution() external onlyOwner whenPaused {
        _unpause();
    }
}
