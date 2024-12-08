<template>
    <div class="toast-container position-fixed bottom-0 end-0 p-3">
      <div class="toast" role="alert" aria-live="assertive"
        aria-atomic="true" ref="toast">
        <div class="toast-header">
          <strong class="me-auto">{{ header }}</strong>
          <button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close" />
        </div>
        <div class="toast-body">{{ body }}</div>
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
      default: '',
    },
  },
  mounted() {
    const toastEl = this.$refs.toast;
    this.toastInstance = new Toast(toastEl, {
      autohide: true,
      delay: 3000,
    });
    this.toastInstance.show();
  },
};
</script>
