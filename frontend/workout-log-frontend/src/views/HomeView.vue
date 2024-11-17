<template>
  <div id="app">
    <component :is="component" :workoutId="id" @change-component="changeComponent"/>
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
    id: '',
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
        case 'edit':
          return Update;
        default:
          return undefined;
      }
    },
  },
  methods: {
    changeComponent(payload) {
      this.componentId = payload.component;
      this.id = payload.workoutId;
    },
  },
};
</script>
