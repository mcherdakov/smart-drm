<script>
    import { ethers } from "ethers";
    import { onDestroy, onMount } from "svelte";
    import { navigate } from "svelte-navigator";
    import { isConnected, provider, signer, error } from "./stores.js";
    import Channel from "./Channel.svelte";
    import { config } from "./config.js";

    let amount = "1000";
    let timeout = 30;

    let isProcessing = false;
    let chanExists = false;

    $: disabled = isProcessing || chanExists;

    let contract, channel;

    intervalFetch(() => fetchChannel(), 1000);

    onMount(async () => {
        if (!$isConnected) {
            navigate("/connect");
            return;
        }

        isProcessing = true;

        const res = await fetch(`${config.url}/drm`);
        const data = await res.json();

        if (data.error !== undefined) {
            error.set(data.error);
            return;
        }

        contract = new ethers.Contract(data.address, data.abi, $signer);

        await fetchChannel();

        isProcessing = false;
    });

    function intervalFetch(callback, milliseconds) {
        const interval = setInterval(callback, milliseconds);

        onDestroy(() => {
            clearInterval(interval);
        });
    }

    async function fetchChannel() {
        if (contract === undefined) {
            return;
        }

        const chanAddr = await contract.getUserChannel($signer.getAddress());
        if (
            ethers.BigNumber.from(chanAddr).toString() ==
            ethers.BigNumber.from("0x0").toString()
        ) {
            return;
        }

        const chanInfo = await contract.getChannelProof(chanAddr);
        const chanBalance = await $provider.getBalance(chanAddr);

        const res = await fetch(
            `http://127.0.0.1:8000/proof?address=${chanAddr}`
        );
        const data = await res.json();

        if (data.error !== undefined) {
            error.set(data.error);
        }

        channel = {
            ...chanInfo,
            address: chanAddr,
            balance: chanBalance,
            offChainValue: data.value,
            offChainDate: data.date,
        };

        chanExists = true;
    }

    async function createContract() {
        isProcessing = true;

        const tx = await contract.createChannel(timeout * 86400, {
            value: ethers.utils.parseUnits(amount, "wei"),
        });

        await tx.wait(1);

        await fetchChannel();

        isProcessing = false;
    }

    async function close() {
        const res = await fetch(`${config.url}/close`, {
            method: "POST",
            body: JSON.stringify({ address: await $signer.getAddress() }),
        });

        const data = await res.json();
        if (data.error !== undefined) {
            error.set(data.error);
            return;
        }

        contract = undefined;
    }
</script>

<div>
    <h1>Pay</h1>
    <div class="page">
        <div class="channel">
            <form on:submit|preventDefault={createContract}>
                <label>
                    <input bind:value={amount} {disabled} />
                    wei
                </label>
                <label>
                    <input bind:value={timeout} {disabled} />
                    days
                </label>
                <button {disabled}>
                    {#if isProcessing}
                        Processing
                    {:else if chanExists}
                        Channel created
                    {:else}
                        Create channel
                    {/if}
                </button>
            </form>

            {#if chanExists}
                <Channel {channel} />
            {:else}
                <p>Created channel will be displayed here</p>
            {/if}
        </div>
        {#if chanExists}
            <div>
                <button on:click={close}> Close channel </button>
            </div>
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
