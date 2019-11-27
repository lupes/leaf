import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import(/* webpackChunkName: "Problems" */ '../views/Problems.vue')
  },
  {
    path: '/problems',
    name: 'problems',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "Problems" */ '../views/Problems.vue')
  },
  {
    path: '/problem/:id',
    name: 'problem',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "Problem" */ '../views/Problem.vue'),
  },
  {
    path: '/solution/add/:id',
    name: 'solutionAdd',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "SolutionAdd" */ '../views/SolutionAdd.vue'),
  },
  {
    path: '/solution/detail/:id',
    name: 'solutionDetail',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "SolutionDetail" */ '../views/SolutionDetail.vue'),
  },
  {
    path: '/solution/edit/:id',
    name: 'solutionEdit',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "SolutionEdit" */ '../views/SolutionEdit.vue'),
  },
  {
    path: '/ace',
    name: 'ace',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "ace" */ '../views/Ace.vue'),
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
]

const router = new VueRouter({
  mode: 'history',
  routes: routes
})

export default router
