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

struct CreatorClicks {
    address creator;
    uint32 clicks;
}

contract SmartDRM {
    address i_creator;

    mapping(address => Channel) private s_user_channel;
    mapping(Channel => Proof) private s_proofs;

    address[] private s_creators;
    mapping(address => uint32) private s_creator_clicks;

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

    function setCreatorsClicks(CreatorClicks[] memory clicks) public {
        if (msg.sender != i_creator) {
            revert("sender must be creator");
        }

        for (uint256 i = 0; i < clicks.length; i++) {
            if (s_creator_clicks[clicks[i].creator] == 0) {
                s_creators.push(clicks[i].creator);
            }

            s_creator_clicks[clicks[i].creator] = clicks[i].clicks;
        }
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

    function getCreatorClicks(address creator) public view returns (uint32) {
        return s_creator_clicks[creator];
    }

    function getChannelProof(Channel channel)
        public
        view
        returns (Proof memory)
    {
        return s_proofs[channel];
    }
}
