<template>
  {{ data.user }}
  <h3>Workout {{ data.workout.muscle_group }}</h3>
  <table class="table table-hover align-middle" v-if="data">
    <thead>
      <tr>
          <th colspan="3">Exercise</th>
          <th v-for="item in data.workout.sets_count"
            v-bind:key="item">
              SET {{ item }}
          </th>
      </tr>
    </thead>
    <tbody v-for="(ex, index)  in data.workout.exercises" v-bind:key="ex">
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
  <table class="table  table-hover align-middle" v-if="data.workout.cardio">
    <thead>
      <tr>
          <th>Cardio</th>
          <th>Speed/Level</th>
          <th>Distance</th>
          <th>Time</th>
          <th>Calories</th>
      </tr>
    </thead>
    <tbody v-for="cardio  in data.workout.cardio" v-bind:key="cardio">
        <tr>
            <td >{{cardio.type}}</td>
            <td >{{cardio.speed}}</td>
            <td >{{cardio.distance}}</td>
            <td >{{cardio.time}}</td>
            <td >{{cardio.calories}}</td>
        </tr>
    </tbody>
  </table>
  {{ data.publishedAt }}
</template>

<script>
export default {
  name: 'WorkoutTable',
  props: ['data'],
};
</script>
