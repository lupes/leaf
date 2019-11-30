<template>
    <a-icon type="rollback" theme="outlined" :style="{ fontSize: '20px', color: '#08c' }" @click="goBack"/>
</template>

<script>
  export default {
    name: "Back",
    methods: {
      goBack: function () {
        this.$router.back();
      }
    }
  }
</script>

<style scoped>

</style>