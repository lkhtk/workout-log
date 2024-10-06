import axios from 'axios';

export default axios.create({
  withCredentials: false,
  timeout: 3000,
});
