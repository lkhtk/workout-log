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
                    <h2 class="accordion-header" id="faqHeadingOne">
                        <button class="accordion-button" type="button" data-bs-toggle="collapse"
            data-bs-target="#faqCollapseOne" aria-expanded="true" aria-controls="faqCollapseOne">
                            What is this service about?
                        </button>
                    </h2>
<div id="faqCollapseOne" class="accordion-collapse collapse show"
aria-labelledby="faqHeadingOne" data-bs-parent="#faqAccordion">
                        <div class="accordion-body">
            This service provides a platform to answer your
            frequently asked questions and help you navigate our features.
                        </div>
                    </div>
                </div>
                <div class="accordion-item">
                    <h2 class="accordion-header" id="faqHeadingTwo">
                <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse"
            data-bs-target="#faqCollapseTwo" aria-expanded="false" aria-controls="faqCollapseTwo">
                            How do I create an account?
                        </button>
                    </h2>
<div id="faqCollapseTwo" class="accordion-collapse collapse"
aria-labelledby="faqHeadingTwo" data-bs-parent="#faqAccordion">
                        <div class="accordion-body">
                    You can create an account by clicking on the "Log in" button at the top-right
                            corner of the page.
                        </div>
                    </div>
                </div>
                <div class="accordion-item">
                    <h2 class="accordion-header" id="faqHeadingThree">
                <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse"
        data-bs-target="#faqCollapseThree" aria-expanded="false" aria-controls="faqCollapseThree">
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
            </div>
          </div>
        </main>
      </div>
    </div>
  </template>
