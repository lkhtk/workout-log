<template>
  <div v-if="edit" class="mb-5">
    <div class="card shadow-sm rounded p-4 mb-2">
      <h2 class="display-6 mb-4 d-flex align-items-center gap-2">
        <font-awesome-icon icon="fa-solid fa-arrows-to-circle" />
        {{ $t('workout.focus.header') }}
      </h2>

      <div class="d-flex flex-wrap gap-2">
        <label
          v-for="option in options"
          :key="option"
          :for="`focus-${option}`"
          class="btn"
          :class="selectedValues.includes(option) ? 'btn-dark' : 'btn-outline-dark'"
        >
          <input
            type="checkbox"
            class="btn-check"
            :id="`focus-${option}`"
            :value="option"
            v-model="selectedValues"
            autocomplete="off"
          />
          {{ $t(`workout.focus.${option}`) }}
        </label>
      </div>
    </div>
  </div>
  <div v-else class="d-flex flex-wrap gap-2 mt-3">
    <h1 class="lead"
      v-for="option in selectedValues"
      :key="option">
      <font-awesome-icon icon="fa-solid fa-chevron-right" />
      {{ $t(`workout.focus.${option}`) }}
    </h1>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';

const props = defineProps({
  modelValue: {
    type: Array,
    default: () => [],
  },
  edit: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(['update:modelValue', 'update:isValid']);
const options = ['legs', 'biceps', 'triceps', 'back', 'chest'];
const selectedValues = props.modelValue ? ref([...props.modelValue]) : ref([]);
watch(
  selectedValues,
  (newValues, oldValues) => {
    const changed = newValues.length !== oldValues.length
      || newValues.some((val, index) => val !== oldValues[index]);

    if (changed) {
      emit('update:modelValue', newValues);
      emit('update:isValid', newValues.length > 0);
    }
  },
  { deep: true },
);
watch(
  () => props.modelValue,
  (newValue) => {
    if (Array.isArray(newValue)) {
      selectedValues.value = [...newValue];
    } else {
      selectedValues.value = [];
    }
  },
  { immediate: true },
);
</script>
