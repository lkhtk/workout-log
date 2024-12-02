<template>
   <div class="container-lg p-3 mb-5 bg-body-tertiary rounded">
      <workoutComponent :workoutData="userData" v-if="userData" :edit='true'/>
  </div>
  <button type="button" class="btn btn-primary"
    @click="changeComponent('list', '')">
    <font-awesome-icon icon="fa-solid fa-angles-left" />
  </button>
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
