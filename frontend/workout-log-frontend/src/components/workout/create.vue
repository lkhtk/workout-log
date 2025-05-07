<template>
  <div class="container-lg p-3 bg-body-tertiary rounded">
    <workoutComponent :edit="true" :workoutData="newWorkout" />
  </div>
</template>
<script>
import changeComponent from '../../mixin/changeComponent';
import workoutComponent from './workoutPage.vue';

export default {
  name: 'CreateWorkout',
  mixins: [changeComponent],
  components: {
    workoutComponent,
  },
  data: () => ({
    newWorkout: {
      id: null,
      coach: false,
      review: {},
      workout: {
        cardio: [{
          heart: 60,
        }],
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
