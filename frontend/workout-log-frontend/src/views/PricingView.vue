<template>
  <div class="d-flex h-100 text-center">
    <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">

        <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
      <h1 class="display-1">Pricing</h1>
      <p class="fs-5 text-body-secondary">Dead simple plan unlocks everything.
        No bundles. No add-ons.</p><button type="button" class="btn btn-link"
        @click="this.$router.push('/features')">All Features ></button>
    </div>
    <main>
    <div class="row row-cols-1 row-cols-md-3 mb-3 text-center justify-content-around">
      <div class="col">
        <div class="card mb-4 rounded-3 shadow-sm">
          <div class="card-header py-3">
            <h4 class="my-0 fw-normal">Free</h4>
          </div>
          <div class="card-body">
            <h1 class="card-title pricing-card-title">$0<small
              class="text-body-secondary fw-light">/mo</small></h1>
            <ul class="list-unstyled mt-3 mb-4">
              <li>'As is' provided! <font-awesome-icon icon="fa-solid fa-check" /></li>
              <li>No support! <font-awesome-icon icon="fa-solid fa-check" /></li>
              <li>No guarantees! <font-awesome-icon icon="fa-solid fa-check" /></li>
              <li>No gamification! <font-awesome-icon icon="fa-solid fa-check" /></li>
            </ul>
            <button type="button" class="w-100 btn btn-lg btn-outline-primary">
              Sounds great, I'm in!</button>
          </div>
        </div>
      </div>
    </div>
  </main>
</div>
  </div>
</template>
