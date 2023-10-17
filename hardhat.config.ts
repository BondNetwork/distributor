import '@nomiclabs/hardhat-ethers'
import '@nomiclabs/hardhat-etherscan'
import '@nomiclabs/hardhat-waffle'
import '@openzeppelin/hardhat-upgrades'
import '@typechain/hardhat'
import 'hardhat-abi-exporter'
import 'hardhat-contract-sizer'
import 'solidity-coverage'

const COMPILER_SETTINGS = {
  optimizer: {
    enabled: true,
    runs: 500,
    },
  metadata: {
    bytecodeHash: 'none',
  },
  viaIR: true,
}

const ACCOUNT_PRIVATE_KEY = "b588f6bf79507310840ba922e1a28c8cd16a5db34ac8161a8c5932692d0addfc";

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
export default {
  abiExporter: {
    path: './abi',
  },
  paths: {
    artifacts: './artifacts',
    cache: './cache',
    sources: './contracts',
    tests: './test',
  },
  typechain: {
    outDir: './typechain',
    target: 'ethers-v5',
  },
  networks: {
    hardhat: {
    },
    localhost: {
      url: "http://127.0.0.1:8545",
      chainId: 31337
    },
    sepolia: {
      url: "https://twilight-magical-snowflake.ethereum-sepolia.discover.quiknode.pro/d88ffbe6a53853e28f56142c912a288eb298dada/",
      chainId: 11155111,
      accounts: [ACCOUNT_PRIVATE_KEY]
    }
  },
  solidity: {
    compilers: [
      {
        version: '0.4.24',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.5.0',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.6.6',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.7.6',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.8.6',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.8.15',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.8.20',
        settings: COMPILER_SETTINGS,
      },
    ],
  },
  contractSizer: {
    alphaSort: true,
    runOnCompile: false,
    disambiguatePaths: false,
  },
  mocha: {
    timeout: 100000,
    forbidOnly: Boolean(process.env.CI),
  },
  /*etherscan: {
    // Your API key for Etherscan
    // Obtain one at https://etherscan.io/
    apiKey: "BZPRYWQJZYBIN5YR459V1NBG8A14KPWRQ9"
  },*/
  etherscan: {
    apiKey: {
      sepolia: "BZPRYWQJZYBIN5YR459V1NBG8A14KPWRQ9"
    },
    customChains: [
      {
        network: "sepolia",
        chainId: 11155111,
        urls: {
          apiURL: "http://api-sepolia.etherscan.io/api",  // https => http
          browserURL: "https://sepolia.etherscan.io"
        }
      }
    ]
  }
}
