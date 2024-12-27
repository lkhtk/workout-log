import { createApp } from 'vue';
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { fas } from '@fortawesome/free-solid-svg-icons';
import GoogleSignInPlugin from 'vue3-google-signin';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './router';

import 'bootstrap/dist/css/bootstrap.css';

library.add(fas);
createApp(App)
  .use(createPinia())
  .component('font-awesome-icon', FontAwesomeIcon)
  .use(router)
  .use(GoogleSignInPlugin, {
    clientId: process.env.VUE_APP_GOOGLE_CLIENT_ID,
  })
  .mount('#app');
