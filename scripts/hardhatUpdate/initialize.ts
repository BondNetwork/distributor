
import { ethers, upgrades } from "hardhat";
import { EthersData } from "../helpers/ethers-data";
import { ProtocolState } from "../helpers/utils";


export async function initializeBond(ethsData: EthersData) {
    
    console.log('initialize bond......');
    const contractAddress = ethsData.contractAddress;
    const bondHub = await ethers.getContractAt("BondHub",
        contractAddress.bondHub.proxy,
        ethsData.governance);
    console.log("bondHub address", await bondHub.getAddress());
    console.log("bondHub version", await bondHub.getVersion());
    const implAddress = await upgrades.erc1967.getImplementationAddress(
        contractAddress.bondHub.proxy
    );
    console.log("bondHub impl", implAddress);
    const state = await bondHub.getState();
    console.log("bondHub state", state.toString());
    if (state !== BigInt(ProtocolState.Unpaused)) {
        await bondHub.setState(ProtocolState.Unpaused);
    }
    
    const bondPeripheryAddress = await bondHub.getBondPeriphery();
    console.log("bondPeriphery ", bondPeripheryAddress, contractAddress.bondPeriphery.address);
    if (bondPeripheryAddress !== contractAddress.bondPeriphery.address) {
        await bondHub.setBondPeriphery(contractAddress.bondPeriphery.address);
    }
    if (contractAddress.lzBondHub) {
        const lzBondHub = await bondHub.getLzBondHub();
        console.log("LzbondHub ", lzBondHub, contractAddress.lzBondHub.address);
        if (lzBondHub !== contractAddress.lzBondHub.address) {
            await bondHub.setLzBondHub(contractAddress.lzBondHub.address);
        }
    }
   
}
