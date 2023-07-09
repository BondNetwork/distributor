import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const ACCOUNT_PRIVATE_KEY = "b588f6bf79507310840ba922e1a28c8cd16a5db34ac8161a8c5932692d0addfc";

const COMPILER_SETTINGS = {
  optimizer: {
    enabled: true,
    runs: 1000000,
  },
  metadata: {
    bytecodeHash: 'none',
  },
}

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: "0.8.15",
        settings: {
          optimizer: {
            enabled: true,
            runs: 1000000,
            details: {
              yul: true,
            },
          },
        },
      },
    ],
  },
  networks: {
    localhost: {
    url: "http://127.0.0.1:8545",
    chainId: 31337,
    },
    sepolia: {
      url: "https://twilight-magical-snowflake.ethereum-sepolia.discover.quiknode.pro/d88ffbe6a53853e28f56142c912a288eb298dada/",
      chainId: 11155111,
      accounts: [ACCOUNT_PRIVATE_KEY]
    }
  },
  etherscan: {
    // Your API key for Etherscan
    // Obtain one at https://etherscan.io/
    apiKey: "BZPRYWQJZYBIN5YR459V1NBG8A14KPWRQ9"
  },
};

export default config;
