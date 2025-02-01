import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { fas } from '@fortawesome/free-solid-svg-icons';
import GoogleSignInPlugin from 'vue3-google-signin';
import { createI18n } from 'vue-i18n';
import App from './App.vue';
import router from './router';
import 'bootstrap/dist/css/bootstrap.css';
import { useUserStore } from './stores/userStore';

const pinia = createPinia();
const app = createApp(App);
app.use(pinia);

const userStore = useUserStore();
const savedLang = userStore.lang || 'en';

async function loadLocaleMessages(locale) {
  const messages = await fetch(`locales/${locale}.json`).then((res) => res.json());
  return messages;
}

const i18n = createI18n({
  legacy: false,
  locale: savedLang,
  fallbackLocale: 'en',
  messages: {},
});

library.add(fas);

app.use(i18n)
  .component('font-awesome-icon', FontAwesomeIcon)
  .use(router)
  .use(GoogleSignInPlugin, {
    clientId: process.env.VUE_APP_GOOGLE_CLIENT_ID,
  });

loadLocaleMessages(savedLang).then((messages) => {
  i18n.global.setLocaleMessage(savedLang, messages);
  i18n.global.locale.value = savedLang;
  app.mount('#app');
});
