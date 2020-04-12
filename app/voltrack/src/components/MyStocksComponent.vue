<template>
  <div>
    <h2>Portfolio</h2>
    <b-table :data="ownedStocks" :columns="columns">
      <template slot-scope="props">
        <b-table-column field="code" label="Stock Code">
          {{ props.row.code }}
        </b-table-column>

        <b-table-column field="quanity" label="Quantity">
          {{ props.row.quanity }}
        </b-table-column>

        <b-table-column field="currentValue" label="Value">
          ${{ parseFloat(props.row.currentValue).toFixed(2) }}
        </b-table-column>

        <b-table-column field="difference" label="Difference">
          ${{ parseFloat(props.row.difference).toFixed(2) }}
        </b-table-column>
      </template></b-table
    >
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { OwnedStockViewModel } from "../models/OwnedStockViewModel";
import { ApiClient } from "../api/apiClient";

@Component
export default class MyStocksComponent extends Vue {
  private ownedStocks: OwnedStockViewModel[] = [];
  private columns = [];
  private apiClient: ApiClient;

  constructor() {
    super();
    this.apiClient = new ApiClient("http://localhost:3000/api");
  }

  mounted() {
    this.getMyStocks();
  }

  async getMyStocks() {
    this.ownedStocks = await this.apiClient.GetCurrentStocks();
  }
}
</script>

<style scoped lang="scss">
</style>
