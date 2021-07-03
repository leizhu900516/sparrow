
// import store from '../store/store'
import Vue from 'vue'
import Vuex from 'vuex'

import {
  Button,
  Row,
  Container,
  Header,
  Main,
  Footer,
  Input,
  checkbox,
  Form,
  FormItem,
  Message,
  Menu,
  MenuItem,
  Submenu,
  Col,
  MenuItemGroup,
  Tabs,
  TabPane,
  Table,
  TableColumn,
  Dropdown,
  DropdownItem,
  DropdownMenu,
  Breadcrumb,
  BreadcrumbItem,
  Card,
  Select,
  Option,
  Popover,
  MessageBox,
  Dialog,
  Divider,
  Aside,
  Avatar,
  Upload,
  Pagination,
  Tooltip,
  Loading
} from 'element-ui'
import infiniteScroll from 'vue-infinite-scroll'

// Vue.use(store)
Vue.use(Vuex)
Vue.use(Button)
Vue.use(Tooltip)
Vue.use(Loading)
Vue.use(Row)
Vue.use(Col)
Vue.use(Container)
Vue.use(Header)
Vue.use(Popover)
Vue.use(Main)
Vue.use(Footer)
Vue.use(Input)
Vue.use(checkbox)
Vue.use(Form)
Vue.use(FormItem)
Vue.use(Menu)
Vue.use(MenuItem)
Vue.use(Submenu)
Vue.use(MenuItemGroup)
Vue.use(Tabs)
Vue.use(TabPane)
Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Dropdown)
Vue.use(DropdownMenu)
Vue.use(DropdownItem)
Vue.use(Breadcrumb)
Vue.use(BreadcrumbItem)
Vue.use(Card)
Vue.use(Select)
Vue.use(Option)
Vue.use(Dialog)
Vue.use(Divider)
Vue.use(Aside)
Vue.use(Avatar)
Vue.use(Upload)
Vue.use(Pagination)
Vue.use(infiniteScroll)
Vue.prototype.$message = Message
Vue.prototype.$msgbox = MessageBox
Vue.prototype.$confirm = MessageBox.confirm
