<template>
  <div v-if="!user">
    <button :disabled="!isReady" @click="() => login()">
      Trigger One Tap Login Manually
    </button>
  </div>
  <div v-else>
    <h2>Welcome, {{ user.name }}</h2>
    <img :src="user.picture" alt="User picture" />
    <p>Email: {{ user.email }}</p>
    <button @click="revoke(user.id)">Sign out</button>
  </div>
</template>

<script setup>
import { storeToRefs } from 'pinia';
import {
  useOneTap,
  decodeCredential,
  idRevoke,
} from 'vue3-google-signin';

import {
  checkToken,
  logOut,
} from '../api/api';

import { useUserStore } from '../stores/userStore';

const userStore = useUserStore();
const { user } = storeToRefs(userStore);

const { isReady, login } = useOneTap({
  disableAutomaticPrompt: true,
  onSuccess: async (response) => {
    try {
      const { credential } = response;
      const result = await checkToken(credential);
      if (result.status !== 200) {
        throw new Error('Server validation failed');
      }
      const decodedCredential = decodeCredential(credential);
      const userData = {
        id: decodedCredential.id,
        name: decodedCredential.name,
        email: decodedCredential.email,
        picture: decodedCredential.picture,
      };
      userStore.setUser(userData);
    } catch (error) {
      console.error('Error in One Tap Login:', error.message);
    }
  },
  onError: () => console.error('Error with One Tap Login'),
});

const revoke = (id) => {
  idRevoke(id, () => {
    userStore.clearUser();
  });
  logOut();
};
</script>
