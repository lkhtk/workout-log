<template>
  <div class="container" v-if="user">
    <form @submit.prevent="" class="mb-4">
      <div class="d-grid gap-2">
        <div class="container d-grid gap-2 col-6 mx-auto">
          <h1 class="display-5">{{ $t('userActions.title') }}</h1>
          <button type="button" class="btn btn-outline-dark" @click="exportData">
            <font-awesome-icon icon="fa-solid fa-file-arrow-down" />
            {{ $t('userActions.export') }}
          </button>
          <div class="h4 pb-2 mb-4 text-danger border-bottom border-danger">
            {{ $t('userActions.danger') }}
          </div>
          <button class="btn btn-outline-dark" @click="clearData">
            <font-awesome-icon icon="fa-solid fa-triangle-exclamation" />
            {{ $t('userActions.wipe') }}
          </button>
          <button class="btn btn-outline-danger" @click="deleteAccount">
            <font-awesome-icon icon="fa-solid fa-skull" />
           {{ $t('userActions.delete') }}
          </button>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import { storeToRefs } from 'pinia';
import { exportData as apiExportData, wipeData, deleteUser } from '../api/api';
import { useUserStore } from '../stores/userStore';

export default {
  name: 'userActions',
  setup() {
    const userStore = useUserStore();
    const { user } = storeToRefs(userStore);
    return {
      user,
    };
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
                <button type="button" class="btn btn-dark" id="confirmBtn">Confirm</button>
              </div>
            </div>
          </div>
        </div>
      `;
      document.body.appendChild(modalTemplate);

      const modalElement = modalTemplate.querySelector('.modal');
      const modalInstance = new window.bootstrap.Modal(modalElement);
      modalInstance.show();

      const confirmButton = modalTemplate.querySelector('#confirmBtn');
      confirmButton.addEventListener('click', () => {
        onConfirm();
        modalInstance.hide();
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
        window.$toast?.showToast(this.$t('info.actionSuccess'));
      } catch (error) {
        window.$toast?.showToast(this.$t('errorsMsg.exportFailed'));
      }
    },
    clearData() {
      this.showConfirmationModal(
        this.$t('info.cleanUpQuestion'),
        async () => {
          try {
            await wipeData();
            window.$toast?.showToast(this.$t('info.actionSuccess'));
          } catch (error) {
            window.$toast?.showToast(this.$t('errorsMsg.cleanupFailed'));
          }
        },
      );
    },
    deleteAccount() {
      this.showConfirmationModal(
        this.$t('info.deleteQuestion'),
        async () => {
          try {
            await deleteUser();
            this.userStore.logout();
            window.$toast?.showToast(this.$t('info.deleteSuccess'));
            setTimeout(() => {
              window.location.replace('/about');
            }, 2000);
          } catch (error) {
            window.$toast?.showToast(this.$t('errorsMsg.deleteFailed'));
            this.userStore.logout();
            setTimeout(() => {
              window.location.replace('/about');
            }, 2000);
          }
        },
      );
    },
  },
};
</script>
