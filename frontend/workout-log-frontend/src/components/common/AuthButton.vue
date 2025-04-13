<template>
  <div>
    <button
      v-if="!user"
      :class="loginButtonClass"
      @click="handleLogin"
    >
      <font-awesome-icon icon="fa-solid fa-arrow-right-to-bracket" />
      {{ $t('buttons.login') }}
    </button>
    <button
      v-else-if="user && loginButtonClass !== 'btn btn-outline-light'"
      :class="logoutButtonClass"
      @click="handleLogout"
    >
      <font-awesome-icon icon="fa-solid fa-arrow-right-from-bracket" />
      {{ $t('buttons.logout') }}
    </button>
  </div>
</template>

<script setup>
import { storeToRefs } from 'pinia';
import { useOneTap, idRevoke, decodeCredential } from 'vue3-google-signin';
import { useRouter } from 'vue-router';
import { checkToken, logOut } from '../../api/api';
import { useUserStore } from '../../stores/userStore';

defineProps({
  loginButtonClass: {
    type: String,
    default: 'btn btn-outline-primary',
  },
  logoutButtonClass: {
    type: String,
    default: 'dropdown-item',
  },
});

const userStore = useUserStore();
const { user } = storeToRefs(userStore);
const router = useRouter();
const { login } = useOneTap({
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
      window.$toast?.showToast(error.message);
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
          window.$toast?.showToast(done.error);
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
  } catch (error) {
    window.$toast?.showToast(error.message);
  }
  userStore.logout();
  router.push('/about');
};
</script>
