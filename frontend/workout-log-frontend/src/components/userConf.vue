<template>
  <div class="container mt-5" v-if="user">
    <form @submit.prevent="submitLang" class="mb-4">
      <div class="d-grid gap-2">
        <div class="container d-grid gap-2 col-6 mx-auto">
          <h1 class="display-5">{{ $t('userConf.title') }}</h1>
          <div class="dropdown">
            <button class="btn btn-primary dropdown-toggle" type="button"
            data-bs-toggle="dropdown" aria-expanded="false">
              <font-awesome-icon icon="fa-solid fa-language" />
              {{ $t('userConf.selectTitle') }}
            </button>
            <ul class="dropdown-menu">
              <li><a class="dropdown-item" href="#" @click.prevent="loadLanguage('ru')">
                🇷🇺 {{ $t('userConf.ru') }}</a></li>
              <li><a class="dropdown-item" href="#" @click.prevent="loadLanguage('en')">
                🇺🇸 {{ $t('userConf.eng') }}</a></li>
            </ul>
          </div>
          <button class="btn btn-primary" type="submit">
            <font-awesome-icon icon="fa-solid fa-floppy-disk" />
            {{ $t('buttons.save') }}
          </button>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import { storeToRefs } from 'pinia';
import { useI18n } from 'vue-i18n';
import { useUserStore } from '../stores/userStore';

export default {
  name: 'UserConfigs',
  setup() {
    const userStore = useUserStore();
    const { user, language } = storeToRefs(userStore);
    const { locale, setLocaleMessage } = useI18n();

    const loadLanguage = async (lang) => {
      const messages = await fetch(`/locales/${lang}.json`).then((res) => res.json());
      setLocaleMessage(lang, messages);
      locale.value = lang;
      userStore.setLanguage(lang);
    };

    return {
      user,
      language,
      loadLanguage,
    };
  },
};
</script>
