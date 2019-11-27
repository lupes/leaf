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
        solution: {
          problem_id: parseInt(this.$route.params.id),
          title: "",
          language: "golang",
          caption: "",
          content: "",
        },
        now: "",
        readOnly: false,
      }
    },
    methods: {
      addSolution() {
        let app = this;
        this.solution.content = this.now;
        axios.post("http://localhost/api/v1/solution", this.solution,
        ).then(function (res) {
          app.$message.info("题解添加成功");
          console.log(res)
        }).catch(function (error) {
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