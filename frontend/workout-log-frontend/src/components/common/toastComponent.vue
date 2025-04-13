<template>
  <div
    v-if="visible"
    class="toast-container position-fixed bottom-0 end-0 p-3"
    style="z-index: 1100;"
  >
    <div class="toast show align-items-center text-white bg-danger border-0">
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
import { ref, defineExpose } from 'vue';

const visible = ref(false);
const message = ref('');

function showToast(msg) {
  message.value = msg;
  visible.value = true;

  setTimeout(() => {
    visible.value = false;
  }, 4000);
}

defineExpose({
  showToast,
});
</script>
