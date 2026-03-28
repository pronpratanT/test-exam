<template>
  <div class="max-w-7xl mx-auto p-4 md:p-8">
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center mb-8 gap-4">
      <div>
        <h1 class="text-3xl font-extrabold text-slate-800 tracking-tight">Approval System</h1>
        <p class="text-slate-500 text-sm mt-1">Review and manage pending requests.</p>
      </div>
      <div class="flex gap-3">
        <button :disabled="selectedIds.length === 0" 
          class="px-6 py-2.5 bg-gradient-to-r from-emerald-500 to-green-600 hover:from-emerald-600 hover:to-green-700 shadow-lg shadow-green-500/20 text-white rounded-xl font-bold text-sm transition-all duration-300 transform hover:-translate-y-0.5 active:scale-95 disabled:opacity-40 disabled:hover:translate-y-0 disabled:cursor-not-allowed flex items-center gap-2" 
          @click="Approve = true">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
          อนุมัติ ({{ selectedIds.length }})
        </button>
        <button :disabled="selectedIds.length === 0" 
          class="px-6 py-2.5 bg-white border-2 border-rose-200 hover:border-rose-400 text-rose-500 hover:bg-rose-50 rounded-xl font-bold text-sm transition-all duration-300 transform hover:-translate-y-0.5 active:scale-95 disabled:opacity-40 disabled:hover:translate-y-0 disabled:cursor-not-allowed flex items-center gap-2" 
          @click="Reject = true">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
          ไม่อนุมัติ
        </button>
      </div>
    </div>

  <!-- Modal Approve -->
  <div v-if="Approve" class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm flex items-center justify-center z-50 p-4">
    <div class="bg-white rounded-[2rem] shadow-2xl w-full max-w-xl relative border border-slate-100 overflow-hidden animate-in fade-in zoom-in duration-300">
      <div class="bg-gradient-to-r from-emerald-500 to-green-600 px-8 py-6 text-white">
          <h2 class="text-2xl font-bold">ยืนยันการอนุมัติ</h2>
          <p class="text-emerald-50/80 text-sm mt-1">คุณกำลังจะอนุมัติรายการที่เลือกทั้งหมด {{ selectedIds.length }} รายการ</p>
      </div>
      <div class="p-8">
        <div class="space-y-2 mb-8">
          <label for="reason_approve" class="text-sm font-bold text-slate-700 ml-1">ระบุเหตุผล (ไม่บังคับ)</label>
          <textarea id="reason_approve" v-model="actionReason" rows="3" placeholder="ระบุเหตุผลในการอนุมัติ..." class="border border-slate-200 focus:border-emerald-500 focus:ring-4 focus:ring-emerald-500/10 rounded-2xl px-5 py-4 outline-none transition-all bg-slate-50 w-full resize-none font-medium"></textarea>
        </div>
        <div class="flex justify-end gap-3">
          <button class="px-6 py-2.5 border-2 border-slate-200 hover:border-slate-300 text-slate-600 rounded-xl font-bold transition-all" @click="closeModal">ยกเลิก</button>
          <button class="px-8 py-2.5 bg-emerald-500 hover:bg-emerald-600 text-white rounded-xl font-bold shadow-lg shadow-emerald-500/20 transition-all hover:-translate-y-0.5" @click="ApproveDataTest3">ยืนยันการอนุมัติ</button>
        </div>
      </div>
      <button class="absolute top-6 right-6 text-white/60 hover:text-white transition-colors" @click="Approve = false">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
      </button>
    </div>
  </div>

  <!-- Modal Reject -->
  <div v-if="Reject" class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm flex items-center justify-center z-50 p-4">
    <div class="bg-white rounded-[2rem] shadow-2xl w-full max-w-xl relative border border-slate-100 overflow-hidden animate-in fade-in zoom-in duration-300">
      <div class="bg-gradient-to-r from-rose-500 to-red-600 px-8 py-6 text-white">
          <h2 class="text-2xl font-bold">ยืนยันการไม่อนุมัติ</h2>
          <p class="text-rose-50/80 text-sm mt-1">คุณกำลังจะปฏิเสธรายการที่เลือกทั้งหมด {{ selectedIds.length }} รายการ</p>
      </div>
      <div class="p-8">
        <div class="space-y-2 mb-8">
          <label for="reason_reject" class="text-sm font-bold text-slate-700 ml-1">ระบุเหตุผล</label>
          <textarea id="reason_reject" v-model="actionReason" rows="3" placeholder="โปรดระบุเหตุผลในการไม่อนุมัติ..." class="border border-slate-200 focus:border-rose-500 focus:ring-4 focus:ring-rose-500/10 rounded-2xl px-5 py-4 outline-none transition-all bg-slate-50 w-full resize-none font-medium"></textarea>
        </div>
        <div class="flex justify-end gap-3">
          <button class="px-6 py-2.5 border-2 border-slate-200 hover:border-slate-300 text-slate-600 rounded-xl font-bold transition-all" @click="closeModal">ยกเลิก</button>
          <button class="px-8 py-2.5 bg-rose-500 hover:bg-rose-600 text-white rounded-xl font-bold shadow-lg shadow-rose-500/20 transition-all hover:-translate-y-0.5" @click="RejectDataTest3">ยืนยันไม่อนุมัติ</button>
        </div>
      </div>
      <button class="absolute top-6 right-6 text-white/60 hover:text-white transition-colors" @click="Reject = false">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
      </button>
    </div>
  </div>

  <!-- Table Section --> 
  <div class="bg-white rounded-[2rem] shadow-xl shadow-slate-200/50 border border-slate-100 overflow-hidden">
    <div class="overflow-x-auto">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-slate-50/50">
            <th class="px-6 py-5 text-center w-20 border-b border-slate-100">
              <div class="flex items-center justify-center">
                <input type="checkbox" class="w-5 h-5 rounded-lg border-slate-300 text-indigo-600 focus:ring-indigo-500 transition-all cursor-pointer" @change="selectAll" :checked="selectedIds.length > 0 && selectedIds.length === products.filter(p => p.status === 'pending').length" />
              </div>
            </th>
            <th class="px-6 py-5 text-xs font-bold text-slate-400 uppercase tracking-widest border-b border-slate-100">รายการ / Description</th>
            <th class="px-6 py-5 text-xs font-bold text-slate-400 uppercase tracking-widest border-b border-slate-100">เหตุผล / Reason</th>
            <th class="px-6 py-5 text-xs font-bold text-slate-400 uppercase tracking-widest border-b border-slate-100">สถานะ / Status</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-50 font-medium">
          <tr v-for="product in products" :key="product.id" class="group hover:bg-slate-50/50 transition-colors duration-200">
            <td class="px-6 py-5 text-center">
              <div class="flex items-center justify-center">
                <input v-if="product.status === 'pending'" type="checkbox" :value="product.id" v-model="selectedIds" class="w-5 h-5 rounded-lg border-slate-300 text-indigo-600 focus:ring-indigo-500 transition-all cursor-pointer" />
                <div v-else-if="product.status === 'approved'" class="w-8 h-8 rounded-full bg-emerald-50 flex items-center justify-center text-emerald-500 shadow-sm border border-emerald-100">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7"></path></svg>
                </div>
                <div v-else-if="product.status === 'rejected'" class="w-8 h-8 rounded-full bg-rose-50 flex items-center justify-center text-rose-500 shadow-sm border border-rose-100">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M6 18L18 6M6 6l12 12"></path></svg>
                </div>
              </div>
            </td>
            <td class="px-6 py-5">
              <div class="text-slate-700 font-bold group-hover:text-indigo-600 transition-colors">{{ product.name }}</div>
              <div class="text-xs text-slate-400 mt-1 uppercase font-bold tracking-tight">ID: #{{ product.id }}</div>
            </td>
            <td class="px-6 py-5">
              <span v-if="product.reason" class="text-slate-600 text-sm italic">"{{ product.reason }}"</span>
              <span v-else class="text-slate-300 text-sm">—</span>
            </td>
            <td class="px-6 py-5">
              <span v-if="product.status === 'pending'" class="inline-flex items-center px-4 py-1.5 rounded-full text-xs font-bold bg-amber-50 text-amber-600 border border-amber-100 ring-4 ring-amber-500/5">
                <span class="w-1.5 h-1.5 rounded-full bg-amber-500 mr-2 animate-pulse"></span>
                {{ product.status }}
              </span>
              <span v-else-if="product.status === 'approved'" class="inline-flex items-center px-4 py-1.5 rounded-full text-xs font-bold bg-emerald-50 text-emerald-600 border border-emerald-100 ring-4 ring-emerald-500/5">
                <span class="w-1.5 h-1.5 rounded-full bg-emerald-500 mr-2"></span>
                {{ product.status }}
              </span>
              <span v-else-if="product.status === 'rejected'" class="inline-flex items-center px-4 py-1.5 rounded-full text-xs font-bold bg-rose-50 text-rose-600 border border-rose-100 ring-4 ring-rose-500/5">
                <span class="w-1.5 h-1.5 rounded-full bg-rose-500 mr-2"></span>
                {{ product.status }}
              </span>
            </td>
          </tr>
          <tr v-if="products.length === 0">
            <td colspan="4" class="px-6 py-20 text-center">
              <div class="flex flex-col items-center">
                <div class="w-16 h-16 bg-slate-100 rounded-full flex items-center justify-center text-slate-400 mb-4">
                  <svg class="w-10 h-10" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"></path></svg>
                </div>
                <p class="text-slate-500 font-medium">No items found in system</p>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Test3Page',
  data() {
    return {
      Approve: false,
      Reject: false,
      products: [],
      selectedIds: [], // เก็บรายการ ID ที่ถูกติ๊กเลือก
      actionReason: '' // จัดเก็บเหตุผลของ modal ควบคู่กันไป
    }
  },
  mounted() {
    this.getDataAllTest3()
  },
  methods: {
    closeModal() {
      this.Approve = false;
      this.Reject = false;
      this.actionReason = '';
    },
    selectAll(event) {
      if (event.target.checked) {
        this.selectedIds = this.products.filter(p => p.status === 'pending').map(p => p.id);
      } else {
        this.selectedIds = [];
      }
    },
    async getDataAllTest3() {
      try {
        const response = await axios.get('/api/test/test3/get-all');
        this.products = response.data.data;
      } catch (error) {
        console.error('Error fetching data:', error)
      }
    },

    async ApproveDataTest3() {
      if (this.selectedIds.length === 0) return;

      // ปั้น Payload โดยอิงจาก item ที่กำลังถูกเลือก (selectedIds) 
      const payload = this.selectedIds.map(id => ({
        id: id,
        reason: this.actionReason
      }));
      console.log(payload)
      try {
        await axios.post('/api/test/test3/approve', payload);
        alert('Approve Success!');
        this.closeModal();
        this.selectedIds = [];
        this.getDataAllTest3();
      } catch (error) {
        console.error('Error approving data:', error)
      }
    },

    async RejectDataTest3() {
      if (this.selectedIds.length === 0) return;

      // ปั้น Payload โดยอิงจาก item ที่กำลังถูกเลือก (selectedIds) 
      const payload = this.selectedIds.map(id => ({
        id: id,
        reason: this.actionReason
      }));

      try {
        await axios.post('/api/test/test3/reject', payload);
        alert('Reject Success!');
        this.closeModal();
        this.selectedIds = [];
        this.getDataAllTest3();
      } catch (error) {
        console.error('Error rejecting data:', error)
      }
    }
  }
}
</script>
