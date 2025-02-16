<template>
  <div class="row">
    <div class="col">
      <userProfile />
    </div>
    <div class="col">
      <userActions />
  </div>
  </div>
</template>

<script setup>
import userProfile from '../components/profile.vue';
import userActions from '../components/userActions.vue';
</script>
