<template>
  <div class="grid items-center justify-center h-screen">
      <button class="px-10 py-7 bg-blue-600 hover:bg-blue-700 text-white rounded font-semibold text-lg transition" @click="handleClick">รับบัตรคิว</button>
      <p class="text-center text-8xl font-bold">A5</p>
      <button class="px-10 py-2 bg-gray-600 hover:bg-gray-700 text-white rounded font-semibold text-lg transition" @click="handleClear">ล้างคิว</button>
  </div>
</template>

<script>
import router from '../router';

export default {
  name: 'Test5Page',
  methods: {
    handleClick() {
      alert('คุณได้รับบัตรคิวแล้ว!');
      router.push('/test5-2'); // เปลี่ยนเส้นทางไปยังหน้า Test5-2 หลังจากคลิกปุ่มรับบัตรคิว
    },
    handleClear() {
      alert('คิวถูกล้างแล้ว!');
    }
  }
}
</script>
