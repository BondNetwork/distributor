import { ethers } from "hardhat";
async function main() {
    const [owner, twoSigner] = await ethers.getSigners();
    const tx = await owner.sendTransaction({
      to: "0xF4888aEd11bFA9ADcfa25B42E11Cb6E064A354b8", // address
      value: ethers.utils.parseEther("1.0")
    });
    const receipt = await tx.wait()
    console.log(receipt);

    const balance = await owner.getBalance();
    console.log(balance);
  }
  
  main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });