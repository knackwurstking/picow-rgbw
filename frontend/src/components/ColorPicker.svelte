<script lang="ts">
    import LockOn from "svelte-material-icons/Lock.svelte";
    import LockOff from "svelte-material-icons/LockOff.svelte";

    export let r = 100;
    export let g = 100;
    export let b = 100;
    export let w = 100;

    let chained = false;
</script>

<div class="container">
    <div class="inputs">
        <div class="input">
            <input
                name="r"
                type="range"
                min={0}
                max={100}
                value={r}
                on:input={(ev) => {
                    const newValue = parseInt(ev.target.value, 10)
                    if (chained) {
                        const diff = r - newValue
                        if (
                            (g < 0 || g > 100) &&
                            (b < 0 || b > 100) &&
                            (w < 0 || w > 100)
                        ) {
                            return
                        }
                        // TODO: need to handler min and max (0 - 100)
                        r -= diff
                        g -= diff
                        b -= diff
                        w -= diff
                    } else {
                        r = newValue
                    }
                }}
            />
            <label for="r">R</label>
            <code>{r}</code>
        </div>
        <div class="input">
            <input
                name="g"
                type="range"
                min={0}
                max={100}
                value={g}
                on:change={(ev) => {
                    const newValue = parseInt(ev.target.value, 10)
                    if (chained) {
                        const diff = g - newValue
                        if (
                            (r < 0 || r > 100) &&
                            (b < 0 || b > 100) &&
                            (w < 0 || w > 100)
                        ) {
                            return
                        }
                        r -= diff
                        g -= diff
                        b -= diff
                        w -= diff
                    } else {
                        g = newValue
                    }
                }}
            />
            <label for="g">G</label>
            <code>{g}</code>
        </div>
        <div class="input">
            <input
                name="b"
                type="range"
                min={0}
                max={100}
                bind:value={b}
            />
            <label for="b">B</label>
            <code>{b}</code>
        </div>
        <div class="input">
            <input
                name="w"
                type="range"
                min={0}
                max={100}
                bind:value={w}
            />
            <label for="w">W</label>
            <code>{w}</code>
        </div>
    </div>
    <div class="chain">
        <button on:click={() => (chained = !chained)}>
            {#if chained}
                <LockOn
                    size={28}
                />
            {:else}
                <LockOff
                    size={28}
                />
            {/if}
        </button>
    </div>
</div>

<style>
    button {
        background-color: transparent;
        color: white;
        border: none;
    }

    input[type="range"] {
        background-color: transparent;
        width: 100%;
        margin: 8px 0;
    }

    div.container {
        padding: 8px;
        width: 100%;
        display: flex;
    }

    div.container .inputs {
        width: 100%;
    }

    div.container .chain {
        width: 40px;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    div.container .chain button {
        height: 100%;
    }

    div.input {
        display: flex;
        padding: 4px;
    }

    div.input input {
        font-size: 1rem;
    }

    div.input label {
        display: inline-block;
        width: 1.5rem;
        margin-top: 0.3rem;
        margin-left: 4px;
        text-align: center;
        font-size: 1.25rem;
        font-weight: bolder;
    }

    div.input code {
        margin-top: 0.5rem;
        font-size: 0.95rem;
        width: 4ch;
    }

    /* Firefox */
    input[type="range"]::-moz-range-thumb {
        background-color: var(--theme-primary);
        border-color: var(--theme-border);
    }

    input[type="range"]::-moz-range-track {
        background-color: var(--theme-secondary);
    }

    input[type="range"]::-moz-range-progress {
        background-color: var(--theme-primary);
    }

    /* TODO: Chrome styles */
    input[type="range"]::-webkit-slider-thumb {}

    input[type="range"]::-webkit-slider-runnable-track {}

    /* NOTE: I do not care about microsoft

    input[type="range"]::-ms-thumb {}

    input[type="range"]::-ms-track {}

    */
</style>
