<template>
  <div>
    <button
      v-if="!user"
      class="btn btn-outline-primary"
      :disabled="!isReady"
      @click="handleLogin"
    >
      <font-awesome-icon icon="fa-solid fa-arrow-right-to-bracket" />
      Log in
    </button>
    <button
      v-else
      class="dropdown-item"
      @click="handleLogout"
    >
      <font-awesome-icon icon="fa-solid fa-arrow-right-from-bracket" />
      Sign out
    </button>
  </div>
</template>

<script setup>
import { storeToRefs } from 'pinia';
import { useOneTap, idRevoke, decodeCredential } from 'vue3-google-signin';
import { checkToken, logOut } from '../../api/api';
import { useUserStore } from '../../stores/userStore';

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

const handleLogin = () => {
  login();
};

const handleLogout = async () => {
  try {
    await new Promise((resolve, reject) => {
      idRevoke(user.value.id, (done) => {
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
    this.$router.push('/about');
  } catch (error) {
    console.error('Logout error:', error.message);
    this.$router.push('/about');
  }
};
</script>
