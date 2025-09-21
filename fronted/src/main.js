import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'   //整个应用布局骨架
import Login from './components/Login.vue'
import Register from './components/Register.vue'
import Dashboard from './components/Dashboard.vue'

const routes = [
  { path: '/', redirect: '/login' },   //用户访问根路径，重定向到/logon路径，设置默认初始界面
  { path: '/login', component: Login },   //URL为/login,Vue Router渲染Login组件
  { path: '/register', component: Register },
  { path: '/dashboard', component: Dashboard }
]
// 什么是路由器实例？就是管理页面导航的对象
// 它能监听URL变化，当浏览器地址栏输入不同网址，或点击链接，路由器会检测到这个变化
// 根据在routes数组定义规则，找到与当前URL匹配对象，将规则对应组件渲染到<router-view>占位符的位置上实现页面切换

// 创建一个vue路由器实例，路由模式是HTML5 history api，routes参数是路由配置（数组包含所有路由规则）
// 创建一个路由器，用干净URL，路由规则是routes
const router = createRouter({
  history: createWebHistory(),
  routes
})


// 创建Vue应用实例并传入根组件App，将路由器实例注入Vue应用中，应用中所有组件都可以使用路由功能
// 最后，将整个Vue应用挂在到html文件中id=app的DOM元素上，Vue会接管渲染到这个div内部
const app = createApp(App)
app.use(router)
app.mount('#app')
