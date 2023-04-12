<script lang="ts">
    export let min: number = 0;
    export let max: number = 100;
    export let value: number = 100;
</script>

<div class="container" {...$$restProps}>
    <input type="range" {min} {max} bind:value orient="vertical" />
</div>

<style>
    .container {
        height: 100%;
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    input {
        height: calc(100% - 16px);
        margin: 8px 0;
    }
</style>
