/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  publicRuntimeConfig: {
    contractAddress: "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9",
    abi: [
      {
        inputs: [
          { internalType: "address payable", name: "to", type: "address" },
          { internalType: "uint256", name: "timeout", type: "uint256" },
        ],
        stateMutability: "payable",
        type: "constructor",
      },
      {
        inputs: [],
        name: "channelTimeout",
        outputs: [],
        stateMutability: "nonpayable",
        type: "function",
      },
      {
        inputs: [
          { internalType: "bytes32", name: "h", type: "bytes32" },
          { internalType: "uint8", name: "v", type: "uint8" },
          { internalType: "bytes32", name: "r", type: "bytes32" },
          { internalType: "bytes32", name: "s", type: "bytes32" },
          { internalType: "uint256", name: "value", type: "uint256" },
        ],
        name: "closeChannel",
        outputs: [],
        stateMutability: "nonpayable",
        type: "function",
      },
      {
        inputs: [],
        name: "getChannelRecipient",
        outputs: [{ internalType: "address", name: "", type: "address" }],
        stateMutability: "view",
        type: "function",
      },
      {
        inputs: [],
        name: "getChannelSender",
        outputs: [{ internalType: "address", name: "", type: "address" }],
        stateMutability: "view",
        type: "function",
      },
      {
        inputs: [],
        name: "getChannelTimeout",
        outputs: [{ internalType: "uint256", name: "", type: "uint256" }],
        stateMutability: "view",
        type: "function",
      },
      {
        inputs: [],
        name: "getStartDate",
        outputs: [{ internalType: "uint256", name: "", type: "uint256" }],
        stateMutability: "view",
        type: "function",
      },
    ],
  },
};

module.exports = nextConfig;
