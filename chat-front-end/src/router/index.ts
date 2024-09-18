/*
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-18 12:15:23
 * @LastEditTime: 2024-09-18 12:21:35
 * @LastEditors: freeman7728
 */
import { createRouter, createWebHistory } from 'vue-router'
import LoginComponent from '@/components/LoginComponent.vue'


//配置路由规则
const routes = [
  {path:"/",redirect:"/login"},
  {path:"/login",component:LoginComponent},
]
//创建路由器
const router = createRouter ({
  history:createWebHistory(),
  routes:routes
})
//加载路由器

export default router
