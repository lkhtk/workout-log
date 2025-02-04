<template>
  <div id="app">
    <component :is="component" :workoutData="data" @change-component="changeComponent"/>
  </div>
</template>

<script>
import Create from '@/components/workout/create.vue';
import View from '@/components/workout/view.vue';
import List from '@/components/workout/list.vue';

export default {
  name: 'HomeView',
  data: () => ({
    componentId: 'list',
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
        case 'list':
          return List;
        case 'create':
          return Create;
        case 'view':
          return View;
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
