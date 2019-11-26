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
        <a-input placeholder="题目标题" v-model="title"/>
        <div style="margin: 24px 0"></div>
        <a-input placeholder="题目连接" v-model="url"/>
        <div style="margin: 24px 0"></div>
        <a-textarea
                v-model="content"
                placeholder="题目内容"
                :autosize="{ minRows: 8, maxRows: 8 }"
        />
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
      problemTitle: String,
      problemUrl: String,
      problemContent: String,
    },
    watch: {
      visible(val) {
        this.open = val
      },
      problemTitle(val) {
        this.title = val
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
        open: false,
        title: this.problemTitle,
        url: this.problemUrl,
        content: this.problemContent,
      }
    },
    methods: {
      handleCancel: function handleCancel(e) {
        this.$emit('cancel', e);
        this.$emit('change', false);
      },
      handleOk: function handleOk(e) {
        this.$emit('ok', {title: this.title, url: this.url, content: this.content}, e);
      },
    }
  }
</script>

<style scoped>

</style>