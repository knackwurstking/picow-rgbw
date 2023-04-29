<script lang="ts">
    export let active: boolean = false;
</script>

<div class="outer" {...$$restProps}>
    <div
        class="inner"
        style={`
            background-color: ${active ? "green" : "red"};
        `}
    />
</div>

<style>
    .outer {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 1.5em;
        height: 1.5em;
        border-radius: var(--theme-border-radius);
    }

    .outer .inner {
        width: 1em;
        height: 1em;
        filter: blur(2px);
        border-radius: 50%;
    }
</style>
