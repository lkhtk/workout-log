<template>
    <div class="d-flex h-100 text-center" >
        <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
            <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
                <h1 class="display-5">
                  <font-awesome-icon icon="fa-solid fa-weight-scale" />
                  Measurements
                </h1>
            </div>
            <main>
              <div class="container my-4">
                <form @submit.prevent="submitData">
                  <div class="table-responsive">
                    <table class="table table-bordered">
                      <thead class="table-dark text-center">
                        <tr>
                          <th colspan="3">BODY</th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr>
                          <td>DATE</td>
                          <td colspan="2">
                            <input
                              type="date"
                              class="form-control"
                              v-model="formData.measurement_date"
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
                    <button type="submit" class="btn btn-primary">
                      <font-awesome-icon icon="fa-solid fa-floppy-disk" inverse />
                      Submit
                    </button>
                  </div>
                </form>
              </div>
            </main>
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
import { useUserStore } from '../stores/userStore';
import ToastComponent from './common/toastComponent.vue';
import { createMeasurement, getMeasurement } from '../api/api';

export default {
  name: 'userMeasurement',
  components: {
    ToastComponent,
  },
  data() {
    return {
      toastTitle: '',
      toastMessage: '',
      formData: {
        measurement_date: this.getCurrentDate(),
        body_weight: 0,
        body_fat: 0,
        neck: 0,
        chest: 0,
        waist: 0,
        hips: 0,
        upperarm: 0,
        forearm: 0,
        thighs: 0,
        calves: 0,
      },
      measurements: {
        body_weight: 'Weight',
        body_fat: 'Body fat %',
        neck: 'Neck',
        chest: 'Chest',
        waist: 'Waist',
        hips: 'Hips',
        upperarm: 'Upper arm',
        forearm: 'Forearm',
        thighs: 'Thighs',
        calves: 'Calves',
      },
    };
  },
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
    showError(title, message) {
      this.toastTitle = title;
      this.toastMessage = message;
    },
    async loadUserMeasurement() {
      const response = await getMeasurement();
      if (response.status === 200) {
        if (response.data) {
          [this.formData] = response.data;
          this.formData.measurement_date = this.getCurrentDate();
        } else {
          this.resetForm();
        }
      }
    },
    async submitData() {
      try {
        this.formData.measurement_date = dayjs(this.formData.measurement_date).format('YYYY-MM-DDTHH:mm:ssZ');
        const response = await createMeasurement(this.formData);
        if (response.status === 201 || response.status === 200) {
          this.showError('Success', 'Data submitted successfully!');
        } else {
          this.showError('Error', 'Failed to submit data.');
          return;
        }
      } catch (error) {
        this.showError('Error', 'Error submitting data');
        console.error(error);
      }
    },
    resetForm() {
      this.formData = {};
      this.formData.measurement_date = this.getCurrentDate();
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
