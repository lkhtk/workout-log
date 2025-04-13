<template>
  <div class="d-flex h-100 text-center">
    <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
      <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
        <h1 class="display-1">{{ $t('links.features') }}</h1>
      </div>
      <main>
        <div class="row row-cols-1 row-cols-md-2 align-items-md-center g-5 py-5">
          <div class="col d-flex flex-column align-items-start gap-2">
            <h2 class="fw-bold text-body-emphasis">{{ $t('features.title') }}</h2>
            <p class="text-body-secondary">{{ $t('features.header') }}</p>
            <div>
              <button class="btn btn-dark btn-lg"
                @click="$router.push('/profile')">
                <font-awesome-icon icon="fa-solid fa-dumbbell" />
                {{ $t('buttons.go') }}
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
              <h4 class="fw-semibold mb-0 text-body-emphasis">{{ $t('features.feature1_h') }}</h4>
              <p class="text-body-secondary">{{ $t('features.feature1_d') }}
                <a href="/pricing" class="icon-link link-body-emphasis">
                  {{ $t('features.read_more') }}</a>
              </p>
            </div>
            <div class="col d-flex flex-column gap-2">
              <div class="feature-icon-small d-inline-flex align-items-center
              justify-content-center fs-4 rounded-3">
                <font-awesome-icon icon="fa-solid fa-code" />
              </div>
              <h4 class="fw-semibold mb-0 text-body-emphasis">{{ $t('features.feature2_h') }}</h4>
              <p class="text-body-secondary">{{ $t('features.feature2_d') }}
                <a href="https://github.com/lkhtk/workout-log" target="_blank" class="icon-link link-body-emphasis">
                  {{ $t('features.read_more') }}</a>
              </p>
            </div>
            <div class="col d-flex flex-column gap-2">
              <div class="feature-icon-small d-inline-flex align-items-center
              justify-content-center fs-4 rounded-3">
                <font-awesome-icon icon="fa-solid fa-baby" />
              </div>
              <h4 class="fw-semibold mb-0 text-body-emphasis">{{ $t('features.feature3_h') }}</h4>
              <p class="text-body-secondary">{{ $t('features.feature3_d') }}</p>
            </div>
            <div class="col d-flex flex-column gap-2">
              <div class="feature-icon-small d-inline-flex align-items-center
              justify-content-center fs-4 rounded-3">
              <font-awesome-icon icon="fa-solid fa-ban" />
              </div>
              <h4 class="fw-semibold mb-0 text-body-emphasis">{{ $t('features.feature4_h') }}</h4>
              <p class="text-body-secondary">{{ $t('features.feature4_d') }}</p>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>
