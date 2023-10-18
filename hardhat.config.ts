import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "hardhat-contract-sizer";
import '@openzeppelin/hardhat-upgrades';

import { ETHERSCAN_BLOCK_EXPLORER_KEY } from "./utils/config";
import { Sepolia, getNetworkConfig } from "./utils/network";

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.19",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      }
    }
  },
  networks: {
    sepolia: getNetworkConfig(Sepolia.name),
  },
  etherscan: {
    apiKey: {
      sepolia: ETHERSCAN_BLOCK_EXPLORER_KEY,
    },
  },
  gasReporter: {
    enabled: true,
    currency: 'USD',
  },
};

export default config;
