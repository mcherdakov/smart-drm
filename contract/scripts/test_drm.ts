import { ethers } from "hardhat";

async function main() {
  const factory = await ethers.getContractFactory("Channel");
  const channel = factory.attach("0xcafac3dd18ac6c6e92c921884f9e4176737c052c");
  console.log(await channel.getChannelSender());
  console.log(await channel.getChannelRecipient());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
