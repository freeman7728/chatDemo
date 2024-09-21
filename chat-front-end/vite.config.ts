/*
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-18 12:15:23
 * @LastEditTime: 2024-09-21 16:45:06
 * @LastEditors: freeman7728
 */
import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
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
