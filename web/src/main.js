import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import {Button, Col, List, Row} from 'ant-design-vue';
import ace from 'ace-builds'
import componentsInstall from './components/install'

Vue.use(ace);
Vue.use(componentsInstall);
Vue.component(Button.name, Button);
Vue.component(Row.name, Row);
Vue.component(Col.name, Col);
Vue.component(List.name, List);
Vue.component(List.Item.name, List.Item);
Vue.component(List.Item.Meta.name, List.Item.Meta);
Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
