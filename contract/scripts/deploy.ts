import { ethers } from "hardhat";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";

async function main() {
  const contractFactory = await ethers.getContractFactory("SmartDRM");

  console.log("deploying contract");

  const contract = await contractFactory.deploy();
  await contract.deployed();

  const signers = await ethers.getSigners();
  const account = signers[0];

  const connected = contract.connect(account);
  const transactionResponse = await connected.registerCreator();
  await transactionResponse.wait(1);

  console.log(`Deployed, contract address: ${contract.address}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
