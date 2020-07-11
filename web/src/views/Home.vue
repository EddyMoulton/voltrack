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
          @click="isUploadLogsModalActive = true"
        >
          Add manual logs
        </button>
        <button
          class="button is-primary is-medium"
          @click="isUploadTranscationsModalActive = true"
        >
          Add manual transactions
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
    <b-modal :active.sync="isUploadLogsModalActive">
      <div class="card">
        <div class="card-content">
          <upload-stock-logs-component />
        </div>
      </div>
    </b-modal>
    <b-modal :active.sync="isUploadTranscationsModalActive">
      <div class="card">
        <div class="card-content">
          <upload-transactions-component />
        </div>
      </div>
    </b-modal>
  </div>
</template>

<script>
// @ is an alias to /src

import MyStocksComponent from '../components/MyStocksComponent.vue';
import AddTransactionComponent from '../components/AddTransactionComponent.vue';
import UploadStockLogsComponent from '../components/UploadStockLogsComponent.vue';
import UploadTransactionsComponent from '../components/UploadTransactionsComponent.vue';

export default {
  name: 'Home',
  components: {
    AddTransactionComponent,
    MyStocksComponent,
    UploadStockLogsComponent,
    UploadTransactionsComponent
  },
  data() {
    return {
      isModalActive: false,
      isUploadLogsModalActive: false,
      isUploadTranscationsModalActive: false
    };
  }
};
</script>
