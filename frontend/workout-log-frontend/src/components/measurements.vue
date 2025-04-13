<template>
    <div class="row">
      <div class="col">
        <div class="d-flex h-100 text-center card shadow">
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
</template>
<script>
import dayjs from 'dayjs';
import { computed } from 'vue';
import { useI18n } from 'vue-i18n';
// eslint-disable-next-line import/no-cycle
import { createMeasurement, getLastMeasurement } from '../api/api';
import userProgress from './analytics/measurementsChart.vue';

export default {
  name: 'userMeasurement',
  components: {
    userProgress,
  },
  data() {
    return {
      toastTitle: '',
      toastMessage: '',
      formData: {
        measurementDate: this.getCurrentDate(),
        bodyWeight: 0,
        neck: 0,
        chest: 0,
        waist: 0,
        bodyFat: 0,
        hips: 0,
        upperarm: 0,
        forearm: 0,
        thighs: 0,
        calves: 0,
      },
    };
  },
  setup() {
    const { t } = useI18n();
    const measurements = computed(() => ({
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
    }));

    return {
      measurements,
    };
  },
  methods: {
    validateForm() {
      const {
        bodyWeight, bodyFat, neck, chest,
        waist, hips, upperarm, forearm, thighs, calves,
      } = this.formData;

      if (bodyFat >= 100 || bodyFat < 0) {
        window.$toast?.showToast(this.$t('errorsMsg.invalidbodyFat'));
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
          window.$toast?.showToast(this.$t('errorsMsg.invalidMeasurement'));
          return true;
        }
        return false;
      });
      if (invalidMeasurement) {
        return false;
      }
      return true;
    },
    async loadUserMeasurement() {
      await getLastMeasurement().then((response) => {
        if (response.data) {
          this.formData = response.data;
          this.formData.measurementDate = this.getCurrentDate();
        } else {
          window.$toast?.showToast(this.$t('errorsMsg.noDataAvailable'));
        }
      }).catch((error) => {
        window.$toast?.showToast(error.message);
      });
    },
    async submitData() {
      if (!this.validateForm()) {
        return;
      }
      try {
        this.formData.measurementDate = dayjs(this.formData.measurementDate)
          .format('YYYY-MM-DDTHH:mm:ssZ');
        const response = await createMeasurement(this.formData);
        if (response.status === 201 || response.status === 200) {
          window.$toast?.showToast(this.$t('info.okMsg'));
        } else {
          window.$toast?.showToast(this.$t('errorsMsg.failedMsg'));
          return;
        }
      } catch (error) {
        window.$toast?.showToast(this.$t('errorsMsg.failedMsg'));
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
