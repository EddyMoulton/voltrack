<template>
  <div>
    <add-transaction-component></add-transaction-component>
  </div>
</template>

<script>
import AddTransactionComponent from '../components/AddTransactionComponent.vue';

export default {
  name: 'AddTransactionView',
  components: { AddTransactionComponent },
};
</script>
