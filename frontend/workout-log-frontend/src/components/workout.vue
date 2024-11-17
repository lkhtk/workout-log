<template>
    <h1 class="display-5">Group {{ workoutData.workout.muscle_group }}</h1>
    <p class="fs-6 fw-light">{{ workoutData.publishedAt }}</p>
    <table class="table table-hover align-middle table-bordered" v-if="workoutData">
    <thead class="table-dark">
        <tr>
            <th colspan="3">Exercise</th>
            <th v-for="item in workoutData.workout.sets_count"
                v-bind:key="item">SET {{ item }}
            </th>
        </tr>
    </thead>
    <tbody v-for="(ex, index) in workoutData.workout.exercises" v-bind:key="ex">
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
    <table class="table table-hover align-middle table-bordered" v-if="workoutData.workout.cardio">
    <thead>
        <tr>
            <th>Cardio</th>
            <th>Speed/Level</th>
            <th>Distance</th>
            <th>Time</th>
            <th>Calories</th>
        </tr>
    </thead>
    <tbody v-for="cardio in workoutData.workout.cardio" v-bind:key="cardio">
        <tr>
            <td >{{cardio.type}}</td>
            <td >{{cardio.speed}}</td>
            <td >{{cardio.distance}}</td>
            <td >{{cardio.time}}</td>
            <td >{{cardio.calories}}</td>
        </tr>
    </tbody>
    <caption>
        <p class="font-monospace">{{ workoutData.id }}</p>
    </caption>
    </table>
</template>

<script>
export default {
  name: 'WorkoutComponent',
  props: {
    workoutData: {
      required: true,
    },
  },
};
</script>
