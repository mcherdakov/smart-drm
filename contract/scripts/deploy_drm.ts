import { ethers } from "hardhat";

async function main() {
  const contractFactory = await ethers.getContractFactory("SmartDRM");

  console.log("deploying drm contract");

  const signers = await ethers.getSigners();

  console.log(`Contract creator: ${signers[0].address}`);

  contractFactory.connect(signers[0]);

  const contract = await contractFactory.deploy();

  await contract.deployed();

  console.log(`Deployed drm contract, address: ${contract.address}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
