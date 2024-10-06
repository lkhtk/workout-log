<template>
    <div v-for="workout in data.data" v-bind:key="workout">
      <table border="1" v-if="workout">
        <caption>
          <h3>Workout {{ workout.workout.muscle_group }}</h3>
          {{ workout.user }} - {{ workout.publishedAt }}
        </caption>
        <thead>
          <tr>
              <th colspan="3">Exercise</th>
              <th v-for="item in workout.workout.sets_count"
                v-bind:key="item">
                  SET {{ item }}
              </th>
          </tr>
        </thead>
        <tbody v-for="(ex, index)  in workout.workout.exercises" v-bind:key="ex">
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
      <table border="1" v-if="workout.workout.cardio">
        <thead>
          <tr>
              <th>Cardio</th>
              <th>Speed/Level</th>
              <th>Distance</th>
              <th>Time</th>
              <th>Calories</th>
          </tr>
        </thead>
        <tbody v-for="cardio  in workout.workout.cardio" v-bind:key="cardio">
            <tr>
                <td >{{cardio.type}}</td>
                <td >{{cardio.speed}}</td>
                <td >{{cardio.distance}}</td>
                <td >{{cardio.time}}</td>
                <td >{{cardio.calories}}</td>
            </tr>
        </tbody>
      </table>
      <br>
    </div>
</template>

<script>
import * as api from '@/api/api';

export default {
  name: 'HomeView',
  components: {

  },
  data: () => ({
    data: 1,
  }),
  async beforeMount() {
    try {
      this.loadUserInfo();
    } catch (e) {
      console.error(e);
    }
  },
  methods: {
    async loadUserInfo() {
      this.data = await api.getAllWorkouts();
    },
  },
};
</script>
