import axios from 'axios';
import router from '@/router';
import { useUserStore } from '../stores/userStore';

const api = axios.create({
  host: '127.0.0.1',
  port: '8000',
  withCredentials: true,
  timeout: 3000,
});
api.interceptors.response.use(
  (response) => response,
  (error) => {
    const userStore = useUserStore();
    const status = error.response?.status;
    switch (status) {
      case 401:
        userStore.logout?.();
        router.push('/about');
        break;
      case 400:
        window.$toast?.showToast('Bad request', 'danger');
        break;
      case 500:
        window.$toast?.showToast('Internal Server Error', 'danger');
        break;
      default:
        window.$toast?.showToast(status, 'info');
        break;
    }
    return Promise.reject(error);
  },
);

export default api;
