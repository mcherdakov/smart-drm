<script>
    import { error } from "./stores.js";

    function resetError() {
        error.set("");
    }
</script>

{#if $error}
    <div class="error">
        <div class="value">{$error}</div>
        <button class="ok" on:click={resetError}>ok</button>
    </div>
{/if}

<style>
    .error {
        display: flex;
        justify-content: center;
        padding-top: 1em;
    }

    .value {
        color: red;
        padding-right: 1em;
    }

    .ok {
        color: black;
    }
</style>
