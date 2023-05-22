<script lang="ts">
  export let items: Color[] = [];

  // TODO: load stored items from the `localStorage`
</script>

<div class="container">
  <div class="actions" />
  <div class="content">
    <ul class="custom-list" style="height: 100%;">
      {#each items as item}
        <li class="custom-list-item" style="width: 50px;">
          <div class="color-container__outer">
            <div class="color-container__inner" />
          </div>
        </li>
      {/each}
    </ul>
  </div>
</div>

<style>
  .container {
    position: relative;
    width: 100%;
    height: fit-content;
  }

  .container .actions {
    display: absolute;
    top: 0;
    right: 0;
    height: calc(1rem + 8px);

    border: 1px solid green;
  }

  .container .content {
    height: 65px;
    width: 100%;
    overflow-x: auto;
    overflow-y: hidden;
  }

  .container .content .custom-list {
    display: flex;
    align-items: flex-start;
    overflow-x: auto;

    border: 1px solid red;
  }

  .container .content .custom-list .custom-list-item {
    float: left;
    min-width: 50px;
    width: 150px;
    max-width: 50px;
    height: calc(100% - 8px);
    margin: 4px;
  }

  .container .content .color-container__outer {
    position: relative;

    width: 100%;
    height: 100%;

    border: 1px solid var(--theme-border-color);
    border-radius: var(--theme-border-radius);
  }

  .container .content .color-container__inner {
    position: absolute;
    left: 4px;
    right: 4px;
    top: 4px;
    bottom: 4px;

    border: 1px solid var(--theme-border-color);
    border-radius: var(--theme-border-radius);
  }
</style>
