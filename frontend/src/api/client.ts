import axios from 'axios'

export default axios.create({
  baseURL: import.meta.env.VITE_SERVICE_API,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
  withCredentials: true,
  timeout: 5000,
})
