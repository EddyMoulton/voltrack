<template>
  <div class="home">
    <my-stock-component></my-stock-component>
  </div>
</template>

<script>
// @ is an alias to /src

import MyStockComponent from "../components/MyStocksComponent.vue";

export default {
  name: "Home",
  components: { MyStockComponent }
};
</script>
