<template>
  <div class="max-w-md mx-auto mt-16 p-8 border border-gray-300 rounded-lg bg-white shadow">
    <h2 class="text-2xl font-bold mb-6 text-center text-black">Login</h2>
    <form @submit.prevent="handleLogin" class="space-y-4">
      <div>
        <label for="username" class="block mb-1 font-medium">Username:</label>
        <input id="username" v-model="username" type="text" required
          class="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400" />
      </div>
      <div>
        <label for="password" class="block mb-1 font-medium">Password:</label>
        <input id="password" v-model="password" type="password" required
          class="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400" />
      </div>
      <button type="submit"
        class="w-full py-2 bg-blue-600 hover:bg-blue-700 text-white rounded font-semibold text-lg transition">Login
      </button>
    </form>
    <a href="/Test2-2" class="text-blue-600 hover:underline">Register</a>
    <div v-if="error" class="text-red-600 mt-4 text-center">{{ error }}</div>
  </div>
</template>

<script>
import router from '../router';

export default {
  name: 'LoginPage',
  data() {
    return {
      username: '',
      password: '',
      error: ''
    }
  },
  methods: {
    handleLogin() {
      // ตัวอย่างการตรวจสอบ username/password แบบง่าย
      if (this.username === 'admin' && this.password === '1234') {
        alert('Login Success!');
        this.error = '';
        router.push('/test2-3'); // เปลี่ยนเส้นทางไปยังหน้า Test2-3 หลังจากล็อกอินสำเร็จ
      } else {
        this.error = 'Username หรือ Password ไม่ถูกต้อง';
      }
    }
  }
}
</script>
