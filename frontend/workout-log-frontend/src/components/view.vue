<template>
  {{ workoutId }} ==
  user: {{userData}}
   <div class="container-lg  p-3 mb-5 bg-body-tertiary rounded">
    <h1 class="display-5">View workout</h1>
    <button type="button" class="btn btn-primary" @click="changeComponent('list', '')">Back</button>
  </div>
</template>
<script>
import changeComponent from '../mixin/changeComponent';
import { getWorkout } from '../api/api';

export default {
  name: 'ViewWorkout',
  mixins: [changeComponent],
  inject: ['workoutId'],
  data: () => ({
    userData: {
      name: '',
      email: '',
      birthday: '',
      country: '',
      phone: '',
    },
  }),
  async beforeMount() {
    console.log(this.workoutId);
    await this.getWorkoutById();
  },
  methods: {
    async getWorkoutById() {
      console.log('this.workoutId ==>', this.workoutId);
      const { data } = await getWorkout(this.workoutId);
      this.userData = data;
    },
  },
};
</script>
