// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Channel {
    address payable private i_channelSender;
    address payable private i_channelRecipient;
    uint256 private i_startDate;
    uint256 private i_channelTimeout;

    mapping(bytes32 => address) private s_signatures;

    constructor(address payable to, uint256 timeout) payable {
        i_channelRecipient = to;
        i_channelSender = payable(msg.sender);
        i_startDate = block.timestamp;
        i_channelTimeout = timeout;
    }

    function closeChannel(
        uint8 v,
        bytes32 r,
        bytes32 s,
        uint256 value,
        string memory date
    ) public {
        address signer;

        bytes32 proof = keccak256(abi.encode(this, value, date));

        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 prefixedHash = keccak256(abi.encodePacked(prefix, proof));

        signer = ecrecover(prefixedHash, v, r, s);

        if (signer != i_channelSender && signer != i_channelRecipient) {
            revert("invalid signer or hash");
        }

        if (s_signatures[proof] == address(0x0)) {
            s_signatures[proof] = signer;
            return;
        }

        if (s_signatures[proof] == signer) {
            return;
        }

        if (!i_channelRecipient.send(value)) {
            revert("payment error");
        }

        destruct();
    }

    function channelTimeout() public {
        require(i_startDate + i_channelTimeout <= block.timestamp);
        destruct();
    }

    function getChannelSender() public view returns (address) {
        return i_channelSender;
    }

    function getChannelRecipient() public view returns (address) {
        return i_channelRecipient;
    }

    function getStartDate() public view returns (uint256) {
        return i_startDate;
    }

    function getChannelTimeout() public view returns (uint256) {
        return i_channelTimeout;
    }

    function getSignatures(bytes32 proof) public view returns (address) {
        return s_signatures[proof];
    }

    function destruct() private {}
}
