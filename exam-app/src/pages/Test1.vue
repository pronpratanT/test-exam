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
        <input type="text" id="firstname" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full">
        <input type="text" id="lastname" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full">
      </div>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="birthday" class="col-span-2 text-left">วันเกิด</label>
        <input type="date" id="birthday" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full">
      </div>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="age" class="col-span-2 text-left">อายุ</label>
        <input type="number" id="age" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full" :disabled="true">
      </div>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="address" class="col-span-2 text-left">ที่อยู่</label>
        <input type="text" id="address" class="col-span-10 border border-gray-300 rounded px-3 py-2 w-full">
      </div>
      <button class="absolute top-2 right-2 text-gray-400 hover:text-gray-700 text-2xl" @click="AddModal = false">&times;</button>
      <div class="flex justify-end space-x-4">
        <button class="px-4 py-2 bg-green-600 text-white rounded" @click="AddModal = false">บันทึก</button>
        <button class="px-4 py-2 bg-red-600 text-white rounded" @click="AddModal = false">ยกเลิก</button>
      </div>
    </div>
  </div>

  <div v-if="DetailModal" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg shadow-lg pt-1 pr-6 pl-6 pb-3 w-full max-w-3xl relative">
      <h2 class="text-xl font-bold mb-4">Modal Title</h2>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="firstname" class="col-span-2 text-left">ชื่อ - สกุล</label>
        <input type="text" id="firstname" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full">
        <input type="text" id="lastname" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full">
      </div>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="birthday" class="col-span-2 text-left">วันเกิด</label>
        <input type="date" id="birthday" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full">
      </div>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="age" class="col-span-2 text-left">อายุ</label>
        <input type="number" id="age" class="col-span-5 border border-gray-300 rounded px-3 py-2 w-full" :disabled="true">
      </div>
      <div class="grid grid-cols-12 gap-4 items-center mb-6">
        <label for="address" class="col-span-2 text-left">ที่อยู่</label>
        <input type="text" id="address" class="col-span-10 border border-gray-300 rounded px-3 py-2 w-full">
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
        <tr v-for="product in products" :key="product.code" class="hover:bg-gray-50">
          <td class="px-4 py-2 border-b">{{ product.code }}</td>
          <td class="px-4 py-2 border-b">{{ product.name }}</td>
          <td class="px-4 py-2 border-b">{{ product.category }}</td>
          <td class="px-4 py-2 border-b">01/01/2000</td>
          <td class="px-4 py-2 border-b">20</td>
          <td class="px-4 py-2 border-b text-center">
            <button class="px-2 py-1 bg-green-500 text-white rounded hover:bg-green-600 mr-2" @click="openDetailModal(product)">View</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'Test1Page',
  data() {
    return {
      AddModal: false,
      DetailModal: false,
      products: [
        { code: 'P001', name: 'Product 1', category: 'Category 1', quantity: 10 },
        { code: 'P002', name: 'Product 2', category: 'Category 2', quantity: 20 },
        { code: 'P003', name: 'Product 3', category: 'Category 3', quantity: 30 },
        { code: 'P004', name: 'Product 4', category: 'Category 4', quantity: 40 }
      ]
    }
  },
  methods: {
    openDetailModal(product) {
      this.DetailModal = true;
      // สามารถเพิ่ม logic เพื่อโหลดข้อมูล detail ของ product ได้ที่นี่
    }
  }
}
</script>
