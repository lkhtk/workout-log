<template>
  <div v-if="!user">
    <GoogleSignInButton
      @success="handleLoginSuccess"
      @error="handleLoginError"
    />
  </div>
  <div v-else>
    <h2>Welcome, {{ user.name }}</h2>
    <img :src="user.picture" alt="User picture" />
    <p>Email: {{ user.email }}</p>
    <button @click="revoke(user.id)">Sign out</button>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { GoogleSignInButton, decodeCredential, idRevoke } from 'vue3-google-signin';
import { checkToken, logOut } from '../api/api';

const user = ref(null);
const handleLoginSuccess = async (response) => {
  const { credential } = response;
  const result = await checkToken(credential);
  if (result.status !== 200) {
    throw new Error('Server validation failed');
  }
  const decodedCredential = decodeCredential(credential);
  user.value = {
    id: decodedCredential.id,
    name: decodedCredential.name,
    email: decodedCredential.email,
    picture: decodedCredential.picture,
  };
};
const handleLoginError = () => {
  console.error('Login failed');
};
const revoke = (id) => {
  idRevoke(id, (done) => {
    console.log(done.error);
    user.value = null;
  });
  logOut();
};
</script>
