import { ethers } from "hardhat";
import {MerkleDistributor__factory} from "../typechain-types";

async function main() {
  /*const merkleDistributorF = await ethers.getContractFactory("MerkleDistributor");
  const merkleDistributorC = await merkleDistributorF.deploy("0xBEc49fA140aCaA83533fB00A2BB19bDdd0290f25", "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853", 12);
  console.log('merkleDistributor address', merkleDistributorC.address);*/

  let provider = ethers.getDefaultProvider("ws://127.0.0.1:8545");

  let privateKey = '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80';
  let wallet = new ethers.Wallet(privateKey, provider);

  //const merkleDistributorF = await ethers.getContractFactory("MerkleDistributor");
  const merkleDistributor = await MerkleDistributor__factory.connect('0x5FeaeBfB4439F3516c74939A9D04e95AFE82C4ae', provider);
  let contractWithSigner = await merkleDistributor.connect(wallet);

  const ret = await contractWithSigner.claim(2310, 1, '0x431aa3467889D68f26E975ED2246b1E2cAd2B3B2', 3000000000,
  ["0x5258ab39ecc6cdbfdf63da24222c77541e38fe285cdfcd3013b3582e76586186", "0x9ad36e3dff1361f966b07e21a31a9b1d66ccb2f7cee48f8eacc00a3c855feafd", "0x5b62384877fb8be24c7c4d12e1367809cc6b0f9a848a4799a4d3bd08136b2a1a"],
  {gasPrice:10000000, gasLimit:2100000})
  console.log('calim', ret)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
