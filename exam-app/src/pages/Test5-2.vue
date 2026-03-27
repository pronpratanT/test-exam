<template lang="html">
    <div class="grid items-center justify-center h-screen">
      <p class="text-center text-2xl font-bold text-white">หมายเลขคิว</p>
      <p class="text-center text-8xl font-bold text-white">A5</p>
      <p class="text-center text-sm text-white">วันที่ 24-02-2025</p>
      <button class="px-10 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded font-semibold text-lg transition" @click="handleClear">กลับไปยังหน้ารับบัตรคิว</button>
  </div>
</template>

<script>
import router from '../router';

export default {
    methods: {
        handleClear() {
            router.push('/test5'); // กลับไปยังหน้ารับบัตรคิว
        }
    }
}
</script>
<style lang="">
    
</style>