<template>
  <div v-if="(cardios && cardios.length > 0) || edit" class="mb-5">
    <div class="card shadow-sm rounded p-4">
      <h2 class="display-6 mb-4 d-flex align-items-center gap-2">
        <font-awesome-icon icon="fa-solid fa-heart-pulse" />
        {{ $t('cardio.title') }}
      </h2>

      <div class="table-responsive">
        <table class="table table-bordered table-striped align-middle text-center">
          <thead class="table-light">
            <tr>
              <th>#</th>
              <th>
                <font-awesome-icon icon="fa-solid fa-person-swimming" />
                {{ $t('cardio.activity') }}
              </th>
              <th>
                <font-awesome-icon icon="fa-solid fa-gauge-high" />
                {{ $t('cardio.speed') }}
              </th>
              <th>
                <font-awesome-icon icon="fa-solid fa-wave-square" />
                {{ $t('cardio.heart') }}
              </th>
              <th>
                <font-awesome-icon icon="fa-solid fa-infinity" />
                {{ $t('cardio.distance') }}
              </th>
              <th>
                <font-awesome-icon icon="fa-solid fa-stopwatch-20" />
                {{ $t('cardio.time') }}
              </th>
              <th>
                <font-awesome-icon icon="fa-solid fa-fire" />
                {{ $t('cardio.calories') }}
              </th>
              <th v-if="edit" />
            </tr>
          </thead>
          <tbody>
            <tr v-for="(cardio, index) in cardios" :key="index">
              <td>{{ index + 1 }}</td>
              <td>
                <input v-if="edit" v-model="cardio.type" type="text"
                  maxlength="150" class="form-control"
                  :class="{ 'is-invalid': cardio.errors?.type }"
                  :placeholder="$t('cardio.activity')" required
                  @input="updateValue" />
                <span v-else>{{ cardio.type }}</span>
              </td>

              <td>
                <input v-if="edit" v-model.number="cardio.speed"
                type="number" min="0" step="0.1" maxlength="6"
                  class="form-control" :class="{ 'is-invalid': cardio.errors?.speed }"
                  :placeholder="$t('cardio.speed')"
                  required @input="updateValue" />
                <span v-else>{{ cardio.speed }}</span>
              </td>

              <td>
                <input v-if="edit" v-model.number="cardio.heart"
                type="number" min="60" max="220" class="form-control"
                  :class="{ 'is-invalid': cardio.errors?.heart }"
                  :placeholder="$t('cardio.heart')" required
                  @input="updateValue" />
                <span v-else>{{ cardio.heart }}</span>
              </td>
              <td>
                <input v-if="edit" v-model.number="cardio.distance" type="number"
                min="0" step="0.1" maxlength="6"
                  class="form-control" :class="{ 'is-invalid': cardio.errors?.distance }"
                  :placeholder="$t('cardio.distance')" required @input="updateValue" />
                <span v-else>{{ cardio.distance }}</span>
              </td>

              <td>
                <input v-if="edit" v-model.number="cardio.time"
                type="number" min="0" maxlength="6" class="form-control"
                  :class="{ 'is-invalid': cardio.errors?.time }"
                  :placeholder="$t('cardio.time')" required
                  @input="updateValue" />
                <span v-else>{{ cardio.time }}</span>
              </td>

              <td>
                <input v-if="edit" v-model.number="cardio.calories"
                type="number" min="0" maxlength="6"
                  class="form-control" :class="{ 'is-invalid': cardio.errors?.calories }"
                  :placeholder="$t('cardio.calories')" required @input="updateValue" />
                <span v-else>{{ cardio.calories }}</span>
              </td>

              <td v-if="edit">
                <button type="button" class="btn btn-sm btn-outline-danger" aria-label="Remove"
                  @click="removeCardio(index)">
                  <font-awesome-icon icon="fa-solid fa-trash" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="edit" class="mt-3 text-center">
        <button type="button" class="btn btn-outline-dark" @click="addCardio">
          <font-awesome-icon icon="fa-solid fa-circle-plus" />
          {{ $t('cardio.new') }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
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

const cardios = ref(props.modelValue);
const emit = defineEmits(['update:modelValue', 'update:isValid']);

watch(() => props.modelValue, (newValue) => {
  cardios.value = newValue;
});

function addCardio() {
  if (cardios.value.length >= 10) {
    window.$toast?.showToast(t('errorsMsg.cardiosLimit'), 'dark');
    return;
  }
  cardios.value.push({});
}

function removeCardio(index) {
  cardios.value.splice(index, 1);
}

function validateCardios() {
  let isValid = true;
  const updated = cardios.value.map((entry) => {
    const newEntry = { ...entry, errors: {} };

    if (!newEntry.type?.trim() || newEntry.type.length < 3) {
      newEntry.errors.type = true;
      isValid = false;
    }

    if (typeof newEntry.speed !== 'number' || newEntry.speed <= 0) {
      newEntry.errors.speed = true;
      isValid = false;
    }

    if (typeof newEntry.heart !== 'number' || newEntry.heart < 60 || newEntry.heart > 220) {
      newEntry.errors.heart = true;
      isValid = false;
    }

    if (typeof newEntry.distance !== 'number' || newEntry.distance <= 0) {
      newEntry.errors.distance = true;
      isValid = false;
    }

    if (typeof newEntry.time !== 'number' || newEntry.time <= 0) {
      newEntry.errors.time = true;
      isValid = false;
    }

    if (typeof newEntry.calories !== 'number' || newEntry.calories <= 0) {
      newEntry.errors.calories = true;
      isValid = false;
    }

    return newEntry;
  });

  cardios.value = updated;
  return isValid;
}

function updateValue() {
  const valid = validateCardios();
  emit('update:modelValue', cardios.value);
  emit('update:isValid', valid);
}

</script>
