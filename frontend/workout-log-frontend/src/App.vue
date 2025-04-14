<template>
  <header data-bs-theme="dark" class="p-4">
  <nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">
        <font-awesome-icon icon="fa-solid fa-dumbbell"/>
        Gym Log
      </a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
        data-bs-target="#navbarCollapse"
        aria-controls="navbarCollapse" aria-expanded="false">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarCollapse">
        <!-- AUTH -->
          <ul class="navbar-nav me-auto mb-2 mb-md-0" v-if="user">
            <li class="nav-item">
              <a aria-current="page" href="#" class="nav-link">
                <router-link to="/dairy" class="link-light">
                  {{ $t('links.journal') }}</router-link>
              </a>
            </li>
            <li class="nav-item">
              <a aria-current="page" href="#" class="nav-link">
                <router-link to="/measurements" class="link-light">
                  {{ $t('measurements.title') }}</router-link>
              </a>
            </li>
            <li class="nav-item">
              <a aria-current="page" href="#" class="nav-link">
                <router-link to="/trends" class="link-light">
                  {{ $t('trends.title') }}</router-link>
              </a>
            </li>
          </ul>
          <ul class="navbar-nav d-flex">
            <li class="nav-item d-flex">
              <auth-button loginButtonClass="btn btn-outline-light"
              logoutButtonClass="btn btn-outline-light"/>
            </li>
          </ul>
        <div
          class="dropdown ms-auto ms-md-0"
          v-if="user">
          <a href="#"
            class="dropdown-toggle d-block text-end"
            data-bs-toggle="dropdown"
            aria-expanded="false">
            <img :src=user.picture alt="pic" width="32" height="32" class="rounded-circle">
          </a>
          <ul
            class="dropdown-menu dropdown-menu-end"
            data-bs-boundary="viewport">
            <li>
              <router-link class="dropdown-item" to="/profile">
                <font-awesome-icon icon="fa-solid fa-address-card" />
                {{ $t('links.profile') }}
              </router-link>
            </li>
            <li>
              <hr class="dropdown-divider">
            </li>
            <li>
              <auth-button />
            </li>
          </ul>
        </div>
      </div>
    </div>
  </nav>
</header>
  <ToastComponent ref="toastRef" />
  <div class="shadow-lg p-3 bg-body-tertiary rounded">
    <div class="container">
      <router-view />
    </div>
  </div>
  <footer class="container d-flex flex-wrap justify-content-between align-items-center
    py-3 my-4 border-top">
    <p class="col-md-4 mb-0 text-body-secondary">Copyright &copy; Gym log</p>
    <nav style="--bs-breadcrumb-divider: '';" aria-label="breadcrumb">
      <ol class="breadcrumb">
        <li class="breadcrumb-item" v-if="user">
          <router-link class="nav-item link-body-emphasis" to="/dairy">
            {{ $t('links.journal') }}</router-link>
        </li>
        <li class="breadcrumb-item" v-if="user">
          <router-link class="nav-item link-body-emphasis" to="/measurements">
            {{ $t('measurements.title') }}
          </router-link>
        </li>
        <li class="breadcrumb-item" v-if="user">
          <router-link class="nav-item link-body-emphasis" to="/trends">
            {{ $t('trends.title') }}
          </router-link>
        </li>
        <!-- UNAUTH -->
        <li class="breadcrumb-item">
          <router-link class="nav-item link-body-emphasis" to="/features">
            {{ $t('links.features') }}</router-link>
        </li>
        <li class="breadcrumb-item">
          <router-link class="nav-item link-body-emphasis" to="/pricing">
            {{ $t('links.pricing') }}</router-link>
        </li>
        <li class="breadcrumb-item">
          <router-link class="nav-item link-body-emphasis" to="/faq">FAQ</router-link>
        </li>
        <li class="breadcrumb-item">
          <router-link class="nav-item link-body-emphasis" to="/about">
            {{ $t('links.about') }}</router-link>
        </li>
        <li class="breadcrumb-item">
          <userLang />
        </li>
      </ol>
    </nav>
  </footer>
</template>
<script setup>
import { storeToRefs } from 'pinia';
import { onMounted, ref } from 'vue';
import { useUserStore } from './stores/userStore';
import AuthButton from './components/common/AuthButton.vue';
import userLang from './components/common/userLang.vue';
import ToastComponent from './components/common/toastComponent.vue';

const toastRef = ref(null);

document.title = 'Gym Log';
const userStore = useUserStore();
const { user } = storeToRefs(userStore);
onMounted(() => {
  window.$toast = toastRef.value;
});
</script>
