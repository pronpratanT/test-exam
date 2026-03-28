<template>
  <div class="max-w-7xl mx-auto p-4 md:p-8">
    <div class="flex justify-between items-center mb-8">
      <h1 class="text-3xl font-extrabold text-slate-800 tracking-tight">Data Management</h1>
      <button class="px-6 py-2.5 bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-500 hover:to-indigo-500 shadow-lg shadow-blue-500/20 text-white rounded-xl font-bold text-sm transition-all duration-300 transform hover:-translate-y-0.5 active:scale-95 flex items-center gap-2" @click="AddModal = true">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path></svg>
        ADD NEW
      </button>
    </div>

  <!-- Modal -->
  <div v-if="AddModal" class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm flex items-center justify-center z-50 p-4">
    <div class="bg-white rounded-[2rem] shadow-2xl p-8 w-full max-w-2xl relative border border-slate-100 animate-in fade-in zoom-in duration-300">
      <div class="mb-8">
        <h2 class="text-2xl font-bold text-slate-800">Add New Entry</h2>
        <p class="text-slate-500 text-sm mt-1">Please fill in the details below to add a new record.</p>
      </div>

      <div class="space-y-5">
        <div class="grid grid-cols-12 gap-4 items-center">
          <label for="firstname" class="col-span-3 text-sm font-semibold text-slate-700">ชื่อ - สกุล</label>
          <div class="col-span-9 flex gap-2">
            <input type="text" id="firstname" v-model="fname" placeholder="ชื่อ" class="flex-1 border border-slate-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 rounded-xl px-4 py-2.5 outline-none transition-all bg-slate-50">
            <input type="text" id="lastname" v-model="lname" placeholder="นามสกุล" class="flex-1 border border-slate-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 rounded-xl px-4 py-2.5 outline-none transition-all bg-slate-50">
          </div>
        </div>

        <div class="grid grid-cols-12 gap-4 items-center">
          <label for="birthday" class="col-span-3 text-sm font-semibold text-slate-700">วันเกิด</label>
          <input type="date" id="birthday" v-model="birthdate" @change="calculateAge" class="col-span-9 border border-slate-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 rounded-xl px-4 py-2.5 outline-none transition-all bg-slate-50 w-full">
        </div>

        <div class="grid grid-cols-12 gap-4 items-center">
          <label for="age" class="col-span-3 text-sm font-semibold text-slate-700">อายุ</label>
          <div class="col-span-9 relative">
            <input type="number" id="age" v-model="age" class="border border-slate-200 rounded-xl px-4 py-2.5 w-full bg-slate-100 text-slate-500 font-medium" readonly>
            <span class="absolute right-4 top-2.5 text-slate-400 text-sm">ปี</span>
          </div>
        </div>

        <div class="grid grid-cols-12 gap-4 items-start">
          <label for="address" class="col-span-3 text-sm font-semibold text-slate-700 mt-2">ที่อยู่</label>
          <textarea id="address" v-model="address" rows="3" class="col-span-9 border border-slate-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 rounded-xl px-4 py-2.5 outline-none transition-all bg-slate-50 w-full resize-none"></textarea>
        </div>
      </div>

      <button class="absolute top-6 right-6 text-slate-400 hover:text-slate-600 transition-colors" @click="closeAddModal">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
      </button>

      <div class="flex justify-end space-x-3 mt-10">
        <button class="px-6 py-2.5 border-2 border-slate-200 hover:border-slate-300 text-slate-600 rounded-xl font-bold transition-all" @click="closeAddModal">ยกเลิก</button>
        <button class="px-8 py-2.5 bg-gradient-to-r from-emerald-500 to-green-600 hover:from-emerald-600 hover:to-green-700 text-white rounded-xl font-bold shadow-lg shadow-green-500/20 transition-all hover:-translate-y-0.5" @click="handleAdd()">บันทึกข้อมูล</button>
      </div>

      <div v-if="error" class="text-rose-600 mt-6 p-3 bg-rose-50 rounded-xl text-center border border-rose-100 text-sm font-medium">{{ error }}</div>
    </div>
  </div>

  <div v-if="DetailModal" class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm flex items-center justify-center z-50 p-4">
    <div class="bg-white rounded-[2rem] shadow-2xl p-8 w-full max-w-2xl relative border border-slate-100 animate-in fade-in zoom-in duration-300">
      <div class="mb-8 flex items-center gap-4">
        <div class="w-12 h-12 bg-indigo-100 rounded-2xl flex items-center justify-center text-indigo-600">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path></svg>
        </div>
        <div>
          <h2 class="text-2xl font-bold text-slate-800">Person Details</h2>
          <p class="text-slate-500 text-sm mt-1">Information about the selected record.</p>
        </div>
      </div>

      <div class="space-y-6">
        <div class="grid grid-cols-12 gap-4">
          <div class="col-span-6">
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">First Name</p>
            <p class="text-lg font-semibold text-slate-700 border-b border-slate-100 pb-2">{{ detail.fname }}</p>
          </div>
          <div class="col-span-6">
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">Last Name</p>
            <p class="text-lg font-semibold text-slate-700 border-b border-slate-100 pb-2">{{ detail.lname }}</p>
          </div>
        </div>

        <div class="grid grid-cols-12 gap-4">
          <div class="col-span-6">
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">Birthday</p>
            <p class="text-lg font-semibold text-slate-700 border-b border-slate-100 pb-2">{{ formatDate(detail.birthdate) }}</p>
          </div>
          <div class="col-span-6">
            <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">Age</p>
            <p class="text-lg font-semibold text-slate-700 border-b border-slate-100 pb-2">{{ detail.age }} Years</p>
          </div>
        </div>

        <div>
          <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">Address</p>
          <p class="text-lg font-semibold text-slate-700 border-b border-slate-100 pb-2">{{ detail.address }}</p>
        </div>
      </div>

      <button class="absolute top-6 right-6 text-slate-400 hover:text-slate-600 transition-colors" @click="DetailModal = false">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
      </button>

      <div class="flex justify-end mt-10">
        <button class="px-8 py-2.5 bg-slate-800 hover:bg-slate-700 text-white rounded-xl font-bold transition-all shadow-lg shadow-slate-200" @click="DetailModal = false">ปิดหน้าต่าง</button>
      </div>
    </div>
  </div>

  <!-- Table Section --> 
  <div class="bg-white rounded-[2rem] shadow-xl shadow-slate-200/50 border border-slate-100 overflow-hidden">
    <div class="overflow-x-auto">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-slate-50/50">
            <th class="px-6 py-4 text-xs font-bold text-slate-400 uppercase tracking-widest border-b border-slate-100">ID</th>
            <th class="px-6 py-4 text-xs font-bold text-slate-400 uppercase tracking-widest border-b border-slate-100">Full Name</th>
            <th class="px-6 py-4 text-xs font-bold text-slate-400 uppercase tracking-widest border-b border-slate-100">Address</th>
            <th class="px-6 py-4 text-xs font-bold text-slate-400 uppercase tracking-widest border-b border-slate-100">Birthday</th>
            <th class="px-6 py-4 text-xs font-bold text-slate-400 uppercase tracking-widest border-b border-slate-100">Age</th>
            <th class="px-6 py-4 text-xs font-bold text-slate-400 uppercase tracking-widest border-b border-slate-100 text-center">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-50">
          <tr v-for="product in products" :key="product.id" class="group hover:bg-indigo-50/30 transition-colors duration-200">
            <td class="px-6 py-4 text-sm font-semibold text-slate-500">#{{ product.id }}</td>
            <td class="px-6 py-4">
              <div class="font-bold text-slate-700 group-hover:text-indigo-600 transition-colors">{{ product.fname }} {{ product.lname }}</div>
            </td>
            <td class="px-6 py-4 text-sm text-slate-600 max-w-xs truncate">{{ product.address }}</td>
            <td class="px-6 py-4 text-sm text-slate-500 font-medium">{{ formatDate(product.birthdate) }}</td>
            <td class="px-6 py-4">
              <span class="px-3 py-1 bg-slate-100 text-slate-600 rounded-full text-xs font-bold">{{ product.age }} Yrs</span>
            </td>
            <td class="px-6 py-4 text-center">
              <button class="p-2 text-indigo-600 hover:bg-indigo-100 rounded-lg transition-colors inline-flex items-center gap-1 font-bold text-sm" @click="openDetailModal(product.id)">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path></svg>
                View
              </button>
            </td>
          </tr>
          <tr v-if="products.length === 0">
            <td colspan="6" class="px-6 py-20 text-center">
              <div class="flex flex-col items-center">
                <div class="w-16 h-16 bg-slate-100 rounded-full flex items-center justify-center text-slate-400 mb-4">
                  <svg class="w-10 h-10" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path></svg>
                </div>
                <p class="text-slate-500 font-medium">No records found</p>
                <button @click="AddModal = true" class="mt-4 text-indigo-600 font-bold text-sm hover:underline">Click here to add the first record</button>
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
  name: 'Test1Page',
  data() {
    return {
      AddModal: false,
      DetailModal: false,
      fname: '',
      lname: '',
      birthdate: '',
      age: '',
      address: '',
      error: '',
      isLoading: false,
      products: [],
      detail: {} // กำหนดค่า default เริ่มต้น
    }
  },
  // Vue Lifecycle hook ที่เทียบเท่ากับ useEffect(..., []) ใน Next/React
  mounted() {
    this.fetchData();
  },
  methods: {
    formatDate(dateStr) {
      if (!dateStr) return '';
      // แปลง String ให้อยู่ในรูป Date Object ของ JS
      const date = new Date(dateStr);
      if (isNaN(date)) return dateStr; // เผื่อกรณีข้อมูลจาก DB ส่งมาผิดรูปแบบ
      
      const d = String(date.getDate()).padStart(2, '0');
      const m = String(date.getMonth() + 1).padStart(2, '0');
      const y = date.getFullYear();
      
      return `${d}/${m}/${y}`;
    },

    async fetchData() {
      try {
        const response = await axios.get('/api/test/test1/get-all');
        this.products = response.data.data;
      } catch (err) {
        console.error('Failed to load data:', err);
      }
    },

    resetForm() {
      this.fname = '';
      this.lname = '';
      this.birthdate = '';
      this.age = '';
      this.address = '';
      this.error = '';
    },

    closeAddModal() {
      this.AddModal = false;
      this.resetForm();
    },

    calculateAge() {
      if (!this.birthdate) {
        this.age = '';
        return;
      }
      const today = new Date();
      const birthDate = new Date(this.birthdate);
      let calculatedAge = today.getFullYear() - birthDate.getFullYear();
      const m = today.getMonth() - birthDate.getMonth();
      if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
          calculatedAge--;
      }
      this.age = calculatedAge;
    },

    async handleAdd() {
        this.error = '';
        this.isLoading = true;
        const payload = {
            fname: this.fname,
            lname: this.lname,
            birthdate: this.birthdate,
            age: this.age,
            address: this.address
        } 
        try {
            const response = await axios.post('/api/test/test1/create', payload);
            alert('Add Success!');
            this.closeAddModal();
            this.fetchData(); // ดึงข้อมูลใหม่มาแสดงในตารางทันที
        } catch (err) {
            if (err.response && err.response.data) {
                this.error = err.response.data.message;
            } else {
                // กรณี API ล่ม หรือเน็ตหลุด
                this.error = 'ไม่สามารถเชื่อมต่อกับเซิร์ฟเวอร์ได้ (Network Error)';
                console.error(err);
            }
        } finally {
            this.isLoading = false;
        }
    },
    async openDetailModal(id) {
      this.DetailModal = true;
      try {
        const response = await axios.get(`/api/test/test1/get/${id}`);
        const data = response.data.data;
        
        // <input type="date"> บังคับว่ารูปแบบจะต้องเป็น YYYY-MM-DD แบบเป๊ะๆ เท่านั้น!
        // ถ้า Backend ส่งมาเป็น "2000-01-01T00:00:00Z" มันจะไม่แสดงผล เราเลยต้องตัด T ออกเอาเฉพาะข้างหน้าครับ
        if (data && data.birthdate) {
            data.birthdate = data.birthdate.split('T')[0];
        }
        
        this.detail = data;
      } catch (err) {
        console.error('Failed to load data:', err);
      }
    }
  }
}
</script>
