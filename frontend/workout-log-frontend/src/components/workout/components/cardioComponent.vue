<template>
  <div class="mb-5" v-if="cardios.length>0 || edit">
    <h1 class="display-6">
      <font-awesome-icon icon="fa-solid fa-heart-pulse" />
      {{ $t('cardio.title') }}
    </h1>
    <table class="table table-hover align-middle table-bordered table-striped table-light">
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
      <tbody v-for="cardio, index in cardios" v-bind:key="cardio">
        <tr>
          <td v-if="edit">
            <button type="button" class="btn-close"
            aria-label="Close" @click="removeCardio(exIndex)" tabindex="-1"></button>
          </td>
          <td v-else>{{ index + 1 }}</td>
          <td v-if="edit">
            <input
            v-model="cardio.type"
            type="text"
            maxlength="150"
            class="form-control"
            :readonly="!edit"
            :placeholder="$t('cardio.activity')"
            required
            @input="updateValue"
          /></td>
          <td v-else>
            {{cardio.type}}
          </td>
          <td v-if="edit">
            <input
            v-model.number="cardio.speed"
            type="number"
            min="0"
            step="0.1"
            maxlength="6"
            required
            class="form-control"
            :placeholder="$t('cardio.speed')"
            @input="updateValue"
          /></td>
          <td v-else>
            {{cardio.speed}}
          </td>
          <td v-if="edit">
            <input
            v-model.number="cardio.heart"
            type="number"
            min="60"
            maxlength="3"
            step="1"
            required
            class="form-control"
            :readonly="!edit"
            :placeholder="$t('cardio.heart')"
            @input="updateValue"
            />
          </td>
          <td v-else>
            {{cardio.heart}}
          </td>
          <td v-if="edit">
            <input
            v-model.number="cardio.distance"
            type="number"
            min="0"
            required
            maxlength="6"
            step="0.1"
            class="form-control"
            :readonly="!edit"
            :placeholder="$t('cardio.distance')"
            @input="updateValue"
            />
          </td>
          <td v-else>
            {{cardio.distance}}
          </td>
          <td v-if="edit">
            <input
            v-model.number="cardio.time"
            type="number"
            min="0"
            required
            maxlength="6"
            class="form-control"
            :readonly="!edit"
            :placeholder="$t('cardio.time')"
            @input="updateValue"
            />
          </td>
          <td v-else>
            {{cardio.time}}
          </td>
          <td v-if="edit">
            <input
            v-model.number="cardio.calories"
            type="number"
            min="0"
            maxlength="6"
            required
            class="form-control"
            :readonly="!edit"
            :placeholder="$t('cardio.calories')"
            @input="updateValue"
            />
          </td>
          <td v-else>
            {{cardio.calories}}
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
const cardios = ref(props.modelValue);

watch(() => props.modelValue, (newValue) => {
  cardios.value = newValue;
});

function updateValue() {
  emit('update:modelValue', cardios.value);
}

function addCardio() {
  if (cardios.value.length >= 10) {
    window.$toast?.showToast(this.$t('errorsMsg.cardiosLimit'), 'info');
    return;
  }
  cardios.value.push({});
}

function removeCardio(index) {
  cardios.value.splice(index, 1);
}
</script>
