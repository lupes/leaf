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
                <a-textarea
                        v-model="solution.caption"
                        placeholder="解题思路"
                        :autosize="{ minRows: 4, maxRows: 8 }"
                />
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="18">
                <ace v-bind:value="solution.content" v-bind:readOnly="readOnly"></ace>
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
          content: ""
        },
        readOnly: true,
      }
    },
    mounted() {
      this.getSolution(this.$route.params.id, res => {
        this.solution = res.data.data;
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
      }
    }
  }
</script>

<style scoped>

</style>