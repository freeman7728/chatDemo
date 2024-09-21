/*
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-18 12:15:23
 * @LastEditTime: 2024-09-21 16:51:09
 * @LastEditors: freeman7728
 */
import { createRouter, createWebHistory } from 'vue-router'
import LoginComponent from '@/components/LoginComponent.vue'
import HomeComponent from '@/components/HomeComponent.vue'
import ChatComponent from '@/components/ChatComponent.vue'
import AddFriendComponent from '@/components/AddNewFriendComponent.vue'
import DelFriendComponent from '@/components/DelFriendComponent.vue'
import RegisterComponent from '@/components/RegisterComponent.vue'
//配置路由规则
const routes = [
  {path:"/",redirect:"/login"},
  {path:"/login",component:LoginComponent},
  {path:"/register",component:RegisterComponent},
  {
    path:"/home",
    component:HomeComponent,
    children:[
      {
        path:"chat",
        name:"ChatPage",
        component:ChatComponent
      },
      {
        path:"add-friend",
        name:"AddNewFriend",
        component:AddFriendComponent
      },
      {
        path:"del-friend",
        name:"DelFriend",
        component:DelFriendComponent
      }
    ]
  },
]
//创建路由器
const router = createRouter ({
  history:createWebHistory(),
  routes:routes
})
//加载路由器

export default router
