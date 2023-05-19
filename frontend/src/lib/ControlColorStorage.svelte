<script lang="ts">
  import List, { Item, Meta } from "svelteui/src/list";

  export let items: Color[] = [];

  // TODO: load stored items from the `localStorage`
</script>

<div class="container">
  <div class="actions" />
  <div class="content">
    <List>
      {#each items as item}
        <Item>
          <Meta>
            <!-- TODO: color element here... -->
          </Meta>
        </Item>
      {/each}
    </List>
  </div>
</div>

<style>
  .container {
  }

  .container .actions {
  }

  .container .content {
  }
</style>
