// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;
import "hardhat/console.sol";
import {IMerkleDistributor} from "contracts/interfaces/IMerkleDistributor.sol";
import {AggregatorMerkleInterface} from "contracts/interfaces/AggregatorMerkleInterface.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import {Pausable} from "@openzeppelin/contracts/security/Pausable.sol";
import {Types} from "contracts/libraries/Types.sol";
import {Events} from "contracts/libraries/Events.sol";
import {Errors} from "contracts/libraries/Errors.sol";

contract MerkleDistributor is
    IMerkleDistributor,
    ReentrancyGuard,
    Ownable,
    Pausable
{
    using SafeERC20 for IERC20;
    AggregatorMerkleInterface internal _merkleAggregator;
    address internal _factory;
    address internal _token;
    bytes32 internal _projectId;
    bytes32 internal _taskId;
    uint256 internal _taskAmount;
    uint256 internal _taskStartTimestamp;
    uint256 internal _taskEndTimestamp;
    uint256 internal _rewardPerBatch;

    bytes32 internal _merkleRoot;
    uint256 internal _curBatch;
    uint256 internal _roundId;
    uint256 internal _startedAt;
    uint256 internal _updatedAt;

    mapping(uint256 => Types.IndexValue) internal _claimedBitMap;
    Types.KeyFlag[] internal _claimedkeys;

    constructor(address factory, Types.CreateDistributorParams memory params) {
        _factory = factory;
        _merkleAggregator = AggregatorMerkleInterface(params.aggregatorAddress);
        _token = params.token;
        _projectId = keccak256(bytes(params.projectId));
        _taskId = keccak256(bytes(params.taskId));
        _taskAmount = params.amount;
        _taskStartTimestamp = params.startTimestamp;
        _taskEndTimestamp = params.endTimestamp;
        _rewardPerBatch = params.rewardPerBatch;
    }

    modifier onlyOwnerOrFactory() {
        if (msg.sender != _factory && msg.sender != owner()) {
            revert Errors.NeitherFactoryNorOwner();
        }
        _;
    }

    function getRewardPerBatch() external view returns (uint256) {
        return _rewardPerBatch;
    }

    function token() external view override returns (address) {
        return _token;
    }

    function merkleRoot() external view override returns (bytes32) {
        return _merkleRoot;
    }

    function getTokenBalance() external view returns (uint256) {
        return IERC20(_token).balanceOf(address(this));
    }
  
    function isClaimed(
        uint256 batch,
        uint256 index
    ) external view override returns (bool) {
        return _isClaimed(batch, index);
    }

    function claim(
        uint256 batch,
        uint256 index,
        address account,
        uint256 amount,
        bytes32[] calldata merkleProof
    ) external virtual override whenNotPaused nonReentrant returns (bool) {
        if (msg.sender != account) revert("account is error");
        if (!_updateRoot()) revert("updateRoot is error");
        if (batch != _curBatch) revert("batch is error");
        if (_isClaimed(_curBatch, index)) revert("isClaimed");
        // Verify the merkle proof.
        bytes32 node = keccak256(abi.encodePacked(index, account, amount));
        if (!MerkleProof.verify(merkleProof, _merkleRoot, node))
            revert("proof verify error");

        // Mark it claimed and send the token.
        _setClaimed(index);
        console.log(
            "claim 02 toarr:%s, amount:%d, thisaddr:%s\n",
            account,
            amount,
            address(this)
        );
        IERC20(_token).transfer(account, amount);
        emit Claimed(
            string(abi.encodePacked(_projectId)),
            string(abi.encodePacked(_taskId)),
            batch,
            index,
            account,
            amount
        );
        return true;
    }

    function taskBaseInfo()
        external
        view
        returns (
            address,
            string memory,
            string memory,
            uint256,
            uint256,
            uint256
        )
    {
        return (
            _token,
            string(abi.encodePacked(_projectId)),
            string(abi.encodePacked(_taskId)),
            _taskAmount,
            _taskStartTimestamp,
            _taskEndTimestamp
        );
    }

    function deposit(uint256 amount) external onlyOwnerOrFactory nonReentrant {
        require(amount > 0, "Deposit: Amount must be > 0");
        IERC20(_token).safeTransferFrom(msg.sender, address(this), amount);
        emit Events.DistributorDeposit(
            string(abi.encodePacked(_projectId)),
            string(abi.encodePacked(_taskId)),
            msg.sender,
            amount
        );
    }

    function withdraw(uint256 amount) external onlyOwnerOrFactory nonReentrant {
        IERC20(_token).safeTransfer(msg.sender, amount);
        emit Events.DistributorWithdraw(
            string(abi.encodePacked(_projectId)),
            string(abi.encodePacked(_taskId)),
            msg.sender,
            amount
        );
    }

    function pauseDistribution() external onlyOwner whenNotPaused {
        _pause();
    }

    function unpauseDistribution() external onlyOwner whenPaused {
        _unpause();
    }

    function _updateRoot() private returns (bool) {
        if (_merkleAggregator.isLocked()) {
            return false;
        }
        (
            uint80 roundIdNew,
            uint256 curBatchNew,
            bytes32 merkleRootNew,
            uint256 startedAtNew,
            uint256 updatedAtNew
        ) = _merkleAggregator.latestMerkleRoundData(
                string(abi.encodePacked(_taskId))
            );

        if (curBatchNew != _curBatch) {
            uint256 arrayLength = _claimedkeys.length;
            for (uint i = 0; i < arrayLength; i++) {
                delete _claimedBitMap[_claimedkeys[i].key];
            }
            while (_claimedkeys.length > 0) {
                _claimedkeys.pop();
            }
            _curBatch = curBatchNew;
            _merkleRoot = merkleRootNew;
            _roundId = roundIdNew;
            _startedAt = startedAtNew;
            _updatedAt = updatedAtNew;
            //console.log("updateRoot curBatch:%d roundId:%d\n", _curBatch,roundId);
            //console.logBytes32(_merkleRoot);
        }
        return true;
    }

    function _setClaimed(uint256 index) private {
        Types.IndexValue memory claimedIndexValue = _claimedBitMap[index];
        if (claimedIndexValue.value == false) {
            uint256 keyIndex = _claimedkeys.length;
            _claimedBitMap[index].value = true;
            _claimedBitMap[index].keyIndex = keyIndex;
            _claimedkeys.push(Types.KeyFlag({key: index, deleted: false}));
        }
    }

    function _isClaimed(
        uint256 batch,
        uint256 index
    ) private view returns (bool) {
        if (_curBatch != batch) return true;
        return _claimedBitMap[index].value;
    }
}
