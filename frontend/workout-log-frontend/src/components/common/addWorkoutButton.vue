<template>
    <button
      type="button"
      class="btn btn-success floating-button"
      @click="handleClick">
      <font-awesome-icon icon="fa-solid fa-plus" />
      {{ label }}
    </button>
  </template>

<script>
export default {
  name: 'CreateButton',
  props: {
    label: {
      type: String,
      default: 'Click Me',
    },
  },
  methods: {
    handleClick() {
      this.$emit('action');
    },
  },
};
</script>
<style scoped>
  .floating-button {
    position: fixed;
    bottom: 100px;
    right: 100px;
    z-index: 1000;
    padding: 10px 20px;
    border-radius: 50px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  }
  .floating-button:hover {
    box-shadow: 0 6px 10px rgba(0, 0, 0, 0.15);
  }
</style>
