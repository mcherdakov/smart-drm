<script>
    import { Router, Route, Link } from "svelte-navigator";
    import Home from "./Home.svelte";
    import Connect from "./Connect.svelte";
    import Topup from "./Topup.svelte";
</script>

<main id="main">
    <div id="page">
        <Router>
            <div id="top">
                <nav id="nav">
                    <Link to="/">Home</Link>
                    <Link to="connect">Connect</Link>
                    <Link to="topup">Top Up</Link>
                </nav>
            </div>
            <div>
                <Route path="/">
                    <Home />
                </Route>

                <Route path="connect">
                    <Connect />
                </Route>

                <Route path="topup">
                    <Topup />
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
