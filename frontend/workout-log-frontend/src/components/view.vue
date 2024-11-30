<template>
   <div class="container-lg p-3 mb-5 bg-body-tertiary rounded">
      <workoutComponent :workoutData="userData" v-if="userData" :edit='true'/>
    <div class="btn-group"
      role="group" aria-label="Basic mixed styles example">
      <button type="button" class="btn btn-success"
        @click="saveWorkout()">Update</button>
      <button type="button" class="btn btn-danger"
        @click="deleteWorkoutById(userData.id)">Delete</button>
      <button type="button" class="btn btn-primary"
        @click="changeComponent('list', '')">Back</button>
    </div>
  </div>
</template>
<script>
import changeComponent from '../mixin/changeComponent';
import { getWorkout, updateWorkout, deleteWorkout } from '../api/api';
import workoutComponent from './workout.vue';
import router from '../router';

export default {
  name: 'ViewWorkout',
  mixins: [changeComponent],
  components: {
    workoutComponent,
  },
  props: {
    workoutId: {
      type: String,
      required: true,
    },
  },
  data: () => ({
    userData: {
      workout: {
        type: Object,
        required: true,
      },
    },
  }),
  async beforeMount() {
    await this.getWorkoutById();
  },
  methods: {
    async getWorkoutById() {
      const { data } = await getWorkout(this.workoutId);
      this.userData = data;
    },
    saveWorkout() {
      delete this.localWorkoutData.workout.sets_count;
      updateWorkout(this.localWorkoutData).then(router.push(''));
    },
    async deleteWorkoutById(id) {
      await deleteWorkout(id);
      this.changeComponent('list', '');
    },
  },
};
</script>
