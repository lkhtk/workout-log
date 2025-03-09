<template>
  <div>
    <apexchart v-if="chartOptions.series.length"
    :options="chartOptions" :series="chartOptions.series" type="line" height="500" />
    <p v-else>Loading chart data...</p>
  </div>
</template>

<script>
import { defineComponent, onMounted, ref } from 'vue';
import { getTrends } from '@/api/api';
import VueApexCharts from 'vue3-apexcharts';

export default defineComponent({
  components: {
    apexchart: VueApexCharts,
  },
  setup() {
    const chartOptions = ref({
      chart: {
        type: 'line',
        height: 500,
        animations: {
          enabled: false,
        },
        zoom: {
          enabled: false,
        },
      },
      xaxis: {
        type: 'datetime',
      },
      stroke: {
        curve: 'smooth',
        width: 2,
      },
      fill: {
        opacity: 0.8,
        type: 'pattern',
        pattern: {
          style: ['verticalLines', 'horizontalLines'],
          width: 5,
          height: 6,
        },
      },
      markers: {
        size: 5,
        hover: {
          size: 9,
        },
      },
      legend: {
        position: 'right',
        horizontalAlign: 'center',
      },
      series: [],
    });

    const N_DAYS = 7;

    const fetchData = async () => {
      try {
        const response = await getTrends(100);
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

        const aggregatedData = Object.keys(exerciseMap).map((exercise) => {
          const dataPoints = Object.keys(exerciseMap[exercise])
            .map((date) => {
              const weights = exerciseMap[exercise][date];
              const avgWeight = (weights.reduce((sum, w) => sum + w, 0)
              / weights.length).toFixed(2);
              return { x: new Date(date).getTime(), y: parseFloat(avgWeight) };
            })
            .sort((a, b) => a.x - b.x)
            .slice(-N_DAYS);
          return dataPoints.length > 1 ? { name: exercise, data: dataPoints } : null;
        }).filter(Boolean);

        chartOptions.value = {
          ...chartOptions.value,
          series: aggregatedData,
        };
      } catch (error) {
        console.error('Error fetching workout data:', error.message);
      }
    };

    onMounted(fetchData);

    return {
      chartOptions,
    };
  },
});
</script>
