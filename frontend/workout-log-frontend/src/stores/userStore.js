import { defineComponent } from 'vue';
import { defineStore, mapStores } from 'pinia';

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null,
  }),
  actions: {
    setUser(userData) {
      this.user = userData;
    },
    clearUser() {
      this.user = null;
    },
  },
});

export default defineComponent({
  computed: {
    ...mapStores(useUserStore),
  },
});
