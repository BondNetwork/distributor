

import {  MNEMONIC, MNEMONIC_PATH,  SEPOLIA_RPC_URL } from "./config";


export type ChainNetwork = {
  name: string,
  chainId: number,
  url: string,
};

export const Hardhat: ChainNetwork = {
  name: "hardhat",
  chainId: 31337,
  url: "http://localhost:8545"
}

export const Sepolia: ChainNetwork = {
  name: "sepolia",
  chainId: 11155111,
  url: SEPOLIA_RPC_URL
}


export const networks = new Map<string, ChainNetwork>([
  [Hardhat.name, Hardhat],
  [Sepolia.name, Sepolia],
]);

export const getNetworkConfig = (name: string) => {
  let chainNetwork = networks.get(name);
  //console.log(chainNetwork);
  return {
    url: chainNetwork?.url,
    accounts: {
      mnemonic: MNEMONIC,
      path: MNEMONIC_PATH,
      initialIndex: 0,
      count: 10,
    },
  };
};
