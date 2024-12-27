<template>
  <nav class="navbar navbar-expand-lg bg-body-tertiary">
  <div class="container-fluid">
    <a class="navbar-brand" href="#"><font-awesome-icon icon="fa-solid fa-dumbbell" />
      Workout log
    </a>
    <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
      <div class="navbar-nav">
        <a class="nav-link active" aria-current="page" href="#">
          <router-link to="/">Home</router-link>
        </a>
        <a class="nav-link" href="#">
          <router-link to="/about">About</router-link>
        </a>
        <a class="nav-link" href="#">
          <router-link to="/login">LogIn</router-link>
        </a>
      </div>
    </div>
  </div>
</nav>
<div class="container-lg shadow-lg p-3 mb-5 bg-body-tertiary rounded">
  <router-view/>
</div>

<div class="container">
  <footer class="d-flex flex-wrap justify-content-between align-items-center py-3 my-4 border-top">
    <p class="col-md-4 mb-0 text-body-secondary">&copy; 2024 Company, Inc</p>
    <ul class="nav col-md-4 justify-content-end">
      <li class="nav-item"><a href="/Login" class="nav-link px-2 text-body-secondary">Login</a></li>
      <li class="nav-item"><a href="/" class="nav-link px-2 text-body-secondary">Home</a></li>
      <li class="nav-item"><a href="/features" class="nav-link px-2 text-body-secondary">
        Features</a></li>
      <li class="nav-item"><a href="/pricing" class="nav-link px-2 text-body-secondary">
        Pricing</a></li>
      <li class="nav-item"><a href="/faq" class="nav-link px-2 text-body-secondary">FAQs</a></li>
      <li class="nav-item"><a href="/about" class="nav-link px-2 text-body-secondary">About</a></li>
    </ul>
  </footer>
</div>
</template>
