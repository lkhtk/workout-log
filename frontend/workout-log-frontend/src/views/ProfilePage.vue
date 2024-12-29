<template>
  <div v-if="!user">
    Unauthorized
  </div>
  <div v-else class="card" style="width: 18rem;">
    <img :src="user.picture" alt="avatar" class="card-img-top">
    <div class="card-body">
      <h5 class="card-title">{{ user.name }}</h5>
      <p class="card-text">{{ user.email }}</p>
      <button class="btn btn-danger" @click="revoke(user.id)">Logout</button>
    </div>
  </div>
</template>

<script setup>
import { storeToRefs } from 'pinia';
import {
  idRevoke,
} from 'vue3-google-signin';

import {
  logOut,
} from '../api/api';

import { useUserStore } from '../stores/userStore';

const userStore = useUserStore();
const { user } = storeToRefs(userStore);
const revoke = async (id) => {
  idRevoke(id, () => {
    userStore.clearUser();
  });
  const result = await logOut();
  if (result.status !== 200) {
    throw new Error('Logout failed');
  }
};
</script>
