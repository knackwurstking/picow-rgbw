<script lang="ts">
    import { onMount } from "svelte";

    onMount(() => {});
</script>

<svelte:head>
    <!-- SMUI Light Theme -->
    <link rel="stylesheet" href="/theme.css" />
    <title>Pico Web | Home</title>
</svelte:head>

<main>
    <div class="devices-container debug" />
    <div class="spacer debug" />
    <div class="control-container debug" />
</main>

<style>
    main {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;

        overflow: hidden;

        display: flex;
        flex-direction: row;
    }

    main > div:not(.spacer) {
        width: calc(100% - 16px);
        height: 100%;
    }

    main > div.spacer {
        width: 16px;
        height: 100%;
    }

    main > div:not(.spacer):last-child {
        margin-right: 8px;
    }

    main > div:not(.spacer):first-child {
        margin-left: 8px;
    }
</style>
