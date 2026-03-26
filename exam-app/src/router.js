import { createRouter, createWebHistory } from 'vue-router'
import Test1 from './pages/Test1.vue'
import Test2 from './pages/Test2.vue'
import Test3 from './pages/Test3.vue'
import Test4 from './pages/Test4.vue'
import Test5 from './pages/Test5.vue'
import App from './App.vue'

const routes = [
  { path: '/', component: App },
  { path: '/test1', component: Test1 },
  { path: '/test2', component: Test2 },
  { path: '/test3', component: Test3 },
  { path: '/test4', component: Test4 },
  { path: '/test5', component: Test5 }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
