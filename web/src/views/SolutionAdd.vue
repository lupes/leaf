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
                <ace v-bind:value="solution.content" v-bind:readOnly="readOnly" v-on:update:value="now = $event"></ace>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="2">
                <a-button type="primary" @click="addSolution">添加</a-button>
            </a-col>
        </a-row>
    </div>
</template>

<script>

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
        this.solution.content = this.now;
        this.http.addSolution(this, this.solution, () => {
          this.$router.push('/problem/' + this.$route.params.id);
        });
      },
    }
  }
</script>

<style scoped>

</style>