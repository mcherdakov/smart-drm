import { ethers } from "hardhat";

async function main() {
  const contractFactory = await ethers.getContractFactory("Channel");

  console.log("deploying channel contract");

  const signers = await ethers.getSigners();
  const account = signers[0];

  const contract = await contractFactory.deploy(account.address, 1000000, {
    value: 10000,
  });
  await contract.deployed();

  console.log(`Deployed, channel contract address: ${contract.address}`);
  console.log(`Reciever address: ${account.address}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
