<template>
  <a-row :gutter="16">
    <a-col :span="6"/>
    <a-col :span="12">
      <div class="ace-container">
        <div class="ace-editor" ref="ace"></div>
      </div>
    </a-col>
    <a-col :span="6"></a-col>
  </a-row>
</template>

<script>
  import ace from 'ace-builds'
  import 'ace-builds/webpack-resolver'
  import 'ace-builds/src-noconflict/theme-solarized_light'
  import 'ace-builds/src-noconflict/mode-golang'
  import 'ace-builds/src-noconflict/snippets/golang'
  import 'ace-builds/src-noconflict/ext-beautify'
  import 'ace-builds/src-noconflict/ext-language_tools'

  export default {
    props: {
      value: String,
      readOnly: Boolean,
    },
    mounted() {
      this.aceEditor = ace.edit(this.$refs.ace, {
        value: this.value,
        maxLines: 30,
        minLines: 30,
        fontSize: 14,
        theme: this.theme,
        mode: this.mode,
        readOnly: this.readOnly,
        tabSize: 4,
        enableSnippets: true,
        enableLiveAutocompletion: true,
        enableBasicAutocompletion: true
      });
      this.aceEditor.getSession().on("change", this.change)
    },
    data() {
      return {
        aceEditor: null,
        theme: 'ace/theme/solarized_light',
        mode: 'ace/mode/golang'
      }
    },
    name: "Ace",
    methods: {
      change() {
        this.$emit('input', this.aceEditor.getSession().getValue());
      }
    }
  }
</script>

<style lang='css' scoped>
  .scrollmargin {
    height: 80px;
  }
</style>
