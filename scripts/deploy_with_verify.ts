import { ethers } from "hardhat";
import { deployWithVerify, waitForTx } from "./helpers/utils";
async function main() {
  /*const merkleDistributorF = await ethers.getContractFactory("MerkleDistributor");
  const [deployer] = await ethers.getSigners();
  const merkleDistributorC = await (
    merkleDistributorF.deploy("0xB30BA212D5C84E0d471d0059722997dFd88D39c1", "0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e", 13),
    ["0xB30BA212D5C84E0d471d0059722997dFd88D39c1", "0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e", 13],
    'contracts/MerkleDistributor.sol:MerkleDistributor'
    );
    console.log('merkleDistributor address', merkleDistributorC.address);*/
    const [deployer] = await ethers.getSigners();
    const merkleDistributorF = await ethers.getContractFactory("DistributorFactory");
    const merkleDistributorC = await deployWithVerify(await merkleDistributorF.deploy(), [], 'contracts/DistributorFactory.sol:DistributorFactory');
    console.log('merkleDistributorF address', merkleDistributorC.address);
    const transparentUpgradeableProxy = await ethers.getContractFactory("TransparentUpgradeableProxy");
    const merkleDistributorFProxy = await deployWithVerify(
      await transparentUpgradeableProxy.deploy(
          merkleDistributorC.address,
          "0xF4888aEd11bFA9ADcfa25B42E11Cb6E064A354b8",
          [],
      ),
      [
        merkleDistributorC.address,
        "0xF4888aEd11bFA9ADcfa25B42E11Cb6E064A354b8",
        [],
      ],
      'contracts/TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy'
  );
  console.log('merkleDistribmerkleDistributorFProxyutorF address', merkleDistributorFProxy.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
