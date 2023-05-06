<script>
    import { navigate } from "svelte-navigator";
    import { isConnected } from "./stores.js";
    import { onMount } from "svelte";

    onMount(() => {
        if (!$isConnected) {
            navigate("/connect");
        }
    });
</script>

<div>
    <h1>Browse</h1>
</div>

<style>
    h1 {
        display: flex;
        justify-content: center;
    }
</style>
