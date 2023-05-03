<script>
    import { ethers } from "ethers";
    import { abi } from "./abi";

    let amount = 0;
    let timeout = 0;

    async function fetchContractBytecode() {
        const res = await fetch("http://127.0.0.1:8000/bytecode");
        const data = await res.json();
        return data.bytecode;
    }

    async function createContract() {
        const provider = new ethers.providers.Web3Provider(window["ethereum"]);

        await provider.send("eth_requestAccounts", []);
        const signer = provider.getSigner();

        const bytecode = await fetchContractBytecode();
        const factory = new ethers.ContractFactory(abi, bytecode, signer);
        factory.deploy(signer.getAddress(), timeout, {
            value: amount,
        });
    }
</script>

<div>
    <h1>Top Up</h1>
    <div class="page">
        <form on:submit|preventDefault={createContract}>
            <input bind:value={amount} />
            <input bind:value={timeout} />
            <button>Create contract</button>
        </form>
    </div>
</div>

<style>
    h1 {
        display: flex;
        justify-content: center;
    }

    .page {
        display: flex;
        justify-content: center;
    }
</style>
