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
    componentIs: 'list',
    userId: 0,
  }),
  provide() {
    const base = {};

    Object.defineProperty(base, 'userId', {
      enumerable: true,
      get: () => Number(this.userId),
    });

    return base;
  },
  computed: {
    component() {
      switch (this.componentIs) {
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
      this.componentIs = payload.component;
      this.userId = Number(payload.userId);
    },
  },
};
</script>
