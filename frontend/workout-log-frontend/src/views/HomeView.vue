<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <LogIn msg="login page"/>
  </div>
</template>

<script>
// @ is an alias to /src
import LogIn from '@/components/Login.vue';

export default {
  name: 'HomeView',
  components: {
    LogIn,
  },
};
</script>
