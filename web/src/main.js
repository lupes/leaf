import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import {Button, Col, Dropdown, Icon, Input, Layout, List, Menu, message, Modal, Popconfirm, Row} from 'ant-design-vue';
import ace from 'ace-builds'
import componentsInstall from './components/install'
import moment from "moment";
import http from "./components/http";

Vue.use(ace);
Vue.use(http);
Vue.use(componentsInstall);
Vue.component(Layout.name, Layout);
Vue.component(Layout.Header.name, Layout.Header);
Vue.component(Layout.Sider.name, Layout.Sider);
Vue.component(Layout.Content.name, Layout.Content);
Vue.component(Layout.Footer.name, Layout.Footer);
Vue.component(Menu.name, Menu);
Vue.component(Menu.Item.name, Menu.Item);
Vue.component(Modal.name, Modal);
Vue.component(Input.name, Input);
Vue.component(Input.TextArea.name, Input.TextArea);
Vue.component(Button.name, Button);
Vue.component(Row.name, Row);
Vue.component(Col.name, Col);
Vue.component(List.name, List);
Vue.component(List.Item.name, List.Item);
Vue.component(List.Item.Meta.name, List.Item.Meta);
Vue.component(Icon.name, Icon);
Vue.component(Dropdown.name, Dropdown);
Vue.component(Popconfirm.name, Popconfirm);
Vue.prototype.$message = message;
Vue.config.productionTip = false;

Vue.filter('datetime', function (value) {
  if(value !== "") {
    return moment(value).format('YYYY-MM-DD HH:mm:ss')
  }
  return ""
});

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
