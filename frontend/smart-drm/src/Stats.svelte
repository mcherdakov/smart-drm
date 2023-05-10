<script>
    import { error } from "./stores.js";
    import { onMount } from "svelte";

    let contentStats = [];
    let authorStats = [];

    onMount(async () => {
        const res = await fetch(`http://127.0.0.1:8000/stats`);

        const data = await res.json();
        if (data.error !== undefined) {
            error.set(data.error);
            return;
        }

        contentStats = data.content;
        authorStats = data.author;
    });
</script>

<div class="page">
    <h1>Stats</h1>
    <h3>Author Stats</h3>
    {#each authorStats as stat (stat.author)}
        <div class="stat-item">
            <p><b>{stat.author}</b></p>
            <p>Clicks commited: {stat.clicks_commited}</p>
            <p>Clicks uncommited: {stat.clicks_uncommited}</p>
        </div>
    {/each}
    <h3>Content Stats</h3>
    {#each contentStats as stat (stat.id)}
        <div class="stat-item">
            <p><b>{stat.header}</b></p>
            <p>Author: {stat.author}</p>
            <p>Clicks commited: {stat.clicks_commited}</p>
            <p>Clicks uncommited: {stat.clicks_uncommited}</p>
        </div>
    {/each}
</div>

<style>
    h1 {
        display: flex;
        justify-content: center;
    }

    .page {
        display: flex;
        flex-direction: column;
    }

    .stat-item {
        border: solid;
        margin: 5px;
        display: flex;
        flex-direction: column;
    }

    .stat-item p {
        margin: 2px 5px;
    }
</style>
