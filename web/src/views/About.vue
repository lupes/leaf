<template>
    <div>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <Back/>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <div class="home">
                    <img alt="Vue logo" src="../assets/logo.png">
                    <HelloWorld v-bind:msg="msg"/>
                    <a-button v-on:click="ping">ping</a-button>
                </div>
            </a-col>
        </a-row>
    </div>
</template>

<script>
  // @ is an alias to /src
  import HelloWorld from '@/components/HelloWorld.vue'

  export default {
    name: 'home',
    data() {
      return {
        msg: ""
      }
    },
    components: {
      HelloWorld
    },
    methods: {
      ping: function () {
        this.http.ping(this, data => {
          this.msg = data
        })
      }
    }
  }
</script>

<style>
    .home {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
    }
</style>
