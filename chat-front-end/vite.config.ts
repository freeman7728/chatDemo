/*
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-18 12:15:23
 * @LastEditTime: 2024-09-24 12:29:03
 * @LastEditors: freeman7728
 */
import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import pxtorem from 'postcss-pxtorem';

// https://vitejs.dev/config/

export default defineConfig({
  plugins: [
    vue(),
    pxtorem({ // 把px单位换算成rem单位
      rootValue: 37.5, //换算基数， 
      unitPrecision: 3, //允许REM单位增长到的十进制数字,小数点后保留的位数。
      propList: ['*'],
      exclude: /(node_module)/,  //默认false，可以（reg）利用正则表达式排除某些文件夹的方法，例如/(node_module)/ 。如果想把前端UI框架内的px也转换成rem，请把此属性设为默认值
      selectorBlackList: ['.van'], //要忽略并保留为px的选择器，本项目我是用的vant ui框架，所以忽略他
      mediaQuery: false,  //（布尔值）允许在媒体查询中转换px。
      minPixelValue: 1 //设置要替换的最小像素值
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    host: '0.0.0.0', // 允许外部访问
    port: 5173, // 您可以选择其他端口
  }
})
