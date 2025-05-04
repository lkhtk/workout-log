<template>
  <div class="container-lg p-3 bg-body-tertiary rounded">
    <workoutComponent :workoutData="userData" v-if="userData" :edit='true' />
  </div>
</template>
<script>
import changeComponent from '../../mixin/changeComponent';
import workoutComponent from './workoutPage.vue';

export default {
  name: 'ViewWorkout',
  mixins: [changeComponent],
  components: {
    workoutComponent,
  },
  props: {
    workoutData: {
      type: Object,
      required: true,
    },
  },
  data: () => ({
    userData: {
      type: Object,
      required: true,
    },
  }),
  async beforeMount() {
    await this.getWorkoutById();
  },

  methods: {
    async getWorkoutById() {
      this.userData = this.workoutData;
    },
  },
};
</script>
