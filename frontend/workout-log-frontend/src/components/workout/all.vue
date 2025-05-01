<template>
  <div id="overlay" class="show" v-if="loading">
      <div id="overlay-container">
        <div class="spinner-border align-items-center" role="status">
          <span class="visually-hidden align-items-center">{{ $t('info.loading') }}</span>
        </div>
      </div>
    </div>
  <div>
    <CreateButton :label="$t('workout.newButtonTitle')" @action="changeComponent('create', {})"/>
    <div v-if="workoutsList && workoutsList.length>0">
      <div class="container-lg p-3 bg-body-tertiary rounded shadow mb-5"
        v-for="workoutItem in workoutsList" v-bind:key="workoutItem.id">
        <workoutComponent :workoutData="workoutItem" />
        <div class="d-flex flex-row-reverse p-3">
          <div class="btn-group" role="group">
            <button type="button" class="btn btn-outline-dark"
              @click="cloneWorkout(workoutItem)">
              <font-awesome-icon icon="fa-solid fa-clone" />
              {{ $t('buttons.clone') }}
            </button>
            <button type="button" class="btn btn-outline-dark"
              @click="changeComponent('edit', workoutItem)">
              <font-awesome-icon icon="fa-solid fa-sliders" />
              {{ $t('buttons.edit') }}
            </button>
          </div>
        </div>
      </div>
      <nav aria-label="Page navigation" v-if="pagination.last > 1" class="p-3">
        <ul class="pagination justify-content-center" data-bs-theme="light">
          <li
            class="page-item"
            :class="{ disabled: pagination.current === 1 }">
            <button
              class="page-link"
              aria-label="first"
              @click="goToPage(pagination.current - 1)"
              :disabled="pagination.current === 1">
              <span aria-hidden="true">&laquo;</span>
            </button>
          </li>
          <li
            v-for="pageNumber in visiblePages"
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
              aria-label="last"
              @click="goToPage(pagination.current + 1)"
              :disabled="pagination.current === pagination.last">
              <span aria-hidden="true">&raquo;</span>
            </button>
          </li>
        </ul>
      </nav>
    </div>
    <div v-else>
      <div class="alert alert-danger" role="alert" v-if="errorData.msg">
        {{ errorData.msg }}
      </div>
      <div v-else-if="!loading" class="container-lg p-3 bg-body-tertiary rounded text-center">
        <h1 class="display-1">{{ $t('errorsMsg.noDataAvailable') }}</h1>
        <h2 class="display-1">😞</h2>
      </div>
    </div>
  </div>
</template>

<script>
import { getAllWorkouts } from '@/api/api';
import changeComponent from '@/mixin/changeComponent';
import workoutComponent from './workout.vue';
import CreateButton from '../common/addWorkoutButton.vue';

export default {
  name: 'ListWorkouts',
  mixins: [changeComponent],
  components: {
    workoutComponent,
    CreateButton,
  },
  props: {
    workoutData: {
      type: Object,
      required: false,
    },
  },
  async beforeMount() {
    this.loadAllWorkouts(1);
  },
  data: () => ({
    loading: false,
    toastMessage: '',
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
    visiblePages() {
      const { last: total, current } = this.pagination;
      const range = 4;
      const start = Math.max(1, current - range);
      const end = Math.min(total, current + range);

      const pages = Array.from({ length: end - start + 1 }, (_, i) => start + i);

      if (start > 1) {
        pages.unshift(1);
        if (start > 2) {
          pages.splice(1, 0, '...');
        }
      }

      if (end < total) {
        if (end < total - 1) {
          pages.push('...');
        }
        pages.push(total);
      }

      return pages;
    },
  },
  methods: {
    cloneWorkout(w) {
      const cloned = w;
      cloned.id = null;
      this.changeComponent('edit', cloned);
    },
    goToPage(pageNumber) {
      if (pageNumber === '...' || pageNumber < 1 || pageNumber > this.pagination.last) {
        return;
      }
      this.pagination.current = pageNumber;
      this.loadAllWorkouts(pageNumber);
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
        window.$toast?.showToast(this.$t('errorsMsg.failedMsg'), 'danger');
      } finally {
        this.loading = false;
        window.scrollTo(0, 0);
      }
    },
  },
};
</script>
<style scoped>
body {
  overflow: hidden;
}

#overlay {
  background-color: rgba(255,255,255,0.7);
  display: none;
  position: absolute;
  bottom: 0;
  top: 0;
  left: 0;
  right: 0;
  z-index: 2;
}

#overlay.show {
  display: block;
}

#overlay-container {
  align-items: center;
  justify-content: center;
  display: flex;
  flex-direction: column;
  height: 100%;
}
</style>
