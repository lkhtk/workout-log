<template>
  <div class="d-flex h-100 text-center">
    <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
      <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
        <h1 class="display-1">{{ $t('links.pricing') }}</h1>
        <p class="fs-5 text-body-secondary">Dead simple plan unlocks everything.
          No bundles. No add-ons.</p>
      </div>
      <main>
        <div class="row row-cols-1 row-cols-md-3 mb-3 text-center justify-content-around">
          <div class="col">
            <div class="card mb-4 rounded-3 shadow-sm">
              <div class="card-header py-3">
                <h4 class="my-0 fw-normal">Free</h4>
              </div>
              <div class="card-body">
                <h1 class="card-title pricing-card-title">$0
                  <small class="text-body-secondary fw-light">/mo</small>
                </h1>
                <ul class="list-unstyled mt-3 mb-4">
                  <li>'As is' provided! <font-awesome-icon icon="fa-solid fa-check" /></li>
                  <li>No support! <font-awesome-icon icon="fa-solid fa-check" /></li>
                  <li>No guarantees! <font-awesome-icon icon="fa-solid fa-check" /></li>
                  <li>No gamification! <font-awesome-icon icon="fa-solid fa-check" /></li>
                </ul>
                <div v-if="!user">
                  <auth-button
                  loginButtonClass="btn btn-outline-dark"
                  logoutButtonClass="btn btn-outline-dark" />
                </div>
                <div v-else>
                  <button class="btn btn-dark btn-lg"
                    @click="$router.push('/')">
                    <font-awesome-icon icon="fa-solid fa-dumbbell" />
                    {{ $t('buttons.go') }}
                  </button>
                </div>
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

import { useUserStore } from '../stores/userStore';
import AuthButton from './common/AuthButton.vue';

document.title = 'Gym Log';
const userStore = useUserStore();
const { user } = storeToRefs(userStore);
</script>
