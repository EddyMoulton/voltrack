<template>
  <div class="home">
    <div class="level">
      <div class="level-right">
        <button
          class="button is-primary is-medium"
          @click="isModalActive = true"
        >
          Add transaction
        </button>
        <button
          class="button is-primary is-medium"
          @click="isUploadModalActive = true"
        >
          Add manual logs
        </button>
      </div>
    </div>
    <my-stocks-component />
    <b-modal :active.sync="isModalActive">
      <div class="card">
        <div class="card-content">
          <add-transaction-component />
        </div>
      </div>
    </b-modal>
    <b-modal :active.sync="isUploadModalActive">
      <div class="card">
        <div class="card-content">
          <upload-stock-logs-component />
        </div>
      </div>
    </b-modal>
  </div>
</template>

<script>
// @ is an alias to /src

import MyStocksComponent from "../components/MyStocksComponent.vue";
import AddTransactionComponent from "../components/AddTransactionComponent.vue";
import UploadStockLogsComponent from "../components/UploadStockLogsComponent.vue";

export default {
  name: "Home",
  components: {
    AddTransactionComponent,
    MyStocksComponent,
    UploadStockLogsComponent
  },
  data() {
    return {
      isModalActive: false,
      isUploadModalActive: false
    };
  }
};
</script>
