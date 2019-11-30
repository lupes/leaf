<template>
    <div>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="12">
                <Back/>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="12">
                <ace v-model="code" v-bind:readOnly="readOnly" v-on:update:value="value= $event"></ace>
            </a-col>
        </a-row>
        <a-row type="flex" justify="center" style="margin: 10px 0">
            <a-col :span="8">
                <a-button type="primary" v-on:click="clear">清空</a-button>
                <a-button type="primary" v-on:click="update" style="float: right;">Submit</a-button>
            </a-col>
        </a-row>
    </div>
</template>

<script>
  export default {
    name: 'AceView',
    data() {
      return {
        value: "",
        code: `func main() {

}`,
        readOnly: false,
      }
    },
    methods: {
      update: function () {
        this.code = this.value;
        console.log('update:' + this.code);
      },
      clear: function () {
        console.log('clear:' + this.value);
        this.code = "func main(){}";
        console.log('clear:' + this.value);
      }
    }
  }
</script>
