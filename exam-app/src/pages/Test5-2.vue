<template lang="html">
    <div class="grid items-center justify-center h-screen">
      <p class="text-center text-2xl font-bold text-white">หมายเลขคิว</p>
    <p class="text-center text-8xl font-bold text-white">{{ queNumber }}</p>
      <p class="text-center text-sm text-white">วันที่ {{ createdAt }}</p>
      <button class="px-10 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded font-semibold text-lg transition" @click="handleClear">กลับไปยังหน้ารับบัตรคิว</button>
  </div>
</template>

<script>
import router from '../router';
import { useRoute } from 'vue-router';

export default {
    data() {
        return {
            queNumber: '',
            createdAt: ''
        }
    },
    created() {
        const route = this.$route;
        this.queNumber = route.query.queNumber || '00';

        const rawTime = route.query.createTime;
        if (rawTime) {
            const date = new Date(rawTime);
            const day = String(date.getDate()).padStart(2, '0');
            const month = String(date.getMonth() + 1).padStart(2, '0');
            const year = date.getFullYear();
            const hours = String(date.getHours()).padStart(2, '0');
            const minutes = String(date.getMinutes()).padStart(2, '0');
            this.createdAt = `${day}/${month}/${year} เวลา ${hours}:${minutes} น.`;
        }
    },
    methods: {
        handleClear() {
            router.push('/test5'); // กลับไปยังหน้ารับบัตรคิว
        }
    }
}
</script>
<style lang="">
    
</style>