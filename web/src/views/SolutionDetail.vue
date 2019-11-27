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