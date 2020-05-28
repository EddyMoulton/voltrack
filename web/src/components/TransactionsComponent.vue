<template>
  <div>
    <h2>Transactions</h2>
    <b-tabs v-model="activeTab">
      <b-tab-item label="Capital">
        <b-table
          :data="transactions"
          :row-class="(row, index) => 'is-disabled'"
        >
          <template slot-scope="props">
            <b-table-column field="purchaseDate" label="Purchase Date">
              {{ props.row.purchaseDate | formatDate("DD/MM/YYYY") }}
            </b-table-column>

            <b-table-column field="code" label="Stock Code">
              {{ props.row.code }}
            </b-table-column>

            <b-table-column field="simpleReturn" label="Simple Return">
              ${{ parseFloat(props.row.capital.value).toFixed(2) }} ({{
                parseFloat(props.row.capital.simpleReturn).toFixed(1)
              }}%)
            </b-table-column>

            <b-table-column field="annualReturn" label="Annual Return">
              {{ parseFloat(props.row.capital.annualReturn).toFixed(1) }}%
            </b-table-column>

            <b-table-column
              field="scaledAnnualReturn"
              label="Annual Return (Scaled)"
            >
              <span v-if="props.row.capital.scaledAnnualReturn">
                {{
                  parseFloat(props.row.capital.scaledAnnualReturn).toFixed(2)
                }}%
              </span>
            </b-table-column>
          </template>
        </b-table>
      </b-tab-item>
      <b-tab-item label="Dividends">
        <b-table
          :data="transactions"
          :row-class="(row, index) => 'is-disabled'"
        >
          <template slot-scope="props">
            <b-table-column field="purchaseDate" label="Purchase Date">
              {{ props.row.purchaseDate | formatDate("DD/MM/YYYY") }}
            </b-table-column>

            <b-table-column field="code" label="Stock Code">
              {{ props.row.code }}
            </b-table-column>

            <b-table-column field="value" label="Value">
              ${{ parseFloat(props.row.dividends.value).toFixed(2) }} ({{
                parseFloat(props.row.dividends.simpleReturn).toFixed(2)
              }}%)
            </b-table-column>

            <b-table-column field="annualReturn" label="Annual Return">
              {{ parseFloat(props.row.dividends.annualReturn).toFixed(2) }}%
            </b-table-column>

            <b-table-column
              field="scaledAnnualReturn"
              label="Annual Return (Scaled)"
            >
              <span v-if="props.row.dividends.scaledAnnualReturn">
                {{
                  parseFloat(props.row.dividends.scaledAnnualReturn).toFixed(2)
                }}%
              </span>
            </b-table-column>
          </template>
        </b-table>
      </b-tab-item>
      <b-tab-item label="Totals">
        <b-table
          :data="transactions"
          :row-class="(row, index) => 'is-disabled'"
        >
          <template slot-scope="props">
            <b-table-column field="purchaseDate" label="Purchase Date">
              {{ props.row.purchaseDate | formatDate("DD/MM/YYYY") }}
            </b-table-column>

            <b-table-column field="code" label="Stock Code">
              {{ props.row.code }}
            </b-table-column>

            <b-table-column field="simpleReturn" label="Simple Return">
              ${{ parseFloat(props.row.total.value).toFixed(2) }} ({{
                parseFloat(props.row.total.simpleReturn).toFixed(2)
              }}%)
            </b-table-column>

            <b-table-column field="annualReturn" label="Annual Return">
              {{ parseFloat(props.row.total.annualReturn).toFixed(2) }}%
            </b-table-column>

            <b-table-column
              field="scaledAnnualReturn"
              label="Annual Return (Scaled)"
            >
              <span v-if="props.row.total.scaledAnnualReturn">
                {{ parseFloat(props.row.total.scaledAnnualReturn).toFixed(2) }}%
              </span>
            </b-table-column>
          </template>
        </b-table>
      </b-tab-item>
    </b-tabs>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { TransactionSummaryViewModel } from "../models/TransactionSummaryViewModel";
import { ApiClient } from "../api/apiClient";
import formatDate from "../filters/FormatDateFilter";

@Component({
  name: "TransactionsComponent",
  filters: {
    formatDate
  }
})
export default class TransactionsComponent extends Vue {
  private transactions: TransactionSummaryViewModel[] = [];
  private apiClient: ApiClient;
  private activeTab = 0;

  constructor() {
    super();
    this.apiClient = new ApiClient("http://localhost:3000/api");
  }

  mounted() {
    this.getMyStocks();
  }

  async getMyStocks() {
    this.transactions = await this.apiClient.GetTransactionSummaries();
    const total = new TransactionSummaryViewModel();
    total.summarise(this.transactions);
    console.log(total);

    this.transactions.push(total);
  }
}
</script>

<style scoped lang="scss"></style>
