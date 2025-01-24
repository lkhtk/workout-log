<template>
  <div class="container mt-5" v-if="user">
    <h1 class="display-5">User Actions</h1>
    <form @submit.prevent="exportData" class="mb-4">
      <div class="d-grid gap-2">
        <button type="submit" class="btn btn-outline-primary">
          <font-awesome-icon icon="fa-solid fa-file-arrow-down" />
          Export Data</button>
        <button
        class="btn btn-outline-warning"
        @click="clearData"
      >
      <font-awesome-icon icon="fa-solid fa-triangle-exclamation" />
        Clear All Data
      </button>
      <button
        class="btn btn-outline-danger"
        @click="deleteAccount"
      >
        <font-awesome-icon icon="fa-solid fa-skull" />
        Delete Account
      </button>
      </div>
    </form>
  </div>
</template>

<script>
import { storeToRefs } from 'pinia';
import { Modal } from 'bootstrap';
import { useUserStore } from '../stores/userStore';

export default {
  name: 'userConfigs',
  setup() {
    const userStore = useUserStore();
    const { user } = storeToRefs(userStore);
    return {
      user,
    };
  },
  computed: {
    currentUser() {
      return this.user || {
        id: 0,
      };
    },
  },
  data() {
    return {
      formData: {
        username: '',
        email: '',
        description: '',
      },
    };
  },
  methods: {
    exportData() {
      // Export form data to JSON
      const jsonData = JSON.stringify(this.formData, null, 2);
      console.log('Exported Data:', jsonData);
    },
    clearData() {
      this.showConfirmationModal(
        'Are you sure you want to clear all data?',
        () => {
          this.formData = {
            username: '',
            email: '',
            description: '',
          };
        },
      );
    },
    deleteAccount() {
      // Simulate account deletion
      this.showConfirmationModal(
        'Are you sure you want to delete your account? This action cannot be undone.',
        () => {
          console.log('Account deleted.');
          // Reset the form data as a placeholder for account deletion
          this.formData = {
            username: '',
            email: '',
            description: '',
          };
        },
      );
    },
    showConfirmationModal(message, onConfirm) {
      const modalElement = document.createElement('div');
      modalElement.innerHTML = `
        <div class="modal fade" tabindex="-1">
          <div class="modal-dialog">
            <div class="modal-content">
              <div class="modal-header">
                <h5 class="modal-title">Confirmation</h5>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
              </div>
              <div class="modal-body">
                <p>${message}</p>
              </div>
              <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
                <button type="button" class="btn btn-primary" id="confirmBtn">Confirm</button>
              </div>
            </div>
          </div>
        </div>
      `;

      document.body.appendChild(modalElement);
      const modal = new Modal(modalElement.querySelector('.modal'));
      modal.show();

      modalElement.querySelector('#confirmBtn').addEventListener('click', () => {
        onConfirm();
        modal.hide();
        document.body.removeChild(modalElement);
      });

      modalElement.querySelector('[data-bs-dismiss="modal"]').addEventListener('click', () => {
        modal.hide();
        document.body.removeChild(modalElement);
      });
    },
  },
};
</script>
