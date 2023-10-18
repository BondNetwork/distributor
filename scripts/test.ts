import { ethers } from "hardhat";
import { loadEthersData } from "./helpers/ethers-data";
import { ZERO_ADDRESS } from "./helpers/utils";
import { Types } from "../typechain-types/contracts/DistributorFactory";

async function main() {

    const ethsData = await loadEthersData();

    const contractAddress = ethsData.contractAddress;
    
    let testAccount = ethsData.accounts[0];

    const distributorFactory = await ethers.getContractAt("DistributorFactory",
        contractAddress.distributorFactory.proxy,
        ethsData.governance);
    console.log("distributorFactory address", await distributorFactory.getAddress());
    console.log("distributorFactory version", (await distributorFactory.getVersion()).toString());
    
    
    let tokenAddress = ZERO_ADDRESS;
    if (ethsData.network == 'localhost') {
        tokenAddress = contractAddress.myToken.address;
        const myToken = await ethers.getContractAt("MyToken",
            tokenAddress,
            testAccount);
        //console.log("totalSupply....", (await myToken.totalSupply()).toString());
        //await myToken.transfer(testAccount.address, 20000);
        console.log("testAccount....", (await myToken.balanceOf(testAccount.address)).toString());
        await myToken.approve(contractAddress.distributorFactory.proxy, 10000);
        console.log("approve....", (await myToken.allowance(testAccount.address, contractAddress.distributorFactory.proxy)).toString());
    } else {
        tokenAddress = '0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e';
    }

    const taskId = "12345678947";
    let params: Types.CreateDistributorParamsStruct = {
        aggregatorAddress: ZERO_ADDRESS,
        token: contractAddress.myToken.address,
        account: testAccount.address,
        ethValue: 0,
        projectId: "xxxxx",
        taskId: taskId,
        amount: 1000,
        startTimestamp: 1697558400,
        endTimestamp: 1698422400,
        rewardPerBatch: 200
    }
    await distributorFactory.createDistributor(params);

    const distributorAddress = await distributorFactory.getDistributorAddress(taskId);
    console.log("distributorAddress", distributorAddress);
    const distributor = await ethers.getContractAt("MerkleDistributor",
        distributorAddress,
        testAccount);
    
    console.log("rewardPerBatch ", await distributor.getRewardPerBatch());
    console.log("balance ", await distributor.getContractBalance());
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
