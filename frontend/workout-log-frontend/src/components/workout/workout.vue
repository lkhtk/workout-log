<template>
  <div class="container p-3 shadow">
    <p class="font-monospace" v-if="edit && localWorkoutData.id">
      <font-awesome-icon icon="fa-solid fa-fingerprint" />
      {{ localWorkoutData.id }}
    </p>
    <div>
      <div class="mb-3">
        <input v-if="edit" v-model="localWorkoutData.muscle_group" type="text"
          class="form-control" maxlength="150"
          :placeholder="$t('workout.name_subtitle')" required />
        <div v-else class="grid text-center">
          <h1 class="display-2">
            {{ formatDate(localWorkoutData.PublishedAt) }}
          </h1>
          <h1 class="display-6">
            <font-awesome-icon icon="fa-solid fa-dumbbell" />
            {{ localWorkoutData.muscle_group }}
          </h1>
        </div>
      </div>
      <div class="mb-4">
        <div class="btn-group" role="group" v-if="edit">
          <label :class="(localWorkoutData.coach) ? 'btn btn-primary' : 'btn btn-outline-primary'"
            for="withCoach">
            <input type="checkbox" id="withCoach" class="btn-check" role="switch" aria-checked
              v-model="localWorkoutData.coach" />
            {{ $t('workout.instructor') }}
          </label>
          <label :class="(!localWorkoutData.coach) ? 'btn btn-primary' : 'btn btn-outline-primary'"
            for="withCoach">
            <input type="checkbox" id="withCoach" class="btn-check" role="switch" aria-checked />
            {{ $t('workout.alone') }}
          </label>
        </div>
        <div v-else>
          <span class="badge rounded-pill text-bg-secondary" v-if="localWorkoutData.coach">
            {{ $t('workout.instructor') }}
          </span>
        </div>
      </div>
    </div>
    <exercisesComponent v-model="localWorkoutData.workout.exercises" :edit="edit" />
    <cardioComponent v-model="localWorkoutData.workout.cardio" :edit="edit"/>
    <reviewComponent v-model="localWorkoutData.review" :edit="edit"/>
    <!-- Buttons -->
    <div class="container mx-auto p-2" style="width: 200px;" v-if="edit">
      <div class="btn-group" role="group" aria-label="Basic mixed styles example">
        <button type="button" class="btn btn-outline-dark" @click="saveWorkout()">
          <font-awesome-icon icon="fa-solid fa-floppy-disk" />
        </button>
        <button type="button" class="btn btn-outline-dark"
          @click="deleteWorkoutById(localWorkoutData.id)"
          :disabled="localWorkoutData.id == null">
          <font-awesome-icon icon="fa-solid fa-trash-can" />
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs';
import { ref } from 'vue';
import { updateWorkout, deleteWorkout, createWorkout } from '../../api/api';
import ExercisesComponent from './components/exercisesComponent.vue';
import ReviewComponent from './components/reviewComponent.vue';
import CardioComponent from './components/cardioComponent.vue';
import { useUserStore } from '../../stores/userStore';

export default {
  name: 'WorkoutComponent',
  components: {
    exercisesComponent: ExercisesComponent,
    reviewComponent: ReviewComponent,
    cardioComponent: CardioComponent,
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
  data() {
    let localData;
    const userStore = useUserStore();
    try {
      localData = JSON.parse(JSON.stringify(this.workoutData));
    } catch (e) {
      localData = {};
    }
    return {
      isAddButtonDisabled: ref(false),
      localWorkoutData: localData,
      userStore,
    };
  },
  methods: {
    formatDate(dateString) {
      return dayjs(dateString).format('DD.MM.YYYY');
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
          this.localWorkoutData = {};
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
        this.localWorkoutData = {};
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
