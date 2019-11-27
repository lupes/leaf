<template>
    <div>
        <a-row>
            <a-col :offset="2" :span="18">
                <a-input placeholder="标题" v-model="solution.title"/>
            </a-col>
        </a-row>
        <a-row style="margin: 10px"></a-row>
        <a-row>
            <a-col :offset="2" :span="18">
                <a-textarea
                        v-model="solution.caption"
                        placeholder="解题思路"
                        :autosize="{ minRows: 4, maxRows: 8 }"
                />
            </a-col>
        </a-row>
        <a-row style="margin: 10px"></a-row>
        <a-row>
            <a-col :offset="2" :span="18">
                <ace v-bind:value="solution.content" v-bind:readOnly="readOnly" v-on:input="getValue"></ace>
            </a-col>
        </a-row>
        <a-row style="margin: 10px"></a-row>
        <a-row>
            <a-col :offset="7" :span="2">
                <a-button type="primary" @click="addSolution">Submit</a-button>
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