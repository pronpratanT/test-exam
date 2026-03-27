<template lang="html">
    <div>
        <h2>Welcome User : {{ username }}</h2>
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