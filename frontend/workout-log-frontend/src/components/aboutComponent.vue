<template>
  <div class="row flex-lg-row-reverse align-items-center g-5 py-5">
    <div class="col-10 col-sm-8 col-lg-6">
      <img src="../assets/logo.jpg" class="d-block mx-lg-auto img-fluid"
      alt="Bootstrap Themes" width="500" height="300" loading="lazy">
    </div>
    <div class="col-lg-6">
      <h1 class="display-1">Gym log</h1>
      <p class="lead">{{ $t('info.description') }}</p>
      <div class="d-grid gap-2 d-md-flex justify-content-md-start">
        <button type="button" class="btn btn-dark btn-lg px-4 me-md-2"
        @click="this.$router.push('/pricing')">
          <font-awesome-icon icon="fa-solid fa-money-check-dollar" />
          {{ $t('links.pricing') }}
        </button>
        <button type="button" class="btn btn-outline-secondary btn-lg px-4"
        @click="this.$router.push('/features')">
          <font-awesome-icon icon="fa-solid fa-user-ninja" />
          {{ $t('links.features') }}</button>
      </div>
    </div>
  </div>
</template>
