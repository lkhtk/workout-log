<template>
  <div class="d-flex h-100 text-center">
    <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
      <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
        <h1 class="display-1">{{ $t('links.pricing') }}</h1>
        <p class="fs-5 text-body-secondary">{{ $t('price.description') }}</p>
      </div>
      <main>
        <div class="row row-cols-1 row-cols-md-3 mb-3 text-center justify-content-around">
          <div class="col">
            <div class="card mb-4 rounded-3 shadow-sm">
              <div class="card-header py-3">
                <h4 class="my-0 fw-normal">{{ $t('price.p1_header') }}</h4>
              </div>
              <div class="card-body">
                <h1 class="card-title pricing-card-title">{{ $t('price.p1_price') }}
                  <small class="text-body-secondary fw-light">{{ $t('price.p1_period') }}</small>
                </h1>
                <ul class="list-unstyled mt-3 mb-4">
                  <li>{{ $t('price.p1_b1') }}</li>
                  <li>{{ $t('price.p1_b2') }}</li>
                  <li>{{ $t('price.p1_b3') }}</li>
                  <li>{{ $t('price.p1_b4') }}</li>
                </ul>
                <div>
                  <button class="btn btn-dark btn-lg"
                    @click="$router.push('/profile')">
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
