import { assert } from "chai";
import { ethers } from "hardhat";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { Channel, Channel__factory } from "../typechain-types";

describe("Channel", function () {
  const initialBalance = ethers.utils.parseEther("10");

  let channel: Channel;
  let factory: Channel__factory;
  let sender: SignerWithAddress;
  let receiver: SignerWithAddress;

  beforeEach(async () => {
    const signers = await ethers.getSigners();

    sender = signers[0];
    receiver = signers[1];

    factory = await ethers.getContractFactory("Channel");
    channel = await factory.deploy(receiver.getAddress(), 1000000, {
      value: initialBalance,
    });
  });

  describe("fund and recieve", function () {
    it("one fund", async function () {
      const balanceBefore = await receiver.getBalance();

      const contract = channel.connect(receiver);
      const contractSender = await contract.getChannelSender();
      assert.equal(contractSender, sender.address);

      const value = 125;
      const date = "01-01-2023";

      const payload = ethers.utils.defaultAbiCoder.encode(
        ["address", "uint256", "string"],
        [contract.address, value, date]
      );

      const payloadHash = ethers.utils.keccak256(payload);

      const signedSender = await sender.signMessage(
        ethers.utils.arrayify(payloadHash)
      );

      const splitSender = ethers.utils.splitSignature(signedSender);

      const verifyAddr = ethers.utils.verifyMessage(
        ethers.utils.arrayify(payloadHash),
        splitSender
      );
      assert.equal(verifyAddr, sender.address);

      await channel.closeChannel(
        splitSender.v,
        splitSender.r,
        splitSender.s,
        value,
        date
      );

      const sigAddr = await channel.getSignatures(payloadHash);
      assert.equal(sigAddr, sender.address);

      const signedReciever = await receiver.signMessage(
        ethers.utils.arrayify(payloadHash)
      );

      const splitReciever = ethers.utils.splitSignature(signedReciever);
      await channel.closeChannel(
        splitReciever.v,
        splitReciever.r,
        splitReciever.s,
        value,
        date
      );

      const balanceAfter = await receiver.getBalance();
      assert.equal(
        balanceAfter.sub(balanceBefore).toString(),
        value.toString()
      );
    });
  });
});
