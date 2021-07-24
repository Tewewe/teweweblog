import Vue from 'vue'
import axios from 'axios'

let Url = 'http://localhost:3000/api/v1'

axios.defaults.baseURL = Url

// 添加拦截器
axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})

Vue.prototype.$http = axios

export { Url }
