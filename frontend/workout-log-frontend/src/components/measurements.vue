<template>
    <div class="row">
      <div class="col">
        <div class="d-flex h-100 text-center" v-if="user">
          <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
              <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
                  <h1 class="display-5">
                    <font-awesome-icon icon="fa-solid fa-weight-scale" />
                    {{ $t('measurements.title') }}
                  </h1>
              </div>
              <main>
                <div class="container my-4">
                  <form @submit.prevent="submitData">
                    <div class="table-responsive">
                      <table class="table table-bordered">
                        <thead class="table-dark text-center"></thead>
                        <tbody>
                          <tr>
                            <td>{{ $t('measurements.date') }}</td>
                            <td colspan="2">
                              <input
                                type="date"
                                class="form-control"
                                v-model="formData.measurementDate"
                                required
                              />
                            </td>
                          </tr>
                          <tr v-for="(label, key) in measurements" :key="key">
                            <td>{{ label }}</td>
                            <td colspan="2">
                              <input
                                type="number"
                                min="0"
                                step="0.1"
                                max="350"
                                class="form-control"
                                v-model="formData[key]"
                                :placeholder="label"
                              />
                            </td>
                          </tr>
                        </tbody>
                      </table>
                    </div>
                    <div class="text-center mt-3">
                      <button type="submit" class="btn btn-outline-dark">
                        <font-awesome-icon icon="fa-solid fa-floppy-disk"/>
                        {{ $t('buttons.submit') }}
                      </button>
                    </div>
                  </form>
                </div>
              </main>
          </div>
        </div>
      </div>
      <div class="col">
        <userProgress />
      </div>
    </div>
    <!-- Toast container -->
    <ToastComponent
      v-if="toastMessage"
      :header="toastTitle"
      :body="toastMessage"
      @close-toast="clearToast"
    />
</template>
<script>
import { storeToRefs } from 'pinia';
import dayjs from 'dayjs';
import { useI18n } from 'vue-i18n';
import { useUserStore } from '../stores/userStore';
import userProgress from './analytics/measurementsChart.vue';
import ToastComponent from './common/toastComponent.vue';
import { createMeasurement, getMeasurement } from '../api/api';

export default {
  name: 'userMeasurement',
  components: {
    ToastComponent,
    userProgress,
  },
  data() {
    return {
      toastTitle: '',
      toastMessage: '',
      formData: {
        measurementDate: this.getCurrentDate(),
        bodyWeight: 0,
        bodyWat: 0,
        neck: 0,
        chest: 0,
        waist: 0,
        hips: 0,
        upperarm: 0,
        forearm: 0,
        thighs: 0,
        calves: 0,
      },
    };
  },
  setup() {
    const userStore = useUserStore();
    const { user } = storeToRefs(userStore);
    const { t } = useI18n();
    const measurements = {
      bodyWeight: t('measurements.bodyWeight'),
      bodyFat: t('measurements.bodyFat'),
      neck: t('measurements.neck'),
      chest: t('measurements.chest'),
      waist: t('measurements.waist'),
      hips: t('measurements.hips'),
      upperarm: t('measurements.upperarm'),
      forearm: t('measurements.forearm'),
      thighs: t('measurements.thighs'),
      calves: t('measurements.calves'),
    };
    return {
      user,
      measurements,
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
    clearToast() {
      this.toastTitle = '';
      this.toastMessage = '';
    },
    validateForm() {
      const {
        bodyWeight, bodyFat, neck, chest,
        waist, hips, upperarm, forearm, thighs, calves,
      } = this.formData;

      if (bodyFat >= 100 || bodyFat < 0) {
        this.showError('', this.$t('errorsMsg.invalidbodyFat'));
        return false;
      }

      const measurementsToValidate = {
        bodyWeight,
        neck,
        chest,
        waist,
        hips,
        upperarm,
        forearm,
        thighs,
        calves,
      };
      const invalidMeasurement = Object.entries(measurementsToValidate).some(([value]) => {
        if (value < 0 || value > 350) {
          this.showError('', this.$t('errorsMsg.invalidMeasurement'));
          return true;
        }
        return false;
      });
      if (invalidMeasurement) {
        return false;
      }
      return true;
    },

    showError(title, message) {
      this.toastTitle = title;
      this.toastMessage = message;
    },
    async loadUserMeasurement() {
      await getMeasurement().then((response) => {
        if (response.data) {
          [this.formData] = response.data;
          this.formData.measurementDate = this.getCurrentDate();
        } else {
          this.showError('', this.$t('errors.noDataAvailable'));
        }
      }).catch((error) => {
        if (error.response?.status === 401) {
          this.user = null;
          localStorage.removeItem('user');
          this.$router.push('/about');
        }
        this.showError(error.code, error.message);
      });
    },
    async submitData() {
      if (!this.validateForm()) {
        return;
      }
      try {
        this.formData.measurementDate = dayjs(this.formData.measurementDate).format('YYYY-MM-DDTHH:mm:ssZ');
        const response = await createMeasurement(this.formData);
        if (response.status === 201 || response.status === 200) {
          this.showError('', this.$t('info.okMsg'));
        } else {
          this.showError('Error', this.$t('errorsMsg.failedMsg'));
          return;
        }
      } catch (error) {
        this.showError('Error', this.$t('errorsMsg.failedMsg'));
        console.error(error);
      }
    },
    resetForm() {
      this.formData = {};
      this.formData.measurementDate = this.getCurrentDate();
    },
    getCurrentDate() {
      return dayjs().format('YYYY-MM-DD');
    },
  },
  mounted() {
    this.loadUserMeasurement();
  },
};
</script>
