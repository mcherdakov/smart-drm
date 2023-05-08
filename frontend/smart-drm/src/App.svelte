<script>
    import { Router, Route, Link } from "svelte-navigator";
    import Browse from "./Browse.svelte";
    import Connect from "./Connect.svelte";
    import Pay from "./Pay.svelte";
    import Error from "./Error.svelte";
    import { isConnected } from "./stores.js";
</script>

<main id="main">
    <div id="page">
        <Router>
            <div id="top">
                <nav id="nav">
                    {#if $isConnected}
                        <Link to="browse">Browse</Link>
                        <Link to="pay">Pay</Link>
                    {:else}
                        <Link to="connect">Connect</Link>
                    {/if}
                </nav>
            </div>
            <Error />
            <div>
                <Route path="/" />

                <Route path="connect">
                    <Connect />
                </Route>

                <Route path="browse">
                    <Browse />
                </Route>

                <Route path="pay">
                    <Pay />
                </Route>
            </div>
        </Router>
    </div>
</main>

<style>
    #main {
        display: flex;
        justify-content: center;
    }

    #page {
        display: flex;
        width: 50%;
        flex-direction: column;
    }

    #top {
        display: flex;
        justify-content: center;
    }

    #nav {
        display: flex;
        width: 50%;
        justify-content: space-evenly;
    }
</style>
