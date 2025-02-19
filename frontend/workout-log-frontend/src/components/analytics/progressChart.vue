<template>
  <div>
    <Line v-if='chartData.datasets.length' :data='chartData' :options='chartOptions' />
    <p v-else>Loading chart data...</p>
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
    const chartData = ref({ datasets: [] });
    const chartOptions = ref({
      responsive: true,
      plugins: {
        legend: {
          position: 'top',
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
            text: 'Weight (kg)',
          },
        },
      },
    });

    const N_DAYS = 365;

    const fetchData = async () => {
      try {
        const data = await getTrends('year');
        const exerciseMap = {};
        data.data.data.forEach((entry) => {
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
      } catch (error) {
        console.error('Error fetching workout data:', error.message);
      }
    };

    onMounted(fetchData);

    return {
      chartData,
      chartOptions,
    };
  },
});
</script>
