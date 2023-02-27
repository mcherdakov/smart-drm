import { ethers } from "hardhat";

async function main() {
  const contractFactory = await ethers.getContractFactory("SmartDRM");

  console.log("deploying drm contract");

  const contract = await contractFactory.deploy();
  await contract.deployed();

  console.log(`Deployed drm contract, address: ${contract.address}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
