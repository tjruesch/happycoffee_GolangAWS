import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import './assets/css/main.css'

import Vue from 'vue'

import VueRouter from 'vue-router'
import VueDialog from "vuejs-dialog"
import 'vuejs-dialog/dist/vuejs-dialog.min.css';
import VueToast from 'vue-toast-notification'
import 'vue-toast-notification/dist/theme-sugar.css'

import Buefy from 'buefy'

import Index from './pages/Index'
import Login from './pages/Login'

Vue.use(VueRouter)
Vue.use(Buefy)

Vue.use(VueDialog, {
    html: true,
    loader: false,
    okText: 'Proceed',
    cancelText: 'Cancel',
    animation: 'fade'
})

Vue.use(VueToast, {
    position: 'top-right',
    duration: 5000,
})

Vue.config.productionTip = false


const routes = [
    {
        path: '/',
        component: Index,
    },
    {
        path: '/login',
        component: Login,
    },
]

const router = new VueRouter({
    routes,
    mode: 'history',
})

const app = new Vue({
    router,
})

app.$mount('#app')