<template>
    <div class="d-flex h-100 text-center" v-if="user">
        <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
            <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
                <h1 class="display-5">
                    Configs
                </h1>
            </div>
            <main>
              <form>
                <div class="mb-3">
                  <label for="exampleInputEmail1" class="form-label">Email address
                    <input type="email" class="form-control" id="exampleInputEmail1"
                    aria-describedby="emailHelp">
                  </label>
                  <div id="emailHelp" class="form-text">
                    We'll never share your email with anyone else.
                  </div>
                </div>
                <div class="mb-3">
                  <label for="exampleInputPassword1" class="form-label">Password
                  <input type="password" class="form-control" id="exampleInputPassword1">
                  </label>
                </div>
                <div class="mb-3 form-check">
                  <label class="form-check-label" for="exampleCheck1">Check me out
                    <input type="checkbox" class="form-check-input" id="exampleCheck1">
                  </label>
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
              </form>
            </main>
        </div>
    </div>
</template>

<script>
import { storeToRefs } from 'pinia';
import { useUserStore } from '../stores/userStore';

export default {
  name: 'userConfigs',
  setup() {
    const userStore = useUserStore();
    const { user } = storeToRefs(userStore);
    return {
      user,
    };
  },
  computed: {
    currentUser() {
      return this.user || {
        id: 0,
      };
    },
  },
};
</script>
