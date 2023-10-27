
import { EthersError } from 'ethers';
import fs from 'fs';
export const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000';
export const MAX_UINT_AMOUNT =
    '115792089237316195423570985008687907853269984665640564039457584007913129639935';

export async function delay(ms: number) {
    return new Promise((resolve) => setTimeout(resolve, ms));
}

export const loadFile = (file: string) => {
    let path = process.cwd().concat('/', file);
    console.log(path);
    try {
        const data = fs.readFileSync(path, {
            encoding: 'utf-8',
        });
        const addresses = JSON.parse(data);
        return addresses;
    } catch (error: any) {
        console.log(error.message);
    }
    return {};
};

export async function loadContactsAddresses(network: string) {
    let data = loadFile(`address-${network}.json`);
    console.log(data);
    return data;
}
export async function saveContractsAddresses(network: string, addrs: any) {
    const json = JSON.stringify(addrs, null, 2);
    fs.writeFileSync(`address-${network}.json`, json, 'utf-8');
    console.log(addrs);
}

 export function getErrorMessage(e: Error): string {
    const ethersError = e as EthersError
    if (ethersError.info) {
        return ethersError.info.error.message;
    }
    return e.message;
};