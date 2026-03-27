<template lang="html">
    <div class="max-w-md mx-auto mt-16 p-8 border border-gray-300 rounded-lg bg-white shadow">
        <h2 class="text-black text-xl font-bold mb-4">Register</h2>
        <form @submit.prevent="handleRegister" class="space-y-4">
          <div>
            <label for="username" class="block mb-1 font-medium">Username:</label>
            <input id="username" v-model="username" type="text" required
              class="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400" />
          </div>
          <div>
            <label for="password" class="block mb-1 font-medium">Password:</label>
            <input id="password" v-model="password" type="password" required
              class="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400" />
          </div>
          <div>
            <label for="confirm-password" class="block mb-1 font-medium">Confirm Password:</label>
            <input id="confirm-password" v-model="confirmPassword" type="password" required
              class="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400" />
          </div>
          <button type="submit"
            class="w-full py-2 bg-blue-600 hover:bg-blue-700 text-white rounded font-semibold text-lg transition">Register</button>
        </form>
        <div v-if="error" class="text-red-600 mt-4 text-center">{{ error }}</div>
    </div>
</template>

<script>
export default {
  name: 'RegisterPage',
  data() {
    return {
      username: '',
      password: '',
      confirmPassword: '',
      error: ''
    }
  },
  methods: {
    handleRegister() {
      if (this.password !== this.confirmPassword) {
        this.error = 'Passwords do not match';
      } else {
        alert('Register Success!');
        this.error = '';
      }
    }
  }
}
</script>