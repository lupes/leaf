<template>
    <div class="ace-container">
        <div class="ace-editor" ref="ace"></div>
    </div>
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
    name: "Ace",
    model: {
      prop: 'value',
      event: 'change'
    },
    props: {
      value: String,
      readOnly: Boolean,
    },
    data() {
      return {
        aceEditor: null,
        theme: 'ace/theme/solarized_light',
        mode: 'ace/mode/golang'
      }
    },
    mounted() {
      this.aceEditor = ace.edit(this.$refs.ace, {
        value: this.value,
        maxLines: 25,
        minLines: 25,
        fontSize: 14,
        theme: this.theme,
        mode: this.mode,
        readOnly: this.readOnly,
        tabSize: 4,
        enableSnippets: true,
        enableLiveAutocompletion: true,
        enableBasicAutocompletion: true
      });
      this.aceEditor.getSession().on("change", this.change);
    },
    watch: {
      value: function (value) {
        this.aceEditor.getSession().setValue(value)
      }
    },
    methods: {
      change() {
        this.$emit('input', this.aceEditor.getSession().getValue());
      }
    }
  }
</script>

<style lang='css' scoped>
</style>
