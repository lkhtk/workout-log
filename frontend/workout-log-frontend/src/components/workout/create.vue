<template>
 <div class="container-lg p-3 mb-5 bg-body-tertiary rounded">
    <workoutComponent :edit="true" :workoutData="newWorkout" />
    <button type="button" class="btn btn-primary"
      @click="changeComponent('list', '')">
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
      muscle_group: 'Leg day',
      workout: {
        sets_count: 3,
        cardio: [{
          type: 'Jogging',
          speed: 10,
          distance: 2.5,
          time: 20,
          calories: 200,
        }],
        exercises: [{
          name: 'Bulgarian Split Squat',
          sets: [
            {
              reps: 15,
              weight: 32,
            },
            {
              reps: 15,
              weight: 32,
            },
            {
              reps: 15,
              weight: 32,
            },
          ],
        },
        ],
      },
    },
  }),

};
</script>
