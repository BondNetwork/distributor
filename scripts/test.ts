import { ethers } from "hardhat";
import { EthersData, loadEthersData } from "./helpers/ethers-data";
import { MAX_UINT_AMOUNT, ZERO_ADDRESS, getErrorMessage, } from "./helpers/utils";
import { Types } from "../typechain-types/contracts/DistributorFactory";
import { Contract,  parseUnits } from "ethers";
import { string } from "hardhat/internal/core/params/argumentTypes";


async function createDistributorByToken(ethsData: EthersData) {
    const wallet = ethsData.deployer;
    console.log("wallet ", wallet.address);
    const contractAddress = ethsData.contractAddress;
    const distributorFactory = await ethers.getContractAt("DistributorFactory",
        contractAddress.distributorFactory.proxy,
        wallet);
    console.log("distributorFactory address", await distributorFactory.getAddress());
    console.log("distributorFactory version", (await distributorFactory.getVersion()).toString());

    let tokenAddress = ZERO_ADDRESS;
    let tokenName = "MyToken";
    if (ethsData.network === 'localhost') {       
        if (tokenName === "MyToken") {
            tokenAddress = ethsData.contractAddress.myToken.address;
            const myToken = await ethers.getContractAt("MyToken",
                tokenAddress,
                wallet);
            let balance = await myToken.balanceOf(wallet.address);
            if (balance < 10000) {
                const transaction = await myToken.transfer(wallet.address, 100000);
                await transaction.wait();
                balance = await myToken.balanceOf(wallet.address);
            }
            console.log("balance", balance);
            await myToken.approve(contractAddress.distributorFactory.proxy, 10000);
            console.log("my token approve....", (await myToken.allowance(wallet.address,
                contractAddress.distributorFactory.proxy)).toString());
        } else if (tokenName === "FaucetToken") {
            tokenAddress = ethsData.contractAddress.wbtc.address;
            const faucetToken = await ethers.getContractAt("FaucetToken",
                tokenAddress,
                wallet);
            let balance = await faucetToken.balanceOf(wallet.address);
            if (balance < 10000) {
                const transaction = await faucetToken.mint(wallet.address, 100000);
                await transaction.wait();
                balance = await faucetToken.balanceOf(wallet.address);
            }
            console.log("balance", balance);
            await faucetToken.approve(contractAddress.distributorFactory.proxy, 10000);
            console.log("faucet token approve....", (await faucetToken.allowance(wallet.address,
                contractAddress.distributorFactory.proxy)).toString());
        }
    } else {
        if (tokenName === "FaucetToken") {
            //wbtc
            tokenAddress = '0x3461461C676381FDD65e419a4c41f21fdB4e9031';
            const tokenABI = '[{"inputs":[{"internalType":"string","name":"name_","type":"string"},{"internalType":"string","name":"symbol_","type":"string"},{"internalType":"uint8","name":"decimals_","type":"uint8"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"subtractedValue","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"addedValue","type":"uint256"}],"name":"increaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"mint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"}]';
            const wbtc = new Contract(tokenAddress, tokenABI, wallet);
            let balance = await wbtc.balanceOf(wallet.address);
            if (balance < 10000) {
                const transaction = await wbtc.mint(wallet.address, 100000);
                await transaction.wait();
                balance = await wbtc.balanceOf(wallet.address);
            }
            console.log("balance", balance);
            await wbtc.approve(contractAddress.distributorFactory.proxy, 10000);
            console.log("approve....", (await wbtc.allowance(wallet.address,
                contractAddress.distributorFactory.proxy)).toString());
            
        } else {
            //bondToken
            tokenAddress = '0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e';
            const tokenABI = '[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"},{"name":"_data","type":"bytes"}],"name":"transferAndCall","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_subtractedValue","type":"uint256"}],"name":"decreaseApproval","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_addedValue","type":"uint256"}],"name":"increaseApproval","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"},{"indexed":false,"name":"data","type":"bytes"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]';
            const onlineToken = new Contract(tokenAddress, tokenABI, wallet);
            //await onlineToken.approve(contractAddress.distributorFactory.proxy, 10000);
            console.log("approve....", (await onlineToken.allowance(wallet.address,
                contractAddress.distributorFactory.proxy)).toString());
        }
    }
    const taskId = (new Date()).getTime().toString();
    console.log("taskId ", taskId);
     let params: Types.CreateDistributorParamsStruct = {
         aggregatorAddress: ZERO_ADDRESS,
         token: tokenAddress,
         projectId: "xxxxx",
         taskId: taskId,
         amount: 200,
         startTimestamp: 1697558400,
         endTimestamp: 1698422400,
         rewardPerBatch: 2
     }
    /*const params: Types.CreateDistributorParamsStruct = {
        aggregatorAddress: '0xcD03049f5C5F55b17ba28821fe07d0E4ED8f4dD6',
        token: tokenAddress,
        projectId: "P63d9JlA21mKbznKDBaWbgQepZ",
        taskId: taskId,
        amount: 200,
        startTimestamp: 1697558400,
        endTimestamp: 1698422400,
        rewardPerBatch: 2
    }*/
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

async function updateDistributorByToken(ethsData: EthersData) {
    const wallet = ethsData.deployer;
    console.log("wallet ", wallet.address);
    const contractAddress = ethsData.contractAddress;
    const distributorFactory = await ethers.getContractAt("DistributorFactory",
        contractAddress.distributorFactory.proxy,
        wallet);
    console.log("distributorFactory address", await distributorFactory.getAddress());
    console.log("distributorFactory version", (await distributorFactory.getVersion()).toString());

    let tokenAddress = ZERO_ADDRESS;
    let tokenName = "FaucetToken";
    if (ethsData.network === 'localhost') {
        if (tokenName === "MyToken") {
            tokenAddress = ethsData.contractAddress.myToken.address;
           
        } else if (tokenName === "FaucetToken") {
            tokenAddress = ethsData.contractAddress.wbtc.address;
         
        }
    } else {
        if (tokenName === "FaucetToken") {
            //wbtc
            tokenAddress = '0x3461461C676381FDD65e419a4c41f21fdB4e9031';
            

        } else {
            //bondToken
            tokenAddress = '0x6c51561c4F5e4ba30209732FF7499a1e4AdE052e';
        }
    }

    const taskId = '1698336774599';
    console.log("taskId ", taskId);

    const params: Types.CreateDistributorParamsStruct = {
        aggregatorAddress: '0xcD03049f5C5F55b17ba28821fe07d0E4ED8f4dD6',
        token: tokenAddress,
        projectId: "P63d9JlA21mKbznKDBaWbgQepZ",
        taskId: taskId,
        amount: 200,
        startTimestamp: 1697558400,
        endTimestamp: 1698422400,
        rewardPerBatch: 2
    }
    const transaction = await distributorFactory.updateDistributor(params);

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

    const taskId = (new Date()).getTime().toString();
    console.log("taskId ", taskId);
    let params: Types.CreateDistributorParamsStruct = {
        aggregatorAddress: ZERO_ADDRESS,
        token: ZERO_ADDRESS,
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
        console.log("token balance ", await distributor.getTokenBalance());

        //
    }

    const withdrawAddress = ethsData.accounts[0].address;
    console.log("withdrawAddress ", withdrawAddress);

    const txs = await distributorFactory.withdrawETH(distributorAddress, MAX_UINT_AMOUNT, withdrawAddress);
    console.log("txs ", txs.hash);
    await txs.wait();
    
    console.log(await ethers.provider.getBalance(withdrawAddress));


}

async function depositETH(ethsData: EthersData, distributorAddress: string) {
    const wallet = ethsData.deployer;
    console.log("wallet ", wallet);
    const contractAddress = ethsData.contractAddress;
    const distributorFactory = await ethers.getContractAt("DistributorFactory",
        contractAddress.distributorFactory.proxy,
        wallet);
    console.log("distributorFactory address", await distributorFactory.getAddress());
    console.log("distributorFactory version", (await distributorFactory.getVersion()).toString());
    const distributor = await ethers.getContractAt("MerkleDistributor",
    distributorAddress,
    wallet);
    console.log("rewardPerBatch ", await distributor.getRewardPerBatch());
    console.log("token balance ", await distributor.getTokenBalance());
    const value = parseUnits('300', 'wei') //parseEther('0.0003'); // 0.001 ETH
    const gasPrice = parseUnits('5', 'gwei')
    console.log("gasPrice ", gasPrice);
    const transaction = await distributorFactory.depositETH(distributorAddress,
        {
            value: value, 
            gasPrice: gasPrice
        });
    await transaction.wait();
    console.log("txs ", transaction.hash);
    console.log("token new balance ", await distributor.getTokenBalance());
}

async function deposit(ethsData: EthersData, distributorAddress: string){
    const wallet = ethsData.deployer;
    console.log("wallet ", wallet);
    const distributor = await ethers.getContractAt("MerkleDistributor",
    distributorAddress,
    wallet);
    console.log("rewardPerBatch ", await distributor.getRewardPerBatch());
    console.log("token balance ", await distributor.getTokenBalance());
    const value = parseUnits('300', 'wei') //parseEther('0.0003'); // 0.001 ETH
    const gasPrice = parseUnits('5', 'gwei')
    console.log("gasPrice ", gasPrice);
    const myToken = await ethers.getContractAt("MyToken",
    ethsData.contractAddress.myToken.address,
        wallet);
    let balance = await myToken.balanceOf(wallet.address);
    if (balance < 10000) {
        const transaction = await myToken.transfer(wallet.address, 100000);
        await transaction.wait();
        balance = await myToken.balanceOf(wallet.address);
    }
    console.log("balance", balance);
    await myToken.approve(distributorAddress, 30000);
    console.log("my token approve....", (await myToken.allowance(wallet.address,
        distributorAddress)).toString());

    const transaction = await distributor.deposit(value);
    await transaction.wait();
    console.log("txs ", transaction.hash);
    console.log("token new balance ", await distributor.getTokenBalance());
}

async function main() {
    const ethsData = await loadEthersData();
    //await updateDistributorByToken(ethsData);
    //await createDistributorByToken(ethsData);
    //await createDistributorByEth(ethsData);
    //await depositETH(ethsData, "0x1F708C24a0D3A740cD47cC0444E9480899f3dA7D");
    await deposit(ethsData, "0x8aCd85898458400f7Db866d53FCFF6f0D49741FF");
}


main().catch((error) => {
    console.log(error);
    console.error(getErrorMessage(error));
    process.exitCode = 1;
});
