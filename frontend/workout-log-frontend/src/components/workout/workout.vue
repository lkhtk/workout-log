<template>
  <p v-if="localWorkoutData.id && edit" class="font-monospace">
    <font-awesome-icon icon="fa-solid fa-fingerprint"/>
    {{ localWorkoutData.id }}
  </p>
  <div>
    <label v-if="edit" for="muscleGroup" class="form-label">Muscle group or name of workout
      <input type="text" class="form-control" id="muscleGroup"
        maxlength="250"
        v-model="localWorkoutData.muscle_group"
        placeholder="Legs day">
    </label>
    <h1 v-else class="display-3">
      <font-awesome-icon icon="fa-solid fa-hand-fist" />
      {{ localWorkoutData.muscle_group }}
    </h1>
    <p class="lead">
      {{ formatDate(localWorkoutData.PublishedAt) }}
    </p>
    <div v-if="localWorkoutData.workout.sets_count>0 || edit">
      <h1 class="display-6"><font-awesome-icon icon="fa-solid fa-dumbbell" />Strength training </h1>
      <table class="table table-hover align-middle table-bordered table-striped table-light">
        <thead >
          <tr>
            <th scope="row" colspan="3"><font-awesome-icon icon="fa-solid fa-medal" />Exercise</th>
            <th scope="col" v-for="(item, index) in localWorkoutData.workout.sets_count"
              :key="index">
              <font-awesome-icon icon="fa-solid fa-fire-flame-curved" /> Set {{ index + 1 }}
              <button v-if="edit" type="button" class="btn-close"
              aria-label="Close" @click="removeSet(index)"></button>
            </th>
            <th v-if="edit">
              <button
                type="button"
                class="btn btn-success"
                @click="addSet()">
                <font-awesome-icon icon="fa-solid fa-circle-plus" inverse/>
              </button>
            </th>
          </tr>
        </thead>
        <tbody v-for="(ex, exIndex) in localWorkoutData.workout.exercises" :key="exIndex">
          <tr>
            <td rowspan="2">
              <button v-if="edit" type="button" class="btn-close"
              aria-label="Close" @click="removeExercise(exIndex)"></button>
              <p v-else>
                {{ exIndex + 1 }}
              </p>
            </td>
            <td rowspan="2">
              <template v-if="edit">
                <input v-model="ex.name" class="form-control" maxlength="250"/>
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
              <label for="reps"><input
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
          <font-awesome-icon icon="fa-solid fa-circle-plus" inverse/> new exercise
        </button>
      </div>
    </div>
    <div v-if="(typeof localWorkoutData.workout.cardio !== 'undefined'
      && localWorkoutData.workout.cardio.length > 0) || edit">
      <h1 class="display-6"><font-awesome-icon icon="fa-solid fa-heart-pulse" />Cardio</h1>
      <table class="table table-hover align-middle table-bordered table-light">
        <thead>
          <tr>
              <th>#</th>
              <th><font-awesome-icon icon="fa-solid fa-solid fa-person-swimming" /> Activity</th>
              <th><font-awesome-icon icon="fa-solid fa-gauge-high" /> Speed/Level</th>
              <th><font-awesome-icon icon="fa-solid fa-infinity" /> Distance</th>
              <th><font-awesome-icon icon="fa-solid fa-stopwatch-20" /> Time</th>
              <th><font-awesome-icon icon="fa-solid fa-fire" /> Calories</th>
          </tr>
        </thead>
        <tbody v-for="cardio in localWorkoutData.workout.cardio" v-bind:key="cardio">
          <tr>
            <td v-if="edit">
              <button v-if="edit" type="button" class="btn-close"
              aria-label="Close" @click="removeCardio(exIndex)"></button>
            </td>
            <td v-else>1</td>
            <td >
              <label for="type">
                <input v-if="edit"
                maxlength="250"
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
            <td>
              <label for="distance">
                <input v-if="edit"
                  v-model="cardio.distance"
                  id="distance"
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
            <font-awesome-icon icon="fa-solid fa-circle-plus" inverse/> new cardio
          </button>
        </div>
      </div>
    </div>
    <div class="container mx-auto p-2" style="width: 200px;">
      <div class="btn-group" v-if="edit"
        role="group" aria-label="Basic mixed styles example">
        <button type="button" class="btn btn-success" @click="saveWorkout()">
          <font-awesome-icon icon="fa-solid fa-floppy-disk" inverse />
        </button>
        <button type="button" class="btn btn-danger"
          @click="deleteWorkoutById(localWorkoutData.id)">
          <font-awesome-icon icon="fa-solid fa-trash-can" inverse />
        </button>
      </div>
    </div>
    <!-- Toast container -->
    <div class="toast-container position-fixed bottom-0 end-0 p-3">
      <div class="toast" role="alert" aria-live="assertive"
        aria-atomic="true" ref="toast">
        <div class="toast-header">
          <strong class="me-auto">{{ toastTitle }}</strong>
          <button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close" />
        </div>
        <div class="toast-body">{{ toastMessage }}</div>
      </div>
    </div>
</template>

<script>
import dayjs from 'dayjs';
import { Toast } from 'bootstrap';
import { ref } from 'vue';
import { updateWorkout, deleteWorkout, createWorkout } from '../../api/api';

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
      isAddButtonDisabled: ref(false),
      localWorkoutData: locationData,
      toastInstance: null,
      toastTitle: 'Default Title',
      toastMessage: 'Default Message',
      toastTime: '',
    };
  },
  mounted() {
    const toastEl = this.$refs.toast;
    this.toastInstance = new Toast(toastEl, {
      autohide: true,
      delay: 3000,
    });
  },
  methods: {
    formatDate(dateString) {
      return dayjs(dateString).format('DD.MM.YYYY hh:mm');
    },

    addSet() {
      if (this.localWorkoutData.workout.sets_count >= 10) {
        this.showToast('limit of sets', 'info');
        return;
      }
      this.localWorkoutData.workout.exercises.forEach((exercise) => {
        exercise.sets.push({ weight: 0, reps: 0 });
      });
      this.localWorkoutData.workout.sets_count += 1;
      this.syncData();
    },
    showToast(message, title = 'Notification') {
      this.toastTitle = title;
      this.toastMessage = message;
      this.toastInstance.show();
    },
    addExercise() {
      if (this.localWorkoutData.workout.exercises.length >= 10) {
        this.showToast('limit of exercises', 'info');
        return;
      }
      const emptySets = [];
      if (this.localWorkoutData.workout.sets_count === 0) {
        this.localWorkoutData.workout.sets_count += 1;
      }
      for (let i = 0; i < this.localWorkoutData.workout.sets_count; i += 1) {
        emptySets.push({ weight: 0, reps: 0 });
      }
      this.localWorkoutData.workout.exercises.push({
        name: 'New Exercise',
        sets: emptySets,
      });
      this.syncData();
    },

    addCardio() {
      if (this.localWorkoutData.workout.cardio.length >= 10) {
        this.showToast('limit of cardio', 'info');

        return;
      }
      this.localWorkoutData.workout.cardio.push({
        type: 'New Cardio Exercise',
        speed: 0,
        distance: 0,
        time: 0,
        calories: 0,
      });
      this.syncData();
    },

    removeCardio(index) {
      this.localWorkoutData.workout.cardio.splice(index, 1);
      this.syncData();
    },

    removeExercise(index) {
      this.localWorkoutData.workout.exercises.splice(index, 1);
      this.syncData();
    },

    removeSet(index) {
      if (this.localWorkoutData.workout.sets_count > 1) {
        this.localWorkoutData.workout.exercises.forEach((exercise) => {
          if (exercise.sets && index < exercise.sets.length) {
            exercise.sets.splice(index, 1);
          }
        });
        if (Array.isArray(this.localWorkoutData.workout.sets_count)) {
          this.localWorkoutData.workout.sets_count.splice(index, 1);
        } else if (typeof this.localWorkoutData.workout.sets_count === 'number') {
          this.localWorkoutData.workout.sets_count -= 1;
        }
        this.syncData();
      }
    },

    syncData() {
      this.$emit('update:workoutData', this.localWorkoutData);
    },
    async saveWorkout() {
      this.localWorkoutData.workout.cardio.forEach((exercise, index) => {
        this.localWorkoutData.workout.cardio[index].distance = parseFloat(exercise.distance);
        this.localWorkoutData.workout.cardio[index].speed = parseFloat(exercise.speed);
        this.localWorkoutData.workout.cardio[index].time = parseFloat(exercise.time);
        this.localWorkoutData.workout.cardio[index].calories = parseFloat(exercise.calories);
      });
      if (this.localWorkoutData.id) {
        await updateWorkout(this.localWorkoutData).then(() => {
          this.showToast('Saved', 'Success');
        }).catch((error) => {
          this.showToast(error.code, error.message);
        });
      } else {
        await createWorkout(this.localWorkoutData).then(() => {
          this.showToast('Created');
          this.errorData = {};
          this.localWorkoutData.workout = {};
          window.location.reload();
        })
          .catch((error) => {
            this.showToast(error.message, error.code);
          });
      }
    },
    async deleteWorkoutById(id) {
      await deleteWorkout(id).then(() => {
        this.showToast('Deleted', 'Info');
        this.errorData = {};
        this.localWorkoutData.workout = {};
        window.location.reload();
      })
        .catch((error) => {
          this.showToast(error.message, error.code);
        });
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
<style>
.toast-container {
  z-index: 1050;
}
</style>
