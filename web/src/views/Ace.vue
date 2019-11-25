<template>
  <div>
    <a-row :gutter="8">
      <a-col xs>
        <ace v-model="code" v-bind:readOnly="readOnly" v-on:input="getValue"></ace>
      </a-col>
    </a-row>
    <a-row :gutter="8">
      <a-col xs>
      </a-col>
    </a-row>
    <a-row :gutter="8">
      <a-col xs>
        <a-button type="primary" v-on:click="test">Show</a-button>
        <a-button type="primary">Submit</a-button>
      </a-col>
    </a-row>
  </div>
</template>

<script>
  export default {
    name: 'AceView',
    data() {
      return {
        code: `func main() {

}`,
        readOnly: false,
      }
    },
    methods: {
      getValue: function (value) {
        console.log('parent:' + value)
      },
      test: function () {
        this.code = ""
      }
    }
  }
</script>
