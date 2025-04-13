<template>
  <form @submit.prevent="submitLang" class="mb-4">
    <div class="container d-grid gap-2 col-6 mx-auto">
      <div class="dropdown">
        <button class="btn btn-outline-dark dropdown-toggle" type="button"
        data-bs-toggle="dropdown" aria-expanded="false">
          <font-awesome-icon icon="fa-solid fa-language" />
          {{ $t('userConf.selectTitle') }}
        </button>
        <ul class="dropdown-menu">
          <li><a class="dropdown-item" href="#" @click.prevent="loadLanguage('ru')">
            🇷🇺 {{ $t('userConf.ru') }}</a></li>
          <li><a class="dropdown-item" href="#" @click.prevent="loadLanguage('en')">
            🇺🇸 {{ $t('userConf.eng') }}</a></li>
          <li><a class="dropdown-item" href="#" @click.prevent="loadLanguage('esp')">
            🇪🇸 {{ $t('userConf.esp') }}</a></li>
          <li><a class="dropdown-item" href="#" @click.prevent="loadLanguage('ch')">
            🇨🇳 {{ $t('userConf.ch') }}</a></li>
        </ul>
      </div>
    </div>
  </form>
</template>

<script>
import { useI18n } from 'vue-i18n';
import { useUserStore } from '../../stores/userStore';

export default {
  name: 'userLang',
  setup() {
    const userStore = useUserStore();
    const { locale, setLocaleMessage } = useI18n();
    const loadLanguage = async (lang) => {
      const messages = await fetch(`/locales/${lang}.json`).then((res) => res.json());
      setLocaleMessage(lang, messages);
      locale.value = lang;
      userStore.setLanguage(lang);
    };
    return {
      loadLanguage,
    };
  },
};
</script>
