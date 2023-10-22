
import { ethers } from "hardhat";
import { deployDistributorFactory, deployMyToken, deployWETH } from "./contracts/deploy";
import { loadEthersData } from "./helpers/ethers-data";
import { saveContractsAddresses } from "./helpers/utils";

async function main() {

    const ethsData = await loadEthersData();

    if (ethsData.network == 'localhost') {

        await deployWETH(ethsData);
        await deployMyToken(ethsData);

    }

    await deployDistributorFactory(ethsData, ethsData.network === 'localhost' ? false : true);
  
    saveContractsAddresses(ethsData.network, ethsData.contractAddress);

    console.log('contracts deploy success!');
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
