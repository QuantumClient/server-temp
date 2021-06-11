import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://localhost:8080/api/v1/',
  timeout: 1500,
  headers: {
    common: {
      'X-Requested-With': 'XMLHttpRequest'
    }
  }
});

export default instance;
