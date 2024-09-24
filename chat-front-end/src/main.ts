/*
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-18 12:15:23
 * @LastEditTime: 2024-09-24 11:12:05
 * @LastEditors: freeman7728
 */
import './assets/main.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { initRem } from './plugins/rem'

import App from './App.vue'
import router from './router'

initRem()
const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')

