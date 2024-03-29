// SPDX-License-Identifier: MIT

pragma solidity >=0.6.0;

library Errors {

    error NotOwner();
    error TaskItemAlreadyExist();
    error TaskItemNoExist();
    error ETHTransferFailed();
    error InvalidWETHAddress();
    error NeitherFactoryNorOwner();
 
}
