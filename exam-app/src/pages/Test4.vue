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
        save data success ID : {{ toast.id }}
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
          <input type="text" id="email" v-model="email" class="border border-gray-300 rounded px-3 py-2 w-full" />
        </div>
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-1 font-medium text-left" for="phone">Phone</label>
          <input type="text" id="phone" v-model="phone" class="border border-gray-300 rounded px-3 py-2 w-full" />
        </div>
      </div>

      <div class="flex flex-wrap gap-6 mb-4 justify-center">
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-1 font-medium text-left" for="profile">Profile</label>
          <input type="file" id="profile" ref="profileInput" @change="handleFileUpload" class="border border-gray-300 rounded px-3 py-2 w-full" accept="image/*,.pdf" />
        </div>
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-1 font-medium text-left" for="birthDay">Birth Day</label>
          <input type="date" id="birthDay" v-model="birthDay" class="border border-gray-300 rounded px-3 py-2 w-full" />
        </div>
      </div>

      <div class="flex flex-wrap gap-6 mb-4 justify-center">
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-1 font-medium text-left" for="occupation">Occupation</label>
          <select id="occupation" v-model="occupation" class="border border-gray-300 rounded px-3 py-2 w-full">
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
              <input type="radio" name="sex" value="male" v-model="sex" class="border border-gray-300" />
              <span class="ml-2">Male</span>
            </label>
            <label class="inline-flex items-center">
              <input type="radio" name="sex" value="female" v-model="sex" class="border border-gray-300" />
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
import axios from 'axios';

export default {
  name: 'Test4Page',
  data() {
    return {
      error: '',
      toasts: [],
      firstName: '',
      lastName: '',
      email: '',
      phone: '',
      profileFile: null,
      birthDay: '',
      occupation: '',
      sex: ''
    }
  },
  methods: {
    handleFileUpload(event) {
      if (event.target.files.length > 0) {
        this.profileFile = event.target.files[0];
      } else {
        this.profileFile = null;
      }
    },
    async handleSave() {
      // ตรวจสอบว่ามีช่องไหนเว้นว่างไว้หรือเปล่า (ถ้าเป็นค่าว่าง จะมีปัญหาเวลาส่งวันที่เข้าฐานข้อมูล)
      if (!this.firstName || !this.lastName || !this.email || !this.phone || !this.birthDay || !this.occupation || !this.sex) {
        this.error = 'กรุณาตรวจสอบและกรอกข้อมูลให้ครบถ้วนทุกช่องครับ';
        return;
      }
      this.error = '';

      const formData = new FormData();
      formData.append('fname', this.firstName);
      formData.append('lname', this.lastName);
      formData.append('email', this.email);
      formData.append('phone', this.phone);
      formData.append('birthdate', this.birthDay);
      formData.append('occupation', this.occupation);
      formData.append('sex', this.sex);
      
      if (this.profileFile) {
        formData.append('profile', this.profileFile); // แนบไฟล์ต้นฉบับไปเลย
      }
      console.log(formData)
      try {
        const response = await axios.post('/api/test/test4/save', formData, {
          headers: { 'Content-Type': 'multipart/form-data' }
        });
        const data = response.data;
        this.handleClear();
        this.error = '';
        this.showToast(data.id, 'success');
      } catch (error) {
        console.error(error);
        this.showToast('บันทึกข้อมูลไม่สำเร็จ', 'error');
      }
    },
    showToast(id, type = 'success') {
      this.toasts.push({id, type});
      setTimeout(() => {
        this.toasts = this.toasts.filter(t => t.id !== id);
      }, 3000);
    },
    handleClear() {
      this.firstName = '';
      this.lastName = '';
      this.email = '';
      this.phone = '';
      this.profileFile = null;
      this.birthDay = '';
      this.occupation = '';
      this.sex = '';
      if (this.$refs.profileInput) {
        this.$refs.profileInput.value = '';
      }
    }
  }
}
</script>
