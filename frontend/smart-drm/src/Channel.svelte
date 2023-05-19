<script>
    import { ethers } from "ethers";
    import { signer, error } from "./stores.js";

    export let channel;
    let isProcessing = false;
    const cost = 100;

    async function savePayment(split, date, value, hash) {
        const req = {
            v: split.v,
            r: split.r,
            s: split.s,
            date: date,
            value: value,
            hash: hash,
        };

        const res = await fetch("http://127.0.0.1:8000/pay", {
            method: "POST",
            body: JSON.stringify(req),
        });

        const data = await res.json();
        if (data.error !== undefined) {
            error.set(data.error);
        }

        return data;
    }

    function currentDate() {
        return new Date().toJSON().slice(0, 10);
    }

    async function pay() {
        isProcessing = true;

        const value = parseInt(channel.offChainValue) + cost;
        const date = currentDate();

        const payload = ethers.utils.defaultAbiCoder.encode(
            ["address", "uint256", "string"],
            [channel.address, value, date]
        );

        const hash = ethers.utils.keccak256(payload);
        const signed = await $signer.signMessage(ethers.utils.arrayify(hash));
        const split = ethers.utils.splitSignature(signed);

        await savePayment(split, date, value, hash);

        isProcessing = false;
    }
</script>

<div class="chan">
    <div class="chan-info">
        <p>
            <b>Channel contract address:</b>
            {channel.address}<br />
            <b>Paid through channel (on-chain):</b>
            {channel.value}<br />
            <b>Paid through channel (off-chain):</b>
            {channel.offChainValue}<br />
            <b>Last paid date (on-chain):</b>
            {#if channel.date}
                {channel.date}
            {:else}
                no payments yet
            {/if}
            <br />
            <b>Last paid date (off-chain):</b>
            {#if channel.offChainDate}
                {channel.offChainDate}
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
            {#if isProcessing}
                Processing
            {:else if channel.offChainDate === currentDate()}
                Additional donation<br /><br />
                Doesn't affect content availability
            {:else}
                Pay for today
            {/if}
        </button>
    </div>
</div>

<style>
    .chan {
        display: flex;
    }

    .chan-info {
        display: flex;
        width: 80%;
    }

    .chan-button {
        display: flex;
        justify-content: center;
        width: 20%;
    }
</style>
