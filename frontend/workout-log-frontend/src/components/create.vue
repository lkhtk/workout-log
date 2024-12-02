<template>
 <div class="container-lg p-3 mb-5 bg-body-tertiary rounded">
    <workoutComponent :edit="true" :isNew="true" :workoutData="newWorkout" />
    <button type="button" class="btn btn-primary"
      @click="changeComponent('list', '')">
    <font-awesome-icon icon="fa-solid fa-angles-left" />
  </button>
  </div>
</template>
<script>
import changeComponent from '../mixin/changeComponent';
import workoutComponent from './workout.vue';

export default {
  name: 'CreateUser',
  mixins: [changeComponent],
  components: {
    workoutComponent,
  },
  data: () => ({
    newWorkout: {
      sets_count: 3,
      muscle_group: 'Legs',
      workout: {
        cardio: [{
          type: 'run',
          speed: 10,
          distance: 2.5,
          time: 60,
          calories: 200,
        }],
        exercises: [{
          name: 'squad',
          sets: [
            {
              reps: 0,
              weight: 0,
            },
            {
              reps: 0,
              weight: 0,
            },
            {
              reps: 0,
              weight: 0,
            },
          ],
        },
        ],
      },
    },
  }),

};
</script>
