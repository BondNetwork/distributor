import { ethers, network } from "hardhat";
import { loadContactsAddresses } from "./utils";
import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";

export type EthersData = {
    network: string,
    deployer: SignerWithAddress,
    governance: SignerWithAddress,
    treasury: SignerWithAddress,
    proxyAdmin: SignerWithAddress,
    accounts: SignerWithAddress[],
    contractAddress: any
}
export async function loadEthersData(): Promise<EthersData> {
    let ethsNetWork = await ethers.provider.getNetwork();
    console.log(JSON.stringify(ethsNetWork));
    const accounts = await ethers.getSigners();
    const contractAddress = await loadContactsAddresses(network.name);
    return {
        network: network.name,
        deployer: accounts[0],
        governance: accounts[1],
        treasury: accounts[2],
        proxyAdmin: accounts[3],
        accounts: accounts.slice(4),
        contractAddress: contractAddress,
    };
}
export async function getDeployerNonce(deployer: SignerWithAddress): Promise<number> {
    return ethers.provider.getTransactionCount(
        deployer
    );
}