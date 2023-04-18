<script lang="ts">
    export let r: number = 100;
    $: {
        bValue = Math.min(...[r, g, b]);
        bMax = 100 - (100 - bMin + Math.max(...[r, g, b]));
    }

    export let g: number = 100;
    $: {
        bValue = Math.min(...[r, g, b]);
        bMax = 100 - (100 - bMin + Math.max(...[r, g, b]));
    }

    export let b: number = 100;
    $: {
        bValue = Math.min(...[r, g, b]);
        bMax = 100 - (100 - bMin + Math.max(...[r, g, b]));
    }

    export let w: number = 100;

    let bValue = 0;
    let bMin = 5;
    let bMax = 100 - (100 - bMin + Math.max(...[r, g, b]));

    // TODO: Update range slider styles to preview the current color
</script>

<div class="container" {...$$restProps}>
    <div class="color">
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

    <div class="brightness" {...$$restProps}>
        <input
            type="range"
            min={bMin}
            max={bMax}
            bind:value={bValue}
            orient="vertical"
            on:input={() => {
                const currentMin = Math.min(...[r, g, b]);

                let handleWhite = false;
                if (currentMin === Math.max(...[r, g, b])) handleWhite = true;

                const diff = currentMin - bValue;

                if (
                    r - diff > 100 ||
                    g - diff > 100 ||
                    b - diff > 100
                ) {
                    const rest = 100 - Math.max(...[r, g, b]);
                    r += rest;
                    g += rest;
                    b += rest;
                } else {
                    r -= diff;
                    g -= diff;
                    b -= diff;
                }

                if (handleWhite) {
                    w = b;
                }
            }}
        />
    </div>
</div>

<style>
    div.container {
        width: 100%;
        height: 168px;
        display: flex;
    }

    div.container .color {
        width: 100%;
    }

    div.container .color div.input {
        display: flex;
        padding: 4px;
    }

    div.container .color div.input input {
        width: 100%;
        margin: 8px 0;
    }

    div.container .color  div.input label {
        display: inline-block;
        width: 1.5rem;
        margin-top: 4px;
        margin-left: 4px;
        text-align: center;
        font-size: 1.25rem;
        font-weight: bolder;
    }

    div.container .color  div.input code {
        margin-top: 0.5rem;
        font-size: 0.95rem;
        width: 4ch;
    }

    div.container .brightness {
        height: 100%;
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    div.container .brightness input {
        height: calc(100% - 16px);
        margin: 8px 0;
    }
</style>
