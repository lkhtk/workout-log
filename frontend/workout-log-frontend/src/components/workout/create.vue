<template>
 <div class="container-lg p-3 bg-body-tertiary rounded">
    <workoutComponent :edit="true" :workoutData="newWorkout" />
    <button type="button" class="btn btn-dark"
      @click="changeComponent('list', {})">
    <font-awesome-icon icon="fa-solid fa-angles-left" />
  </button>
  </div>
</template>
<script>
import changeComponent from '../../mixin/changeComponent';
import workoutComponent from './workout.vue';

export default {
  name: 'CreateWorkout',
  mixins: [changeComponent],
  components: {
    workoutComponent,
  },
  data: () => ({
    newWorkout: {
      id: null,
      muscle_group: null,
      workout: {
        cardio: [{}],
        exercises: [{
          name: null,
          sets: [
            {
              reps: null,
              weight: null,
            },
            {
              reps: null,
              weight: null,
            },
            {
              reps: null,
              weight: null,
            },
          ],
        },
        ],
      },
    },
  }),

};
</script>
