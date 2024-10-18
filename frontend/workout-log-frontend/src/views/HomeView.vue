<template>
    <div v-for="workout in data.data" v-bind:key="workout">
      <WorkoutTable :data="workout" />
    </div>
</template>

<script>
import * as api from '@/api/api';
import WorkoutTable from '@/components/WorkoutTable.vue';

export default {
  name: 'HomeView',
  components: {
    WorkoutTable,
  },
  data: () => ({
    data: 1,
  }),
  async beforeMount() {
    try {
      this.loadUserInfo();
    } catch (e) {
      console.error(e);
    }
  },
  methods: {
    async loadUserInfo() {
      this.data = await api.getAllWorkouts();
    },
  },
};
</script>
