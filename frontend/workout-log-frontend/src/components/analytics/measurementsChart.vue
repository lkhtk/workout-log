<template>
  <div class="d-flex h-100 text-center">
    <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column card shadow">
      <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
        <div class="spinner-border align-items-center" role="status" v-if="isLoading"></div>
        <h1 class="display-5" v-if="isLoading">{{ $t('info.loading') }}</h1>
        <h1 class="display-5" v-else>
          <font-awesome-icon icon="fa-solid fa-arrow-trend-up" />
          {{ $t('trends.title') }}
        </h1>
      </div>
      <main v-if="!isLoading">
        <div class="container my-4">
          <div v-if="hasNoData">
            <p>{{ $t('errorsMsg.noDataAvailable') }}</p>
          </div>
          <div v-else-if="chartSeries.length">
            <ApexCharts type="line" :options="chartOptions" :series="chartSeries" height="400" />
          </div>
          <div v-else>
            <p>{{ $t('errorsMsg.noDataAvailable') }}</p>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import ApexCharts from 'vue3-apexcharts';
import { getMeasurements } from '../../api/api';

const measurements = ref([]);
const isLoading = ref(true);
const hasNoData = ref(false);
const { t } = useI18n();

const fetchMeasurements = async () => {
  try {
    const response = await getMeasurements();
    if (!response.data || response.data.total === 0) {
      hasNoData.value = true;
    } else {
      measurements.value = response.data.data;
    }
  } catch (error) {
    window.$toast?.showToast(error.message);
  } finally {
    isLoading.value = false;
  }
};

onMounted(fetchMeasurements);

const measurementFields = [
  { key: 'bodyFat', color: '#03071E' },
  { key: 'bodyWeight', color: '#370617' },
  { key: 'neck', color: '#6A040F' },
  { key: 'chest', color: '#9D0208' },
  { key: 'waist', color: '#D00000' },
  { key: 'upperarm', color: '#DC2F02' },
  { key: 'hips', color: '#E85D04' },
  { key: 'forearm', color: '#F48C06' },
  { key: 'thighs', color: '#FAA307' },
  { key: 'calves', color: '#FFBA08' },
];

const chartSeries = computed(() => {
  if (!measurements.value.length) return [];

  return measurementFields.map(({ key, color }) => {
    const seriesData = measurements.value
      .map((m) => ({ x: new Date(m.measurementDate), y: m[key] }))
      .filter((entry) => entry.y !== 0);
    return seriesData.length > 1
      ? { name: t(`measurements.${key}`), data: seriesData, color }
      : null;
  }).filter((series) => series !== null);
});

const chartOptions = computed(() => ({
  chart: {
    type: 'line',
    height: 400,
    toolbar: { show: false },
  },
  xaxis: {
    type: 'datetime',
    labels: { format: 'dd MMM' },
  },
  stroke: {
    width: 2,
    curve: 'smooth',
  },
  markers: { size: 0 },
  legend: {
    position: 'right',
    verticalAlign: 'middle',
  },
  tooltip: { x: { format: 'dd MMM yyyy' } },
}));
</script>

<style scoped>
canvas {
  max-width: 100%;
}
</style>
