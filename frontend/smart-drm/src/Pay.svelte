<script>
    import { ethers } from "ethers";
    import { onMount } from "svelte";

    let amount = "1";
    let timeout = 1000000;

    let isProcessing = false;
    let chanExists = false;

    $: disabled = isProcessing || chanExists;

    let provider, signer, contract;
    let channel, chanAddr, chanBalance;

    onMount(async () => {
        isProcessing = true;

        const res = await fetch("http://127.0.0.1:8000/drm");
        const data = await res.json();

        provider = new ethers.providers.Web3Provider(window["ethereum"]);

        await provider.send("eth_requestAccounts", []);
        signer = provider.getSigner();

        contract = new ethers.Contract(data.address, data.abi, signer);

        await fetchChannel();

        isProcessing = false;
    });

    async function fetchChannel() {
        chanAddr = await contract.getUserChannel(signer.getAddress());
        if (
            ethers.BigNumber.from(chanAddr).toString() ==
            ethers.BigNumber.from("0x0").toString()
        ) {
            return;
        }

        channel = await contract.getChannelProof(chanAddr);
        chanBalance = await provider.getBalance(chanAddr);

        chanExists = true;
    }

    async function createContract() {
        isProcessing = true;

        const tx = await contract.createChannel(timeout, {
            value: ethers.utils.parseEther(amount),
        });

        await tx.wait(1);

        await fetchChannel();

        isProcessing = false;
    }
</script>

<div>
    <h1>Pay</h1>
    <div class="page">
        <form on:submit|preventDefault={createContract}>
            <input bind:value={amount} {disabled} />
            <input bind:value={timeout} {disabled} />
            <button {disabled}>
                {#if isProcessing}
                    Processing
                {:else if chanExists}
                    Channel created
                {:else}
                    Create contract
                {/if}
            </button>
        </form>
        {#if chanExists}
            <p>
                <b>Channel contract address:</b>
                {chanAddr}<br />
                <b>Paid through channel:</b>
                {channel.value}<br />
                <b>Channel balance:</b>
                {chanBalance} wei
            </p>
        {:else}
            <p>Created channel will be displayed here</p>
        {/if}
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
        flex-direction: column;
    }
</style>
