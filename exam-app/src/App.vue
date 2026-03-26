
<template>
  <nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/test1">Test1</router-link> |
    <router-link to="/test2">Test2</router-link> |
    <router-link to="/test3">Test3</router-link> |
    <router-link to="/test4">Test4</router-link> |
    <router-link to="/test5">Test5</router-link>
  </nav>
  <router-view />
</template>

<script>
export default {
  name: 'App',
}
</script>