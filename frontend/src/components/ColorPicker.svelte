<script lang="ts">
    export let r: number = 100;
    export let g: number = 100;
    export let b: number = 100;
    export let w: number = 100;
</script>

<div class="container">
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

<style>
    div.container {
        width: 100%;
        height: 100%;
        padding: 4px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-evenly;
    }

    input[type="range"] {
        background-color: transparent;
        color: var(--theme-primary);
    }

    input[type="range"]::-moz-range-progress {
      background-color: var(--theme-primary); 
    }

    input[type="range"]::-moz-range-track {  
      background-color: var(--theme-secondary);
    }

    div.input {
        display: flex;
    }

    div.input input {
        font-size: 1.25rem;
        margin-top: 0px;
    }

    div.input label {
        display: inline-block;
        width: 1.5rem;
        margin-left: 4px;
        text-align: center;
        font-size: 1.25rem;
        font-weight: bolder;
    }

    div.input code {
        margin-top: 0.15rem;
        font-size: 0.95rem;
        width: 4ch;
    }
</style>
