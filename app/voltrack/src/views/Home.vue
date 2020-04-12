<template>
  <div class="home">
    <div class="level">
      <div class="level-right">
        <button
          class="button is-primary is-medium"
          @click="isModalActive = true"
        >
          Add Transaction
        </button>
      </div>
    </div>
    <my-stock-component></my-stock-component
    ><b-modal :active.sync="isModalActive">
      <div class="card">
        <div class="card-content">
          <add-transaction-component></add-transaction-component>
        </div>
      </div>
    </b-modal>
  </div>
</template>

<script>
// @ is an alias to /src

import MyStockComponent from "../components/MyStocksComponent.vue";
import AddTransactionComponent from "../components/AddTransactionComponent.vue";

export default {
  name: "Home",
  components: { AddTransactionComponent, MyStockComponent },
  data() {
    return {
      isModalActive: false
    };
  }
};
</script>
