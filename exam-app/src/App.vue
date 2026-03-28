<template>
  <div class="min-h-screen bg-slate-50">
    <header class="bg-white/80 backdrop-blur-md sticky top-0 z-40 border-b border-slate-100 shadow-sm">
      <nav class="max-w-7xl mx-auto px-6 py-4 flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 bg-indigo-600 rounded-xl flex items-center justify-center text-white shadow-lg shadow-indigo-200">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18 18.246 18.477 16.5 18c-1.746 0-3.332.477-4.5 1.253"></path></svg>
          </div>
          <span class="text-xl font-black text-slate-800 tracking-tight">EXAM APP</span>
        </div>
        <div class="flex gap-1 bg-slate-100 p-1 rounded-2xl">
          <router-link to="/" class="px-5 py-2 rounded-xl text-sm font-bold transition-all hover:text-indigo-600 active:scale-95">Home</router-link>
          <router-link to="/test1" class="px-5 py-2 rounded-xl text-sm font-bold transition-all hover:text-indigo-600 active:scale-95">Test1</router-link>
          <router-link to="/test2" class="px-5 py-2 rounded-xl text-sm font-bold transition-all hover:text-indigo-600 active:scale-95">Test2</router-link>
          <router-link to="/test3" class="px-5 py-2 rounded-xl text-sm font-bold transition-all hover:text-indigo-600 active:scale-95">Test3</router-link>
          <router-link to="/test4" class="px-5 py-2 rounded-xl text-sm font-bold transition-all hover:text-indigo-600 active:scale-95">Test4</router-link>
          <router-link to="/test5" class="px-5 py-2 rounded-xl text-sm font-bold transition-all hover:text-indigo-600 active:scale-95">Test5</router-link>
        </div>
      </nav>
    </header>
    
    <main class="max-w-7xl mx-auto">
      <router-view />
    </main>
  </div>
</template>

<style>
/* CSS สำหรับ Active Link */
.router-link-active {
  background-color: white !important;
  color: #4f46e5 !important; /* indigo-600 */
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
}
</style>

<script>
export default {
  name: 'App',
}
</script>