<template>
    <a-row>
        <a-col>
            <div class="home">
                <img alt="Vue logo" src="../assets/logo.png">
                <HelloWorld v-bind:msg="msg"/>
                <a-button v-on:click="ping">ping</a-button>
            </div>
        </a-col>
    </a-row>
</template>

<script>
  // @ is an alias to /src
  import HelloWorld from '@/components/HelloWorld.vue'
  import _ from "lodash"
  import axios from 'axios'

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
        let app = this;
        axios.get('http://localhost/api/ping')
          .then(function (response) {
            app.msg = _.capitalize(response.data.data)
          })
          .catch(function (error) {
            app.msg = 'Error! Could not reach the API. ' + error
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
