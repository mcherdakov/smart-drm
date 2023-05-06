<script>
    import { Router, Route, Link } from "svelte-navigator";
    import Browse from "./Browse.svelte";
    import Connect from "./Connect.svelte";
    import Pay from "./Pay.svelte";
</script>

<main id="main">
    <div id="page">
        <Router>
            <div id="top">
                <nav id="nav">
                    <Link to="/">Browse</Link>
                    <Link to="connect">Connect</Link>
                    <Link to="pay">Pay</Link>
                </nav>
            </div>
            <div>
                <Route path="/">
                    <Browse />
                </Route>

                <Route path="connect">
                    <Connect />
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
