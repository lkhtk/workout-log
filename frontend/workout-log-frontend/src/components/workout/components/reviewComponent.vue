<template>
  <div v-if="review.duration || review.intensity || review.slept || edit">
    <h1 class="display-6">
      <font-awesome-icon icon="fa-solid fa-square-poll-vertical" />
      {{ $t('review.title') }}
    </h1>
    <div class="row g-4 py-5 row-cols-1 row-cols-lg-3">
      <div class="feature col">
        <p class="text-body-emphasis">
          <font-awesome-icon icon="fa-solid fa-stopwatch" />
          {{ $t('review.duration') }}
        </p>
        <input v-if="edit"
          v-model.number="review.duration"
          type="number"
          class="form-control"
          :placeholder="$t('review.duration')"
          @input="updateValue"
        />
        <div v-else-if="review.duration">
          <div class="feature-icon d-inline-flex align-items-center
          justify-content-center">
            {{ review.duration }} {{ $t('units.min') }}
          </div>
        </div>
      </div>
      <div class="feature col" v-if="edit || review.duration">
        <p class="text-body-emphasis">
          <font-awesome-icon icon="fa-solid fa-rocket" />
          {{ $t('review.intensity') }}
        </p>
        <button
          v-for="star in 10"
          id="start"
          :key="star"
          :disabled="!edit"
          type="button"
          class="btn p-0"
          :class="{ 'text-warning': review.intensity >= star }"
          @click="review.intensity = star"
          @keydown.enter="review.intensity = star"
          @keydown.space.prevent="review.intensity = star"
          style="cursor: pointer; font-size: 0.7rem; border: 0;">
          <font-awesome-icon icon="fa-solid fa-star" />
        </button>
      </div>
      <div class="feature col">
        <p class="text-body-emphasis">
          <font-awesome-icon icon="fa-solid fa-bed" />
          {{ $t('review.slept') }}
        </p>
        <input v-if="edit"
          v-model.number="review.slept"
          type="number"
          class="form-control"
          :placeholder="$t('review.slept')"
          @input="updateValue"
        />
        <div v-else-if="review.slept"
        class="feature-icon d-inline-flex align-items-center
        justify-content-center">
          {{ review.slept }} {{ $t('units.hrs') }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';

const props = defineProps({
  modelValue: {
    type: Object,
    required: true,
  },
  edit: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(['update:modelValue']);
const review = ref(props.modelValue);

watch(() => props.modelValue, (newValue) => {
  review.value = newValue;
});

function updateValue() {
  emit('update:modelValue', review.value);
}
</script>
