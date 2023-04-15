<script lang="ts">
    export let r = 100;
    export let g = 100;
    export let b = 100;
    export let w = 100;
</script>

<div class="container" {...$$restProps}>
    <div class="inputs">
        <div class="input">
            <input name="r" type="range" min={0} max={100} bind:value={r} />
            <label for="r">R</label>
            <code>{r}</code>
        </div>
        <div class="input">
            <input name="g" type="range" min={0} max={100} bind:value={g} />
            <label for="g">G</label>
            <code>{g}</code>
        </div>
        <div class="input">
            <input name="b" type="range" min={0} max={100} bind:value={b} />
            <label for="b">B</label>
            <code>{b}</code>
        </div>
        <div class="input">
            <input name="w" type="range" min={0} max={100} bind:value={w} />
            <label for="w">W</label>
            <code>{w}</code>
        </div>
    </div>
</div>

<style>
    div.container {
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
        width: 100%;
        margin: 8px 0;
    }

    div.input label {
        display: inline-block;
        width: 1.5rem;
        margin-top: 4px;
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
</style>
