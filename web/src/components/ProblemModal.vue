<template>
    <a-modal
            :title="modalTitle"
            cancelText="取消"
            okText="确认"
            centered
            :closable="false"
            v-model="open"
            @cancel="handleCancel"
            @ok="handleOk">
        <a-input placeholder="题目编号" v-model="no" :allowClear="true"/>
        <a-input placeholder="题目标题" v-model="title" :allowClear="true" style="margin-top: 24px;"/>
        <a-select defaultValue="简单" style="width: 100%;margin-top: 24px;" v-model="difficulty">
            <a-select-option value="Easy">简单</a-select-option>
            <a-select-option value="Medium">中等</a-select-option>
            <a-select-option value="Hard">困难</a-select-option>
        </a-select>
        <a-select mode="tags" style="width: 100%;margin-top: 24px;" @change="handleChange" :tokenSeparators="[';']"
                  placeholder="标签"
                  v-model="topics" :options="options">
        </a-select>
        <a-input placeholder="题目连接" v-model="url" :allowClear="true" style="margin-top: 24px;"/>
        <a-textarea placeholder="题目内容" v-model="content" :allowClear="true" :autosize="{ minRows: 8, maxRows: 8 }"
                    style="margin-top: 24px;"/>
    </a-modal>
</template>

<script>
  export default {
    name: "ProblemModal",
    model: {
      prop: 'visible',
      event: 'change'
    },
    props: {
      visible: Boolean,
      modalTitle: String,
      problemNo: String,
      problemTitle: String,
      problemDifficulty: String,
      problemTopics: String,
      problemUrl: String,
      problemContent: String,
    },
    watch: {
      visible(val) {
        this.open = val
      },
      problemNo(val) {
        this.no = val
      },
      problemTitle(val) {
        this.title = val
      },
      problemDifficulty(val) {
        this.difficulty = val
      },
      problemTopics(val) {
        this.topics = val.split(';')
      },
      problemUrl(val) {
        this.url = val
      },
      problemContent(val) {
        this.content = val
      },
    },
    data() {
      return {
        tmpTopics: this.problemTopics,
        options: [
          {
            value: '广度优先搜索',
            label: '广度优先搜索',
          }, {
            value: '动态规划',
            label: '动态规划',
          }, {
            value: '贪心算法',
            label: '贪心算法',
          }, {
            value: '哈希表',
            label: '哈希表',
          }, {
            value: '字符串',
            label: '字符串',
          }, {
            value: '数组',
            label: '数组',
          }, {
            value: '数学',
            label: '数学',
          }, {
            value: '队列',
            label: '队列',
          }, {
            value: '栈',
            label: '栈',
          }, {
            value: '树',
            label: '栈',
          }
        ],
        open: false,
        no: this.problemNo,
        title: this.problemTitle,
        difficulty: this.problemDifficulty,
        topics: this.problemTopics,
        url: this.problemUrl,
        content: this.problemContent,
      }
    },
    methods: {
      handleChange(value) {
        this.tmpTopics = value.join(";")
      },
      handleCancel: function handleCancel(e) {
        this.$emit('cancel', e);
        this.$emit('change', false);
      },
      handleOk: function handleOk(e) {
        this.$emit('ok', {
          no: this.no,
          title: this.title,
          url: this.url,
          difficulty: this.difficulty,
          topics: this.tmpTopics,
          content: this.content
        }, e);
      },
    }
  }
</script>

<style scoped>

</style>