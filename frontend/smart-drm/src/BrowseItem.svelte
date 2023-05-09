<script>
    import { Link } from "svelte-navigator";

    export let item;
</script>

<div class="item">
    <div class="info">
        <p>
            <Link to={`/content/${item.id}`}>{item.header}</Link>
        </p>
        <p>by {item.author}</p>
    </div>
    <div class="meta">
        {#if item.click_exists}
            👁️
        {:else}
            💸
        {/if}
    </div>
</div>

<style>
    .item {
        border: solid;
        margin: 5px;
        display: flex;
        justify-content: space-between;
    }

    .info {
        width: 90%;
    }

    .meta {
        text-align: center;
        width: 10%;
    }

    p {
        margin: 3px;
    }
</style>
