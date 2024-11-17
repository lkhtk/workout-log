<template>
  <div v-if="workoutsList.length>0">
    <div class="container-lg p-3 mb-5 bg-body-tertiary rounded"
      v-for="workoutItem in workoutsList" v-bind:key="workoutItem">
        <workoutComponent :workoutData="workoutItem"/>
      <div class="d-flex flex-row-reverse">
        <div class="btn-group"
          role="group" aria-label="Basic mixed styles example">
          <button type="button" class="btn btn-danger"
            @click="deleteWorkoutById(workoutItem.id)">Delete</button>
          <button type="button" class="btn btn-warning"
            @click="changeComponent('view', workoutItem.id)">Edit</button>
          <button type="button" class="btn btn-success"
            @click="changeComponent('create', '')">Add</button>
        </div>
      </div>
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
import {
  getAllWorkouts,
  deleteWorkout,
} from '@/api/api';
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
    async deleteWorkoutById(id) {
      await deleteWorkout(id);
      await this.getAllWorkouts();
    },
  },
};
</script>
