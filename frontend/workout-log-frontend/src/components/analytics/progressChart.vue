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
      <div class="col-12 col-lg-3 p-3">
        <apexchart
          class="card"
          v-if="chartOptionsPie.series.length"
          type="donut"
          :options="chartOptionsPie"
          :series="chartOptionsPie.series"
          height="300"
        />
      </div>

      <div class="col-12 col-lg-3 p-3">
        <card
          v-if="exercisesCount"
          :cardData="{
            icon: 'fa-solid fa-person-swimming',
            count: exercisesCount,
            text: $t('trends.exercisesCountTitle'),
          }"
        />
      </div>

      <div class="col-12 col-lg-3 p-3">
        <card
          v-if="workoutCount"
          :cardData="{
            icon: 'fa-solid fa-person-chalkboard',
            count: workoutCount + ' / ' + coachCount,
            text: $t('trends.countTitle') + '/' + $t('workout.instructor'),
          }"
        />
      </div>

      <div class="col-12 col-lg-3 p-3">
        <card
          v-if="timeTotal"
          :cardData="{
            icon: 'fa-solid fa-clock',
            count: timeTotal.toFixed(2),
            text: $t('trends.time_total') + ' (' + $t('units.hrs') + ')',
          }"
        />
      </div>
    </div>

    <div class="row">
      <div class="col-12 p-3" v-if="workoutCount > 1">
        <apexchart
          class="card"
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

    <div class="row">
      <div class="col-12 col-lg-4 p-3">
        <card
          v-if="avgDuration"
          :cardData="{
            icon: 'fa-solid fa-stopwatch-20',
            count: avgDuration.toFixed(2),
            text: $t('trends.avg_duration')  + ' (' + $t('units.min') + ')',
          }"
        />
      </div>

      <div class="col-12 col-lg-4 p-3">
        <card
          v-if="avgSlept"
          :cardData="{
            icon: 'fa-solid fa-bed',
            count: avgSlept.toFixed(2),
            text: $t('trends.avg_slept') + ' (' + $t('units.hrs') + ')',
          }"
        />
      </div>

      <div class="col-12 col-lg-4 p-3">
        <card
          v-if="avgIntens"
          :cardData="{
            icon: 'fa-solid fa-battery-three-quarters',
            count: avgIntens.toFixed(2) +'/10',
            text: $t('trends.avg_intensity'),
          }"
        />
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
import Card from '../common/NumberCard.vue';

export default defineComponent({
  components: {
    apexchart: VueApexCharts,
    card: Card,
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

    const selectedPeriod = ref('all');
    const workoutCount = ref(0);
    const coachCount = ref(0);
    const avgDuration = ref(0);
    const durationCount = ref(0);
    const avgSlept = ref(0);
    const sleptCount = ref(0);
    const avgIntens = ref(0);
    const intensCount = ref(0);
    const exercisesCount = ref(0);
    const timeTotal = ref(0);

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
        coachCount.value = 0;
        avgDuration.value = 0;
        avgSlept.value = 0;
        avgIntens.value = 0;
        durationCount.value = 0;
        sleptCount.value = 0;
        intensCount.value = 0;

        response.data.data.forEach((entry) => {
          if (entry.coach) {
            coachCount.value += 1;
          }
          if (entry.review.duration != null) {
            avgDuration.value += entry.review.duration;
            durationCount.value += 1;
          }
          if (entry.review.slept != null) {
            avgSlept.value += entry.review.slept;
            sleptCount.value += 1;
          }
          if (entry.review.intensity != null) {
            avgIntens.value += entry.review.intensity;
            intensCount.value += 1;
          }
          const date = new Date(entry.PublishedAt).getTime();
          if (entry.exercises) {
            entry.exercises.forEach((exercise) => {
              if (!exerciseMap[exercise.name]) {
                exerciseMap[exercise.name] = [];
              }
              exerciseMap[exercise.name].push({
                x: date,
                y: parseFloat(exercise.average_weight_per_exercise.toFixed(2)),
              });
            });
          }
        });

        timeTotal.value = (avgDuration.value * 1.0) / 60;
        avgDuration.value = (avgDuration.value * 1.0) / durationCount.value;
        avgSlept.value = (avgSlept.value * 1.0) / sleptCount.value;
        avgIntens.value = (avgIntens.value * 1.0) / intensCount.value;
        exercisesCount.value = Object.keys(exerciseMap).length;

        const aggregatedData = Object.keys(exerciseMap)
          .filter((exercise) => exerciseMap[exercise].length >= 2)
          .map((exercise) => ({
            name: exercise,
            data: exerciseMap[exercise].sort((a, b) => a.x - b.x),
          }));

        chartOptions.value.series = aggregatedData;
      } catch (error) {
        console.error('Ошибка при получении трендов:', error);
        window.$toast?.showToast(error.message || 'Ошибка получения данных', 'danger');
      }
    };

    const fetchTopData = async () => {
      try {
        const response = await getTop5(selectedPeriod.value);
        const topData = response.data.data;
        chartOptionsPie.value.labels = topData.map((item) => item.name);
        chartOptionsPie.value.series = topData.map((item) => item.count);
      } catch (error) {
        console.error('Ошибка при получении ТОП-5:', error);
        window.$toast?.showToast(error.message || 'Ошибка получения данных', 'danger');
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
      coachCount,
      avgDuration,
      avgSlept,
      avgIntens,
      timeTotal,
    };
  },
});
</script>
