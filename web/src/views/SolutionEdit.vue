<template>
    <div>
        <a-row type="flex" justify="start">
            <a-col :offset="3" :span="18">
                <back/>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <a-input placeholder="标题" v-model="solution.title"/>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <a-textarea v-model="solution.caption" placeholder="解题思路" :autosize="{ minRows: 4, maxRows: 8 }"/>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <ace v-bind:value="solution.content" v-bind:readOnly="readOnly" @input="getValue"></ace>
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
        if(res.status === 200 && res.data.code === 1) {
          this.$message.info("请求成功");
          this.solution = res.data.data;
        } else {
          this.$message.error("请求失败 " + res.data.message);
        }
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
          if(res.status === 200 && res.data.code === 1) {
            app.$message.info("题解编辑成功");
          } else {
            app.$message.error("请求失败 " + res.data.message);
          }
        }).catch(function (error) {
          app.$message.error("请求失败 " + error);
        })
      },
      getValue: function (value) {
        this.now = value
      },
    }
  }
</script>

<style scoped>

</style>