import { ethers } from "hardhat";
import {MerkleDistributor__factory} from "../typechain-types";
import { deployWithVerify, waitForTx } from "./helpers/utils";
async function main() {
  const merkleDistributorF = await ethers.getContractFactory("MerkleDistributor");
  const merkleDistributorC = await merkleDistributorF.deploy("0xB30BA212D5C84E0d471d0059722997dFd88D39c1", "0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e", 13);
  console.log('merkleDistributor address', merkleDistributorC.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
