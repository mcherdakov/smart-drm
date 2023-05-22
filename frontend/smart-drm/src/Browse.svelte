<script>
    import { navigate } from "svelte-navigator";
    import { signer, isConnected, error } from "./stores.js";
    import { onMount } from "svelte";
    import BrowseItem from "./BrowseItem.svelte";
    import { config } from "./config.js";

    let items = [];

    onMount(async () => {
        if (!$isConnected) {
            navigate("/connect");
            return;
        }

        const res = await fetch(
            `${config.url}/content?address=${await $signer.getAddress()}`
        );

        const data = await res.json();
        if (data.error !== undefined) {
            error.set(data.error);
        }

        items = data;
    });
</script>

<div>
    <h1>Browse</h1>
    {#each items as item (item.id)}
        <BrowseItem {item} />
    {/each}
</div>

<style>
    h1 {
        display: flex;
        justify-content: center;
    }
</style>
