<template>
  <div v-if="edit" class="mb-5">
    <div class="card shadow-sm rounded p-4 mb-2">
      <h2 class="display-6 mb-4 d-flex align-items-center gap-2">
        <font-awesome-icon icon="fa-solid fa-dumbbell" />
        {{ $t('workout.focus.header') }}
      </h2>
      <div class="d-flex flex-wrap gap-2">
        <label
          v-for="option in musclesGroup"
          :key="option"
          :for="`focus-${option}`"
          class="btn"
          :class="selectedMuscles.includes(option) ? 'btn-dark' : 'btn-outline-dark'"
        >
          <input
            type="checkbox"
            class="btn-check"
            :id="`focus-${option}`"
            :value="option"
            v-model="selectedMuscles"
            autocomplete="off"
          />
          {{ $t(`workout.focus.${option}`) }}
        </label>
      </div>
      <hr class="my-4" />
      <h2 class="display-6 mb-4 d-flex align-items-center gap-2">
        <font-awesome-icon icon="fa-solid fa-bolt" />
        {{ $t('workout.type.header') }}
      </h2>
      <div class="d-flex flex-wrap gap-2">
        <label
          v-for="option in gymTypes"
          :key="option"
          :for="`type-${option}`"
          class="btn"
          :class="selectedGymType === option ? 'btn-dark' : 'btn-outline-dark'"
        >
          <input
            type="radio"
            class="btn-check"
            :id="`type-${option}`"
            :value="option"
            v-model="selectedGymType"
            autocomplete="off"
          />
          {{ $t(`workout.type.${option}`) }}
        </label>
      </div>
    </div>
  </div>
  <div v-else class="d-flex flex-wrap gap-2 mt-3 align-items-center">
  <div class="d-flex flex-wrap gap-2">
    <h1 class="lead"
        v-for="option in selectedMuscles"
        :key="option">
      <font-awesome-icon icon="fa-solid fa-chevron-right" />
      {{ $t(`workout.focus.${option}`) }}
    </h1>
  </div>

  <span
    v-if="selectedGymType"
    class="badge bg-warning text-dark d-flex align-items-center ms-auto">
    <font-awesome-icon icon="fa-solid fa-bolt" class="me-1" />
    {{ $t(`workout.type.${selectedGymType}`) }}
  </span>
</div>

</template>

<script setup>
import { ref, watch, computed } from 'vue';

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({ muscles: [], type: '' }),
  },
  edit: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(['update:modelValue', 'update:isValid']);

const musclesGroup = ['legs', 'biceps', 'triceps', 'back', 'chest'];
const gymTypes = ['recovery', 'cardio', 'hypertrophy', 'strength'];

const selectedMuscles = ref([]);
const selectedGymType = ref('');

const initialized = ref(false);
watch(
  () => props.modelValue,
  (newVal) => {
    if (initialized.value) return;
    if (newVal && typeof newVal === 'object') {
      selectedMuscles.value = Array.isArray(newVal.muscles) ? [...newVal.muscles] : [];
      selectedGymType.value = newVal.type || '';
      initialized.value = true;
    }
  },
  { immediate: true },
);

const isValid = computed(() => selectedMuscles.value.length > 0 && !!selectedGymType.value);

watch(
  [selectedMuscles, selectedGymType],
  ([muscles, type]) => {
    const changed = JSON.stringify(muscles)
      !== JSON.stringify(props.modelValue.muscles)
      || type !== props.modelValue.type;
    if (changed) {
      emit('update:modelValue', { muscles, type });
    }

    emit('update:isValid', isValid.value);
  },
  { deep: true },
);
</script>
