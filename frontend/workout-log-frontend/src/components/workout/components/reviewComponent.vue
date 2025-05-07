<template>
  <div v-if="review.duration || review.intensity || review.slept || edit">
    <div class="row row-cols-1 row-cols-md-3 g-4">
      <!-- Duration -->
      <div class="col">
        <div class="card h-100 border-1 rounded p-4 d-flex flex-column
          justify-content-center align-items-center text-center">
          <div class="mb-2">
            <font-awesome-icon icon="fa-solid fa-stopwatch-20" class="fs-2" />
          </div>
          <div class="mb-2 w-100">
            <div v-if="edit">
              <input v-model.number="review.duration" type="number" min="0" step="1"
                class="form-control text-center" :class="{ 'is-invalid': errors.duration }"
                :placeholder="$t('review.duration')"
                @input="updateValue" />
            </div>
            <div v-else class="fs-4">
              {{ review.duration }}
            </div>
          </div>
          <div class="text-muted">
            {{ $t('review.duration') }} ({{ $t('units.min') }})
          </div>
        </div>
      </div>
      <!-- Intensity -->
      <div class="col">
        <div class="card h-100 border-1 rounded p-4 d-flex flex-column
          justify-content-center align-items-center text-center">
          <div class="mb-2">
            <font-awesome-icon icon="fa-solid fa-battery-three-quarters" class="fs-2" />
          </div>
          <div class="mb-2">
            <div>
              <button v-for="star in 10" :key="star" type="button"
                class="btn p-0" :disabled="!edit" :class="{
                'text-warning': review.intensity >= star,
                'text-muted': review.intensity < star,
              }" @click="setIntensity(star)" @keydown.enter="setIntensity(star)"
                @keydown.space.prevent="setIntensity(star)"
                style="cursor: pointer; font-size: 1.2rem; border: none;">
                <font-awesome-icon icon="fa-solid fa-star" />
              </button>
            </div>
          </div>
          <div class="text-muted">
            {{ $t('review.intensity') }}
          </div>
        </div>
      </div>
      <!-- Slept -->
      <div class="col">
        <div class="card h-100 border-1 rounded p-4 d-flex flex-column
          justify-content-center align-items-center text-center">
          <div class="mb-2">
            <font-awesome-icon icon="fa-solid fa-bed" class="fs-2" />
          </div>
          <div class="mb-2 w-100">
            <div v-if="edit">
              <input v-model.number="review.slept" type="number" min="0" step="0.1" max="24"
                class="form-control text-center" :class="{ 'is-invalid': errors.slept }"
                :placeholder="$t('review.slept')"
                @input="updateValue" />
            </div>
            <div v-else class="fs-4">
              {{ review.slept }}
            </div>
          </div>
          <div class="text-muted">
            {{ $t('review.slept') }} ({{ $t('units.hrs') }})
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';
// import card from '../../common/NumberCard.vue';

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

const emit = defineEmits(['update:modelValue', 'update:isValid']);
const review = ref({ ...props.modelValue });
const errors = ref({
  duration: false,
  intensity: false,
  slept: false,
});

watch(
  () => props.modelValue,
  (newValue) => {
    review.value = { ...newValue };
  },
);

function validateReview() {
  const result = {
    duration: !(typeof review.value.duration === 'number' && review.value.duration > 0 && review.value.duration < 480),
    intensity: !(typeof review.value.intensity === 'number' && review.value.intensity >= 1 && review.value.intensity <= 10),
    slept: !(typeof review.value.slept === 'number' && review.value.slept >= 0 && review.value.slept <= 24),
  };

  errors.value = result;

  return !result.duration && !result.intensity && !result.slept;
}

function updateValue() {
  const valid = validateReview();
  emit('update:modelValue', review.value);
  emit('update:isValid', valid);
}
function setIntensity(star) {
  review.value.intensity = star;
  updateValue();
}
</script>
