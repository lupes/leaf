<template>
    <a-icon type="rollback" theme="outlined" :style="{ fontSize: '20px', color: '#08c' }" @click="goBack"/>
</template>

<script>
  export default {
    name: "Back",
    methods: {
      goBack: function () {
        window.history.length > 1 ? this.$router.go(-1) : this.$router.push('/')
      }
    }
  }
</script>

<style scoped>

</style>