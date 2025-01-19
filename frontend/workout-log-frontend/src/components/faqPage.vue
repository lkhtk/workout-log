<template>
    <div class="d-flex h-100 text-center">
      <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
        <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
          <h1 class="display-1">Frequently Asked Questions</h1>
        </div>
        <main>
          <div class="container">
            <div class="faq-header text-center">
                <p class="text-muted">Find answers to common questions below</p>
            </div>
            <div class="accordion" id="faqAccordion">
                <div class="accordion-item">
                    <h2 class="accordion-header" id="faqHeadingTwo">
                        <button class="accordion-button collapsed" type="button"
                        data-bs-toggle="collapse"
                        data-bs-target="#faqCollapseTwo" aria-expanded="true"
                        aria-controls="faqCollapseTwo">
                        What makes your app different from other workout logs?
                        </button>
                    </h2>
                    <div id="faqCollapseTwo" class="accordion-collapse collapse"
                    aria-labelledby="faqHeadingTwo" data-bs-parent="#faqAccordion">
                        <div class="accordion-body">
                            Our app is minimalistic, completely free and extremely easy to use.
                            There are no adverts, complicated features or distractions - just a
                            convenient way to record your workouts and track your progress.
                        </div>
                    </div>
                </div>
                <div class="accordion-item">
                    <h2 class="accordion-header" id="faqHeadingOne">
                        <button class="accordion-button collapsed"
                        type="button" data-bs-toggle="collapse"
                        data-bs-target="#faqCollapseOne" aria-expanded="false"
                        aria-controls="faqCollapseOne">
                        Is the app really free? Are there any paid features?
                        </button>
                    </h2>
                    <div id="faqCollapseOne" class="accordion-collapse collapse"
                    aria-labelledby="faqHeadingOne" data-bs-parent="#faqAccordion">
                        <div class="accordion-body">
                            Yep, the app is completely free and requires no subscriptions.
                            You get full access to all features with no hidden fees.
                        </div>
                    </div>
                </div>

                <div class="accordion-item">
                    <h2 class="accordion-header" id="faq6">
                        <button class="accordion-button collapsed" type="button"
                        data-bs-toggle="collapse"
                        data-bs-target="#faqCol6" aria-expanded="false"
                        aria-controls="faqCol6">
                            How can I see my progress?
                        </button>
                    </h2>
                    <div id="faqCol6" class="accordion-collapse collapse"
                    aria-labelledby="faq6" data-bs-parent="#faqAccordion">
                        <div class="accordion-body">
                            Only on the scale and in the mirror
                        </div>
                    </div>
                </div>
                <div class="accordion-item">
                    <h2 class="accordion-header" id="faq4">
                        <button class="accordion-button collapsed" type="button"
                        data-bs-toggle="collapse"
                        data-bs-target="#faqCollapseThree" aria-expanded="false"
                        aria-controls="faqCollapseThree">
                            Who can I contact for support?
                        </button>
                    </h2>
                    <div id="faqCollapseThree" class="accordion-collapse collapse"
                    aria-labelledby="faqHeadingThree" data-bs-parent="#faqAccordion">
                        <div class="accordion-body">
                            There's no support here, you're alone.
                        </div>
                    </div>
                </div>
                <div class="accordion-item">
                    <h2 class="accordion-header" id="faq4">
                        <button class="accordion-button collapsed" type="button"
                        data-bs-toggle="collapse"
                        data-bs-target="#faqCol4" aria-expanded="false"
                        aria-controls="faqCol4">
                        Is the app suitable for those who are new to working out?
                        </button>
                    </h2>
                    <div id="faqCol4" class="accordion-collapse collapse"
                    aria-labelledby="faq4" data-bs-parent="#faqAccordion">
                        <div class="accordion-body">
                            Of course it is! The simplicity of the interface makes it ideal for
                            beginners. You can easily record basic exercises and track your results
                            without being overwhelmed by unnecessary details.
                        </div>
                    </div>
                </div>
            </div>
          </div>
        </main>
      </div>
    </div>
  </template>
