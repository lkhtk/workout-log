<template>
  <div>
    <h1 class="display-3">{{ formatDate(localWorkoutData.PublishedAt) }}</h1>
    <h1 class="display-5">{{ localWorkoutData.workout.muscle_group }}</h1>
    <table class="table table-hover align-middle table-bordered" v-if="localWorkoutData">
      <thead class="table-dark">
        <tr>
          <th colspan="3">Exercise</th>
          <th v-for="(item, index) in localWorkoutData.workout.sets_count" :key="index">
            SET {{ index + 1 }}
          </th>
          <th v-if="edit">
            <button class="btn btn-sm btn-success" @click="addSet()">Add Set</button>
          </th>
        </tr>
      </thead>
      <tbody v-for="(ex, exIndex) in localWorkoutData.workout.exercises" :key="exIndex">
        <tr>
          <td rowspan="2">{{ exIndex + 1 }}</td>
          <td rowspan="2">
            <template v-if="edit">
              <input v-model="ex.name" class="form-control" />
            </template>
            <template v-else>
              {{ ex.name }}
            </template>
          </td>
          <td>Weight</td>
          <td v-for="(set, setIndex) in ex.sets" :key="'weight-' + setIndex">
            <template v-if="edit">
              <label for="ss"><input
                v-model="set.weight"
                class="form-control"
                type="number"
                min="0"
              /></label>
            </template>
            <template v-else>
              {{ set.weight }}
            </template>
          </td>
          <td v-if="edit" rowspan="2">
            <button class="btn btn-sm btn-danger" @click="removeExercise(exIndex)">Delete</button>
          </td>
        </tr>
        <tr>
          <td>Reps</td>
          <td v-for="(set, setIndex) in ex.sets" :key="'reps-' + setIndex">
            <template v-if="edit">
            <label for="ssa"><input
                v-model="set.reps"
                class="form-control"
                type="number"
                min="0"
              /></label>
            </template>
            <template v-else>
              {{ set.reps }}
            </template>
          </td>
        </tr>
      </tbody>
      <tbody v-if="edit">
        <tr>
          <td colspan="100">
            <button class="btn btn-sm btn-success" @click="addExercise()">Add Exercise</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import dayjs from 'dayjs';

export default {
  name: 'WorkoutComponent',
  props: {
    workoutData: {
      type: Object,
      required: true,
    },
    edit: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      localWorkoutData: JSON.parse(JSON.stringify(this.workoutData)),
    };
  },
  methods: {
    formatDate(dateString) {
      return dayjs(dateString).format('YYYY-MM-DD hh:mm:ss');
    },
    addSet() {
      this.localWorkoutData.workout.exercises.forEach((exercise) => {
        exercise.sets.push({ weight: 0, reps: 0 });
      });
      this.localWorkoutData.workout.sets_count.push(
        this.localWorkoutData.workout.sets_count.length + 1,
      );
      this.syncData();
    },
    addExercise() {
      this.localWorkoutData.workout.exercises.push({
        name: 'New Exercise',
        sets: this.localWorkoutData.workout.sets_count.map(() => ({ weight: 0, reps: 0 })),
      });
      this.syncData();
    },
    removeExercise(index) {
      this.localWorkoutData.workout.exercises.splice(index, 1);
      this.syncData();
    },
    syncData() {
      this.$emit('update:workoutData', this.localWorkoutData);
    },
  },
  watch: {
    workoutData: {
      deep: true,
      immediate: true,
      handler(newData) {
        this.localWorkoutData = JSON.parse(JSON.stringify(newData));
      },
    },
  },
};
</script>
