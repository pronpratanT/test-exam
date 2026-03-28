<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-900 via-indigo-950 to-slate-900 p-6">
    <div class="max-w-md w-full bg-white rounded-[2.5rem] shadow-2xl p-10 relative overflow-hidden border border-white/10">
      <!-- Decorative circle -->
      <div class="absolute -top-24 -right-24 w-48 h-48 bg-indigo-500/10 rounded-full blur-3xl"></div>
      <div class="absolute -bottom-24 -left-24 w-48 h-48 bg-blue-500/10 rounded-full blur-3xl"></div>

      <div class="relative z-10 w-full">
        <div class="flex flex-col items-center mb-10">
          <div class="w-16 h-16 bg-gradient-to-tr from-indigo-600 to-blue-500 rounded-2xl shadow-lg shadow-indigo-500/30 flex items-center justify-center text-white mb-6">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path></svg>
          </div>
          <h2 class="text-3xl font-extrabold text-slate-800 tracking-tight">Welcome Back</h2>
          <p class="text-slate-500 mt-2 text-sm font-medium">Please enter your details to sign in.</p>
        </div>

        <form @submit.prevent="handleLogin" class="space-y-6">
          <div class="space-y-1.5">
            <label for="username" class="block text-sm font-bold text-slate-700 ml-1">Username</label>
            <div class="relative">
              <span class="absolute left-4 top-3.5 text-slate-400">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path></svg>
              </span>
              <input id="username" v-model="username" type="text" required placeholder="Enter your username"
                class="w-full pl-12 pr-4 py-3.5 border border-slate-200 focus:border-indigo-500 focus:ring-4 focus:ring-indigo-500/10 rounded-2xl outline-none transition-all bg-slate-50 hover:bg-white focus:bg-white text-slate-700 font-medium" />
            </div>
          </div>

          <div class="space-y-1.5">
            <label for="password" class="block text-sm font-bold text-slate-700 ml-1">Password</label>
            <div class="relative">
              <span class="absolute left-4 top-3.5 text-slate-400">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path></svg>
              </span>
              <input id="password" v-model="password" type="password" required placeholder="••••••••"
                class="w-full pl-12 pr-4 py-3.5 border border-slate-200 focus:border-indigo-500 focus:ring-4 focus:ring-indigo-500/10 rounded-2xl outline-none transition-all bg-slate-50 hover:bg-white focus:bg-white text-slate-700 font-medium" />
            </div>
          </div>

          <button type="submit" :disabled="isLoading"
            class="w-full py-4 bg-gradient-to-r from-indigo-600 to-blue-600 hover:from-indigo-500 hover:to-blue-500 disabled:from-slate-400 disabled:to-slate-500 disabled:cursor-not-allowed text-white rounded-2xl font-bold text-lg transition-all shadow-xl shadow-indigo-500/25 hover:-translate-y-0.5 active:scale-95 flex justify-center items-center gap-3 mt-4">
            <svg v-if="isLoading" class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ isLoading ? 'Signing in...' : 'Sign In' }}
          </button>
        </form>

        <div class="mt-8 text-center pt-6 border-t border-slate-100">
          <p class="text-slate-500 text-sm">Don't have an account? 
            <a href="/Test2-2" class="text-indigo-600 font-bold hover:text-indigo-500 transition-colors ml-1">Create Account</a>
          </p>
        </div>

        <div v-if="error" class="text-rose-600 mt-6 p-4 bg-rose-50 rounded-2xl font-medium text-center border border-rose-100 flex items-center justify-center gap-2 text-sm animate-bounce">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
          {{ error }}
        </div>
      </div>
    </div>
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
