<template>
  <div class="container mt-5" v-if="user">
    <form @submit.prevent="" class="mb-4">
      <div class="d-grid gap-2">
        <div class="container d-grid gap-2 col-6 mx-auto">
          <h1 class="display-5">Language</h1>
          <div class="dropdown">
            <button class="btn btn-primary dropdown-toggle" type="button"
            data-bs-toggle="dropdown" aria-expanded="false">
              <font-awesome-icon icon="fa-solid fa-language" />
              Select lang
            </button>
            <ul class="dropdown-menu">
              <li><a class="dropdown-item" href="#">🇷🇺 Ru</a></li>
              <li><a class="dropdown-item" href="#">🇺🇸 Eng</a></li>
            </ul>
          </div>
          <button class="btn btn-primary" @click="submitLang">
            <font-awesome-icon icon="fa-solid fa-floppy-disk" />
            Save
          </button>
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import { storeToRefs } from 'pinia';
import { useUserStore } from '../stores/userStore';

export default {
  name: 'userConfigs',
  setup() {
    const userStore = useUserStore();
    const { user } = storeToRefs(userStore);
    return {
      user,
    };
  },
  computed: {
    currentUser() {
      return this.user || {
        id: 0,
      };
    },
  },
  methods: {
    submitLang() {
      console.log('lang changed');
    },
  },
};
</script>
