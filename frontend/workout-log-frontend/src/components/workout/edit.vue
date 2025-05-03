<template>
   <div class="container-lg p-3 bg-body-tertiary rounded">
      <workoutComponent :workoutData="userData" v-if="userData" :edit='true'/>
  </div>
  <div class="container mt-3 text-center">
      <div class="btn-group" role="group" aria-label="label">
        <button type="button btn-lg"
          class="btn btn-outline-dark btn-lg"
          @click="changeComponent('all', {})">
          <font-awesome-icon icon="fa-solid fa-angles-left" />
        </button>
      </div>
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
