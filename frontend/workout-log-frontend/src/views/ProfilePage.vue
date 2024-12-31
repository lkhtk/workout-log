<template>
  <div class="d-flex h-100 text-center">
    <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
      <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
        <h1 class="display-1">
          User profile</h1>
      </div>
      <main>
        <div class="row row-cols-1 row-cols-md-3 mb-3 text-center justify-content-around">
          <div class="col">
            <div class="card mb-4 rounded-3 shadow-sm">
              <div class="card-header py-3">
                <img
                  :src="user?.picture || defaultUser.picture"
                  alt="avatar"
                  class="card-img-top"
                >
              </div>
              <div class="card-body">
                <h1 class="card-title pricing-card-title">
                  <font-awesome-icon icon="fa-solid fa-circle-user" />
                  {{ user?.name || defaultUser.name }}
                </h1>
                <h6 class="card-subtitle mb-2 text-body-secondary">
                  <font-awesome-icon icon="fa-solid fa-at" />
                  {{ user?.email || defaultUser.email }}
                </h6>
                <button
                  class="btn w-100 btn-lg btn-outline-danger"
                  v-if="user"
                  @click="revoke(user.id)"
                >
                <font-awesome-icon icon="fa-solid fa-arrow-right-from-bracket" />
                Sign Out
                </button>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { storeToRefs } from 'pinia';
import { idRevoke } from 'vue3-google-signin';
import { logOut } from '../api/api';
import { useUserStore } from '../stores/userStore';

const userStore = useUserStore();
const { user } = storeToRefs(userStore);

const defaultUser = {
  id: 0,
  name: 'Chad',
  picture: 'https://upload.wikimedia.org/wikipedia/ru/9/94/%D0%93%D0%B8%D0%B3%D0%B0%D1%87%D0%B0%D0%B4.jpg',
  email: 'root@localhost',
};

const revoke = async (id) => {
  try {
    await new Promise((resolve, reject) => {
      idRevoke(id, (done) => {
        if (done.error) {
          console.error('Error during revoke:', done.error);
          reject(new Error(done.error));
        } else {
          resolve();
        }
      });
    });
    const result = await logOut();
    if (result.status !== 200) {
      throw new Error('Server logout failed');
    }
    userStore.clearUser();
    console.log('User logged out successfully');
  } catch (error) {
    console.error('Logout error:', error.message);
  }
};
</script>
