<template>
  <div class="container-fluid">
    <div class="row">
      <nav aria-label="breadcrumb">
        <ol class="breadcrumb">
          <li class="breadcrumb-item" :class="{ active: selectedPeriod === 'all' }">
            <a href="#" @click.prevent="setPeriod('all')">{{ $t('periodTitles.all') }}</a>
          </li>
          <li class="breadcrumb-item" :class="{ active: selectedPeriod === 'month-12' }">
            <a href="#" @click.prevent="setPeriod('month-12')">{{ $t('periodTitles.year') }}</a>
          </li>
          <li class="breadcrumb-item" :class="{ active: selectedPeriod === 'month-6' }">
            <a href="#" @click.prevent="setPeriod('month-6')">{{ $t('periodTitles.month-6') }}</a>
          </li>
          <li class="breadcrumb-item" :class="{ active: selectedPeriod === 'month-3' }">
            <a href="#" @click.prevent="setPeriod('month-3')">{{ $t('periodTitles.month-3') }}</a>
          </li>
          <li class="breadcrumb-item" :class="{ active: selectedPeriod === 'month-1' }">
            <a href="#" @click.prevent="setPeriod('month-1')">{{ $t('periodTitles.month') }}</a>
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
        <div class="text-center shadow border rounded" style="height: 300px;" v-if="exercisesCount">
          <div class="position-relative top-50 start-50 translate-middle">
            <h2 class="display-1">
            {{ exercisesCount }}
            </h2>
            <p class="card-text">{{ $t('trends.exercisesCountTitle') }}</p>
          </div>
        </div>

      </div>
      <div class="col-12 col-lg-4 p-3">
        <div class="text-center shadow border rounded" style="height: 300px;" v-if="workoutCount">
          <div class="position-relative top-50 start-50 translate-middle">
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
      <div class="col-12 p-3" v-if="workoutCount>1">
        <apexchart
          class="card shadow"
          v-if="chartOptions.series.length"
          :options="chartOptions"
          :series="chartOptions.series"
          type="line"
          height="600"
        />
        <div v-else class="spinner-border" role="status"></div>
      </div>
      <div class="col-12 p-3 shadow" v-else>{{ $t('errorsMsg.invalndChartData') }}</div>
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

    const selectedPeriod = ref('all'); // 'all', 'year','month', or 'week'
    const workoutCount = ref(0);
    const exercisesCount = ref(0);
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
        exercisesCount.value = Object.keys(exerciseMap).length;
        const aggregatedData = Object.keys(exerciseMap)
          .filter((exercise) => exerciseMap[exercise].length >= 2)
          .map((exercise) => ({
            name: exercise,
            data: exerciseMap[exercise].sort((a, b) => a.x - b.x),
          }));
        chartOptions.value.series = aggregatedData;
      } catch (error) {
        window.$toast?.showToast(error.message, 'danger');
      }
    };

    const fetchTopData = async () => {
      try {
        const response = await getTop5(selectedPeriod.value);
        chartOptionsPie.value.labels = response.data.data.map((item) => item.name);
        chartOptionsPie.value.series = response.data.data.map((item) => item.count);
      } catch (error) {
        window.$toast?.showToast(error.message, 'danger');
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
      exercisesCount,
    };
  },
});
</script>
