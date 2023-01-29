// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract SmartDRM {
    address[] private s_creators;

    constructor() {}

    function registerCreator() public {
        s_creators.push(msg.sender);
    }

    function getCreator(uint256 index) public view returns (address) {
        return s_creators[index];
    }
}
