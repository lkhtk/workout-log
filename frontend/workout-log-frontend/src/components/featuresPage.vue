<template>
  <div class="d-flex h-100 text-center">
    <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
      <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
        <h1 class="display-1">Features</h1>
      </div>
      <main>
        <div class="row row-cols-1 row-cols-md-2 align-items-md-center g-5 py-5">
          <div class="col d-flex flex-column align-items-start gap-2">
            <h2 class="fw-bold text-body-emphasis">A dead simple</h2>
            <p class="text-body-secondary">gym journal for anyone and everyone</p>
            <div v-if="!user">
              <auth-button />
            </div>
            <div v-else>
              <button class="btn btn-primary btn-lg"
                @click="this.$router.push('/')">
                <font-awesome-icon icon="fa-solid fa-dumbbell" />
                let's go!
              </button>
            </div>
          </div>
          <div class="col">
            <div class="row row-cols-1 row-cols-sm-2 g-4">
              <div class="col d-flex flex-column gap-2">
                <div class="feature-icon-small d-inline-flex align-items-center
                justify-content-center fs-4 rounded-3">
                <font-awesome-icon icon="fa-solid fa-credit-card" />
              </div>
              <h4 class="fw-semibold mb-0 text-body-emphasis">It's free!</h4>
              <p class="text-body-secondary">Wait, what? - Yeah, it's totally free!
                <a href="/pricing" class="icon-link">Read more ></a>
              </p>
            </div>
            <div class="col d-flex flex-column gap-2">
              <div class="feature-icon-small d-inline-flex align-items-center
              justify-content-center fs-4 rounded-3">
                <font-awesome-icon icon="fa-solid fa-code" />
              </div>
              <h4 class="fw-semibold mb-0 text-body-emphasis">Open Source</h4>
              <p class="text-body-secondary"><a href="https://github.com/lkhtk/workout-log" target="_blank">Source code ></a></p>
            </div>
            <div class="col d-flex flex-column gap-2">
              <div class="feature-icon-small d-inline-flex align-items-center
              justify-content-center fs-4 rounded-3">
                <font-awesome-icon icon="fa-solid fa-baby" />
              </div>
              <h4 class="fw-semibold mb-0 text-body-emphasis">Easy to use!</h4>
              <p class="text-body-secondary">It's simple.
                Quick and easy recording of training sessions</p>
            </div>
            <div class="col d-flex flex-column gap-2">
              <div class="feature-icon-small d-inline-flex align-items-center
              justify-content-center fs-4 rounded-3">
              <font-awesome-icon icon="fa-solid fa-ban" />
              </div>
              <h4 class="fw-semibold mb-0 text-body-emphasis">No bullshit</h4>
              <p class="text-body-secondary">No ads, no opt-out cookies, no pop-ups,
                no chatbots, no clickbait (okay, except for a little bit).</p>
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

document.title = 'Workout Log';
const userStore = useUserStore();
const { user } = storeToRefs(userStore);
</script>
