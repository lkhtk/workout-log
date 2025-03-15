<template>
  <div class="container-fluid">
      <div class="row">
        <div class="col-12 p-3">
          <apexchart
            class="card shadow"
            v-if="chartOptions.series.length"
            :options="chartOptions"
            :series="chartOptions.series"
            type="line"
            height="400"
          />
          <div v-else class="spinner-border" role="status"></div>
        </div>
      <div class="row">
        <div class="col-12 col-lg-4 p-3">
          <apexchart
            class="card shadow"
            v-if="chartOptionsPie.series.length"
            type="donut"
            :options="chartOptionsPie"
            :series="chartOptionsPie.series"
            height="400"
          />
          <div v-else class="spinner-border" role="status"></div>
        </div>
      </div>
    </div>
  </div>
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
        text: 'Top 5',
        align: 'center',
      },
      legend: {
        position: 'bottom',
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
      title: {
        text: 'Progress trends',
        align: 'center',
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
        horizontalAlign: 'top',
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
