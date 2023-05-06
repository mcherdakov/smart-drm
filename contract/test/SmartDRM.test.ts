import { assert } from "chai";
import { ethers } from "hardhat";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { SmartDRM, SmartDRM__factory } from "../typechain-types";

describe("SmartDRM", function () {
  let smartDRM: SmartDRM;
  let factory: SmartDRM__factory;
  let owner: SignerWithAddress;
  let creator: SignerWithAddress;
  let user: SignerWithAddress;

  beforeEach(async () => {
    const signers = await ethers.getSigners();

    owner = signers[0];
    creator = signers[1];
    user = signers[2];

    factory = await ethers.getContractFactory("SmartDRM");
    smartDRM = await factory.deploy();
  });

  describe("create channel", function () {
    it("adds channel to mapping", async function () {
      const contract = smartDRM.connect(user);

      await contract.createChannel(100000);
      const userChannel = await contract.getUserChannel(user.address);
      assert.notEqual(
        ethers.BigNumber.from(userChannel).toString(),
        ethers.BigNumber.from("0x0").toString()
      );

      const creatorChannel = await contract.getUserChannel(creator.address);
      assert.equal(
        ethers.BigNumber.from(creatorChannel).toString(),
        ethers.BigNumber.from("0x0").toString()
      );
    });
  });
});
