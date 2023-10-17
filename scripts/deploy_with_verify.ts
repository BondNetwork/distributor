import { ethers } from "hardhat";
import {deployProxy} from "./hardhatUpdate/utils"
import {verifyContract} from "./hardhatUpdate/utils"
import hre from "hardhat";

async function main() {
    const accounts = await ethers.getSigners();
    const options = {
      signer: accounts[0]
    };
    const contractData = await deployProxy("DistributorFactory",
      {
          options: options,
          args: [
          ]
      },
      true
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
