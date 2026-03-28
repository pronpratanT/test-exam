<template>
  <div class="grid items-center justify-center h-screen">
      <button class="px-10 py-7 bg-blue-600 hover:bg-blue-700 text-white rounded font-semibold text-lg transition" 
      @click="handleCreateTicket">รับบัตรคิว</button>
      <p class="text-center text-8xl font-bold">{{ lastTicket }}</p>
      <button
        class="px-10 py-2 bg-gray-600 hover:bg-gray-700 text-white rounded font-semibold text-lg transition"
        :class="{ 'cursor-not-allowed opacity-50': lastTicket === '00' }"
        @click="handleClear"
        :disabled="lastTicket === '00'"
      >ล้างคิว</button>
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
