<template>
  <div class="row flex-lg-row-reverse align-items-center g-5 py-5">
    <div class="col-10 col-sm-8 col-lg-6">
      <img src="https://upload.wikimedia.org/wikipedia/ru/9/94/%D0%93%D0%B8%D0%B3%D0%B0%D1%87%D0%B0%D0%B4.jpg" class="d-block mx-lg-auto img-fluid"
      alt="Bootstrap Themes" width="500" height="300" loading="lazy">
    </div>
    <div class="col-lg-6">
      <h1 class="display-1">Wokrout log</h1>
      <p class="lead">A free and simple ad-free service for keeping a diary of
        your gym sessions and analyzing your strength training progress. Don't miss
        the opportunity to make your workouts even more effective and interesting!</p>
      <div class="d-grid gap-2 d-md-flex justify-content-md-start">
        <button type="button" class="btn btn-primary btn-lg px-4 me-md-2"
        @click="this.$router.push('/pricing')">Pricing</button>
        <button type="button" class="btn btn-outline-secondary btn-lg px-4"
        @click="this.$router.push('/features')">Features</button>
      </div>
    </div>
  </div>
</template>
