<template>
   <div class="container-lg p-3 bg-body-tertiary rounded">
      <workoutComponent :workoutData="userData" v-if="userData" :edit='true'/>
  </div>
  <button type="button" class="btn btn-outline-dark"
    @click="changeComponent('list', {})">
    <font-awesome-icon icon="fa-solid fa-angles-left" />
  </button>
</template>
<script>
import changeComponent from '../../mixin/changeComponent';
import workoutComponent from './workout.vue';

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
