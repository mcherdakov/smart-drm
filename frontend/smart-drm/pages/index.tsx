import { ethers } from "ethers";
import getConfig from "next/config";
import { useState } from "react";

export default function Home() {
  const { publicRuntimeConfig } = getConfig();
  const [buttonText, setButtonText] = useState("Connect");

  async function fetchBytecode() {
    const res = await fetch("http://127.0.0.1:8080/bytecode");
    const data = await res.json();
    return data.bytecode;
  }

  async function handleConnect() {
    if (typeof window.ethereum === "undefined") {
      alert("please install metamask");
      return;
    }

    try {
      await window.ethereum.request({ method: "eth_requestAccounts" });
    } catch (error) {
      console.log(error);
    }

    setButtonText("Connected");
  }

  async function handleCreate() {
    const provider = new ethers.providers.Web3Provider(window.ethereum);
    await provider.send("eth_requestAccounts", []);
    const signer = provider.getSigner();
    const bytecode = await fetchBytecode();
    const factory = new ethers.ContractFactory(
      publicRuntimeConfig.abi,
      bytecode,
      signer
    );
    factory.deploy(signer.getAddress(), 1000000, {
      value: 10000,
    });
  }

  return (
    <main>
      <button onClick={handleConnect}>{buttonText}</button>
      <button onClick={handleCreate}>Create</button>
    </main>
  );
}
