<script lang="ts">
    import { onMount } from "svelte";
    import { Router, Route, navigate } from "svelte-routing";
    import Devices from "./routes/Devices.svelte";

    export let url = "";

    onMount(() => {
        // TODO: do a sign in here and navigate to /devices or /login
        navigate("/devices", { replace: true }); // NOTE: there is no login yet
    });
</script>

<svelte:head>
    <!-- SMUI Light Theme -->
    <link rel="stylesheet"
        href="/theme.css"
    />
    <title>Pico Web | Home</title>
</svelte:head>

<Router url="{url}">
    <div class="content">
        <!-- TODO: adding route for "login" -->
        <Route path="/devices"><Devices /></Route>
        <!-- Route path="/"><Login /></Route -->
    </div>
</Router>

<style>
    :global(body) {
        overflow: hidden;
    }

    .content {
        width: 100vw;
        height: 100vh;
        overflow: hidden;
    }
</style>
