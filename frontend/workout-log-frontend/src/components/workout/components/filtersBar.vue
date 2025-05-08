<template>
  <nav class="navbar navbar-expand-md navbar-light bg-light py-3">
    <div class="container-fluid justify-content-center">
      <div class="navbar-nav flex-row flex-wrap justify-content-center gap-3">
        <div class="nav-item dropdown">
          <button class="btn btn-outline-secondary dropdown-toggle"
          type="button" id="muscleDropdown"
            data-bs-toggle="dropdown" aria-expanded="false">
            {{ t('workout.focus.header') }}
          </button>
          <ul class="dropdown-menu" aria-labelledby="muscleDropdown" style="z-index: 1150;">
            <li v-for="muscle in musclesGroup" :key="muscle">
              <a class="dropdown-item" href="#" @click.prevent="selectMuscle(muscle)">
                {{ t(`workout.focus.${muscle}`) }}
              </a>
            </li>
          </ul>
        </div>
        <div class="nav-item dropdown">
          <button class="btn btn-outline-secondary dropdown-toggle"
          type="button" id="gymTypeDropdown"
            data-bs-toggle="dropdown" aria-expanded="false">
            {{ t('workout.type.header') }}
          </button>
          <ul class="dropdown-menu" aria-labelledby="gymTypeDropdown" style="z-index: 1150;">
            <li v-for="type in gymTypes" :key="type">
              <a class="dropdown-item" href="#" @click.prevent="selectGymType(type)">
                {{ t(`workout.type.${type}`) }}
              </a>
            </li>
          </ul>
        </div>
        <div class="nav-item">
          <input type="date" class="form-control" v-model="localFilters.date"
          @change="emitFilters" />
        </div>
        <div class="nav-item">
          <button type="button" class="btn btn-outline-dark" @click="resetFilters()">
            <font-awesome-icon icon="fa-solid fa-filter-circle-xmark" />
          </button>
        </div>
      </div>
    </div>
  </nav>
  <div class="container-fluid position-sticky top-0 py-2" style="z-index: 1050;">
    <div class="d-flex flex-wrap justify-content-center">
      <span class="badge bg-secondary m-1" v-if="localFilters.gymType"
        @keydown.enter="resetFilters()" @click="resetFilters()">
      {{ t(`workout.type.${localFilters.gymType}`) }}
      </span>
      <span class="badge bg-secondary m-1" v-if="localFilters.muscle"
      @keydown.enter="resetFilters()" @click="resetFilters()">
        {{ t(`workout.focus.${localFilters.muscle}`) }}
      </span>
      <span class="badge bg-secondary m-1"
      @keydown.enter="resetFilters()" @click="resetFilters()">
        {{ localFilters.date }}
      </span>
    </div>
  </div>
</template>

<script setup>
import {
  reactive,
} from 'vue';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const emit = defineEmits(['update:filters']);

const musclesGroup = ['legs', 'biceps', 'triceps', 'back', 'chest', 'glutes'];
const gymTypes = ['recovery', 'cardio', 'hypertrophy', 'strength'];

const localFilters = reactive({
  muscle: '',
  gymType: '',
  date: '',
});

function emitFilters() {
  emit('update:filters', { ...localFilters });
}

function selectMuscle(muscle) {
  localFilters.muscle = muscle;
  emitFilters();
}

function selectGymType(type) {
  localFilters.gymType = type;
  emitFilters();
}
function resetFilters() {
  const hasFilters = Object.values(localFilters).some(Boolean);
  if (hasFilters) {
    Object.keys(localFilters).forEach((key) => {
      localFilters[key] = '';
    });
    emitFilters();
  }
}
</script>

<style scoped>
.navbar-nav>.nav-item {
  display: flex;
  align-items: center;
}
</style>
