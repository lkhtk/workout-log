<template>
  <div class="container" v-for="workoutItem in workoutsList" v-bind:key="workoutItem">
    {{ workoutItem.user }}
    <h3>Group {{ workoutItem.workout.muscle_group }}</h3>
    <table class="table table-hover align-middle table-bordered" v-if="workoutItem">
      <thead class="table-dark">
        <tr>
            <th colspan="3">Exercise</th>
            <th v-for="item in workoutItem.workout.sets_count"
              v-bind:key="item">
                SET {{ item }}
            </th>
        </tr>
      </thead>
      <tbody v-for="(ex, index) in workoutItem.workout.exercises" v-bind:key="ex">
          <tr>
              <td rowspan="2">{{index + 1}}</td>
              <td rowspan="2">{{ex.name}}</td>
              <td>Weight</td>
              <td v-for="set in ex.sets" v-bind:key="set">
                {{ set.weight }}
              </td>
          </tr>
          <tr>
              <td>Reps</td>
              <td v-for="set in ex.sets" v-bind:key="set">
                {{ set.reps }}
              </td>
          </tr>
      </tbody>
    </table>
    <br>
    <table class="table table-hover align-middle table-bordered" v-if="workoutItem.workout.cardio">
      <thead>
        <tr>
            <th>Cardio</th>
            <th>Speed/Level</th>
            <th>Distance</th>
            <th>Time</th>
            <th>Calories</th>
        </tr>
      </thead>
      <tbody v-for="cardio  in workoutItem.workout.cardio" v-bind:key="cardio">
          <tr>
              <td >{{cardio.type}}</td>
              <td >{{cardio.speed}}</td>
              <td >{{cardio.distance}}</td>
              <td >{{cardio.time}}</td>
              <td >{{cardio.calories}}</td>
          </tr>
      </tbody>
      <caption>
        {{ workoutItem.id }}
      </caption>
    </table>
    {{ workoutItem.publishedAt }}
  </div>
  <button type="button" class="btn btn-primary" @click="changeComponent('create', 0)">
    Create Workout</button>
</template>
<script>
import {
  getAllWorkouts,
  deleteWorkout,
} from '@/api/api';
import changeComponent from '@/mixin/changeComponent';

export default {
  name: 'ListUser',
  mixins: [changeComponent],
  async beforeMount() {
    await this.getAllWorkouts();
  },
  data: () => ({
    workoutsList: [],
  }),
  methods: {
    async getAllWorkouts() {
      const { data } = await getAllWorkouts();
      this.workoutsList = data;
    },
    async deleteUser(id) {
      await deleteWorkout(id);
      await this.getAllWorkouts();
    },
  },
};
</script>
