<script>
    import { ethers } from "ethers";
    import { signer } from "./stores.js";

    export let channel;
    let isProcessing = false;

    async function pay() {
        isProcessing = true;

        const value = parseInt(channel.value) + 100;
        console.log(value);

        const date = new Date().toJSON().slice(0, 10);
        console.log(date);

        const payload = ethers.utils.defaultAbiCoder.encode(
            ["address", "uint256", "string"],
            [channel.address, value, date]
        );

        const payloadHash = ethers.utils.keccak256(payload);

        const signedSender = await $signer.signMessage(
            ethers.utils.arrayify(payloadHash)
        );

        const splitSender = ethers.utils.splitSignature(signedSender);

        console.log(splitSender);

        isProcessing = false;
    }
</script>

<div class="chan">
    <div class="chan-info">
        <p>
            <b>Channel contract address:</b>
            {channel.address}<br />
            <b>Paid through channel:</b>
            {channel.value}<br />
            <b>Last paid date:</b>
            {#if channel.date}
                {channel.date}
            {:else}
                no payments yet
            {/if}
            <br />
            <b>Channel balance:</b>
            {channel.balance} wei
        </p>
    </div>
    <div class="chan-button">
        <button disabled={isProcessing} on:click={pay}>
            {#if !isProcessing}
                Pay for today
            {:else}
                Processing
            {/if}
        </button>
    </div>
</div>

<style>
    .chan {
        display: flex;
    }

    .chan-info {
        width: 80%;
    }

    .chan-button {
        display: flex;
        justify-content: center;
    }
</style>
