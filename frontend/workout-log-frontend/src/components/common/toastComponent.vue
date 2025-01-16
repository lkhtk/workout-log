<template>
  <div
    class="toast-container position-fixed bottom-0 end-0 p-3"
    style="z-index: 1050;">
    <div
      ref="toast"
      class="toast align-items-center text-bg-primary border-0"
      role="alert"
      aria-live="assertive"
      aria-atomic="true">
      <div class="toast-header" v-if="header">
        <strong class="me-auto">{{ header }}</strong>
        <button
          type="button"
          class="btn-close"
          data-bs-dismiss="toast"
          aria-label="Close"
          @click="closeToast"></button>
      </div>
      <div class="toast-body">
        {{ body }}
      </div>
    </div>
  </div>
</template>

<script>
import { Toast } from 'bootstrap';

export default {
  name: 'ToastComponent',
  props: {
    header: {
      type: String,
      required: true,
    },
    body: {
      type: String,
      required: true,
    },
    autohide: {
      type: Boolean,
      default: true,
    },
    delay: {
      type: Number,
      default: 3000,
    },
  },
  data() {
    return {
      toastInstance: null,
    };
  },
  methods: {
    showToast() {
      if (this.toastInstance) {
        this.toastInstance.show();
      }
    },
    closeToast() {
      if (this.toastInstance) {
        this.toastInstance.hide();
      }
      this.$emit('close-toast');
    },
  },
  watch: {
    body: {
      immediate: true,
      handler() {
        this.showToast();
      },
    },
  },
  mounted() {
    const toastEl = this.$refs.toast;
    if (toastEl) {
      this.toastInstance = new Toast(toastEl, {
        autohide: this.autohide,
        delay: this.delay,
      });
      this.showToast();
    }
  },
  beforeUnmount() {
    if (this.toastInstance) {
      this.toastInstance.dispose();
    }
  },
};
</script>

<style scoped>
.toast-container {
  max-width: 300px;
}
</style>
