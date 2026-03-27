<template>
  <div class="flex justify-end">
    <button class="mb-4 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition" @click="AddModal = true">ADD</button>
  </div>

  <!-- Modal -->
  <div v-if="AddModal" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg shadow-lg pt-1 pr-6 pl-6 pb-3 w-full max-w-3xl relative">
      <h2 class="text-xl font-bold mb-4">Modal Title</h2>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="firstname" class="col-span-2 text-left">ชื่อ - สกุล</label>
        <input type="text" id="firstname" v-model="fname" placeholder="ชื่อ" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full">
        <input type="text" id="lastname" v-model="lname" placeholder="นามสกุล" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full border-l-0">
      </div>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="birthday" class="col-span-2 text-left">วันเกิด</label>
        <input type="date" id="birthday" v-model="birthdate" @change="calculateAge" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full">
      </div>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="age" class="col-span-2 text-left">อายุ</label>
        <input type="number" id="age" v-model="age" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full bg-gray-100" readonly>
      </div>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="address" class="col-span-2 text-left">ที่อยู่</label>
        <input type="text" id="address" v-model="address" class="col-span-10 border border-gray-300 rounded px-3 py-2 w-full">
      </div>
      <button class="absolute top-2 right-2 text-gray-400 hover:text-gray-700 text-2xl" @click="closeAddModal">&times;</button>
      <div class="flex justify-end space-x-4">
        <button class="px-4 py-2 bg-green-600 text-white rounded" @click="handleAdd()">บันทึก</button>
        <button class="px-4 py-2 bg-red-600 text-white rounded" @click="closeAddModal">ยกเลิก</button>
      </div>
      <div v-if="error" class="text-red-600 mt-4 text-center">{{ error }}</div>
    </div>
  </div>

  <div v-if="DetailModal" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg shadow-lg pt-1 pr-6 pl-6 pb-3 w-full max-w-3xl relative">
      <h2 class="text-xl font-bold mb-4">Modal Title</h2>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="firstname" class="col-span-2 text-left">ชื่อ - สกุล</label>
        <input type="text" id="firstname" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full" v-model="detail.fname" readonly>
        <input type="text" id="lastname" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full" v-model="detail.lname" readonly>
      </div>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="birthday" class="col-span-2 text-left">วันเกิด</label>
        <input type="date" id="birthday" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full" v-model="detail.birthdate" readonly>
      </div>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="age" class="col-span-2 text-left">อายุ</label>
        <input type="number" id="age" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full" v-model="detail.age" readonly>
      </div>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="address" class="col-span-2 text-left">ที่อยู่</label>
        <input type="text" id="address" class="col-span-10 border border-gray-300 rounded px-3 py-2 w-full" v-model="detail.address" readonly>
      </div>
      <button class="absolute top-2 right-2 text-gray-400 hover:text-gray-700 text-2xl" @click="DetailModal = false">&times;</button>
      <div class="flex justify-end space-x-4">
        <button class="px-4 py-2 bg-red-600 text-white rounded" @click="DetailModal = false">ปิด</button>
      </div>
    </div>
  </div>

  <!-- Table --> 
  <div class="overflow-x-auto">
    <table class="min-w-full bg-white border border-gray-200 rounded-lg shadow">
      <thead class="bg-gray-100">
        <tr>
          <th class="px-4 py-2 border-b">ID</th>
          <th class="px-4 py-2 border-b">ชื่อ-สกุล</th>
          <th class="px-4 py-2 border-b">ที่อยู่</th>
          <th class="px-4 py-2 border-b">วันเกิด</th>
          <th class="px-4 py-2 border-b">อายุ</th>
          <th class="px-4 py-2 border-b">Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="product in products" :key="product.id" class="hover:bg-gray-50">
          <td class="px-4 py-2 border-b">{{ product.id }}</td>
          <td class="px-4 py-2 border-b">{{ product.fname }} {{ product.lname }}</td>
          <td class="px-4 py-2 border-b">{{ product.address }}</td>
          <td class="px-4 py-2 border-b">{{ formatDate(product.birthdate) }}</td>
          <td class="px-4 py-2 border-b">{{ product.age }}</td>
          <td class="px-4 py-2 border-b text-center">
            <button class="px-2 py-1 bg-green-500 text-white rounded hover:bg-green-600 mr-2" @click="openDetailModal(product.id)">View</button>
          </td>
        </tr>
      </tbody>
    </table>
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
                this.error = err.response.data.message || 'Username หรือ Password ไม่ถูกต้อง';
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
