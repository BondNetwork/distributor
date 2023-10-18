

import dotenv from 'dotenv';
dotenv.config({});

export const DEFAULT_BLOCK_GAS_LIMIT = 12450000;
export const MNEMONIC_PATH = "m/44'/60'/0'/0";
export const MNEMONIC = process.env.MNEMONIC || 'test test test test test test test test test test test junk';
export const TRACK_GAS = process.env.TRACK_GAS === 'true';
export const ETHERSCAN_BLOCK_EXPLORER_KEY = process.env.ETHERSCAN_BLOCK_EXPLORER_KEY || '';

export const SEPOLIA_RPC_URL = process.env.SEPOLIA_RPC_URL || "https://ethereum-sepolia.blockpi.network/v1/rpc/public"



