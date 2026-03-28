<template lang="html">
    <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-indigo-50 to-slate-100 p-6">
        <div class="max-w-md w-full bg-white rounded-[3rem] shadow-2xl shadow-indigo-100 p-12 text-center border border-white relative overflow-hidden">
            <div class="absolute top-0 left-0 w-full h-2 bg-gradient-to-r from-indigo-500 to-blue-500"></div>
            
            <div class="mb-8 relative inline-block">
                <div class="w-24 h-24 bg-indigo-50 rounded-full flex items-center justify-center text-indigo-600 mx-auto">
                    <svg class="w-12 h-12" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
                </div>
                <div class="absolute -bottom-1 -right-1 w-8 h-8 bg-emerald-500 rounded-full border-4 border-white flex items-center justify-center text-white">
                    <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20"><path d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"></path></svg>
                </div>
            </div>

            <h1 class="text-4xl font-black text-slate-800 mb-4 tracking-tight">Success!</h1>
            <p class="text-slate-500 font-medium mb-8">You have successfully logged in as:</p>
            
            <div class="bg-slate-50 rounded-2xl p-6 border border-slate-100 mb-10">
                <span class="block text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">Authenticated Account</span>
                <span class="text-2xl font-bold text-indigo-600 tracking-tight">{{ username }}</span>
            </div>

            <a href="/test2" class="inline-flex items-center gap-2 text-slate-400 hover:text-slate-600 font-bold transition-all text-sm group">
                <svg class="w-4 h-4 transition-transform group-hover:-translate-x-1" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
                Sign Out
            </a>
        </div>
    </div>
</template>

<script>
export default {
  data() {
    return {
      username: ''
    }
  },
  created() {
    this.username = localStorage.getItem('username') || '';
  }
}
</script>