<script lang="ts">
    export let r = 100;
    export let g = 100;
    export let b = 100;
    export let w = 100;
</script>

<div class="container">
    <div class="inputs">
        <div class="input">
            <input
                name="r"
                type="range"
                min={0}
                max={100}
                bind:value={r}
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
                bind:value={g}
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
</div>

<style>
    input[type="range"] {
        background-color: transparent;
        width: 100%;
        margin: 8px 0;
        cursor: pointer;
    }

    div.container {
        padding: 8px;
        width: 100%;
        display: flex;
    }

    div.container .inputs {
        width: 100%;
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
