<template>
    <div class="d-flex h-100 text-center">
      <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
        <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
          <h1 class="display-1">{{ $t('faq.title') }}</h1>
        </div>
        <main>
          <div class="container">
            <div class="faq-header text-center">
                <p class="text-muted">{{ $t('faq.header') }}</p>
            </div>
            <div class="accordion" id="faqAccordion">
                <div class="accordion-item">
                    <h2 class="accordion-header" id="faqHeadingTwo">
                        <button class="accordion-button collapsed" type="button"
                        data-bs-toggle="collapse"
                        data-bs-target="#faqCollapseTwo" aria-expanded="true"
                        aria-controls="faqCollapseTwo">{{ $t('faq.q1') }}</button>
                    </h2>
                    <div id="faqCollapseTwo" class="accordion-collapse collapse"
                    aria-labelledby="faqHeadingTwo" data-bs-parent="#faqAccordion">
                        <div class="accordion-body">
                            {{ $t('faq.a1') }}
                        </div>
                    </div>
                </div>
                <div class="accordion-item">
                    <h2 class="accordion-header" id="faqHeadingOne">
                        <button class="accordion-button collapsed"
                        type="button" data-bs-toggle="collapse"
                        data-bs-target="#faqCollapseOne" aria-expanded="false"
                        aria-controls="faqCollapseOne">{{ $t('faq.q2') }}</button>
                    </h2>
                    <div id="faqCollapseOne" class="accordion-collapse collapse"
                    aria-labelledby="faqHeadingOne" data-bs-parent="#faqAccordion">
                        <div class="accordion-body">{{ $t('faq.a2') }}</div>
                    </div>
                </div>

                <div class="accordion-item">
                    <h2 class="accordion-header" id="faq6">
                        <button class="accordion-button collapsed" type="button"
                        data-bs-toggle="collapse"
                        data-bs-target="#faqCol6" aria-expanded="false"
                        aria-controls="faqCol6">{{ $t('faq.q3') }}</button>
                    </h2>
                    <div id="faqCol6" class="accordion-collapse collapse"
                    aria-labelledby="faq6" data-bs-parent="#faqAccordion">
                        <div class="accordion-body">{{ $t('faq.a3') }}</div>
                    </div>
                </div>
                <div class="accordion-item">
                    <h2 class="accordion-header" id="faq4">
                        <button class="accordion-button collapsed" type="button"
                        data-bs-toggle="collapse"
                        data-bs-target="#faqCollapseThree" aria-expanded="false"
                        aria-controls="faqCollapseThree">{{ $t('faq.q4') }}</button>
                    </h2>
                    <div id="faqCollapseThree" class="accordion-collapse collapse"
                    aria-labelledby="faqHeadingThree" data-bs-parent="#faqAccordion">
                        <div class="accordion-body">{{ $t('faq.a4') }}</div>
                    </div>
                </div>
                <div class="accordion-item">
                    <h2 class="accordion-header" id="faq4">
                        <button class="accordion-button collapsed" type="button"
                        data-bs-toggle="collapse"
                        data-bs-target="#faqCol4" aria-expanded="false"
                        aria-controls="faqCol4">{{ $t('faq.q5') }}</button>
                    </h2>
                    <div id="faqCol4" class="accordion-collapse collapse"
                    aria-labelledby="faq4" data-bs-parent="#faqAccordion">
                        <div class="accordion-body">{{ $t('faq.a5') }}</div>
                    </div>
                </div>
            </div>
          </div>
        </main>
      </div>
    </div>
  </template>
