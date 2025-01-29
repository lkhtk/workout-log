<template>
  <p v-if="localWorkoutData.id && edit" class="font-monospace">
    <font-awesome-icon icon="fa-solid fa-fingerprint"/>
    {{ localWorkoutData.id }}
  </p>
  <div>
    <div v-if="edit">
      <div class="input-group input-group-lg">
          <span class="input-group-text" id="inputGroup-sizing-lg">Title</span>
          <input label="rest"
          type="text" class="form-control" id="muscleGroup"
          placeholder="Legs day"
          maxlength="150"
          v-model="localWorkoutData.muscle_group"
          aria-describedby="inputGroup-sizing-lg">
      </div>
      <div class="form-text" id="basic-addon4">
        Muscle group or name of workout
      </div>
    </div>
    <div v-else>
      <div class="grid text-center">
        <div class="g-col-3">
          <h1 class="display-2">
            <font-awesome-icon icon="fa-solid fa-dumbbell" />
            {{ localWorkoutData.muscle_group }}
          </h1>
        </div>
        <div class="g-col-4">
          <h1 class="display-6">
            {{ formatDate(localWorkoutData.PublishedAt) }}
          </h1>
        </div>
      </div>
    </div>
    <div class="form-check form-switch">
      <label class="form-check-label" for="flexSwitchCheckChecked">
      <input class="form-check-input" type="checkbox" role="switch"
      id="flexSwitchCheckChecked" :disabled="!edit"
      v-model="localWorkoutData.coach"
      aria-checked>
      Training with an instructor
    </label>
    </div>
    <div v-if="localWorkoutData.workout.exercises.length > 0 || edit">
      <h1 class="display-6">
        Strength training</h1>
      <table class="table table-hover align-middle table-bordered table-striped table-light">
        <thead >
          <tr>
            <th scope="row" colspan="3">
              <font-awesome-icon icon="fa-solid fa-medal" />
              Exercise</th>
              <th
                scope="col"
                v-for="(set, index) in getMaxSets(localWorkoutData.workout.exercises)"
                :key="'set-' + index"
              >
              <font-awesome-icon icon="fa-solid fa-fire-flame-curved" />
              Set {{ index + 1 }}
            </th>
          </tr>
        </thead>
        <tbody v-for="(ex, exIndex) in localWorkoutData.workout.exercises" :key="exIndex">
          <tr>
            <td rowspan="2">
              <div v-if="edit" style="border: 0px solid #ccc; width: 50px;"
                class="d-flex flex-column align-items-center">
                <font-awesome-icon icon="fa-solid fa-trash"
                  style="color: #dc3545; cursor: pointer;"
                  class="mb-2"
                  @keydown.enter="removeExercise(exIndex)"
                  @click="removeExercise(exIndex)"
                  />
                <font-awesome-icon icon="fa-solid fa-circle-plus"
                  style="color: #198754; cursor: pointer;"
                  class="mb-2"
                  @keydown.enter="addSet(exIndex)"
                  @click="addSet(exIndex)"
                  />
                <font-awesome-icon icon="fa-solid fa-circle-minus"
                  style="color: #fd7e14; cursor: pointer;"
                  @keydown.enter="removeSet(exIndex)"
                  @click="removeSet(exIndex)"
                  />
              </div>
              <p v-else>
                {{ exIndex + 1 }}
              </p>
            </td>
            <td rowspan="2">
              <template v-if="edit">
                <input v-model="ex.name" class="form-control"
                maxlength="150" size="50"/>
              </template>
              <template v-else>
                {{ ex.name }}
              </template>
            </td>
            <td><font-awesome-icon icon="fa-solid fa-weight-hanging" /></td>
            <td v-for="(set, setIndex) in ex.sets" :key="'weight-' + setIndex">
              <template v-if="edit">
                <div class="col-7">
                  <label for="weight">
                    <input
                      id="weight"
                      v-model="set.weight"
                      class="form-control col-sm-10"
                      type="number"
                      min="0"
                      step="0.1"
                      maxlength="6"
                    />
                  </label>
                </div>
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
                <div class="col-7">
                  <label for="reps">
                    <input
                      v-model="set.reps"
                      maxlength="6"
                      class="form-control"
                      type="number"
                      min="0"
                    />
                  </label>
                </div>
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
          New exercise
        </button>
      </div>
    </div>
    <div v-if="(typeof localWorkoutData.workout.cardio !== 'undefined'
      && localWorkoutData.workout.cardio.length > 0) || edit">
      <h1 class="display-6"><font-awesome-icon icon="fa-solid fa-heart-pulse" />
        Cardio</h1>
      <table class="table table-hover align-middle table-bordered table-light">
        <thead>
          <tr>
              <th>#</th>
              <th><font-awesome-icon icon="fa-solid fa-solid fa-person-swimming" />
                Activity</th>
              <th><font-awesome-icon icon="fa-solid fa-gauge-high" />
                Speed/Level</th>
              <th><font-awesome-icon icon="fa-solid fa-wave-square" />
                Heart rate</th>
              <th><font-awesome-icon icon="fa-solid fa-infinity" />
                Distance</th>
              <th><font-awesome-icon icon="fa-solid fa-stopwatch-20" />
                Time</th>
              <th><font-awesome-icon icon="fa-solid fa-fire" />
                Calories</th>
          </tr>
        </thead>
        <tbody v-for="cardio in localWorkoutData.workout.cardio" v-bind:key="cardio">
          <tr>
            <td v-if="edit">
              <button v-if="edit" type="button" class="btn-close"
              aria-label="Close" @click="removeCardio(exIndex)" tabindex="-1"></button>
            </td>
            <td v-else>1</td>
            <td >
              <label for="type">
                <input v-if="edit"
                maxlength="150"
                type="text"
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
                  step="0.1"
                  maxlength="6"
                />
                <p v-else>
                  {{cardio.speed}}
                </p>
              </label>
            </td>
            <td>
              <label for="heart">
                <input v-if="edit"
                  v-model="cardio.heart"
                  id="heart"
                  class="form-control"
                  min="60"
                  type="number"
                  maxlength="3"
                  step="1"
                />
                <p v-else>
                  {{cardio.heart}}
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
                  type="number"
                  maxlength="6"
                  step="0.1"
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
                  type="number"
                  min="0"
                  maxlength="6"
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
                  type="number"
                  maxlength="6"
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
            New cardio
          </button>
        </div>
      </div>
    </div>
    <div class="container mx-auto p-2" style="width: 200px;" v-if="edit">
      <div class="btn-group"
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
    <ToastComponent
      v-if="toastMessage"
      :header="toastTitle"
      :body="toastMessage"
      @close-toast="clearToast"
    />

</template>

<script>
import dayjs from 'dayjs';
import { ref } from 'vue';
import { updateWorkout, deleteWorkout, createWorkout } from '../../api/api';
import ToastComponent from '../common/toastComponent.vue';
import { useUserStore } from '../../stores/userStore';

export default {
  name: 'WorkoutComponent',
  components: {
    ToastComponent,
  },
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
  computed: {
    maxSets() {
      return this.localWorkoutData?.workout?.exercises
        ? this.getMaxSets(this.localWorkoutData.workout.exercises)
        : 0;
    },
  },
  data() {
    let locationData;
    const userStore = useUserStore();
    try {
      locationData = JSON.parse(JSON.stringify(this.workoutData));
    } catch (e) {
      console.error('Invalid JSON in propsName:', e);
      locationData = {};
    }
    return {
      isAddButtonDisabled: ref(false),
      localWorkoutData: locationData,
      toastTitle: '',
      toastMessage: '',
      userStore,
    };
  },
  methods: {
    getMaxSets(exercises) {
      if (!exercises || !Array.isArray(exercises)) {
        return 0;
      }

      return exercises.reduce((max, exercise) => {
        if (exercise.sets && Array.isArray(exercise.sets)) {
          return Math.max(max, exercise.sets.length);
        }
        return max;
      }, 0);
    },
    showError(title, message) {
      this.toastTitle = title;
      this.toastMessage = '';
      this.$nextTick(() => {
        this.toastMessage = message;
        setTimeout(() => {
          this.clearToast();
        }, 3000);
      });
    },
    clearToast() {
      this.toastTitle = '';
      this.toastMessage = '';
    },
    formatDate(dateString) {
      return dayjs(dateString).format('DD.MM.YYYY hh:mm');
    },

    addSet(exerciseIndex) {
      const exercise = this.localWorkoutData.workout.exercises[exerciseIndex];
      if (!exercise) {
        console.error(`Exercise at index ${exerciseIndex} not found.`);
        return;
      }
      if (!exercise.sets) {
        exercise.sets = [];
      }
      if (exercise.sets.length >= 10) {
        this.showError('', 'Maximum number of sets = 10');
        return;
      }
      exercise.sets.push({
        reps: 0,
        weight: 0,
      });
      this.syncData();
    },
    addExercise() {
      if (this.localWorkoutData.workout.exercises.length >= 10) {
        this.showError('', 'Maximum number of exercises = 10');
        return;
      }
      const emptySets = [];
      for (let i = 0; i <= 2; i += 1) {
        emptySets.push({ weight: 10, reps: 12 });
      }
      this.localWorkoutData.workout.exercises.push({
        name: 'New Exercise',
        sets: emptySets,
      });
      this.syncData();
    },

    addCardio() {
      if (this.localWorkoutData.workout.cardio.length >= 10) {
        this.showError('', 'Maximum number of cardio = 10');
        return;
      }
      this.localWorkoutData.workout.cardio.push({
        type: 'New Cardio Exercise',
        speed: 10,
        distance: 100,
        time: 15,
        calories: 100,
        heart: 60,
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

    removeSet(exerciseIndex) {
      const exercise = this.localWorkoutData.workout.exercises[exerciseIndex];
      if (!exercise) {
        console.error(`Exercise at index ${exerciseIndex} not found.`);
        return;
      }
      if (exercise.sets.length <= 1) {
        return;
      }
      exercise.sets.splice(exercise.sets.length - 1, 1);
      this.syncData();
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
        this.localWorkoutData.workout.cardio[index].heart = parseFloat(exercise.heart);
      });
      if (this.localWorkoutData.id) {
        await updateWorkout(this.localWorkoutData).then(() => {
          this.showError('', 'The training has been updated!');
        }).catch((error) => {
          if (error.response?.status === 401) {
            this.userStore.clearUser();
            this.$router.push('/about');
          }
          this.showError(error.code, error.message);
        });
      } else {
        await createWorkout(this.localWorkoutData).then(() => {
          this.showError('', 'The training has been added!');
          this.errorData = {};
          this.localWorkoutData.workout = {};
          window.location.reload();
        })
          .catch((error) => {
            if (error.response?.status === 401) {
              this.userStore.clearUser();
              window.location.replace('/about');
            }
            this.showError(error.message, error.code);
          });
      }
    },
    async deleteWorkoutById(id) {
      await deleteWorkout(id).then(() => {
        this.showError('', 'The training has been deleted');
        this.errorData = {};
        this.localWorkoutData.workout = {};
        window.location.reload();
      })
        .catch((error) => {
          if (error.response?.status === 401) {
            this.userStore.clearUser();
            window.location.replace('/about');
          }
          this.showError(error.message, error.code);
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
