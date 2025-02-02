import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/measurements',
    name: 'measurements',
    component: () => import('../components/measurements.vue'),
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('../components/aboutComponent.vue'),
  },
  {
    path: '/features',
    name: 'features',
    component: () => import('../components/featuresPage.vue'),
  },
  {
    path: '/pricing',
    name: 'pricing',
    component: () => import('../components/pricingPage.vue'),
  },
  {
    path: '/faq',
    name: 'faq',
    component: () => import('../components/faqPage.vue'),
  },
  {
    path: '/profile',
    name: 'profile',
    component: () => import('../views/ProfilePage.vue'),
  },
  {
    path: '/trends',
    name: 'trends',
    component: () => import('../views/TrendsPage.vue'),
  },
  {
    path: '/:pathMatch(.*)*',
    component: () => import('../components/notFound.vue'),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
