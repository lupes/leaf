<template>
    <div>
        <a-row>
            <a-col :offset="4" :span="16">
                <ace v-bind:value="solution.content" v-bind:readOnly="readOnly" v-on:input="getValue"></ace>
            </a-col>
        </a-row>
        <a-row type="flex">
            <a-col>
                <p class="height-50"></p>
            </a-col>
        </a-row>
        <a-row>
            <a-col :offset="7" :span="2">
                <a-button type="primary" v-on:click="test">Show</a-button>
            </a-col>
            <a-col :offset="7" :span="2">
                <a-button type="primary">Submit</a-button>
            </a-col>
        </a-row>
    </div>
</template>

<script>
  import axios from "axios"

  export default {
    name: 'Solution',
    data() {
      return {
        now: "",
        solution: {
          content: "",
        },
        readOnly: false,
      }
    },
    mounted() {
      this.getSolution(this.$route.params.id, res => {
        console.log(res.data.data);
        this.solution = res.data.data;
        this.content = res.data.data.content;
      });
    },
    methods: {
      getSolution(solution_id, callback) {
        let app = this;
        axios.get("http://localhost/api/v1/solution/" + solution_id)
          .then(callback)
          .catch(function (error) {
            app.$message.error("请求失败 " + error);
          })
      },
      getValue: function (value) {
        this.now = value
      },
      test: function (res) {
        console.log(res)
      }
    }
  }
</script>

<style scoped>

</style>