<template>
  <div>
    <Line v-if='chartData.datasets.length' :data='chartData' :options='chartOptions' />
    <p v-else>Loading chart data...</p>
  </div>
</template>

<script>
import { defineComponent, onMounted, ref } from 'vue';
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
import { getAggregationData } from '../../api/api';

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

    const fetchData = async () => {
      try {
        const data = await getAggregationData();
        if (!data || !Array.isArray(data.datasets)) {
          throw new Error('Invalid data format');
        }

        chartData.value = {
          datasets: data.datasets.map((dataset) => ({
            ...dataset,
            data: dataset.data.map((point) => ({ x: new Date(point.x), y: point.y })),
          })),
        };
      } catch (error) {
        console.error('Error fetching workout data:', error);
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
