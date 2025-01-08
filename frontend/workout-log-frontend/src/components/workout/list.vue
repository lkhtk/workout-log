<template>
  <div class="spinner-border align-items-center" role="status" v-if="loading">
    <span class="visually-hidden align-items-center">Loading...</span>
  </div>
  <div v-else>
    <div v-if="workoutsList.length > 0">
      <div class="container-lg p-3 mb-5 bg-body-tertiary rounded"
        v-for="workoutItem in workoutsList" v-bind:key="workoutItem.id">
        <workoutComponent :workoutData="workoutItem" />
        <div class="d-flex flex-row-reverse">
          <div class="btn-group" role="group" aria-label="Basic mixed styles example">
            <button type="button" class="btn btn-primary"
              @click="changeComponent('view', workoutItem.id)">
              <font-awesome-icon inverse icon="fa-solid fa-file-pen" />
            </button>
            <button type="button" class="btn btn-success"
              @click="changeComponent('create', '')">
              <font-awesome-icon icon="fa-solid fa-calendar-plus" />
            </button>
          </div>
        </div>
        <hr>
      </div>
      <nav aria-label="Page navigation" v-if="pagination.total > 0">
        <ul class="pagination justify-content-center">
          <li
            class="page-item"
            :class="{ disabled: pagination.current === 1 }">
            <button
              class="page-link"
              @click="goToPage(pagination.current - 1)"
              :disabled="pagination.current === 1">
              Previous
            </button>
          </li>
          <li
            v-for="pageNumber in pages"
            :key="pageNumber"
            class="page-item"
            :class="{ active: pageNumber === pagination.current }">
            <button
              class="page-link"
              @click="goToPage(pageNumber)">
              {{ pageNumber }}
            </button>
          </li>
          <li
            class="page-item"
            :class="{ disabled: pagination.current === pagination.last }">
            <button
              class="page-link"
              @click="goToPage(pagination.current + 1)"
              :disabled="pagination.current === pagination.last">
              Next
            </button>
          </li>
        </ul>
      </nav>
    </div>
    <div v-else>
      <div class="alert alert-danger" role="alert" v-if="errorData.msg">
        {{ errorData.msg }}
      </div>
      <div v-else class="container-lg p-3 mb-5 bg-body-tertiary rounded text-center">
        <h1 class="display-1">There are no workouts yet</h1>
        <h2 class="display-1">😞</h2>
        <button type="button" class="btn btn-success"
          @click="changeComponent('create', '')">Add your first workout</button>
      </div>
    </div>
  </div>
</template>

<script>
import { getAllWorkouts } from '@/api/api';
import changeComponent from '@/mixin/changeComponent';
import workoutComponent from './workout.vue';

export default {
  name: 'ListWorkouts',
  mixins: [changeComponent],
  components: {
    workoutComponent,
  },
  async beforeMount() {
    this.loadAllWorkouts(1);
  },
  data: () => ({
    loading: false,
    workoutsList: [],
    pagination: {
      last: 0,
      current: 0,
      total: 0,
    },
    errorData: {
      code: '',
      msg: '',
    },
  }),
  computed: {
    pages() {
      return Array.from({ length: this.pagination.last }, (_, i) => i + 1);
    },
  },
  methods: {
    goToPage(pageNumber) {
      if (pageNumber >= 1 && pageNumber <= this.pagination.last) {
        this.pagination.current = pageNumber;
        this.loadAllWorkouts(pageNumber);
      }
    },
    async loadAllWorkouts(pageId) {
      this.loading = true;
      try {
        const data = await getAllWorkouts(pageId);
        this.workoutsList = data.data.data;
        this.pagination = {
          last: data.data.last_page,
          current: data.data.page,
          total: data.data.total,
        };
      } catch (error) {
        if (error.response?.status === 401) {
          localStorage.removeItem('user');
          this.$router.push('/about');
        } else {
          this.errorData.code = error.code;
          this.errorData.msg = error.message;
        }
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>
