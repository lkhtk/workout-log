<template>
  <div class="container" v-if="!isLoading">
    <div v-if="hasNoData">
      <p>{{ $t('errorsMsg.noDataAvailable') }}</p>
    </div>
    <div v-else-if="chartData && chartData.datasets.length">
      <Line :data="chartData" :options="chartOptions" />
    </div>
    <div v-else>
      <p>{{ $t('errorsMsg.noDataAvailable') }}</p>
    </div>
  </div>
  <div v-else>
    <h1>LOADING!</h1>
  </div>
</template>

<script>
import { defineComponent, onMounted, ref } from 'vue';
import { getTrends } from '@/api/api';
import { Line } from 'vue-chartjs';
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  LineElement,
  CategoryScale,
  LinearScale,
  PointElement,
  TimeScale,
} from 'chart.js';
import 'chartjs-adapter-date-fns';

ChartJS.register(
  Title,
  Tooltip,
  Legend,
  LineElement,
  CategoryScale,
  LinearScale,
  PointElement,
  TimeScale,
);

export default defineComponent({
  components: {
    Line,
  },
  setup() {
    const isLoading = ref(true);
    const hasNoData = ref(false);
    const chartData = ref({ datasets: [] });
    const chartOptions = ref({
      responsive: true,
      plugins: {
        legend: {
          position: 'right',
        },
      },
      scales: {
        x: {
          type: 'time',
          time: {
            unit: 'day',
          },
          title: {
            display: true,
            text: 'Date',
          },
        },
        y: {
          title: {
            display: true,
            text: 'Average Weight',
          },
        },
      },
    });

    const N_DAYS = 365;

    const fetchData = async () => {
      try {
        const response = await getTrends(100);
        if (!response.data || response.data.total === 0) {
          hasNoData.value = true;
        } else {
          const exerciseMap = {};
          response.data.data.forEach((entry) => {
            const date = new Date(entry.PublishedAt).toISOString().split('T')[0];
            entry.workout.exercises.forEach((exercise) => {
              if (!exerciseMap[exercise.name]) {
                exerciseMap[exercise.name] = {};
              }
              if (!exerciseMap[exercise.name][date]) {
                exerciseMap[exercise.name][date] = [];
              }
              exercise.sets.forEach((set) => {
                exerciseMap[exercise.name][date].push(set.weight);
              });
            });
          });

          const aggregatedData = Object.keys(exerciseMap).map((exercise, index) => {
            const dataPoints = Object.keys(exerciseMap[exercise])
              .map((date) => {
                const weights = exerciseMap[exercise][date];
                const avgWeight = weights.reduce((sum, w) => sum + w, 0) / weights.length;
                return { x: date, y: avgWeight };
              })
              .sort((a, b) => new Date(a.x) - new Date(b.x))
              .slice(-N_DAYS);

            return {
              label: exercise,
              borderColor: `hsl(${index * 60}, 70%, 50%)`,
              data: dataPoints,
              fill: false,
            };
          });

          chartData.value = { datasets: aggregatedData };
          hasNoData.value = aggregatedData.length === 0;
        }
      } catch (error) {
        console.error('Error fetching workout data:', error.message);
        hasNoData.value = true;
      } finally {
        isLoading.value = false;
      }
    };

    onMounted(fetchData);

    return {
      isLoading,
      hasNoData,
      chartData,
      chartOptions,
    };
  },
});
</script>
