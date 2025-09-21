import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],  //在开发过程使用Vue插件，处理.vue但文件，编译为浏览器可识别的js和css
  server: {
    port:5173,
    host:'localhost',
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
      }
    }
  }
})


// 配置Vite的前端开发服务器
// 运行在5173端口