<template lang="html">
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-tr from-slate-900 via-slate-800 to-indigo-950 p-6">
    <div class="max-w-md w-full bg-white rounded-[2.5rem] shadow-2xl p-10 relative overflow-hidden border border-white/10">
      <!-- Decorative circle -->
      <div class="absolute -top-24 -left-24 w-48 h-48 bg-emerald-500/10 rounded-full blur-3xl"></div>
      <div class="absolute -bottom-24 -right-24 w-48 h-48 bg-indigo-500/10 rounded-full blur-3xl"></div>

      <div class="relative z-10 w-full">
        <div class="flex flex-col items-center mb-10">
          <div class="w-16 h-16 bg-gradient-to-tr from-emerald-500 to-teal-400 rounded-2xl shadow-lg shadow-emerald-500/30 flex items-center justify-center text-white mb-6">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"></path></svg>
          </div>
          <h2 class="text-3xl font-extrabold text-slate-800 tracking-tight">Create Account</h2>
          <p class="text-slate-500 mt-2 text-sm font-medium">Join us today! It only takes a minute.</p>
        </div>

        <form @submit.prevent="handleRegister" class="space-y-5">
          <div class="space-y-1.5">
            <label for="username" class="block text-sm font-bold text-slate-700 ml-1">Username</label>
            <input id="username" v-model="username" type="text" required placeholder="Pick a unique username"
              class="w-full px-5 py-3.5 border border-slate-200 focus:border-emerald-500 focus:ring-4 focus:ring-emerald-500/10 rounded-2xl outline-none transition-all bg-slate-50 hover:bg-white focus:bg-white text-slate-700 font-medium" />
          </div>

          <div class="space-y-1.5">
            <label for="password" class="block text-sm font-bold text-slate-700 ml-1">Password</label>
            <input id="password" v-model="password" type="password" required placeholder="Create a strong password"
              class="w-full px-5 py-3.5 border border-slate-200 focus:border-emerald-500 focus:ring-4 focus:ring-emerald-500/10 rounded-2xl outline-none transition-all bg-slate-50 hover:bg-white focus:bg-white text-slate-700 font-medium" />
          </div>

          <div class="space-y-1.5">
            <label for="confirm-password" class="block text-sm font-bold text-slate-700 ml-1">Confirm Password</label>
            <input id="confirm-password" v-model="confirmPassword" type="password" required placeholder="Repeat your password"
              class="w-full px-5 py-3.5 border border-slate-200 focus:border-emerald-500 focus:ring-4 focus:ring-emerald-500/10 rounded-2xl outline-none transition-all bg-slate-50 hover:bg-white focus:bg-white text-slate-700 font-medium" />
          </div>

          <button type="submit" :disabled="isLoading"
            class="w-full py-4 bg-gradient-to-r from-emerald-500 to-teal-500 hover:from-emerald-600 hover:to-teal-600 disabled:from-slate-400 disabled:to-slate-500 disabled:cursor-not-allowed text-white rounded-2xl font-bold text-lg transition-all shadow-xl shadow-emerald-500/25 hover:-translate-y-0.5 active:scale-95 flex justify-center items-center gap-3 mt-6">
            <svg v-if="isLoading" class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ isLoading ? 'Creating Account...' : 'Register Now' }}
          </button>
        </form>

        <div class="mt-8 text-center pt-6 border-t border-slate-100">
          <p class="text-slate-500 text-sm">Already have an account? 
            <a href="/test2" class="text-emerald-600 font-bold hover:text-emerald-500 transition-colors ml-1">Sign In</a>
          </p>
        </div>

        <div v-if="error" class="text-rose-600 mt-6 p-4 bg-rose-50 rounded-2xl font-medium text-center border border-rose-100 flex items-center justify-center gap-2 text-sm">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
          {{ error }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import router from '../router';

export default {
  name: 'RegisterPage',
  data() {
    return {
      username: '',
      password: '',
      confirmPassword: '',  
      error: '',
      isLoading: false
    }
  },
  methods: {
    async handleRegister() {
      // 1. ตรวจสอบรหัสผ่านให้ตรงกันก่อนส่งไป API
      if (this.password !== this.confirmPassword) {
        this.error = 'Passwords do not match';
        return;
      }

      this.error = '';
      this.isLoading = true;

      const payload = {
        username: this.username,
        password: this.password,
        confirm_password: this.confirmPassword
      }

      try {
        // 2. ใช้ Relative Path ได้เลยเพราะเราจะตั้งค่า Vite Proxy ไว้แล้ว
        const response = await axios.post('/api/test/test2/signup', payload);
        
        alert('Register Success!');
        
        // 3. เมื่อสมัครเสร็จนำทางกลับไปหน้า Login
        router.push('/test2'); 
      } catch (err) {
        if (err.response && err.response.data) {
          // แจ้งเตือนข้อผิดพลาดจาก Backend เช่น "Username already exists"
          this.error = err.response.data.message || 'สมัครสมาชิกไม่สำเร็จ';
        } else {
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