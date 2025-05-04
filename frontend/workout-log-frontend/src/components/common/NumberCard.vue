<template>
  <div class="card h-100 border-1 rounded p-4 d-flex justify-content-center
  align-items-center text-center">
    <div class="mb-2">
      <font-awesome-icon :icon="cardData.icon" class="fs-2" />
    </div>
    <div class="fs-4 fw-bold">
      <h2 class="display-5">{{ props.cardData.count }}</h2>
    </div>
    <div class="text-muted mb-2">
      {{ cardData.text }}
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  cardData: {
    type: Object,
    required: true,
  },
});
</script>
