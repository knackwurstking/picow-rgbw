<script lang="ts"<>
</script>

<div class="container">
    <button class="off">OFF</button>
    <button class="set">SET</button>
</div>

<style>
    div.container {
        width: 100%;
        height: 100%;
        display: flex;
        justify-content: space-evenly;
        align-items: center;
    }

    div.container button {
        height: 90%;
        width: 40%;
        max-width: 100px;
        background-color: var(--theme-surface);
        font-size: 1.5rem;
        font-weight: bolder;
        border-color: var(--theme-border);
        color: var(--theme-on-surface);
        transition: transform 0.1s ease;
    }

    div.container button.off {
        text-shadow: 0 0 0.25rem red;
    }

    div.container button.set{
        text-shadow: 0 0 0.35rem green;
    }

    div.container button:active {
        transition: transform 0.25s ease;
        transform: scale(0.5);
    }
</style>
