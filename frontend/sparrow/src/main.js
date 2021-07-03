import Vue from 'vue'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'
import router from './router'
import './plugins/element.js'
import './assets/css/global.css'
import './assets/icon/iconfont.css'
import store from './store/store'
import axios from 'axios'
Vue.prototype.$http = axios
// axios.defaults.baseURL = config.backendurl
axios.defaults.timeout = 10000
// axios.defaults.withCredentials = true
axios.defaults.headers.post['Content-Type'] =
  'application/x-www-form-urlencoded;charset=UTF-8'
// axios拦截器
axios.interceptors.request.use(config=>{
    config.headers.authorization = window.localStorage.getItem('token')
  return config
},error => {
  return Promise.reject(error)
})
axios.interceptors.response.use(response => {
  return response
},error => {
  if (error.response) {
    switch (error.response.status) {
      // 返回403，清除token信息并跳转到登录页面
      case 403:
        localStorage.removeItem('token')
        localStorage.removeItem('username')
        router.replace({
          path: '/login'
          // 登录成功后跳入浏览的当前页面
  // query: {redirect: router.currentRoute.fullPath}
        })
    }
// 返回接口返回的错误信息
    return Promise.reject(error.response.data)
  }
})
Vue.config.productionTip = false

Vue.prototype.modifyArticleFun = function (aid) {
  this.$router.push({ name: 'wikiedit', params: { docCode:aid } })
}
Vue.prototype.gotodoArtilceDescFun = function (aid) {
  this.$router.push({ name: 'articledesc', params: { docCode: aid } })
}
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
