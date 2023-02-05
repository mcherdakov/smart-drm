// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

error SmartDRMOutOfBounds();

contract SmartDRM {
    address[] private s_creators;

    constructor() {}

    function registerCreator() public {
        s_creators.push(msg.sender);
    }

    function getCreator(uint256 index) public view returns (address) {
        if (index >= s_creators.length) revert SmartDRMOutOfBounds();

        return s_creators[index];
    }
}
