// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "./Channel.sol";

struct Proof {
    uint8 v;
    bytes32 r;
    bytes32 s;
    uint256 value;
    string date;
}

struct ChannelProof {
    Proof proof;
    Channel channel;
}

contract SmartDRM {
    address i_creator;

    mapping(address => Channel) private s_user_channel;
    mapping(Channel => Proof) private s_proofs;

    constructor() {
        i_creator = msg.sender;
    }

    function createChannel(uint256 timeout) public payable {
        address payable contractAddress = payable(address(this));
        Channel channel = (new Channel){value: msg.value}(
            contractAddress,
            timeout
        );
        s_user_channel[msg.sender] = channel;
        s_proofs[channel] = Proof(0, 0, 0, 0, "");
    }

    function setChannelsProofs(ChannelProof[] memory proofs) public {
        if (msg.sender != i_creator) {
            revert("sender must be creator");
        }

        for (uint256 i = 0; i < proofs.length; i++) {
            s_proofs[proofs[i].channel] = proofs[i].proof;
        }
    }

    function getUserChannel(address user) public view returns (Channel) {
        return s_user_channel[user];
    }

    function getChannelProof(Channel channel)
        public
        view
        returns (Proof memory)
    {
        return s_proofs[channel];
    }
}
