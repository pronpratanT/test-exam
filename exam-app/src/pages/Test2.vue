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
      <button type="submit" :disabled="isLoading"
        class="w-full py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-blue-400 disabled:cursor-not-allowed text-white rounded font-semibold text-lg transition flex justify-center items-center">
        <svg v-if="isLoading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        {{ isLoading ? 'กำลังเข้าสู่ระบบ...' : 'Login' }}
      </button>
    </form>
    <a href="/Test2-2" class="text-blue-600 hover:underline">Register</a>
    <div v-if="error" class="text-red-600 mt-4 text-center">{{ error }}</div>
  </div>
</template>

<script>
import router from '../router';
import axios from 'axios';

export default {
  name: 'LoginPage',
  data() {
    return {
      username: '',
      password: '',
      error: '',
      isLoading: false
    }
  },
  methods: {
    async handleLogin() {
      this.error = '';
      this.isLoading = true;

      const payload = {
        username: this.username,
        password: this.password
      }

      try {
        // TODO: เปลี่ยน URL ด้านล่างให้เป็น API ของคุณ
        const response = await axios.post('/api/test/test2/signin', payload);
        localStorage.setItem('username', response.data.username);
        localStorage.setItem('token', response.data.token);
        alert('Login Success!');
        
        // ตัวอย่างการเก็บ Token ที่ได้จาก API (ถ้ามีนำไปใช้งานต่อ)
        // localStorage.setItem('token', response.data.token);
        
        router.push('/test2-3'); // เปลี่ยนเส้นทางไปยังหน้าถัดไปเมื่อล็อกอินสำเร็จ
      } catch (err) {
        // ตรวจสอบว่า API ส่ง Error Message กลับมาหรือไม่
        if (err.response && err.response.data) {
          this.error = err.response.data.message || 'Username หรือ Password ไม่ถูกต้อง';
        } else {
          // กรณี API ล่ม หรือเน็ตหลุด
          this.error = 'ไม่สามารถเชื่อมต่อกับเซิร์ฟเวอร์ได้ (Network Error)';
          console.error(err);
        }
      } finally {
        this.isLoading = false; // ปิดสถานะโหลดเสมอไม่ว่าจะสำเร็จหรือ Error
      }
    }
  }
}
</script>
