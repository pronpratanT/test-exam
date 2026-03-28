<template>
  <div class="flex flex-col items-center justify-center h-screen bg-gradient-to-br from-slate-50 to-indigo-50 gap-10">
      <div class="flex flex-col items-center p-12 bg-white rounded-[2.5rem] shadow-2xl shadow-indigo-100/50 border border-white/50 backdrop-blur-sm min-w-[400px]">
        <button class="px-12 py-6 mb-8 bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-500 hover:to-indigo-500 shadow-xl shadow-blue-500/20 text-white rounded-2xl font-bold text-2xl transition-all duration-300 transform hover:scale-105 active:scale-95 w-full" 
        @click="handleCreateTicket">รับบัตรคิว</button>
        
        <div class="relative flex flex-col items-center justify-center w-64 h-64 bg-slate-50 rounded-full shadow-inner border-8 border-white mb-8">
            <span class="absolute top-8 text-xs font-bold text-slate-400 tracking-widest uppercase">คิวปัจจุบัน</span>
            <p class="text-center text-8xl font-black text-transparent bg-clip-text bg-gradient-to-br from-slate-800 to-slate-500 drop-shadow-sm mt-4">{{ lastTicket }}</p>
        </div>

        <button
          class="px-10 py-3 mt-2 bg-white border-2 border-slate-200 hover:border-rose-400 hover:bg-rose-50 text-slate-600 hover:text-rose-500 rounded-xl font-semibold text-lg transition-all duration-300 w-full max-w-[200px]"
          :class="{ 'cursor-not-allowed opacity-40 hover:border-slate-200 hover:bg-white hover:text-slate-600': lastTicket === '00' }"
          @click="handleClear"
          :disabled="lastTicket === '00'"
        >ล้างคิว</button>
      </div>
  </div>
</template>

<script>
import router from '../router';
import axios from 'axios';

export default {
  name: 'Test5Page',
  data() {
    return {
      lastTicket: '00', // กำหนดค่าเริ่มต้นของหมายเลขคิวล่าสุด
    };
  },
  mounted() {
    // เมื่อหน้าโหลดขึ้นมา ให้เรียก API เพื่อดึงหมายเลขคิวล่าสุด
    this.fetchLastTicket();
  },
  methods: {
    async fetchLastTicket() {
      try {
        const response = await axios.get('/api/test/test5/get-last-ticket');
        const data = response.data;
        // ทำอะไรบางอย่างกับหมายเลขคิวล่าสุด เช่น แสดงในหน้า
        this.lastTicket = data.ticket.que_number;
        console.log('หมายเลขคิวล่าสุด:', data);
      } catch (error) {
        console.error('เกิดข้อผิดพลาดในการดึงหมายเลขคิวล่าสุด:', error);
      }
    },
    async handleCreateTicket() {
      const response = await axios.post('/api/test/test5/create-ticket')
      const data = response.data;
      console.log('API Response:', data); // ตรวจสอบข้อมูลที่ได้รับจาก API
      alert('คุณได้รับบัตรคิวแล้ว!');
      // ส่ง queNumber ไปยังหน้า test5-2 ด้วย query
      router.push({ path: '/test5-2', query: { queNumber: data.ticket.que_number, createTime: data.ticket.created_at } });
    },
    async handleClear() {
      try {
        await axios.post(`/api/test/test5/clear-ticket/${this.lastTicket}`);
        alert('คิวถูกล้างแล้ว!');
        this.lastTicket = '00'; // รีเซ็ตหมายเลขคิวล่าสุด
      } catch (error) {
        console.error('เกิดข้อผิดพลาดในการล้างคิว:', error);
      }
    }
  }
}
</script>
