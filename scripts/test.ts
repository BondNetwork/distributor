import { ethers } from "hardhat";
import { EthersData, loadEthersData } from "./helpers/ethers-data";
import { ZERO_ADDRESS, delay } from "./helpers/utils";
import { Types } from "../typechain-types/contracts/DistributorFactory";
import { Contract, JsonRpcProvider, parseEther, parseUnits } from "ethers";

async function createDistributorByToken(ethsData: EthersData) {
    const wallet = ethsData.deployer;
    const contractAddress = ethsData.contractAddress;
    const distributorFactory = await ethers.getContractAt("DistributorFactory",
        contractAddress.distributorFactory.proxy,
        wallet);
    console.log("distributorFactory address", await distributorFactory.getAddress());
    console.log("distributorFactory version", (await distributorFactory.getVersion()).toString());

    let tokenAddress = ZERO_ADDRESS;
    if (ethsData.network == 'localhost') {
        tokenAddress = contractAddress.weth.address;
       
        const myToken = await ethers.getContractAt("MyToken",
            tokenAddress,
            wallet);
        await myToken.approve(contractAddress.distributorFactory.proxy, 10000);
        console.log("approve....", (await myToken.allowance(wallet.address,
            contractAddress.distributorFactory.proxy)).toString());
    } else {
        tokenAddress = '0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e';
        const tokenABI = '[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"},{"name":"_data","type":"bytes"}],"name":"transferAndCall","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_subtractedValue","type":"uint256"}],"name":"decreaseApproval","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_addedValue","type":"uint256"}],"name":"increaseApproval","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"},{"indexed":false,"name":"data","type":"bytes"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]';
        const myToken = new Contract(tokenAddress, tokenABI, wallet);
        await myToken.approve(contractAddress.distributorFactory.proxy, 1000);
        console.log("approve....", (await myToken.allowance(wallet.address,
            contractAddress.distributorFactory.proxy)).toString());
    }

    const taskId = (new Date()).getTime().toString();
    console.log("taskId ", taskId);
    let params: Types.CreateDistributorParamsStruct = {
        aggregatorAddress: ZERO_ADDRESS,
        token: tokenAddress,
        projectId: "xxxxx",
        taskId: taskId,
        amount: 10,
        startTimestamp: 1697558400,
        endTimestamp: 1698422400,
        rewardPerBatch: 2
    }
    const transaction = await distributorFactory.createDistributorByToken(params);

    await transaction.wait();
   
    const distributorAddress = await distributorFactory.getDistributorAddress(taskId);
    console.log("distributorAddress", distributorAddress);
    if (distributorAddress !== ZERO_ADDRESS) {
        const distributor = await ethers.getContractAt("MerkleDistributor",
            distributorAddress,
            wallet);

        console.log("rewardPerBatch ", await distributor.getRewardPerBatch());
        console.log("token balance ", await distributor.getTokenBalance());
    }

}

async function createDistributorByEth(ethsData: EthersData) {
    const wallet = ethsData.deployer;
    const contractAddress = ethsData.contractAddress;
    const distributorFactory = await ethers.getContractAt("DistributorFactory",
        contractAddress.distributorFactory.proxy,
        wallet);
    console.log("distributorFactory address", await distributorFactory.getAddress());
    console.log("distributorFactory version", (await distributorFactory.getVersion()).toString());

    let tokenAddress = ZERO_ADDRESS;
    if (ethsData.network == 'localhost') {
        tokenAddress = contractAddress.weth.address;
    } else {
        tokenAddress = '0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e';
    }

    const taskId = (new Date()).getTime().toString();
    console.log("taskId ", taskId);
    let params: Types.CreateDistributorParamsStruct = {
        aggregatorAddress: ZERO_ADDRESS,
        token: tokenAddress,
        projectId: "xxxxx",
        taskId: taskId,
        amount: 100,
        startTimestamp: 1697558400,
        endTimestamp: 1698422400,
        rewardPerBatch: 2
    }
    
    const value = parseUnits('1000', 'wei') //parseEther('0.001'); // 0.1 ETH
    const gasPrice = parseUnits('5', 'gwei')
    console.log("gasPrice ", gasPrice);
    const transaction = await distributorFactory.createDistributorByEth(params,
        {
            value: value, 
            gasPrice: gasPrice
        });
    await transaction.wait();
    console.log(await ethers.provider.getBalance(contractAddress.distributorFactory.proxy));

    const distributorAddress = await distributorFactory.getDistributorAddress(taskId);
    console.log("distributorAddress", distributorAddress);
    if (distributorAddress !== ZERO_ADDRESS) {
        const distributor = await ethers.getContractAt("MerkleDistributor",
            distributorAddress,
            wallet);

        console.log("rewardPerBatch ", await distributor.getRewardPerBatch());
        console.log("eth balance ", await distributor.getEthBalance());
        console.log("token balance ", await distributor.getTokenBalance());

        //
    }

    const withdrawAddress = ethsData.accounts[0].address;
    console.log("withdrawAddress ", withdrawAddress);

    const txs = await distributorFactory.withdrawETH(distributorAddress, 100, withdrawAddress);
    console.log("txs ", txs.hash);
    await txs.wait();
    
    console.log(await ethers.provider.getBalance(withdrawAddress));


}
async function main() {

    const ethsData = await loadEthersData();

    //await createDistributorByToken(ethsData);
    await createDistributorByEth(ethsData);

   
}


main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
