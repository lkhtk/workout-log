import { defineComponent } from 'vue';
import { defineStore, mapStores } from 'pinia';

export const useUserStore = defineStore('user', {
  state: () => ({
    user: JSON.parse(localStorage.getItem('user')) || null,
  }),
  actions: {
    setUser(userData) {
      this.user = userData;
      localStorage.setItem('user', JSON.stringify(userData));
    },
    clearUser() {
      this.user = null;
      localStorage.removeItem('user');
    },
  },
});

export default defineComponent({
  computed: {
    ...mapStores(useUserStore),
  },
});
