<template>
  <div class="container-fluid">
    <div class="row">
      <nav aria-label="breadcrumb">
        <ol class="breadcrumb">
          <li class="breadcrumb-item" :class="{ active: selectedPeriod === 'all' }">
            <a href="#" @click.prevent="setPeriod('all')">{{ $t('periodTitles.all') }}</a>
          </li>
          <li class="breadcrumb-item" :class="{ active: selectedPeriod === 'month' }">
            <a href="#" @click.prevent="setPeriod('month')">{{ $t('periodTitles.month') }}</a>
          </li>
          <li class="breadcrumb-item" :class="{ active: selectedPeriod === 'week' }"
          aria-current="page">
            <a href="#" @click.prevent="setPeriod('week')">{{ $t('periodTitles.week') }}</a>
          </li>
        </ol>
      </nav>
    </div>

    <div class="row">
      <div class="col-12 col-lg-4 p-3">
        <apexchart
          class="card shadow"
          v-if="chartOptionsPie.series.length"
          type="donut"
          :options="chartOptionsPie"
          :series="chartOptionsPie.series"
          height="300"
        />
        <div v-else class="spinner-border" role="status"></div>
      </div>
      <div class="col-12 col-lg-4 p-3">
        <div class="card text-center shadow" v-if="workoutCount">
          <div class="card-body">
            <h2 class="display-1">
            {{ workoutCount }}
            </h2>
            <p class="card-text">{{ $t('trends.countTitle') }}</p>
          </div>
        </div>
        <div v-else class="spinner-border" role="status"></div>
      </div>
    </div>

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
    </div>
  </div>
</template>

<script>
import {
  getTop5, getTrends,
} from '@/api/api';
import VueApexCharts from 'vue3-apexcharts';
import {
  defineComponent, onMounted, ref, watch,
} from 'vue';

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

    const selectedPeriod = ref('all'); // 'all', 'month', or 'week'
    const workoutCount = ref(0);
    const setPeriod = (period) => {
      if (selectedPeriod.value !== period) {
        selectedPeriod.value = period;
      }
    };

    const fetchTrendsData = async () => {
      try {
        const response = await getTrends(selectedPeriod.value);
        const exerciseMap = {};
        workoutCount.value = response.data.data.length;

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
        const response = await getTop5(selectedPeriod.value);
        chartOptionsPie.value.labels = response.data.data.map((item) => item.name);
        chartOptionsPie.value.series = response.data.data.map((item) => item.count);
      } catch (error) {
        console.error(error);
      }
    };

    watch(selectedPeriod, () => {
      chartOptions.value.series = [];
      chartOptionsPie.value.series = [];
      fetchTrendsData();
      fetchTopData();
    });

    onMounted(() => {
      fetchTrendsData();
      fetchTopData();
    });

    return {
      chartOptions,
      chartOptionsPie,
      selectedPeriod,
      setPeriod,
      workoutCount,
    };
  },
});
</script>
