<template>
  <div class="d-flex h-100 text-center" v-if="!user">
    <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
      <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
          <h1 class="display-5" v-if="isLoading">{{ $t('info.loading') }}</h1>
          <h1 class="display-5">
            <font-awesome-icon icon="fa-solid fa-chart-column" />
            {{ $t('trends.title') }}
          </h1>
      </div>
      <main>
        <div class="container my-4">
          <div v-if="!isLoading && chartData && chartData.labels.length">
            <LineChart :data="chartData" :options="chartOptions" />
          </div>
          <div v-else>
            <p>{{ $t('errorsMsg.noworkouts') }}</p>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import { Line } from 'vue-chartjs';
import {
  Chart,
  LineElement,
  PointElement,
  LinearScale,
  Title,
  CategoryScale,
  Legend,
  Tooltip,
} from 'chart.js';
import { getMeasurements } from '../../api/api';

Chart.register(LineElement, PointElement, LinearScale, Title, CategoryScale, Legend, Tooltip);

const measurements = ref([]);
const isLoading = ref(true);
const { t } = useI18n();

const fetchMeasurements = async () => {
  try {
    const response = await getMeasurements();
    measurements.value = response.data.data;
  } catch (error) {
    console.error('Error fetching measurements:', error);
    if (error.response?.status === 401) {
      localStorage.removeItem('user');
      window.location.replace('/about');
    }
  } finally {
    isLoading.value = false;
  }
};

onMounted(fetchMeasurements);

const measurementFields = [
  { key: 'body_fat', color: 'rgba(3, 7, 30, 1)' },
  { key: 'body_weight', color: 'rgba(55, 6, 23, 1)' },
  { key: 'neck', color: 'rgba(106, 4, 15, 1)' },
  { key: 'chest', color: 'rgba(157, 2, 8, 1)' },
  { key: 'waist', color: 'rgba(208, 0, 0, 1)' },
  { key: 'upperarm', color: 'rgba(220, 47, 2, 1)' },
  { key: 'hips', color: 'rgba(232, 93, 4, 1)' },
  { key: 'forearm', color: 'rgba(244, 140, 6, 1)' },
  { key: 'thighs', color: 'rgba(250, 163, 7, 1)' },
  { key: 'calves', color: 'rgba(255, 186, 8, 1)' },
];

const chartData = computed(() => {
  if (!measurements.value.length) return null;

  const labels = measurements.value.map((m) => new Date(m.measurement_date).toLocaleDateString());

  const datasets = measurementFields.map(({ key, color }) => ({
    label: t(`measurements.${key}`),
    data: measurements.value.map((m) => m[key]),
    borderColor: color,
    backgroundColor: color.replace('1)', '0.2)'),
    tension: 0.4,
  }));

  return { labels, datasets };
});

const chartOptions = computed(() => ({
  responsive: true,
  plugins: {
    legend: { position: 'top' },
    title: { display: true, text: t('trends.chartTitle') },
  },
}));

const LineChart = Line;
</script>

<style scoped>
canvas {
  max-width: 100%;
}
</style>
