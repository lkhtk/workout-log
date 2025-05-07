<template>
  <div class="container p-3">
    <p class="font-monospace" v-if="edit && localWorkoutData.id">
      <font-awesome-icon icon="fa-solid fa-fingerprint" />
      {{ localWorkoutData.id }}
    </p>

    <div>
      <div class="mb-4 position-relative">
        <span class="position-absolute top-0 start-0 translate-middle badge bg-secondary"
          v-if="!edit && localWorkoutData.coach">
            {{ $t('workout.instructor') }}
          </span>
        <div class="grid text-center" v-if="!edit && localWorkoutData.PublishedAt">
          <h1 class="display-2">{{ formatDate(localWorkoutData.PublishedAt) }}</h1>
        </div>
        <div class="grid mb-3" v-if="edit">
          <label for="workout-date" class="form-label fw-bold mb-2">
            {{ $t('workout.dateLabel') }}
          <input
            id="workout-date"
            type="date"
            class="form-control form-control-lg mx-auto"
            style="max-width: 300px;"
            v-model="publishedAtInput"
            required
          /></label>
        </div>
        <div v-if="edit" class="btn-group btn-group-lg" role="group" aria-label="Тип тренировки">
          <label
            for="withCoach"
            class="btn"
            :class="localWorkoutData.coach ? 'btn-dark' : 'btn-outline-dark'"
          >
            <input
              type="radio"
              class="btn-check"
              name="workoutType"
              id="withCoach"
              autocomplete="off"
              :value="true"
              v-model="localWorkoutData.coach"
            />
            {{ $t('workout.instructor') }}
          </label>
          <label
            for="withoutCoach"
            class="btn"
            :class="!localWorkoutData.coach ? 'btn-dark' : 'btn-outline-dark'"
          >
            <input
              type="radio"
              class="btn-check"
              name="workoutType"
              id="withoutCoach"
              autocomplete="off"
              :value="false"
              v-model="localWorkoutData.coach"
            />
            {{ $t('workout.alone') }}
          </label>
        </div>
      </div>
    </div>

    <FocusType
      v-if="localWorkoutData.workout.muscleGroups || edit"
      v-model="localWorkoutData.workout.muscleGroups"
      v-model:isValid="isFocusValid"
      :edit="edit"
    />
    <ExercisesComponent
      v-if="localWorkoutData.workout.exercises || edit"
      v-model="localWorkoutData.workout.exercises"
      v-model:isValid="isExercisesValid"
      :edit="edit"
    />
    <CardioComponent
      v-if="localWorkoutData.workout.cardio || edit"
      v-model="localWorkoutData.workout.cardio"
      v-model:isValid="isCardioValid"
      :edit="edit"
    />
    <ReviewComponent
      v-if="localWorkoutData.review || edit"
      v-model="localWorkoutData.review"
      v-model:isValid="isReviewValid"
      :edit="edit"
    />

    <div class="container mt-3 text-center" v-if="edit">
      <div class="btn-group" role="group" aria-label="label">
        <button
          type="button"
          class="btn btn-outline-dark btn-lg"
          @click="$router.go($router.currentRoute)"
        >
          <font-awesome-icon icon="fa-solid fa-angles-left" />
          {{ $t('buttons.back')}}
        </button>
          <button
            type="button"
            class="btn btn-outline-dark btn-lg"
            @click="handleSave"
            :disabled="!isFormValid"
          >
            <font-awesome-icon icon="fa-solid fa-floppy-disk" />
            {{ $t('buttons.submit')}}
          </button>

        <button
          type="button"
          class="btn btn-outline-danger btn-lg"
          @click="handleDelete(localWorkoutData.id)"
          :disabled="!localWorkoutData.id"
        >
          <font-awesome-icon icon="fa-solid fa-trash-can" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  ref, reactive, computed, watch,
} from 'vue';
import dayjs from 'dayjs';
import { useI18n } from 'vue-i18n';
import {
  updateWorkout,
  deleteWorkout,
  createWorkout,
} from '../../api/api';

import ExercisesComponent from './components/exercisesComponent.vue';
import ReviewComponent from './components/reviewComponent.vue';
import CardioComponent from './components/cardioComponent.vue';
import FocusType from './components/focusType.vue';

const props = defineProps({
  workoutData: { type: Object, required: true },
  edit: { type: Boolean, default: false },
});

const { t } = useI18n();

const isReviewValid = ref(true);
const isCardioValid = ref(true);
const isFocusValid = ref(true);
const isExercisesValid = ref(true);

const localWorkoutData = reactive(JSON.parse(JSON.stringify(props.workoutData)));

watch(
  () => props.workoutData,
  (newVal) => {
    Object.assign(localWorkoutData, JSON.parse(JSON.stringify(newVal)));
  },
  { immediate: true, deep: true },
);

const publishedAtInput = computed({
  get() {
    return localWorkoutData.PublishedAt
      ? dayjs(localWorkoutData.PublishedAt).format('YYYY-MM-DD')
      : '';
  },
  set(val) {
    localWorkoutData.PublishedAt = val
      ? dayjs(val, 'YYYY-MM-DD').toISOString()
      : '';
  },
});

function formatDate(date) {
  return dayjs(date).format('DD.MM.YYYY');
}

const isFormValid = computed(() => isReviewValid.value
  && isFocusValid.value
  && (isCardioValid.value || localWorkoutData.workout.cardio.length === 0)
  && (isExercisesValid.value || localWorkoutData.workout.exercises.length === 0));

async function handleSave() {
  if (localWorkoutData.workout.cardio) {
    localWorkoutData.workout.cardio = localWorkoutData.workout.cardio.map((exercise) => ({
      ...exercise,
      distance: parseFloat(exercise.distance),
      speed: parseFloat(exercise.speed),
      time: parseFloat(exercise.time),
      calories: parseFloat(exercise.calories),
      heart: parseFloat(exercise.heart),
    }));
  }

  try {
    if (localWorkoutData.id) {
      await updateWorkout(localWorkoutData);
      window.$toast?.showToast(t('info.updated'), 'success');
    } else {
      await createWorkout(localWorkoutData);
      window.$toast?.showToast(t('info.created'), 'success');
      window.location.reload();
    }
  } catch (error) {
    window.$toast?.showToast(error?.response?.data?.error || 'Error', 'danger');
  }
}

async function handleDelete(id) {
  try {
    await deleteWorkout(id);
    window.$toast?.showToast(t('info.deleted'), 'success');
    window.location.reload();
  } catch (error) {
    window.$toast?.showToast(error?.response?.data?.error || 'Error', 'danger');
  }
}
</script>
