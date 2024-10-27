<template>
VIEW
</template>
<script>
import changeComponent from '../mixin/changeComponent';
import { addWorkout } from '../api/api';

export default {
  name: 'WievUser',
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
      this.changeComponent('list', 0);
    },
  },
};
</script>
