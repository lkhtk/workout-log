import axios from 'axios';

export default axios.create({
  withCredentials: false,
  timeout: 3000,
  host: '127.0.0.1',
  port: '8000',
});
