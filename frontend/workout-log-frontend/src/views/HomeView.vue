<template>
  <div id="app">
    <component :is="component" @change-component="changeComponent"/>
  </div>
</template>

<script>
import Create from '@/components/create.vue';
import View from '@/components/view.vue';
import List from '@/components/list.vue';
import Update from '@/components/update.vue';

export default {
  name: 'HomeView',
  data: () => ({
    componentId: 'list',
    workoutId: '',
  }),
  provide() {
    console.log('provide:', this.workoutId);
    const base = {};
    Object.defineProperty(base, 'workoutId', {
      value: this.workoutId,
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
        case 'edit':
          return Update;
        default:
          return undefined;
      }
    },
  },
  methods: {
    changeComponent(payload) {
      console.log('payload --->', payload);
      this.componentId = payload.component;
      this.workoutId = payload.workoutId;
    },
  },
};
</script>
