import { ethers } from 'hardhat'
import {ContractFactory, Signer} from 'ethers'
import {MerkleDistributor__factory} from "../typechain-types";
import type {
    BaseContract,
    BigNumber,
    BytesLike,
    CallOverrides,
    ContractTransaction,
    Overrides,
    PopulatedTransaction,
    Signer,
    utils,
} from "ethers";

before(async () => {
})

describe('claim_test', () => {
    it('one', async () => {
        /*const [ownerSigner] = await ethers.getSigners();
        const contractWithSigner =  MerkleDistributor__factory.connect('0x6b417Cf7066a70D76568F0A6F7F2aACaEb14ea76', ownerSigner);
        const ret = await contractWithSigner.claim('1698710400', '0', 0x6b417Cf7066a70D76568F0A6F7F2aACaEb14ea76, '12121', []);
        utils.BigNumber.from(1698710400)
        console.log('calim', ret)
        const recipt = await ret.wait(1)
        console.log('recipt:', recipt)*/

        let provider = ethers.getDefaultProvider("https://twilight-magical-snowflake.ethereum-sepolia.discover.quiknode.pro/d88ffbe6a53853e28f56142c912a288eb298dada/");
        let privateKey = 'b588f6bf79507310840ba922e1a28c8cd16a5db34ac8161a8c5932692d0addfc';
        let wallet = new ethers.Wallet(privateKey, provider);
        const contractABI = "[\n" +
            "  {\n" +
            "    \"inputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"address\",\n" +
            "        \"name\": \"factory\",\n" +
            "        \"type\": \"address\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"components\": [\n" +
            "          {\n" +
            "            \"internalType\": \"address\",\n" +
            "            \"name\": \"aggregatorAddress\",\n" +
            "            \"type\": \"address\"\n" +
            "          },\n" +
            "          {\n" +
            "            \"internalType\": \"address\",\n" +
            "            \"name\": \"token\",\n" +
            "            \"type\": \"address\"\n" +
            "          },\n" +
            "          {\n" +
            "            \"internalType\": \"string\",\n" +
            "            \"name\": \"projectId\",\n" +
            "            \"type\": \"string\"\n" +
            "          },\n" +
            "          {\n" +
            "            \"internalType\": \"string\",\n" +
            "            \"name\": \"taskId\",\n" +
            "            \"type\": \"string\"\n" +
            "          },\n" +
            "          {\n" +
            "            \"internalType\": \"uint256\",\n" +
            "            \"name\": \"amount\",\n" +
            "            \"type\": \"uint256\"\n" +
            "          },\n" +
            "          {\n" +
            "            \"internalType\": \"uint256\",\n" +
            "            \"name\": \"startTimestamp\",\n" +
            "            \"type\": \"uint256\"\n" +
            "          },\n" +
            "          {\n" +
            "            \"internalType\": \"uint256\",\n" +
            "            \"name\": \"endTimestamp\",\n" +
            "            \"type\": \"uint256\"\n" +
            "          },\n" +
            "          {\n" +
            "            \"internalType\": \"uint256\",\n" +
            "            \"name\": \"rewardPerBatch\",\n" +
            "            \"type\": \"uint256\"\n" +
            "          }\n" +
            "        ],\n" +
            "        \"internalType\": \"struct Types.CreateDistributorParams\",\n" +
            "        \"name\": \"params\",\n" +
            "        \"type\": \"tuple\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"stateMutability\": \"nonpayable\",\n" +
            "    \"type\": \"constructor\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [],\n" +
            "    \"name\": \"NeitherFactoryNorOwner\",\n" +
            "    \"type\": \"error\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"anonymous\": false,\n" +
            "    \"inputs\": [\n" +
            "      {\n" +
            "        \"indexed\": true,\n" +
            "        \"internalType\": \"string\",\n" +
            "        \"name\": \"projectId\",\n" +
            "        \"type\": \"string\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"indexed\": false,\n" +
            "        \"internalType\": \"string\",\n" +
            "        \"name\": \"taskId\",\n" +
            "        \"type\": \"string\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"indexed\": false,\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"batch\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"indexed\": false,\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"index\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"indexed\": false,\n" +
            "        \"internalType\": \"address\",\n" +
            "        \"name\": \"account\",\n" +
            "        \"type\": \"address\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"indexed\": false,\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"amount\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"name\": \"Claimed\",\n" +
            "    \"type\": \"event\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"anonymous\": false,\n" +
            "    \"inputs\": [\n" +
            "      {\n" +
            "        \"indexed\": true,\n" +
            "        \"internalType\": \"address\",\n" +
            "        \"name\": \"previousOwner\",\n" +
            "        \"type\": \"address\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"indexed\": true,\n" +
            "        \"internalType\": \"address\",\n" +
            "        \"name\": \"newOwner\",\n" +
            "        \"type\": \"address\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"name\": \"OwnershipTransferred\",\n" +
            "    \"type\": \"event\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"anonymous\": false,\n" +
            "    \"inputs\": [\n" +
            "      {\n" +
            "        \"indexed\": false,\n" +
            "        \"internalType\": \"address\",\n" +
            "        \"name\": \"account\",\n" +
            "        \"type\": \"address\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"name\": \"Paused\",\n" +
            "    \"type\": \"event\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"anonymous\": false,\n" +
            "    \"inputs\": [\n" +
            "      {\n" +
            "        \"indexed\": false,\n" +
            "        \"internalType\": \"address\",\n" +
            "        \"name\": \"account\",\n" +
            "        \"type\": \"address\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"name\": \"Unpaused\",\n" +
            "    \"type\": \"event\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"batch\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"index\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"internalType\": \"address\",\n" +
            "        \"name\": \"account\",\n" +
            "        \"type\": \"address\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"amount\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"internalType\": \"bytes32[]\",\n" +
            "        \"name\": \"merkleProof\",\n" +
            "        \"type\": \"bytes32[]\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"name\": \"claim\",\n" +
            "    \"outputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"bool\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"bool\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"stateMutability\": \"nonpayable\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"amount\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"name\": \"deposit\",\n" +
            "    \"outputs\": [],\n" +
            "    \"stateMutability\": \"nonpayable\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [],\n" +
            "    \"name\": \"getRewardPerBatch\",\n" +
            "    \"outputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"stateMutability\": \"view\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [],\n" +
            "    \"name\": \"getTokenBalance\",\n" +
            "    \"outputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"stateMutability\": \"view\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"batch\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"index\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"name\": \"isClaimed\",\n" +
            "    \"outputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"bool\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"bool\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"stateMutability\": \"view\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [],\n" +
            "    \"name\": \"merkleRoot\",\n" +
            "    \"outputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"bytes32\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"bytes32\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"stateMutability\": \"view\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [],\n" +
            "    \"name\": \"owner\",\n" +
            "    \"outputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"address\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"address\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"stateMutability\": \"view\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [],\n" +
            "    \"name\": \"pauseDistribution\",\n" +
            "    \"outputs\": [],\n" +
            "    \"stateMutability\": \"nonpayable\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [],\n" +
            "    \"name\": \"paused\",\n" +
            "    \"outputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"bool\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"bool\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"stateMutability\": \"view\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [],\n" +
            "    \"name\": \"renounceOwnership\",\n" +
            "    \"outputs\": [],\n" +
            "    \"stateMutability\": \"nonpayable\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [],\n" +
            "    \"name\": \"taskBaseInfo\",\n" +
            "    \"outputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"address\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"address\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"internalType\": \"string\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"string\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"internalType\": \"string\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"string\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      },\n" +
            "      {\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"stateMutability\": \"view\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [],\n" +
            "    \"name\": \"token\",\n" +
            "    \"outputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"address\",\n" +
            "        \"name\": \"\",\n" +
            "        \"type\": \"address\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"stateMutability\": \"view\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"address\",\n" +
            "        \"name\": \"newOwner\",\n" +
            "        \"type\": \"address\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"name\": \"transferOwnership\",\n" +
            "    \"outputs\": [],\n" +
            "    \"stateMutability\": \"nonpayable\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [],\n" +
            "    \"name\": \"unpauseDistribution\",\n" +
            "    \"outputs\": [],\n" +
            "    \"stateMutability\": \"nonpayable\",\n" +
            "    \"type\": \"function\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"inputs\": [\n" +
            "      {\n" +
            "        \"internalType\": \"uint256\",\n" +
            "        \"name\": \"amount\",\n" +
            "        \"type\": \"uint256\"\n" +
            "      }\n" +
            "    ],\n" +
            "    \"name\": \"withdraw\",\n" +
            "    \"outputs\": [],\n" +
            "    \"stateMutability\": \"nonpayable\",\n" +
            "    \"type\": \"function\"\n" +
            "  }\n" +
            "]\n";
        const tokenContract = new ethers.Contract("0x56b12df69036fff2987d59199e083ec907397a31", contractABI, provider);
        const contractWithSigner = tokenContract.connect(wallet);
        const tx = await contractWithSigner.claim(utils.BigNumber.from("1699574400"), utils.BigNumber.from("35"), '0xf4888aed11bfa9adcfa25b42e11cb6e064a354b8', BigNumber.from("1073696"),
            [
                "0xf176979809ccef791c386aa0bcc509e7f66c2598b2277a400771c6d89332940b",
                "0xf5687f5ffd3f9d353b0ab5c448b6af18f7293e410d33bb872d0368bc8328593e",
                "0x135d5598f411e54b48529f1bc0240e5d40640c34b2b97babe8654c06af40d3f3"
            ],
            {gasPrice:100000000, gasLimit:2100000});
        console.log("claim", tx);

    })
})
