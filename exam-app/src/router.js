import { createRouter, createWebHistory } from 'vue-router'
import Test1 from './pages/Test1.vue'
import Test2 from './pages/Test2.vue'
import Test2_2 from './pages/Test2-2.vue'
import Test2_3 from './pages/Test2-3.vue'
import Test3 from './pages/Test3.vue'
import Test4 from './pages/Test4.vue'
import Test5 from './pages/Test5.vue'
import Test5_2 from './pages/Test5-2.vue'
import Home from './pages/Home.vue'
import App from './App.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/test1', component: Test1 },
  { path: '/test2', component: Test2 },
  { path: '/test2-2', component: Test2_2 },
  { path: '/test2-3', component: Test2_3 },
  { path: '/test3', component: Test3 },
  { path: '/test4', component: Test4 },
  { path: '/test5', component: Test5 },
  { path: '/test5-2', component: Test5_2 }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
