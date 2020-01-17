import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Dashboard from "../views/Dashboard";
import Iniciar from "../components/Iniciar";
import Rank from "../components/Rank";

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path:'/dashboard',
    name:'dashboard',
    component: Dashboard
  },
  {
    path:'/login',
    name: 'login',
    component: Iniciar
  },
  {
    path:'/rank',
    name:'rank',
    component: Rank
  }
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
});

export default router;
