import { ethers } from 'hardhat'
import {ContractFactory, Signer} from 'ethers'
import {MerkleDistributor__factory} from "../typechain-types";

before(async () => {
})

describe('claim_test', () => {
    it('one', async () => {
        const [ownerSigner] = await ethers.getSigners();
        const gasPrice = (await ownerSigner.getGasPrice());
        const contractWithSigner =  MerkleDistributor__factory.connect('0x6b417Cf7066a70D76568F0A6F7F2aACaEb14ea76', ownerSigner);
        const ret = await contractWithSigner.claim(1688391561, 3, '0x32DA87cffC407d29B104cb4dD47dD8e9b6308e0F', '350000000000000000',
        [
            "0x0e4a4d57b66fcc68d5599f0788748cab59da71ab0416bf5fb268aab20db5d270", 
            "0xe61ef60d8d6e563dc7bf3f5c189f7b50d5f1791061ef9a4168cf6302dad6614c", 
            "0xa645c4d41396f2224e3b4012c578c4a7667e103e27749e59dbb36242fd10303d"
        ],
        {gasPrice:1000000000, gasLimit:210000})
        console.log('calim', ret)
        const recipt = await ret.wait(1)
        console.log('recipt:', recipt)
    })
    it('gasprice', async () => {
        console.log('mine is gasprice');
        //let provider = ethers.getDefaultProvider("https://twilight-magical-snowflake.ethereum-sepolia.discover.quiknode.pro/d88ffbe6a53853e28f56142c912a288eb298dada/");
        //let privateKey = 'b588f6bf79507310840ba922e1a28c8cd16a5db34ac8161a8c5932692d0addfc';
        const [ownerSigner] = await ethers.getSigners();
        const gasPrice = await ownerSigner.getGasPrice();
        console.log('gasprice:', gasPrice.toString());
    })
})
