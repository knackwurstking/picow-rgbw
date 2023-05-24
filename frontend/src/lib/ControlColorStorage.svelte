<script lang="ts">
  import { createEventDispatcher } from "svelte";

  import IconButton from "svelteui/src/icon-button";

  const dispatch = createEventDispatcher();

  export let items: Color[] = [];

  function _clickAdd() {
    dispatch("itemadd");
  }

  function _clickDelete() {
    dispatch("itemdelete");
  }
</script>

<div class="container">
  <div class="actions">
    <IconButton style="margin: 4px; margin-right: 0;" on:click={_clickAdd}
      >add</IconButton
    >
    <IconButton style="margin: 4px" on:click={_clickDelete}>delete</IconButton>
    <div class="spacer" />
  </div>
  <div class="content">
    <ul class="custom-list" style="height: 100%;">
      {#each items as item}
        <li class="custom-list-item" style="width: 50px;">
          <div class="color-container__outer">
            <div
              class="color-container__inner"
              style={`
                background-color: rgba(
                  ${(item[0] / 100) * 255},
                  ${(item[1] / 100) * 255},
                  ${(item[2] / 100) * 255},
                1.0);
              `}
            />
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
    display: flex;
    flex-direction: row-reverse;

    position: relative;
    top: 0;
    right: 0;
  }

  .container .actions .spacer {
    width: 100%;
  }

  .container .content {
    height: 65px;
    width: 100%;
    overflow-x: auto;
    overflow-y: hidden;

    border: 1px solid var(--theme-border-color);
    border-radius: var(--theme-border-radius);
  }

  .container .content .custom-list {
    display: flex;
    align-items: flex-start;
    overflow-x: auto;
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
