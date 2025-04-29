<template>
  <div id="app">
    <component :is="component" :workoutData="data" @change-component="changeComponent"/>
  </div>
</template>

<script>
import Create from '@/components/workout/create.vue';
import Edit from '@/components/workout/edit.vue';
import All from '@/components/workout/all.vue';

export default {
  name: 'HomeView',
  data: () => ({
    componentId: 'all',
    data: {},
  }),
  provide() {
    const base = {};
    Object.defineProperty(base, 'workoutId', {
      value: this.id,
    });
    return base;
  },
  computed: {
    component() {
      switch (this.componentId) {
        case 'all':
          return All;
        case 'create':
          return Create;
        case 'edit':
          return Edit;
        default:
          return undefined;
      }
    },
  },
  methods: {
    changeComponent(payload) {
      this.componentId = payload.component;
      this.data = payload.payload;
    },
  },
};
</script>
