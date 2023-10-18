
import { delay } from "../helpers/utils";
import hre, { ethers, upgrades } from "hardhat";
import { FactoryOptions } from "hardhat/types";

export type ContractData = {
    proxy?: string;
    address: string;
    timestamp: number;
};

export function checkDeployContract(contract: ContractData): boolean {
    if (
        contract !== undefined &&
        contract !== null &&
        typeof contract === "object"
    ) {
        if (contract.address !== undefined && contract.address !== "") {
            return false;
        }
    }
    return true;
}
export function checkContractProxy(contract: ContractData): boolean {
    if (
        contract !== undefined &&
        contract !== null &&
        typeof contract === "object"
    ) {
        if (contract.proxy !== undefined && contract.proxy !== "") {
            return false;
        }
    }
    return true;
}



async function verifyContract(
    name: string,
    contractAddress: string,
    args?: any
) {
    let count = 0;
    const maxTries = 3;
    while (true) {
        await delay(10000);
        try {
            console.log("Verifying contract at", contractAddress);
            await hre.run("verify:verify", {
                address: contractAddress,
                constructorArguments: args,
            });
            break;
        } catch (error) {
            if (String(error).includes("Already Verified")) {
                console.log(
                    `Already verified contract ${name} at address ${contractAddress}`
                );
                break;
            }
            if (++count == maxTries) {
                console.log(
                    `Failed to verify contract ${name} at address ${contractAddress}, error: ${error}`
                );
                break;
            }
            console.log(`Retrying... Retry #${count}, last error: ${error}`);
        }
    }
}

export type deployParams = {
    options?: FactoryOptions;
    args?: any[];
    initializer?: string;
};


export async function deployContract(
    name: string,
    params: deployParams,
    verify: boolean = false,
): Promise<ContractData> {
    console.log(``);
    const contract = await ethers.deployContract(name,
        params.args ? params.args : [],
        params.options);
    const address = await contract.getAddress();
    console.log(`deploy contract ${name} at address ${address} `);
    if (verify) {
        await verifyContract(name, address, params.args);
    }
    return {
        address: address,
        timestamp: Date.now(),
    };
}

export async function deployProxy(
    name: string,
    params: deployParams,
    verify: boolean = false,
): Promise<ContractData> {
    console.log(``);
    const contractFactory = await ethers.getContractFactory(name, params.options);
    const proxy = await upgrades.deployProxy(
        contractFactory,
        params.args,
        {
            initializer: params.initializer ? params.initializer : "initialize",
            kind: "uups",
            unsafeAllow: ["external-library-linking"],
        });
    await proxy.waitForDeployment();
    const proxyAddress = await proxy.getAddress();
    await delay(10000);
    const implAddress = await upgrades.erc1967.getImplementationAddress(
        proxyAddress
    );
    console.log(
        `deploy proxy contract ${name} at address ${implAddress}`
    );
    console.log(
        `deploy proxy contract ${name} at proxy address ${proxyAddress}`
    );
    if (verify) {
        await verifyContract(name, implAddress);
    }
    return {
        proxy: proxyAddress,
        address: implAddress,
        timestamp: Date.now(),
    };
}

export async function upgradeProxy(
    proxyAddress: string,
    name: string,
    params: deployParams,
    verify: boolean = false,
): Promise<ContractData> {
    console.log(``);
    const contractFactory = await ethers.getContractFactory(name, params.options);
    const proxy = await upgrades.upgradeProxy(
        proxyAddress,
        contractFactory,
        {
            kind: "uups",
            unsafeAllow: ["external-library-linking"],
            redeployImplementation: 'always' ,
        });
    await proxy.waitForDeployment();
    const tmpProxyAddress = await proxy.getAddress();
    console.log(
        `upgrade proxy contract ${name} at proxy address ${tmpProxyAddress}`
    );
    await delay(10000);
    const implAddress = await upgrades.erc1967.getImplementationAddress(
        proxyAddress
    );
    console.log(
        `upgrade proxy contract ${name} at address ${implAddress}`
    );
   
    if (verify) {
        await verifyContract(name, implAddress);
    }
    return {
        proxy: proxyAddress,
        address: implAddress,
        timestamp: Date.now(),
    };
}


