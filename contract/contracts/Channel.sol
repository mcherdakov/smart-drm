// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Channel {
    address payable private channelSender;
    address payable private channelRecipient;
    uint256 private startDate;
    uint256 private channelTimeout;

    mapping(bytes32 => address) signatures;

    constructor(address payable to, uint256 timeout) payable {
        channelRecipient = to;
        channelSender = payable(msg.sender);
        startDate = block.timestamp;
        channelTimeout = timeout;
    }

    function CloseChannel(
        bytes32 h,
        uint8 v,
        bytes32 r,
        bytes32 s,
        uint256 value
    ) public {
        address signer;
        bytes32 proof;

        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 prefixedHash = keccak256(abi.encodePacked(prefix, h));

        signer = ecrecover(prefixedHash, v, r, s);

        if (signer != channelSender && signer != channelRecipient) {
            revert();
        }

        proof = keccak256(abi.encodePacked(this, value));

        require(proof != h);

        if (signatures[proof] == address(0x0)) {
            signatures[proof] = signer;
            return;
        }

        if (signatures[proof] == signer) {
            return;
        }

        if (!channelRecipient.send(value)) {
            revert();
        }

        selfdestruct(channelSender);
    }

    function ChannelTimeout() public {
        require(startDate + channelTimeout <= block.timestamp);
        selfdestruct(channelSender);
    }
}
