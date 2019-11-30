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
                <ace v-bind:value="solution.content" v-bind:readOnly="readOnly"
                     v-on:update:value="now = $event"></ace>
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
      this.getSolution()
    },
    methods: {
      getSolution() {
        this.http.getSolution(this, this.$route.params.id, data => {
          this.solution = data;
        });
      },
      editSolution() {
        this.http.editSolution(this, this.solution, () => {
          this.getSolution()
        })
      },
    }
  }
</script>

<style scoped>

</style>