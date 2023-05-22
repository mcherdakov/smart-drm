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
    address i_owner;

    mapping(address => Channel) private s_user_channel;
    mapping(Channel => Proof) private s_proofs;

    address[] private s_creators;
    mapping(address => uint32) private s_creator_clicks;

    constructor() {
        i_owner = msg.sender;
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

    function closeChannel(Channel channel) public {
        if (msg.sender != i_owner) {
            revert("sender must be owner");
        }

        Proof memory proof = s_proofs[channel];
        channel.closeChannel(
            proof.v,
            proof.r,
            proof.s,
            proof.value,
            proof.date
        );
    }

    function splitBalance() public {
        if (msg.sender != i_owner) {
            revert("sender must be owner");
        }

        uint64 total_clicks = 0;

        for (uint256 i = 0; i < s_creators.length; i++) {
            total_clicks += s_creator_clicks[s_creators[i]];
        }

        uint256 balance = address(this).balance;

        for (uint256 i = 0; i < s_creators.length; i++) {
            uint256 share = (balance * s_creator_clicks[s_creators[i]]) /
                total_clicks;

            if (!payable(s_creators[i]).send(share)) {
                revert("payment error");
            }

            delete s_creator_clicks[s_creators[i]];
        }

        delete s_creators;
    }

    function setCreatorsClicks(CreatorClicks[] memory clicks) public {
        if (msg.sender != i_owner) {
            revert("sender must be owner");
        }
        for (uint256 i = 0; i < clicks.length; i++) {
            if (s_creator_clicks[clicks[i].creator] == 0) {
                s_creators.push(clicks[i].creator);
            }
            s_creator_clicks[clicks[i].creator] += clicks[i].clicks;
        }
    }

    function setChannelsProofs(ChannelProof[] memory proofs) public {
        if (msg.sender != i_owner) {
            revert("sender must be owner");
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

    receive() external payable {}
}
