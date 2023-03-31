<script lang="ts">
    import devices from "./devices";
</script>

<svelte:head>
    <title>Pico Web | Devices</title>
</svelte:head>

<!-- TODO: ...
    ...split view (landscape or horizontal orientation)
    ...checkbox like group of devices for selecting device(s) to control
    ...button group for on/off
    ...input group (range slider) for rgbw (0-100)
    ...color preview somewhere
    ...color history (stored on client side)
-->

<div class="devices container">
    <section class="list">
        <fieldset>
            <legend>Devices</legend>

            <div class="content list">
                <!-- TODO: do some custom checkbox style here (create component) -->
                <label><input type="checkbox" />Device 1</label>
                <label><input type="checkbox" />Device 2</label>
                <label><input type="checkbox" />Device 3</label>
                <label><input type="checkbox" />Device 4</label>
                <label><input type="checkbox" />Device 5</label>
                <label><input type="checkbox" />Device 6</label>
                <label><input type="checkbox" />Device 7</label>
                <label><input type="checkbox" />Device 8</label>
                <label><input type="checkbox" />Device 9</label>
                <label><input type="checkbox" />Device 10</label>
                <label><input type="checkbox" />Device 11</label>
            </div>
        </fieldset>
    </section>
    <section class="ctrl">
    </section>
</div>

<style>
    div.devices.container {
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100%;
    }

    div.devices.container > section {
        width: 100%;
        height: 100%;
        overflow: hidden;
    }

    div.devices.container > section.list {
        height: 60%;
    }

    div.devices.container > section.ctrl {
        height: 40%;
    }

    /* default: mobile (portrait) */

    @media (min-width: 1025px) {
        div.devices.container {
            flex-direction: row;
        }

        div.devices.container > section.list {
            width: 50%;
            height: 100%;
        }

        div.devices.container > section.ctrl {
            width: 50%;
            height: 100%;
        }
    }

    div.devices.container > section.list fieldset {
        margin: 16px;
        border-color: var(--theme-border);
        overflow: hidden;
        height: calc(100% - 32px);
    }

    /* TODO: remove scroll bar */
    div.devices.container > section.list > fieldset > div.content.list {
        display: flex;
        flex-direction: column;
        overflow: hidden;
        overflow-y: auto;
        scroll-behavior: smooth;
        width: 100%;
        height: 100%;
    }

    div.devices.container > section.list > fieldset > div.content.list > label {
        margin: 8px;
        padding: 16px;
        border: 1px solid var(--theme-border);
    }
</style>
