<script>
    import { ethers } from "ethers";
    import { isConnected, provider, signer, error } from "./stores.js";

    async function connect() {
        let eth = window["ethereum"];

        if (eth === undefined) {
            error.set("please install metamask and reload this page");
            return;
        }

        provider.set(new ethers.providers.Web3Provider(window["ethereum"]));

        try {
            await $provider.send("eth_requestAccounts", []);
            signer.set($provider.getSigner());

            isConnected.set(true);
        } catch (err) {
            error.set(err);
            return;
        }
    }
</script>

<div>
    <h1>Connect</h1>
    <div class="connect">
        <div class="connect-button">
            <button on:click={connect} disabled={$isConnected}>
                {#if !$isConnected}
                    Connect to Metamask
                {:else}
                    Connected
                {/if}
            </button>
        </div>
    </div>
</div>

<style>
    h1 {
        display: flex;
        justify-content: center;
    }

    .connect {
        display: flex;
        justify-content: center;
    }

    .connect-button {
        display: flex;
        width: 50%;
        flex-direction: column;
    }

    .error {
        display: flex;
        justify-content: center;
        color: red;
    }
</style>
