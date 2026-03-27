<template>
    <!-- Toast Notification -->
    <div class="fixed top-4 right-4 z-50 space-y-2">
      <div
        v-for="toast in toasts"
        :key="toast.id"
        :class="[
          'px-4 py-2 rounded shadow text-white',
          toast.type === 'success' ? 'bg-green-500' : 'bg-red-500'
        ]"
      >
        {{ toast.message }}
      </div>
    </div>
  <div>
    <form @submit.prevent="handleSave" class="space-y-4">
      <div class="flex flex-wrap gap-6 mb-4 justify-center">
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-1 font-medium text-left" for="firstName">First Name</label>
          <input type="text" id="firstName" v-model="firstName" class="border border-gray-300 rounded px-3 py-2 w-full" />
        </div>
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-1 font-medium text-left" for="lastName">Last Name</label>
          <input type="text" id="lastName" v-model="lastName" class="border border-gray-300 rounded px-3 py-2 w-full" />
        </div>
      </div>

      <div class="flex flex-wrap gap-6 mb-4 justify-center">
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-1 font-medium text-left" for="email">Email</label>
          <input type="text" id="email" class="border border-gray-300 rounded px-3 py-2 w-full" />
        </div>
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-1 font-medium text-left" for="phone">Phone</label>
          <input type="text" id="phone" class="border border-gray-300 rounded px-3 py-2 w-full" />
        </div>
      </div>

      <div class="flex flex-wrap gap-6 mb-4 justify-center">
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-1 font-medium text-left" for="profile">Profile</label>
          <input type="file" id="profile" class="border border-gray-300 rounded px-3 py-2 w-full" />
        </div>
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-1 font-medium text-left" for="birthDay">Birth Day</label>
          <input type="date" id="birthDay" class="border border-gray-300 rounded px-3 py-2 w-full" />
        </div>
      </div>

      <div class="flex flex-wrap gap-6 mb-4 justify-center">
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-1 font-medium text-left" for="occupation">Occupation</label>
          <select id="occupation" class="border border-gray-300 rounded px-3 py-2 w-full">
            <option value="">Select Occupation</option>
            <option value="developer">Developer</option>
            <option value="designer">Designer</option>
            <option value="manager">Manager</option>
          </select>
        </div>
        <div class="flex-1 min-w-[200px] max-w-md"></div>
      </div>

      <div class="flex flex-wrap gap-6 mb-4 justify-center">
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-1 font-medium text-left" for="sex">Sex</label>
          <div class="flex items-center gap-6 pl-2">
            <label class="inline-flex items-center">
              <input type="radio" name="sex" value="male" class="border border-gray-300" />
              <span class="ml-2">Male</span>
            </label>
            <label class="inline-flex items-center">
              <input type="radio" name="sex" value="female" class="border border-gray-300" />
              <span class="ml-2">Female</span>
            </label>
          </div>
        </div>
        <div class="flex-1 min-w-[200px] max-w-md"></div>
      </div>

      <div class="flex gap-6 justify-center">
        <button type="submit" class="px-8 py-2 bg-green-600 hover:bg-green-700 text-white rounded font-semibold text-lg transition">Save</button>
        <button type="button" class="px-8 py-2 bg-gray-600 hover:bg-gray-700 text-white rounded font-semibold text-lg transition" @click="handleClear">Clear</button>
      </div>
    </form>
    <div v-if="error" class="text-red-600 mt-4 text-center">{{ error }}</div>
  </div>
</template>

<script>
export default {
  name: 'Test4Page',
  data() {
    return {
      error: '',
      toasts: [],
      firstName: '',
      lastName: ''
    }
  },
  methods: {
    handleSave() {
      // ตัวอย่างการตรวจสอบข้อมูลแบบง่าย
      if (!this.firstName || !this.lastName) {
        this.error = 'กรุณากรอกข้อมูลให้ครบถ้วน';
      } else {
        this.showToast('Save Success!', 'success');
        this.error = '';
      }
    },
    showToast(message, type = 'success') {
      const id = Date.now() + Math.random();
      this.toasts.push({ id, message, type });
      setTimeout(() => {
        this.toasts = this.toasts.filter(t => t.id !== id);
      }, 3000);
    }
  }
}
</script>
