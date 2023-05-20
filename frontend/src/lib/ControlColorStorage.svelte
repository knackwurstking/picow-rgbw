<script lang="ts">
  import List, { Item, Meta } from "svelteui/src/list";

  export let items: Color[] = [];

  // TODO: load stored items from the `localStorage`
</script>

<div class="container debug">
  <div class="actions debug2" />
  <div class="content debug2">
    <List class="debug" style="height: 100%;" orientation="vertical">
      {#each items as item}
        <Item class="debug2" style="width: 50px;" value={item}>
          <Meta>
            <div class="color-container__outer debug">
              <div class="color-container__inner debug">
                {item[0]}
                {item[1]}
                {item[2]}
              </div>
            </div>
          </Meta>
        </Item>
      {/each}
    </List>
  </div>
</div>

<style>
  .container {
    position: relative;
    overflow: hidden;
  }

  .container .actions {
    display: absolute;
    top: 0;
    right: 0;
  }

  .container .content {
    height: 65px;
    width: 100%;
  }

  .container .content .color-container__outer {
    width: 100%;
    height: 100%;
  }

  .container .content .color-container__inner {
    width: 46px;
  }

  :global(.debug) {
    border: 1px solid red;
  }

  :global(.debug2) {
    border: 1px solid blue;
  }

  :global(.debug3) {
    border: 1px solid green;
  }
</style>
