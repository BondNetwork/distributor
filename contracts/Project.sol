contract Project {
    uint256 chainId
    uint256 projectType
    uint256 startBlock
    uint256 endBlock
    uint256 totalRewardAmount
    IMerkleDistributor merkleDistributor
}

contract ProjectManager {
    uint256 chainId
    maping(uint256=>Project) Projects
    function newProject(uint256 chainId, uint256 projectType...)
}