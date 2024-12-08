<template>
  <div v-if="workoutsList.length>0">
    <div class="container-lg p-3 mb-5 bg-body-tertiary rounded"
      v-for="workoutItem in workoutsList" v-bind:key="workoutItem">
        <workoutComponent :workoutData="workoutItem"/>
      <div class="d-flex flex-row-reverse">
        <div class="btn-group"
          role="group" aria-label="Basic mixed styles example">
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
  </div>
  <div v-else>
    <div class="container-lg p-3 mb-5 bg-body-tertiary rounded">
      <h1 class="display-1">There are no workouts yet :(</h1>
      <button type="button" class="btn btn-success"
      @click="changeComponent('create', '')">Add your first workout</button>
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
    await this.getAllWorkouts();
  },
  data: () => ({
    workoutsList: [],
  }),
  methods: {
    async getAllWorkouts() {
      const { data } = await getAllWorkouts();
      this.workoutsList = data;
    },
  },
};
</script>
