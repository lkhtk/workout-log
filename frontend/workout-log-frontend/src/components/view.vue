<template>
   <div class="container-lg  p-3 mb-5 bg-body-tertiary rounded">
    <h1 class="display-5">Edit workout</h1>
      <workoutComponent :workoutData="userData" v-if="userData"/>
    <button type="button" class="btn btn-primary" @click="changeComponent('list', '')">Back</button>
  </div>
</template>
<script>
import changeComponent from '../mixin/changeComponent';
import { getWorkout } from '../api/api';
import workoutComponent from './workout.vue';

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
  },
};
</script>
