<template>
  <p v-if="localWorkoutData.id != null && edit" class="font-monospace">
    <font-awesome-icon icon="fa-solid fa-fingerprint"/>
    {{ localWorkoutData.id }}
  </p>
  <div>
    <div v-if="edit">
      <div class="input-group input-group-lg">
          <span class="input-group-text" id="inputGroup-sizing-lg">{{ $t('workout.label') }}</span>
          <input label="rest"
            type="text" class="form-control" id="muscleGroup"
            placeholder="Legs day"
            maxlength="150"
            v-model="localWorkoutData.muscle_group"
            aria-describedby="inputGroup-sizing-lg"
            required
          />
      </div>
      <div class="form-text">
        {{ $t('workout.name_subtitle') }}
      </div>
    </div>
    <div v-else>
      <div class="grid text-center">
        <div class="g-col-3">
          <h1 class="display-2">
            {{ formatDate(localWorkoutData.PublishedAt) }}
          </h1>
        </div>
        <div class="g-col-4">
          <h1 class="display-6">
            <font-awesome-icon icon="fa-solid fa-dumbbell" />
            {{ localWorkoutData.muscle_group }}
          </h1>
        </div>
      </div>
    </div>
    <div class="form-check form-switch">
      <label class="form-check-label" for="flexSwitchCheckChecked">
      <input class="form-check-input" type="checkbox" role="switch"
        id="flexSwitchCheckChecked" :disabled="!edit"
        v-model="localWorkoutData.coach"
        aria-checked
      />
      {{ $t('workout.instructor') }}
    </label>
    </div>
    <div v-if="edit || (typeof localWorkoutData.workout.exercises !== 'undefined'
      && localWorkoutData.workout.exercises.length > 0)">
      <h1 class="display-6">
        {{ $t('workout.training') }}</h1>
      <table class="table table-hover align-middle table-bordered table-striped table-light">
        <thead >
          <tr>
            <th scope="row" colspan="3">
              <font-awesome-icon icon="fa-solid fa-medal" />
              {{ $t('workout.exercise') }}</th>
              <th
                scope="col"
                v-for="(set, index) in getMaxSets(localWorkoutData.workout.exercises)"
                :key="'set-' + index"
              >
              <font-awesome-icon icon="fa-solid fa-fire-flame-curved" />
              {{ $t('workout.set') }} {{ index + 1 }}
            </th>
          </tr>
        </thead>
        <tbody v-for="(ex, exIndex) in localWorkoutData.workout.exercises" :key="exIndex">
          <tr>
            <td rowspan="2">
              <div v-if="edit" style="border: 0px solid #ccc; width: 50px;"
                class="d-flex flex-column align-items-center">
                <font-awesome-icon icon="fa-solid fa-trash"
                  style="color: #838383; cursor: pointer;"
                  class="mb-2"
                  @keydown.enter="removeExercise(exIndex)"
                  @click="removeExercise(exIndex)"
                  />
                <font-awesome-icon icon="fa-solid fa-circle-plus"
                  style="color: #454545; cursor: pointer;"
                  class="mb-2"
                  @keydown.enter="addSet(exIndex)"
                  @click="addSet(exIndex)"
                  />
                <font-awesome-icon icon="fa-solid fa-circle-minus"
                  style="color: #454545; cursor: pointer;"
                  @keydown.enter="removeSet(exIndex)"
                  @click="removeSet(exIndex)"
                  />
              </div>
              <div class="d-flex flex-column align-items-center" v-else>
                {{ exIndex + 1 }}
              </div>
            </td>
            <td rowspan="2">
              <template v-if="edit">
                <input
                  v-model="ex.name"
                  class="form-control"
                  maxlength="150"
                  size="50"
                  required
                />
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
                      required
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
                      required
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
        <button type="button" class="btn btn-outline-dark" @click="addExercise()">
          <font-awesome-icon icon="fa-solid fa-circle-plus"/>
          {{ $t('workout.new') }}
        </button>
      </div>
    </div>
    <div v-if="edit || (typeof localWorkoutData.workout.cardio !== 'undefined'
      && localWorkoutData.workout.cardio.length > 0)">
      <h1 class="display-6"><font-awesome-icon icon="fa-solid fa-heart-pulse" />
        {{ $t('cardio.title') }}</h1>
      <table class="table table-hover align-middle table-bordered table-light">
        <thead>
          <tr>
              <th>#</th>
              <th><font-awesome-icon icon="fa-solid fa-solid fa-person-swimming" />
                {{ $t('cardio.activity') }}</th>
              <th><font-awesome-icon icon="fa-solid fa-gauge-high" />
                {{ $t('cardio.speed') }}</th>
              <th><font-awesome-icon icon="fa-solid fa-wave-square" />
                {{ $t('cardio.heart') }}</th>
              <th><font-awesome-icon icon="fa-solid fa-infinity" />
                {{ $t('cardio.distance') }}</th>
              <th><font-awesome-icon icon="fa-solid fa-stopwatch-20" />
                {{ $t('cardio.time') }}</th>
              <th><font-awesome-icon icon="fa-solid fa-fire" />
                {{ $t('cardio.calories') }}</th>
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
                <input
                  v-if="edit"
                  maxlength="150"
                  type="text"
                  v-model="cardio.type"
                  class="form-control"
                  required
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
                  required
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
                  required
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
                  required
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
                  required
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
                  required
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
          <button type="button" class="btn btn-outline-dark" @click="addCardio()">
            <font-awesome-icon icon="fa-solid fa-circle-plus"/>
            {{ $t('cardio.new') }}
          </button>
        </div>
      </div>
    </div>
    <div class="container mx-auto p-2" style="width: 200px;" v-if="edit">
      <div class="btn-group"
        role="group" aria-label="Basic mixed styles example">
        <button type="button" class="btn btn-outline-dark" @click="saveWorkout()">
          <font-awesome-icon icon="fa-solid fa-floppy-disk" />
        </button>
        <button type="button" class="btn btn-outline-dark"
          @click="deleteWorkoutById(localWorkoutData.id)">
          <font-awesome-icon icon="fa-solid fa-trash-can" />
        </button>
      </div>
    </div>
</template>

<script>
import dayjs from 'dayjs';
import { ref } from 'vue';
import { updateWorkout, deleteWorkout, createWorkout } from '../../api/api';

import { useUserStore } from '../../stores/userStore';

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
  computed: {
    maxSets() {
      return this.localWorkoutData?.workout?.exercises
        ? this.getMaxSets(this.localWorkoutData.workout.exercises)
        : 0;
    },
  },
  data() {
    let localData;
    const userStore = useUserStore();
    try {
      localData = JSON.parse(JSON.stringify(this.workoutData));
    } catch (e) {
      window.$toast?.showToast('Invalid JSON', 'danger');
      localData = {};
    }
    return {
      isAddButtonDisabled: ref(false),
      localWorkoutData: localData,
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
    formatDate(dateString) {
      return dayjs(dateString).format('DD.MM.YYYY');
    },

    addSet(exerciseIndex) {
      const exercise = this.localWorkoutData.workout.exercises[exerciseIndex];
      if (!exercise) {
        window.$toast?.showToast(`Exercise at index ${exerciseIndex} not found.`, 'danger');
        return;
      }
      if (!exercise.sets) {
        exercise.sets = [];
      }
      if (exercise.sets.length >= 10) {
        window.$toast?.showToast(this.$t('errorsMsg.setsLimit'), 'info');
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
        window.$toast?.showToast(this.$t('errorsMsg.exercisesLimit'), 'info');
        return;
      }
      const emptySets = [];
      for (let i = 0; i <= 2; i += 1) {
        emptySets.push({});
      }
      this.localWorkoutData.workout.exercises.push({
        name: '',
        sets: emptySets,
      });
      this.syncData();
    },

    addCardio() {
      if (this.localWorkoutData.workout.cardio.length >= 10) {
        window.$toast?.showToast(this.$t('errorsMsg.cardiosLimit'), 'info');
        return;
      }
      this.localWorkoutData.workout.cardio.push({});
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
        window.$toast?.showToast(`Exercise at index ${exerciseIndex} not found.`, 'danger');
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
          window.$toast?.showToast(this.$t('info.updated'), 'success');
        }).catch((error) => {
          window.$toast?.showToast(error.message, 'danger');
        });
      } else {
        await createWorkout(this.localWorkoutData).then(() => {
          window.$toast?.showToast(this.$t('info.created'), 'success');
          this.errorData = {};
          this.localWorkoutData.workout = {};
          window.location.reload();
        })
          .catch((error) => {
            window.$toast?.showToast(error.message, 'danger');
          });
      }
    },
    async deleteWorkoutById(id) {
      await deleteWorkout(id).then(() => {
        window.$toast?.showToast(this.$t('info.deleted'), 'success');
        this.errorData = {};
        this.localWorkoutData.workout = {};
        window.location.reload();
      })
        .catch((error) => {
          window.$toast?.showToast(error.message, 'danger');
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
