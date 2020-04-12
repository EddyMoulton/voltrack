<template>
  <div>
    <div class="block">
      <b-radio v-model="action" name="action" native-value="buy">
        Buy
      </b-radio>
      <b-radio v-model="action" name="action" native-value="sell">
        Sell
      </b-radio>
    </div>
    <b-field label="Code">
      <b-input v-model="code"></b-input>
    </b-field>
    <b-field label="Date">
      <b-datepicker
        v-model="date"
        ref="datepicker"
        expanded
        placeholder="Select a date"
      >
      </b-datepicker>
      <b-button
        @click="$refs.datepicker.toggle()"
        icon-left="calendar-today"
        type="is-primary"
      />
    </b-field>
    <b-field label="Quantity">
      <b-numberinput v-model="quantity" />
    </b-field>
    <b-field label="Total Cost (ex fees)">
      <b-numberinput v-model="cost" step="0.01" />
    </b-field>
    <b-field label="Fees">
      <b-numberinput v-model="fees" step="0.01" />
    </b-field>
    <b-button type="is-primary" @click="addTransaction">Add</b-button>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import { AddTransactionDto } from "../api/models/AddTransactionDto";
import { ApiClient } from "../api/apiClient";

@Component
export default class AddTransactionComponent extends Vue {
  private code = "";
  private quantity = 0;
  private date = new Date();
  private action = "buy";
  private cost = 0;
  private fees = 0;

  private apiClient: ApiClient;

  constructor() {
    super();
    this.apiClient = new ApiClient("http://localhost:3000/api");
  }

  clearForm() {
    this.code = "";
    this.quantity = 0;
    this.date = new Date();
    this.action = "buy";
    this.cost = 0;
    this.fees = 0;
  }

  async addTransaction() {
    const dto = new AddTransactionDto();

    if (this.action === "buy") {
      dto.buySell = 1;
    } else if (this.action === "sell") {
      dto.buySell = -1;
    } else {
      throw Error("Invalid action");
    }

    dto.stockCode = this.code;
    dto.quantity = this.quantity;
    dto.date = this.date;
    dto.cost = this.cost * 10000;
    dto.fee = this.fees * 10000;

    try {
      await this.apiClient.AddTransaction(dto);
      this.clearForm();
      (this.$parent as any).close();
    } catch (e) {
      console.error("Failed to add transaction");
      console.error(e);
    }
  }
}
</script>

<style scoped lang="scss">
</style>
