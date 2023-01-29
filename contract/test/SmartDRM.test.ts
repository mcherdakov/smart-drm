import { assert, expect } from "chai";
import { ethers } from "hardhat";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { SmartDRM, SmartDRM__factory } from "../typechain-types";

describe("SmartDRM", function () {
  let smartDRM: SmartDRM;
  let factory: SmartDRM__factory;
  let owner: SignerWithAddress;
  let creator: SignerWithAddress;

  beforeEach(async () => {
    const signers = await ethers.getSigners();

    owner = signers[0];
    creator = signers[1];

    factory = await ethers.getContractFactory("SmartDRM");
    smartDRM = await factory.deploy();
  });

  describe("Constructor", function () {
    it("Empty constructor for now", async function () {
      const contract = smartDRM.attach(owner.address);
      await expect(contract.getCreator(0)).to.be.reverted;
    });
  });

  describe("Register creator", function () {
    it("adds creator to creators array", async function () {
      const contract = smartDRM.connect(creator);
      await contract.registerCreator();

      const creatorFromArray = await contract.getCreator(0);
      assert.equal(creatorFromArray, creator.address);
    });
  });
});
