<template>
  <div class="flex gap-4">
    <button class="mb-4 px-4 py-2 bg-green-600 text-white rounded hover:bg-green-700 transition" @click="Approve = true">อนุมัติ</button>
    <button class="mb-4 px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 transition" @click="Reject = true">ไม่อนุมัติ</button>
  </div>

  <!-- Modal -->
  <div v-if="Approve" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
    <div class="rounded-lg bg-blue-600 w-full max-w-3xl relative">
      <button class="absolute top-0 right-2 text-white hover:text-gray-700 text-2xl" @click="Approve = false">&times;</button>
      <h2 class="text-md text-white text-left pl-6 pt-2">ยืนยันการอนุมัติ</h2>
      <div class="bg-white rounded-b-lg shadow-lg p-8 w-full max-w-3xl relative">
        <div class="grid grid-cols-12 gap-4 items-center mb-6">
          <label for="reason" class="col-span-2 text-left">เหตุผล</label>
          <input type="text" id="reason" class="col-span-10 border border-gray-300 rounded px-3 py-2 w-full">
        </div>
        <div class="flex justify-end space-x-4">
          <button class="px-4 py-2 bg-green-600 text-white rounded" @click="Approve = false">อนุมัติ</button>
          <button class="px-4 py-2 bg-red-600 text-white rounded" @click="Approve = false">ยกเลิก</button>
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
          <label for="reason" class="col-span-2 text-left">เหตุผล</label>
          <input type="text" id="reason" class="col-span-10 border border-gray-300 rounded px-3 py-2 w-full">
        </div>
        <div class="flex justify-end space-x-4">
          <button class="px-4 py-2 bg-green-600 text-white rounded" @click="Reject = false">ไม่อนุมัติ</button>
          <button class="px-4 py-2 bg-red-600 text-white rounded" @click="Reject = false">ยกเลิก</button>
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
            <input type="checkbox" />
          </th>
          <th class="px-4 py-2 border-b">รายการ</th>
          <th class="px-4 py-2 border-b">เหตุผล</th>
          <th class="px-4 py-2 border-b">สถานะเอกสาร</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="product in products" :key="product.id" class="hover:bg-gray-50">
          <td class="px-4 py-2 border-b text-center">
            <input type="checkbox" :value="product.id" />
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
      products: []
    }
  },
  mounted() {
    this.getDataAllTest3()
  },
  methods: {
    async getDataAllTest3() {
      try {
        const response = await axios.get('/api/test/test3/get-all');
        this.products = response.data.data;
        // console.log(this.products)
      } catch (error) {
        console.error('Error fetching data:', error)
      }
    }
  }
}
</script>
