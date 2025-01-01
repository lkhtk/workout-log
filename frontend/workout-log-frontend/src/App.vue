<template>
  <header data-bs-theme="dark" class="p-4">
    <nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
      <div class="container-fluid">
        <a class="navbar-brand" href="#"><font-awesome-icon icon="fa-solid fa-dumbbell" />
          Workout log
        </a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
          data-bs-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false"
          aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarCollapse">
          <ul class="navbar-nav me-auto mb-2 mb-md-0">
            <li class="nav-item" v-if="user">
              <a class="nav-link active" aria-current="page" href="#">
                <router-link to="/">Journal</router-link>
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">
                <router-link to="/about">About</router-link>
              </a>
            </li>
          </ul>
          <ul class="navbar-nav d-flex" v-if="!user">
            <li class="nav-item d-flex">
              <button class="btn btn-outline-primary" :disabled="!isReady" @click="() => login()">
                <font-awesome-icon icon="fa-solid fa-arrow-right-to-bracket" />
                Log in
              </button>
            </li>
          </ul>
        </div>
        <div class="dropdown dropstart" v-if="user">
          <a href="#" class="d-block link-body-emphasis text-decoration-none dropdown-toggle"
            data-bs-toggle="dropdown" aria-expanded="false">
            <img :src=user.picture alt="pic" width="32" height="32" class="rounded-circle">
          </a>
          <ul class="dropdown-menu" style="">
            <li>
              <router-link class="dropdown-item" to="/profile">
                <font-awesome-icon icon="fa-solid fa-address-card" />
                Profile
              </router-link>
            </li>
            <li><hr class="dropdown-divider"></li>
            <li>
              <button class="dropdown-item" @click="revoke(user.id)">
                <font-awesome-icon icon="fa-solid fa-arrow-right-from-bracket" />
                Sign out
              </button>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  </header>
  <div class="shadow-lg p-3 mb-5 bg-body-tertiary rounded">
    <div class="container">
      <router-view/>
    </div>
  </div>
  <footer class="container d-flex flex-wrap justify-content-between align-items-center
    py-3 my-4 border-top">
    <p class="col-md-4 mb-0 text-body-secondary">&copy; 2025</p>
    <nav style="--bs-breadcrumb-divider: '';" aria-label="breadcrumb">
      <ol class="breadcrumb">
        <li class="breadcrumb-item" v-if="user">
          <router-link class="nav-item" to="/">Journal</router-link>
        </li>
        <li class="breadcrumb-item">
          <router-link class="nav-item" to="/features">Features</router-link>
        </li>
        <li class="breadcrumb-item">
          <router-link class="nav-item" to="/pricing">Pricing</router-link>
        </li>
        <li class="breadcrumb-item">
          <router-link class="nav-item" to="/faq">FAQ</router-link>
        </li>
        <li class="breadcrumb-item">
          <router-link class="nav-item" to="/about">About</router-link>
        </li>
      </ol>
    </nav>
  </footer>
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
} from './api/api';

import { useUserStore } from './stores/userStore';

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
