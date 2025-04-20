<template>
  <div
    v-if="visible"
    class="toast-container position-fixed bottom-0 end-0 p-3"
    style="z-index: 1100;"
  >
    <div
      class="toast show align-items-center text-white border-0"
      :class="toastClass"
    >
      <div class="d-flex">
        <div class="toast-body">
          {{ message }}
        </div>
        <button
          type="button"
          class="btn-close btn-close-white me-2 m-auto"
          @click="visible = false"
        ></button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';

const visible = ref(false);
const message = ref('');
const type = ref('danger');

const toastClass = computed(() => ({
  'bg-danger': type.value === 'danger',
  'bg-success': type.value === 'success',
  'bg-warning': type.value === 'warning',
  'bg-info': type.value === 'info',
  'bg-primary': type.value === 'primary',
  'bg-secondary': type.value === 'secondary',
  'bg-dark': type.value === 'dark',
}));
function showToast(msg, toastType = 'danger') {
  message.value = msg;
  type.value = toastType;
  visible.value = true;

  setTimeout(() => {
    visible.value = false;
  }, 4000);
}

defineExpose({
  showToast,
});
</script>
