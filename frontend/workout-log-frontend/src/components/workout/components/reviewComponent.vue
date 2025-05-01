<template>
  <div v-if="review.duration || review.intensity || review.slept || edit" class="container my-4">
    <div class="row row-cols-1 row-cols-md-3 g-4">
      <!-- Duration -->
      <card v-if="!edit && review.duration"
        :cardData="{
          icon: 'fa-solid fa-stopwatch-20',
          count: review.duration,
          text: $t('review.duration') + ' (' + $t('units.min') + ')',
        }"
      />
      <div v-else>
        <div class="text-muted mb-2">
          {{ $t('review.duration') }} ({{ $t('units.min') }})
        </div>
        <input v-model.number="review.duration" type="number" min="0"
          step="1" class="form-control text-center"
          :placeholder="$t('review.duration')" @input="updateValue" />
      </div>
      <!-- Intensity -->
      <div class="col" v-if="edit || review.intensity">
        <div class="card h-100 border-1 rounded p-4 d-flex justify-content-center
  align-items-center text-center">
          <div>
            <div class="mb-2">
              <font-awesome-icon icon='fa-solid fa-battery-three-quarters' class="fs-2" />
            </div>
              <button v-for="star in 10" :key="star" :disabled="!edit"
              type="button" class="btn p-0" :class="{
                'text-warning': review.intensity >= star,
                'text-muted': review.intensity < star
              }" @click="review.intensity = star"
                @keydown.enter="review.intensity = star"
                @keydown.space.prevent="review.intensity = star"
                style="cursor: pointer; font-size: 1.2rem; border: none;">
                <font-awesome-icon icon="fa-solid fa-star" />
              </button>
          </div>
          <div class="text-muted mb-2">
            {{ $t('review.intensity') }}
          </div>
        </div>
      </div>
      <!-- Sleep -->
      <card  v-if="!edit && review.slept"
      :cardData="{
        count: review.slept,
        text: $t('review.slept') + ' (' + $t('units.hrs') + ')',
        icon: 'fa-solid fa-bed',
      }"/>
      <div class="col" v-else-if="edit">
        <div class="text-muted mb-2">
          {{ $t('review.slept') }} ({{ $t('units.hrs') }})
        </div>
        <input v-model.number="review.slept" type="number" min="0"
          step="0.1" max="24" class="form-control text-center"
          :placeholder="$t('review.slept')" @input="updateValue" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';
import card from '../../common/NumberCard.vue';

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
