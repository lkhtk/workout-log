<template>
  <div v-if="exercises.length > 0 || edit">
    <div class="card shadow-sm rounded p-4 mb-2">
      <h2 class="display-6 mb-4 d-flex align-items-center gap-2">
        <font-awesome-icon icon="fa-solid fa-anchor" />
        {{ $t('workout.training') }}
      </h2>
      <div class="table-responsive">
        <table class="table table-bordered table-striped align-middle text-center">
          <thead>
            <tr>
              <th scope="row" colspan="3">
                <font-awesome-icon icon="fa-solid fa-medal" />
                {{ $t('workout.exercise') }}
              </th>
              <th scope="col" v-for="(set, index) in getMaxSets(exercises)" :key="'set-' + index">
                <font-awesome-icon icon="fa-solid fa-fire-flame-curved" />
                {{ $t('workout.set') }} {{ index + 1 }}
              </th>
            </tr>
          </thead>
          <tbody v-for="(ex, exIndex) in exercises" :key="exIndex">
            <tr>
              <td rowspan="2" v-if="edit">
                <div v-if="edit" style="border: 0px solid #ccc; width: 50px;"
                  class="d-flex flex-column align-items-center">
                  <font-awesome-icon icon="fa-solid fa-trash"
                    class="btn btn-sm btn-outline-danger mb-2"
                    style="cursor: pointer;"
                    @keydown.enter="removeExercise(exIndex)"
                    @click="removeExercise(exIndex)" />
                  <font-awesome-icon icon="fa-solid fa-circle-plus" style="color: #454545;
                    cursor: pointer;" class="mb-2" @keydown.enter="addSet(exIndex)"
                    @click="addSet(exIndex)" />
                  <font-awesome-icon icon="fa-solid fa-circle-minus"
                    style="color: #454545; cursor: pointer;"
                    @keydown.enter="removeSet(exIndex)" @click="removeSet(exIndex)" />
                </div>
              </td>
              <td rowspan="2" v-else>
                {{ exIndex + 1 }}
              </td>
              <td rowspan="2" v-if="edit">
                <input type="text"
                  class="form-control mb-2"
                  :class="{ 'is-invalid': ex.errors?.name }"
                  :placeholder="$t('workout.default_ex')" v-model="ex.name"
                  maxlength="150" size="50" required
                  @input="updateValue"/>
              </td>
              <td rowspan="2" v-else>
                {{ ex.name }}
              </td>
              <td><font-awesome-icon icon="fa-solid fa-weight-hanging" /></td>
              <td v-for="(set, setIndex) in ex.sets" :key="'weight-' + setIndex">
                <div class="col-7" v-if="edit">
                  <label for="weight">
                    <input id="weight" :placeholder="$t('workout.weight')"
                      v-model="set.weight" class="form-control col-sm-10"
                      :class="{ 'is-invalid': ex.errors?.weight[setIndex] }"
                      type="number" min="0" step="0.1" max="999" required
                      @input="updateValue"/>
                  </label>
                </div>
                <span v-else>{{ set.weight }}</span>
              </td>
            </tr>
            <tr>
              <td><font-awesome-icon icon="fa-solid fa-repeat" /></td>
              <td v-for="(set, setIndex) in ex.sets" :key="'reps-' + setIndex">
                <div class="col-7" v-if="edit">
                  <label for="reps">
                    <input v-model="set.reps" :placeholder="$t('workout.reps')"
                      max="99" class="form-control col-sm-10"
                      :class="{ 'is-invalid': ex.errors?.reps[setIndex] }"
                      type="number" min="0" required
                      @input="updateValue"/>
                  </label>
                </div>
                <span v-else>{{ set.reps }}</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-if="edit" class="mt-3 text-center">
        <button
          type="button"
          class="btn btn-outline-dark"
          @click="addExercise()">
          <font-awesome-icon icon="fa-solid fa-circle-plus" />
        {{ $t('workout.new') }}
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

const emit = defineEmits(['update:modelValue']);
const exercises = ref(props.modelValue);

watch(() => props.modelValue, (newValue) => {
  exercises.value = newValue;
});
function getMaxSets(exSet) {
  if (!exSet || !Array.isArray(exSet)) {
    return 0;
  }
  return exSet.reduce((max, ex) => {
    if (ex.sets && Array.isArray(ex.sets)) {
      return Math.max(max, ex.sets.length);
    }
    return max;
  }, 0);
}

function syncData() {
  emit('update:modelValue', exercises.value);
}
function addSet(exerciseIndex) {
  const exercise = exercises.value[exerciseIndex];
  if (!exercise) {
    window.$toast?.showToast(`Exercise at index ${exerciseIndex} not found.`, 'danger');
    return;
  }
  if (!exercise.sets) {
    exercise.sets = [];
  }
  if (exercise.sets.length >= 10) {
    window.$toast?.showToast(t('errorsMsg.setsLimit'), 'dark');
    return;
  }
  exercise.sets.push({
    reps: 0,
    weight: 0,
  });
  syncData();
}
function addExercise() {
  if (exercises.value.length >= 10) {
    window.$toast?.showToast(t('errorsMsg.exercisesLimit'), 'info');
    return;
  }
  const emptySets = [];
  for (let i = 0; i <= 2; i += 1) {
    emptySets.push({});
  }
  exercises.value.push({
    name: '',
    sets: emptySets,
  });
  syncData();
}

function removeExercise(index) {
  exercises.value.splice(index, 1);
  syncData();
}

function removeSet(exerciseIndex) {
  const exercise = exercises.value[exerciseIndex];
  if (!exercise) {
    window.$toast?.showToast(`Exercise at index ${exerciseIndex} not found.`, 'danger');
    return;
  }
  if (exercise.sets.length <= 1) {
    return;
  }
  exercise.sets.splice(exercise.sets.length - 1, 1);
  syncData();
}

function validateEx() {
  let isValid = true;

  const updated = exercises.value.map((entry) => {
    const newEntry = {
      ...entry,
      errors: {
        name: false,
        reps: [],
        weight: [],
      },
    };

    if (!newEntry.name?.trim() || newEntry.name.length < 3) {
      newEntry.errors.name = true;
      isValid = false;
    }

    newEntry.sets = newEntry.sets.map((set, index) => {
      const reps = Number(set.reps) || 0;
      const weight = Number(set.weight) || 0;
      if (!(typeof reps === 'number' && reps > 0 && reps < 99)) {
        newEntry.errors.reps[index] = true;
        isValid = false;
      } else {
        newEntry.errors.reps[index] = false;
      }
      if (!(typeof weight === 'number' && weight >= 0 && weight < 999)) {
        newEntry.errors.weight[index] = true;
        isValid = false;
      } else {
        newEntry.errors.weight[index] = false;
      }

      return {
        ...set,
        reps,
        weight,
      };
    });

    return newEntry;
  });

  exercises.value = updated;
  return isValid;
}

function updateValue() {
  validateEx();
  emit('update:modelValue', exercises.value);
}
</script>
