import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// bootstrapvue
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import './assets/scss/index.scss'
// Install BootstrapVue
Vue.use(BootstrapVue)
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin)

// vuelidate
import Vuelidate from 'vuelidate'
Vue.use(Vuelidate)

// axios
import axios from 'axios'
import VueAxios from 'vue-axios'
Vue.use(VueAxios, axios)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
