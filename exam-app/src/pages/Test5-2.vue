<template lang="html">
  <div class="flex flex-col items-center justify-center h-screen bg-gradient-to-br from-indigo-900 via-blue-900 to-slate-900 gap-10 p-6">
      <div class="flex flex-col items-center p-12 bg-white rounded-[3rem] shadow-2xl w-full max-w-md relative overflow-hidden">
        <!-- Ticket cutouts effect -->
        <div class="absolute -left-4 top-1/2 -translate-y-1/2 w-8 h-8 bg-indigo-900 rounded-full"></div>
        <div class="absolute -right-4 top-1/2 -translate-y-1/2 w-8 h-8 bg-indigo-900 rounded-full"></div>
        <div class="absolute inset-x-0 top-1/2 border-t-2 border-dashed border-slate-100 -translate-y-1/2"></div>

        <div class="mb-12 text-center">
            <div class="w-16 h-16 bg-indigo-50 rounded-2xl flex items-center justify-center text-indigo-600 mx-auto mb-4">
                <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
            </div>
            <p class="text-slate-400 text-xs font-black tracking-[0.2em] uppercase">คิวของคุณคือ</p>
        </div>
        
        <div class="mb-12 relative">
            <p class="text-center text-9xl font-black text-transparent bg-clip-text bg-gradient-to-br from-indigo-600 to-blue-500 drop-shadow-sm">{{ queNumber }}</p>
        </div>

        <div class="w-full space-y-6">
            <div class="flex justify-between items-center px-2">
                <span class="text-slate-400 text-sm font-bold uppercase tracking-wider">วันที่เข้าคิว</span>
                <span class="text-slate-700 font-bold text-sm">{{ createdAt }}</span>
            </div>

            <button class="w-full px-10 py-4 bg-slate-900 hover:bg-slate-800 text-white rounded-2xl font-bold text-lg transition-all duration-300 transform hover:-translate-y-1 active:scale-95 shadow-xl shadow-slate-200" @click="handleClear">
                กลับไปยังหน้าหลัก
            </button>
        </div>
      </div>

      <p class="text-blue-200/50 text-sm font-medium">กรุณารอเรียกหมายเลขคิวของท่านที่จุดบริการ</p>
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