
import { EthersData } from "../helpers/ethers-data";
import { ZERO_ADDRESS } from "../helpers/utils";
import { checkContractProxy, checkDeployContract, deployContract, deployProxy, upgradeProxy } from "./utils";


export const deployDistributorFactory = async (ethsData: EthersData, verify: boolean = false) => {
    const name = 'DistributorFactory';
    let contractAddress = ethsData.contractAddress;
    let wethAddress = ZERO_ADDRESS;
    if (ethsData.network == 'sepolia') {
        wethAddress = "0xd0df82de051244f04bff3a8bb1f62e1cd39eed92";
    } else if (contractAddress.weth){
        wethAddress = contractAddress.weth.address;

    }
    console.log("WETH ", wethAddress);

    if (checkDeployContract(contractAddress.distributorFactory)) {
        console.log(`\n\t--- Deploying ${name} ---`);        
        if (checkContractProxy(contractAddress.distributorFactory)) {
            console.log('deploy proxy distributor factory');
            contractAddress.distributorFactory = await deployProxy(name,
                {
                    args: [
                        wethAddress
                    ]
                },
                verify
            );

        } else {
            console.log('upgrade proxy distributor factory');
            contractAddress.distributorFactory = await upgradeProxy(
                contractAddress.distributorFactory.proxy,
                name,
                {
                    
                },
                verify
            );
        }
    }
}

export const deployMyToken = async (ethsData: EthersData) => {
    let contractAddress = ethsData.contractAddress;
    if (checkDeployContract(contractAddress.myToken)) {
        console.log("\n\t--- Deploying My Token ---");
        contractAddress.myToken = await deployContract(
            "MyToken",
            {
                args: [                    
                ]
            },
            false);
    }
}

export const deployWETH = async (ethsData: EthersData) => {
    let contractAddress = ethsData.contractAddress;
    if (checkDeployContract(contractAddress.weth)) {
        console.log("\n\t--- Deploying WETH---");
        contractAddress.weth = await deployContract(
            "WETH",
            {
                args: [
                ]
            },
            false);
    }
}