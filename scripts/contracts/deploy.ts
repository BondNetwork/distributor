
import { EthersData } from "../helpers/ethers-data";
import { checkContractProxy, checkDeployContract, deployContract, deployProxy, upgradeProxy } from "./utils";


export const deployDistributorFactory = async (ethsData: EthersData, verify: boolean = false) => {
    let contractAddress = ethsData.contractAddress;
    const name = 'DistributorFactory';
    let wethAddress = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2";
    if (ethsData.network == 'sepolia') {
        wethAddress = "0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9";
    }

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