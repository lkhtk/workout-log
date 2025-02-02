<template>
    <div class="d-flex h-100">
        <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
            <div class="pricing-header p-3 pb-md-4 mx-auto">
                <h1 class="display-5">
                  {{ $t('links.profile') }}
                </h1>
            </div>
            <main>
                <div class="row row-cols-1 row-cols-md-3 mb-3 justify-content-around">
                    <div class="col">
                        <div class="card mb-4 rounded-3 shadow-sm" style="width: 18rem;">
                            <div class="card-header py-3">
                                <img :src="currentUser.picture" alt="avatar"
                                class="card-img-top img-fluid">
                            </div>
                            <div class="card-body">
                                <h1 class="card-title pricing-card-title">
                                    <font-awesome-icon icon="fa-solid fa-circle-user" />
                                    {{ currentUser.name }}
                                </h1>
                                <h6 class="card-subtitle mb-2 text-body-secondary">
                                    <font-awesome-icon icon="fa-solid fa-envelope" />
                                    {{ currentUser.email }}
                                </h6>
                                <auth-button
                                loginButtonClass="btn btn-outline-dark"
                                logoutButtonClass="btn btn-outline-dark"/>
                            </div>
                        </div>
                    </div>
                </div>
            </main>
        </div>
    </div>
</template>

<script>
import { storeToRefs } from 'pinia';
import { useUserStore } from '../stores/userStore';
import AuthButton from './common/AuthButton.vue';

export default {
  name: 'userProfile',
  components: {
    AuthButton,
  },
  setup() {
    const userStore = useUserStore();
    const { user, language } = storeToRefs(userStore);
    return {
      user,
      language,
    };
  },
  computed: {
    currentUser() {
      return this.user || {
        id: 0,
        name: 'Chad',
        picture: 'https://upload.wikimedia.org/wikipedia/ru/9/94/%D0%93%D0%B8%D0%B3%D0%B0%D1%87%D0%B0%D0%B4.jpg',
        email: 'chad@localhost',
      };
    },
  },
};
</script>
