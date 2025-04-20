<template>
  <div class="container py-5">
    <div class="row justify-content-center">
      <div class="col-md-8">
        <div class="card shadow rounded">
          <div class="card-body text-center p-5">
            <h2 class="card-title text-center mb-4">
              {{ $t('links.profile') }}
            </h2>
            <div class="mb-4">
              <img
                :src="currentUser.picture"
                alt="User Avatar"
                class="rounded-circle img-thumbnail shadow"
                style="width: 150px; height: 150px; object-fit: cover"
              />
            </div>
            <h2 class="card-title mb-3">
              <font-awesome-icon icon="fa-solid fa-circle-user" class="me-2" />
              {{ currentUser.name }}
            </h2>
            <p class="card-text text-muted mb-4">
              <font-awesome-icon icon="fa-solid fa-envelope" class="me-2" />
              {{ currentUser.email }}
            </p>

            <auth-button
              loginButtonClass="btn btn-outline-primary me-2"
              logoutButtonClass="btn btn-outline-danger"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { storeToRefs } from 'pinia';
import fallbackAvatar from '@/assets/logo.jpg';
import { useUserStore } from '../stores/userStore';
import AuthButton from './common/AuthButton.vue';

export default {
  name: 'userProfile',
  components: {
    AuthButton,
  },
  setup() {
    const userStore = useUserStore();
    const { user } = storeToRefs(userStore);
    return {
      user,
    };
  },
  computed: {
    currentUser() {
      return this.user ?? {
        id: 0,
        name: 'Chad',
        picture: fallbackAvatar,
        email: 'chad@localhost',
      };
    },
  },
};
</script>
<style scoped>
.card-title {
  font-size: 1.75rem;
  font-weight: 600;
}
</style>
