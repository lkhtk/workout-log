import { defineComponent } from 'vue';
import { defineStore, mapStores } from 'pinia';

export const useUserStore = defineStore('user', {
  state: () => ({
    lang: localStorage.getItem('lang') ? JSON.parse(localStorage.getItem('lang')) : 'en',
    user: JSON.parse(localStorage.getItem('user')) || null,
  }),
  actions: {
    setLanguage(lang) {
      this.lang = lang;
      localStorage.setItem('lang', JSON.stringify(lang));
    },
    setUser(userData) {
      this.user = userData;
      localStorage.setItem('user', JSON.stringify(userData));
    },
    clearUser() {
      this.user = null;
      localStorage.removeItem('user');
      localStorage.setItem('lang', JSON.stringify(this.lang));
    },
  },
});

export default defineComponent({
  computed: {
    ...mapStores(useUserStore),
  },
});
