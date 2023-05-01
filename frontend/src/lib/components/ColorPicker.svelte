<script lang="ts">
    import Slider from "./custom/slider/Slider.svelte";

    export let r: number = 0;
    $: r >= 0 && updateBrightnessValue();

    export let g: number = 0;
    $: g >= 0 && updateBrightnessValue();

    export let b: number = 0;
    $: b >= 0 && updateBrightnessValue();

    export let w: number = 0;

    let bValue = Math.min(...[r, g, b]);

    function updateBrightnessValue() {
        const rgb: number[] = [];
        if (r > 0) rgb.push(r);
        if (g > 0) rgb.push(g);
        if (b > 0) rgb.push(b);
        bValue = Math.min(...rgb);

        if (r === g && r === b) w = r;
        else w = 0;
    }
</script>

<div
    class="container"
    {...$$restProps}
    style={`--rgb: rgba(${(r / 100) * 255}, ${(g / 100) * 255}, ${
        (b / 100) * 255
    }, 0.75);`}
>
    <div class="color">
        <div class="input">
            <Slider min={0} max={100} bind:value={r} />
            <label for="r">R</label>
            <code>{r}</code>
        </div>
        <div class="input">
            <Slider min={0} max={100} bind:value={g} />
            <label for="g">G</label>
            <code>{g}</code>
        </div>
        <div class="input">
            <Slider min={0} max={100} bind:value={b} />
            <label for="b">B</label>
            <code>{b}</code>
        </div>
    </div>

    <div class="brightness">
        <Slider
            min={5}
            max={100}
            bind:value={bValue}
            orient="vertical"
            on:input={() => {
                const rgb = [];
                if (r > 0) rgb.push(r);
                if (g > 0) rgb.push(g);
                if (b > 0) rgb.push(b);

                const currentMin = Math.min(...rgb);

                // Ignore white color if rgb is not the same
                let handleWhite = false;
                if (currentMin === Math.max(...rgb)) handleWhite = true;

                const diff = currentMin - bValue;

                if (rgb.find((c) => c - diff > 100)) {
                    const rest = 100 - Math.max(...rgb);
                    if (r > 0) r += rest;
                    if (g > 0) g += rest;
                    if (b > 0) b += rest;
                } else {
                    if (r > 0) r -= diff;
                    if (g > 0) g -= diff;
                    if (b > 0) b -= diff;
                }

                if (handleWhite) w = b;
                else w = 0;
            }}
        />
    </div>
</div>

<style>
    div.container {
        position: relative;
        width: 100%;
        height: 168px;
        display: flex;
    }

    div.container::after {
        content: "";
        position: absolute;
        z-index: -1;
        top: 25px;
        right: 25px;
        bottom: 25px;
        left: 25px;
        background: var(--rgb);
        filter: blur(50px);
    }

    div.container .color {
        width: 100%;
    }

    div.container .color div.input {
        display: flex;
        margin: 20px 8px;
    }

    div.container .color div.input label {
        display: inline-block;
        width: 1.5rem;
        margin-left: 12px;
        text-align: center;
        font-size: 1.25rem;
        font-weight: bolder;
    }

    div.container .color div.input code {
        margin-top: 0.25rem;
        font-size: 0.95rem;
        width: 4ch;
    }

    div.container .brightness {
        height: calc(100% - 20px);
        width: 48px;
        padding-top: 30px;
        padding-bottom: 6px;
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>
