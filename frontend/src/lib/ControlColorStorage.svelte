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
            <div class="color-container__inner">
              {item[0]}
              {item[1]}
              {item[2]}
            </div>
          </div>
        </li>
      {/each}
    </ul>
  </div>
</div>

<style>
  .container {
    position: relative;
    height: 250px;
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

  .container .content .custom-list .custom-list-item {
    display: block;
    float: left;
    width: fit-content;
    height: 100%;
  }

  .container .content .color-container__outer {
    width: 100%;
    height: 100%;
  }

  .container .content .color-container__inner {
    width: calc(100% - 4px);
    height: clac(100% - 4px);
    margin: 2px;
  }
</style>
