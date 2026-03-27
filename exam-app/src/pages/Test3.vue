<template>
  <div class="flex gap-4">
    <button :disabled="selectedIds.length === 0" class="mb-4 px-4 py-2 bg-green-600 text-white rounded hover:bg-green-700 transition disabled:opacity-50 disabled:cursor-not-allowed" @click="Approve = true">อนุมัติ</button>
    <button :disabled="selectedIds.length === 0" class="mb-4 px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 transition disabled:opacity-50 disabled:cursor-not-allowed" @click="Reject = true">ไม่อนุมัติ</button>
  </div>

  <!-- Modal -->
  <div v-if="Approve" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
    <div class="rounded-lg bg-blue-600 w-full max-w-3xl relative">
      <button class="absolute top-0 right-2 text-white hover:text-gray-700 text-2xl" @click="Approve = false">&times;</button>
      <h2 class="text-md text-white text-left pl-6 pt-2">ยืนยันการอนุมัติ</h2>
      <div class="bg-white rounded-b-lg shadow-lg p-8 w-full max-w-3xl relative">
        <div class="grid grid-cols-12 gap-4 items-center mb-6">
          <label for="reason_approve" class="col-span-2 text-left">เหตุผล</label>
          <input type="text" id="reason_approve" v-model="actionReason" class="col-span-10 border border-gray-300 rounded px-3 py-2 w-full">
        </div>
        <div class="flex justify-end space-x-4">
          <button class="px-4 py-2 bg-green-600 text-white rounded hover:bg-green-700" @click="ApproveDataTest3">อนุมัติ</button>
          <button class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700" @click="closeModal">ยกเลิก</button>
        </div>
      </div>
    </div>
  </div>

  <div v-if="Reject" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
    <div class="rounded-lg bg-blue-600 w-full max-w-3xl relative">
      <button class="absolute top-0 right-2 text-white hover:text-gray-700 text-2xl" @click="Reject = false">&times;</button>
      <h2 class="text-md text-white text-left pl-6 pt-2">ยืนยันการไม่อนุมัติ</h2>
      <div class="bg-white rounded-b-lg shadow-lg p-8 w-full max-w-3xl relative">
        <div class="grid grid-cols-12 gap-4 items-center mb-6">
          <label for="reason_reject" class="col-span-2 text-left">เหตุผล</label>
          <input type="text" id="reason_reject" v-model="actionReason" class="col-span-10 border border-gray-300 rounded px-3 py-2 w-full">
        </div>
        <div class="flex justify-end space-x-4">
          <button class="px-4 py-2 bg-green-600 text-white rounded hover:bg-green-700" @click="RejectDataTest3">ไม่อนุมัติ</button>
          <button class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700" @click="closeModal">ยกเลิก</button>
        </div>
      </div>
    </div>
  </div>

  <!-- Table --> 
  <div class="overflow-x-auto">
    <table class="min-w-full bg-white border border-gray-200 rounded-lg shadow">
      <thead class="bg-gray-100">
        <tr>
          <th class="px-4 py-2 border-b w-12 text-center">
            <input type="checkbox" @change="selectAll" :checked="selectedIds.length > 0 && selectedIds.length === products.filter(p => p.status === 'pending').length" />
          </th>
          <th class="px-4 py-2 border-b">รายการ</th>
          <th class="px-4 py-2 border-b">เหตุผล</th>
          <th class="px-4 py-2 border-b">สถานะเอกสาร</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="product in products" :key="product.id" class="hover:bg-gray-50">
          <td class="px-4 py-2 border-b text-center">
            <input v-if="product.status === 'pending'" type="checkbox" :value="product.id" v-model="selectedIds" />
            <span v-else-if="product.status === 'approved'" class="text-green-500 text-sm" title="อนุมัติแล้ว">✔️</span>
            <span v-else-if="product.status === 'rejected'" class="text-red-500 text-sm" title="ไม่อนุมัติ">❌</span>
          </td>
          <td class="px-4 py-2 border-b">{{ product.name }}</td>
          <td class="px-4 py-2 border-b">{{ product.reason }}</td>
          <td class="px-4 py-2 border-b">{{ product.status }}</td>
        </tr>
      </tbody>
    </table>
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
