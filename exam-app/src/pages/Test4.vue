<template>
    <!-- Toast Notification -->
    <div class="fixed top-6 right-6 z-50 space-y-3">
      <div
        v-for="toast in toasts"
        :key="toast.id"
        :class="[
          'px-6 py-4 rounded-xl shadow-xl text-white font-medium transition-all duration-300 flex items-center gap-3',
          toast.type === 'success' ? 'bg-gradient-to-r from-emerald-500 to-green-500' : 'bg-gradient-to-r from-red-500 to-rose-500'
        ]"
      >
        save data success ID : {{ toast.id }}
      </div>
    </div>
  <div class="max-w-4xl mx-auto my-12 p-10 bg-white rounded-[2rem] shadow-2xl shadow-slate-200/50 border border-slate-100">
    <form @submit.prevent="handleSave" class="space-y-6">
      <div class="flex flex-wrap gap-6 mb-4 justify-center">
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-2 font-semibold text-slate-700 text-left text-sm" for="firstName">First Name</label>
          <input type="text" id="firstName" v-model="firstName" class="border border-slate-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 rounded-xl px-4 py-3 w-full transition-all outline-none bg-slate-50 hover:bg-white focus:bg-white" />
        </div>
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-2 font-semibold text-slate-700 text-left text-sm" for="lastName">Last Name</label>
          <input type="text" id="lastName" v-model="lastName" class="border border-slate-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 rounded-xl px-4 py-3 w-full transition-all outline-none bg-slate-50 hover:bg-white focus:bg-white" />
        </div>
      </div>

      <div class="flex flex-wrap gap-6 mb-4 justify-center">
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-2 font-semibold text-slate-700 text-left text-sm" for="email">Email</label>
          <input type="text" id="email" v-model="email" class="border border-slate-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 rounded-xl px-4 py-3 w-full transition-all outline-none bg-slate-50 hover:bg-white focus:bg-white" />
        </div>
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-2 font-semibold text-slate-700 text-left text-sm" for="phone">Phone</label>
          <input type="text" id="phone" v-model="phone" class="border border-slate-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 rounded-xl px-4 py-3 w-full transition-all outline-none bg-slate-50 hover:bg-white focus:bg-white" />
        </div>
      </div>

      <div class="flex flex-wrap gap-6 mb-4 justify-center">
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-2 font-semibold text-slate-700 text-left text-sm" for="profile">Profile</label>
          <input type="file" id="profile" ref="profileInput" @change="handleFileUpload" class="border border-slate-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 rounded-xl px-4 py-2.5 w-full transition-all outline-none bg-slate-50 hover:bg-white focus:bg-white file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-semibold file:bg-indigo-50 file:text-indigo-700 hover:file:bg-indigo-100" accept="image/*,.pdf" />
        </div>
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-2 font-semibold text-slate-700 text-left text-sm" for="birthDay">Birth Day</label>
          <input type="date" id="birthDay" v-model="birthDay" class="border border-slate-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 rounded-xl px-4 py-3 w-full transition-all outline-none bg-slate-50 hover:bg-white focus:bg-white" />
        </div>
      </div>

      <div class="flex flex-wrap gap-6 mb-4 justify-center">
        <div class="flex-1 min-w-[200px] max-w-md">
          <label class="block mb-2 font-semibold text-slate-700 text-left text-sm" for="occupation">Occupation</label>
          <select id="occupation" v-model="occupation" class="border border-slate-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 rounded-xl px-4 py-3 w-full transition-all outline-none bg-slate-50 hover:bg-white focus:bg-white appearance-none">
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
          <label class="block mb-3 font-semibold text-slate-700 text-left text-sm" for="sex">Sex</label>
          <div class="flex items-center gap-8 pl-4 py-2 bg-slate-50 rounded-xl border border-slate-100">
            <label class="inline-flex items-center hover:cursor-pointer group">
              <input type="radio" name="sex" value="male" v-model="sex" class="w-5 h-5 text-indigo-600 border-slate-300 focus:ring-indigo-500" />
              <span class="ml-3 text-slate-700 group-hover:text-indigo-600 transition-colors font-medium">Male</span>
            </label>
            <label class="inline-flex items-center hover:cursor-pointer group">
              <input type="radio" name="sex" value="female" v-model="sex" class="w-5 h-5 text-rose-500 border-slate-300 focus:ring-rose-400" />
              <span class="ml-3 text-slate-700 group-hover:text-rose-500 transition-colors font-medium">Female</span>
            </label>
          </div>
        </div>
        <div class="flex-1 min-w-[200px] max-w-md"></div>
      </div>

      <div class="flex gap-6 justify-center mt-10">
        <button type="submit" class="px-10 py-3.5 bg-gradient-to-r from-emerald-500 to-green-600 hover:from-emerald-600 hover:to-green-700 shadow-lg shadow-green-500/30 text-white rounded-xl font-bold text-lg transition-all duration-300 hover:-translate-y-1 active:scale-95 min-w-[140px]">Save</button>
        <button type="button" class="px-10 py-3.5 bg-white border-2 border-slate-200 hover:border-slate-300 hover:bg-slate-50 text-slate-600 rounded-xl font-bold text-lg transition-all duration-300 hover:-translate-y-1 active:scale-95 min-w-[140px]" @click="handleClear">Clear</button>
      </div>
    </form>
    <div v-if="error" class="text-rose-600 mt-6 p-4 bg-rose-50 rounded-xl font-medium text-center border border-rose-100 flex items-center justify-center gap-2">
      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
      {{ error }}
    </div>
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
