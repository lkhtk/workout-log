<template>
 <div class="container-lg p-3 mb-5 bg-body-tertiary rounded">
  <h1 class="display-1">Add a new workout log </h1>
      <button type="button" class="btn btn-success"
        @click="changeComponent('create', '')">create</button>
      <button type="button" class="btn btn-primary" @click="changeComponent('list', '')">
        back</button>
        <div class="btn-group" v-if="edit"
      role="group" aria-label="Basic mixed styles example">
      <button type="button" class="btn btn-success"
      @click="saveWorkout()">Update</button>
      <button type="button" class="btn btn-danger"
        @click="deleteWorkoutById(workoutData.id)">Delete</button>
      <button type="button" class="btn btn-success"
        @click="saveWorkout()">Cancel</button>
    </div>
  </div>
</template>
<script>
import changeComponent from '../mixin/changeComponent';
import { addWorkout } from '../api/api';

export default {
  name: 'CreateUser',
  mixins: [changeComponent],
  // components: {
  //   userForm,
  // },
  data: () => ({
    userData: {
      name: '',
      email: '',
      birthday: '',
      country: '',
      phone: '',
    },
  }),
  methods: {
    async createUser() {
      await addWorkout(`${window.location.href}api/users`, {
        data: {
          ...this.userData,
        },
      });
      this.changeComponent('list', '');
    },
  },
};
</script>
