<template>
  <p v-if="localWorkoutData.id" class="font-monospace">
    <font-awesome-icon icon="fa-solid fa-fingerprint"/>
    {{ localWorkoutData.id }}
  </p>
  <div>
  <h1 class="display-3">
    <font-awesome-icon icon="fa-solid fa-person-swimming" />
    {{ formatDate(localWorkoutData.PublishedAt) }}
  </h1>
  <label v-if="edit" for="inputAddress" class="form-label">Name of workout
    <input type="text" class="form-control" id="inputAddress"
      v-model="localWorkoutData.muscle_group"
      placeholder="Legs day">
  </label>
  <h1 v-else class="display-5">{{ localWorkoutData.muscle_group }}</h1>
  <table class="table table-hover align-middle table-bordered"
    v-if="localWorkoutData.workout.exercises || edit">
    <thead class="table-dark">
      <tr>
        <th colspan="3"><font-awesome-icon icon="fa-solid fa-medal" /> Exercise</th>
        <th v-for="(item, index) in localWorkoutData.workout.sets_count" :key="index">
          SET {{ index + 1 }}
        </th>
        <th v-if="edit">
          <button type="button" class="btn btn-success" @click="addSet()">
            <font-awesome-icon icon="fa-solid fa-circle-plus" inverse/>
          </button>
        </th>
      </tr>
    </thead>
    <tbody v-for="(ex, exIndex) in localWorkoutData.workout.exercises" :key="exIndex">
      <tr>
        <td rowspan="2">
          <button v-if="edit" type="button" class="btn-close"
          data-bs-dismiss="alert" aria-label="Close" @click="removeExercise(exIndex)"></button>
          <p v-else>
            {{ exIndex + 1 }}
          </p>
        </td>
        <td rowspan="2">
          <template v-if="edit">
            <input v-model="ex.name" class="form-control" />
          </template>
          <template v-else>
            {{ ex.name }}
          </template>
        </td>
        <td><font-awesome-icon icon="fa-solid fa-weight-hanging" /></td>
        <td v-for="(set, setIndex) in ex.sets" :key="'weight-' + setIndex">
          <template v-if="edit">
            <label for="weight">
              <input
                id="weight"
                v-model="set.weight"
                class="form-control"
                type="number"
                min="0"
              />
            </label>
          </template>
          <template v-else>
            {{ set.weight }}
          </template>
        </td>
      </tr>
      <tr>
        <td><font-awesome-icon icon="fa-solid fa-repeat" /></td>
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
  </table>
  <div v-if="edit">
    <button type="button" class="btn btn-success" @click="addExercise()">
      <font-awesome-icon icon="fa-solid fa-circle-plus" inverse/>
    </button>
  </div>
  <h1 class="display-5"><font-awesome-icon icon="fa-solid fa-person-running" />Cardio</h1>
  <table
    class="table table-hover align-middle table-bordered"
    v-if="localWorkoutData.workout.cardio">
    <thead>
      <tr>
          <th><font-awesome-icon icon="fa-solid fa-heart-pulse" /> Type</th>
          <th><font-awesome-icon icon="fa-solid fa-gauge-high" /> Speed/Level</th>
          <th><font-awesome-icon icon="fa-solid fa-infinity" /> Distance</th>
          <th><font-awesome-icon icon="fa-solid fa-stopwatch-20" /> Time</th>
          <th><font-awesome-icon icon="fa-solid fa-fire" /> Calories</th>
      </tr>
    </thead>
    <tbody v-for="cardio in localWorkoutData.workout.cardio" v-bind:key="cardio">
        <tr>
            <td >
              <label for="weight"><input v-if="edit"
                v-model="cardio.type"
                class="form-control"
                />
              <p v-else>
                {{cardio.type}}
              </p>
              </label>
            </td>
            <td >
              <label for="speed">
                <input v-if="edit"
                  v-model="cardio.speed"
                  id="speed"
                  class="form-control"
                  type="number"
                  min="0"
                />
                <p v-else>
                  {{cardio.speed}}
                </p>
              </label>
            </td>
            <td >
              <label for="speed">
                <input v-if="edit"
                  v-model="cardio.distance"
                  id="speed"
                  class="form-control"
                  min="0"
                />
                <p v-else>
                  {{cardio.distance}}
                </p>
              </label>
            </td>
            <td >
              <label for="time">
                <input v-if="edit"
                  v-model="cardio.time"
                  id="time"
                  class="form-control"
                  min="0"
                />
                <p v-else>
                  {{cardio.time}}
                </p>
              </label>
            </td>
            <td >
              <label for="calories">
                <input v-if="edit"
                  v-model="cardio.calories"
                  id="calories"
                  class="form-control"
                  min="0"
                />
                <p v-else>
                  {{cardio.calories}}
                </p>
              </label>
            </td>
        </tr>
    </tbody>
  </table>
  <div v-if="edit">
      <button type="button" class="btn btn-success" @click="addCardio()">
        <font-awesome-icon icon="fa-solid fa-circle-plus" inverse/>
      </button>
    </div>
  </div>
  <div class="alert alert-info mx-auto p-2 m-2" role="alert" v-if="alertData.msg">
    <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"
    @click="hideAlert()"></button>
    {{ alertData.msg }}
  </div>
  <div class="container mx-auto p-2" style="width: 200px;">
    <div class="btn-group" v-if="edit"
      role="group" aria-label="Basic mixed styles example">
      <button type="button" class="btn btn-success" @click="saveWorkout()">
        <font-awesome-icon icon="fa-solid fa-floppy-disk" inverse />
      </button>
      <button type="button" class="btn btn-danger" @click="deleteWorkoutById(localWorkoutData.id)">
        <font-awesome-icon icon="fa-solid fa-trash-can" inverse />
      </button>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs';
import { updateWorkout, deleteWorkout, createWorkout } from '../api/api';

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
    let locationData;
    try {
      locationData = JSON.parse(JSON.stringify(this.workoutData));
    } catch (e) {
      console.error('Invalid JSON in propsName:', e);
      locationData = {};
    }
    return {
      localWorkoutData: locationData,
      alertData: { code: '', msg: '' },
    };
  },
  methods: {
    formatDate(dateString) {
      return dayjs(dateString).format('YYYY-MM-DD hh:mm');
    },

    addSet() {
      this.localWorkoutData.workout.sets_count += 1;
      this.localWorkoutData.workout.exercises.forEach((exercise) => {
        exercise.sets.push({ weight: 0, reps: 0 });
      });
      this.syncData();
    },

    addExercise() {
      const emptySets = [];
      for (let i = 0; i <= this.localWorkoutData.workout.sets_count; i += 1) {
        emptySets.push({ weight: 0, reps: 0 });
      }
      this.localWorkoutData.workout.exercises.push({
        name: 'New Exercise',
        sets: emptySets,
      });
      this.syncData();
    },

    addCardio() {
      this.localWorkoutData.workout.cardio.push({
        type: 'New Cardio Exercise',
        speed: 0,
        distance: 0,
        time: 0,
        calories: 0,
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
    hideAlert() {
      this.alertData = {};
    },
    async saveWorkout() {
      this.localWorkoutData.workout.cardio.forEach((exercise, index) => {
        this.localWorkoutData.workout.cardio[index].distance = parseFloat(exercise.distance);
        this.localWorkoutData.workout.cardio[index].speed = parseFloat(exercise.speed);
        this.localWorkoutData.workout.cardio[index].time = parseFloat(exercise.time);
        this.localWorkoutData.workout.cardio[index].calories = parseFloat(exercise.calories);
      });
      if (this.localWorkoutData.id) {
        await updateWorkout(this.localWorkoutData);
        this.alertData.msg = 'saved';
      } else {
        await createWorkout(this.localWorkoutData);
        this.alertData.msg = 'Created';
        window.location.reload();
      }
    },

    async deleteWorkoutById(id) {
      await deleteWorkout(id);
      this.alertData.msg = 'Deleted!';
      this.localWorkoutData.workout = {};
      window.location.reload();
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
