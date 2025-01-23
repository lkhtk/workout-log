<template>
  <header data-bs-theme="dark" class="p-4">
  <nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#"><font-awesome-icon icon="fa-solid fa-dumbbell" />
        Workout log
      </a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
        data-bs-target="#navbarCollapse"
        aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul class="navbar-nav me-auto mb-2 mb-md-0">
          <li class="nav-item" v-if="user">
            <a class="nav-link active" aria-current="page" href="#">
              <router-link to="/">Journal</router-link>
            </a>
          </li>
          <li class="nav-item" v-if="user">
            <a class="nav-link active" aria-current="page" href="#">
              <router-link to="/measurements">Measurements</router-link>
            </a>
          </li>
        </ul>
        <ul class="navbar-nav d-flex">
          <li class="nav-item d-flex">
            <auth-button />
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
                Profile
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
  <div class="shadow-lg p-3 bg-body-tertiary rounded">
    <div class="container">
      <router-view />
    </div>
  </div>
  <footer class="container d-flex flex-wrap justify-content-between align-items-center
    py-3 my-4 border-top">
    <p class="col-md-4 mb-0 text-body-secondary">Copyright &copy; Workout log</p>
    <nav style="--bs-breadcrumb-divider: '';" aria-label="breadcrumb">
      <ol class="breadcrumb">
        <li class="breadcrumb-item" v-if="user">
          <router-link class="nav-item" to="/">Journal</router-link>
        </li>
        <li class="breadcrumb-item" v-if="user">
          <router-link class="nav-item" to="/measurements">Measurements</router-link>
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

import { useUserStore } from './stores/userStore';
import AuthButton from './components/common/AuthButton.vue';

document.title = 'Workout Log';
const userStore = useUserStore();
const { user } = storeToRefs(userStore);
</script>
