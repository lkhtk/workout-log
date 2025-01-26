<template>
  <div class="container mt-5" v-if="user">
    <h1 class="display-5">User Actions</h1>
    <form @submit.prevent="" class="mb-4">
      <div class="d-grid gap-2">
        <button type="button" class="btn btn-outline-primary" @click="exportData">
          <font-awesome-icon icon="fa-solid fa-file-arrow-down" />
          Export Data
        </button>
        <button class="btn btn-outline-warning" @click="clearData">
          <font-awesome-icon icon="fa-solid fa-triangle-exclamation" />
          Clear All Data
        </button>
        <button class="btn btn-outline-danger" @click="deleteAccount">
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
import { exportData as apiExportData, wipeData, deleteUser } from '../api/api';

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
  methods: {
    showConfirmationModal(message, onConfirm) {
      const modalTemplate = document.createElement('div');
      modalTemplate.innerHTML = `
        <div class="modal fade" tabindex="-1" aria-hidden="true">
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
      document.body.appendChild(modalTemplate);
      const modalElement = modalTemplate.querySelector('.modal');
      const modalInstance = new Modal(modalElement);
      modalInstance.show();
      const confirmButton = modalTemplate.querySelector('#confirmBtn');
      confirmButton.addEventListener('click', () => {
        onConfirm();
        modalInstance.hide();
        modalElement.addEventListener('hidden.bs.modal', () => {
          document.body.removeChild(modalTemplate);
        });
      });

      modalElement.addEventListener('hidden.bs.modal', () => {
        document.body.removeChild(modalTemplate);
      });
    },
    async exportData() {
      try {
        const response = await apiExportData();
        if (response.status !== 200) {
          throw new Error(`Error: ${response.status}`);
        }
        const url = window.URL.createObjectURL(new Blob([response.data]));
        const contentDisposition = response.headers['content-disposition'];
        const filenameMatch = contentDisposition?.match(/filename="?([^"]+)"?/);
        const filename = filenameMatch ? filenameMatch[1] : 'download.zip';
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', filename);
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
        window.URL.revokeObjectURL(url);
      } catch (error) {
        console.error('Export failed:', error.message);
        this.showError('Export Error', 'Failed to export data. Please try again.');
      }
    },
    clearData() {
      this.showConfirmationModal(
        'Are you sure you want to clear all account data?',
        async () => {
          try {
            await wipeData();
            alert('Data cleared successfully.');
          } catch (error) {
            this.showError('Clear Data Error', 'Failed to clear data. Please try again.');
          }
        },
      );
    },
    deleteAccount() {
      this.showConfirmationModal(
        'Are you sure you want to delete your account? This action cannot be undone.',
        async () => {
          try {
            await deleteUser();
            this.userStore.clearUser();
            window.location.replace('/about');
          } catch (error) {
            console.log(error);
            console.error('delete failed:', error.message);
            window.location.replace('/about');
          }
        },
      );
    },
    showError(title, message) {
      console.error(`${title}: ${message}`);

      alert(`${title}: ${message}`);
    },
  },
};
</script>
