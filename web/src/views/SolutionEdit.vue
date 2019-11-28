<template>
    <div>
        <a-row type="flex" justify="start">
            <a-col :offset="3" :span="18">
                <back />
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <a-input placeholder="标题" v-model="solution.title"/>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <a-textarea
                        v-model="solution.caption"
                        placeholder="解题思路"
                        :autosize="{ minRows: 4, maxRows: 8 }"
                />
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <ace v-model="solution.content" v-bind:readOnly="readOnly"></ace>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="2">
                <a-button type="primary" @click="editSolution">更新</a-button>
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
        this.$message.info("请求成功");
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
      editSolution() {
        let app = this;
        this.solution.content = this.now;
        axios.put("http://localhost/api/v1/solution", this.solution,
        ).then(function (res) {
          console.log(res)
          app.$message.info("题解编辑成功");
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