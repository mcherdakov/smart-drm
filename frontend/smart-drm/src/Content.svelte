<script>
    import { navigate } from "svelte-navigator";
    import { onMount } from "svelte";
    import { signer, error } from "./stores.js";
    import { config } from "./config.js";

    export let id;
    let content;

    onMount(async () => {
        const res = await fetch(
            `${
                config.url
            }/content/detail?id=${id}&address=${await $signer.getAddress()}`
        );
        const data = await res.json();

        if (data.error !== undefined) {
            error.set(data.error);

            navigate("/browse");
            return;
        }

        content = data;
    });
</script>

<div class="page">
    <div class="content">
        {#if content}
            <h1>{content.header}</h1>
            <p>{content.content}</p>
        {/if}
    </div>
</div>

<style>
    .page {
        display: flex;
        justify-content: center;
    }

    .content {
        display: flex;
        flex-direction: column;
    }

    h1 {
        text-align: center;
    }
</style>
