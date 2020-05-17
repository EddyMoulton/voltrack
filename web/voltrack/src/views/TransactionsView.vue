<template>
  <div><transactions-component></transactions-component></div>
</template>

<script>
import TransactionsComponent from "../components/TransactionsComponent.vue";

export default {
  name: "Home",
  components: { TransactionsComponent },
  data() {
    return {
      isModalActive: false
    };
  }
};
</script>
