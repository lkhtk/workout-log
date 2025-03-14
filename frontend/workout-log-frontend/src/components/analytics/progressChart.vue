<template>
  <div v-if="chartOptions.series.length && chartOptionsPie.series.length">
    <apexchart
      :options="chartOptions"
      :series="chartOptions.series"
      type="line"
      height="500"
    />
    <apexchart
      type="donut"
      :options="chartOptionsPie"
      :series="chartOptionsPie.series"
      height="300"
    />
  </div>
  <div class="spinner-border align-items-center" role="status" v-else></div>
</template>

<script>
import { defineComponent, onMounted, ref } from 'vue';
import { getTop5, getTrends } from '@/api/api';
import VueApexCharts from 'vue3-apexcharts';

export default defineComponent({
  components: {
    apexchart: VueApexCharts,
  },

  setup() {
    const chartOptionsPie = ref({
      chart: {
        type: 'donut',
        height: 400,
      },
      labels: [],
      title: {
        text: 'Top',
        align: 'center',
      },
      legend: {
        position: 'right',
      },
      series: [],
    });

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
    const fetchTrendsData = async () => {
      try {
        const response = await getTrends(100);
        const exerciseMap = {};
        response.data.data.forEach((entry) => {
          const date = new Date(entry.PublishedAt).getTime();

          entry.exercises.forEach((exercise) => {
            if (!exerciseMap[exercise.name]) {
              exerciseMap[exercise.name] = [];
            }
            exerciseMap[exercise.name].push({
              x: date,
              y: parseFloat(exercise.average_weight_per_exercise.toFixed(2)),
            });
          });
        });
        const aggregatedData = Object.keys(exerciseMap)
          .filter((exercise) => exerciseMap[exercise].length >= 2)
          .map((exercise) => ({
            name: exercise,
            data: exerciseMap[exercise].sort((a, b) => a.x - b.x),
          }));
        chartOptions.value.series = aggregatedData;
      } catch (error) {
        console.error(error.message);
      }
    };
    const fetchTopData = async () => {
      try {
        const response = await getTop5(100);
        chartOptionsPie.value.labels = response.data.data.map((item) => item.name);
        chartOptionsPie.value.series = response.data.data.map((item) => item.count);
      } catch (error) {
        console.error(error);
      }
    };
    onMounted(() => {
      fetchTrendsData();
      fetchTopData();
    });

    return {
      chartOptions,
      chartOptionsPie,
    };
  },
});
</script>
