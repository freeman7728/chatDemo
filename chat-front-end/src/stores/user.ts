/*
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-16 14:22:51
 * @LastEditTime: 2024-09-18 13:21:48
 * @LastEditors: freeman7728
 */
// src/stores/userStore.ts
import { defineStore } from 'pinia'
import apiClient from '@/plugins/axios'
interface UserState {
  userInfo: { name: string } | null;
  token: string | null;
  data: any;
  resp:any;
}
export const useUserStore = defineStore('userStore', {
  state: ():UserState => ({
    userInfo: null,
    token: null,
    data:null,
    resp:null
  }),
  actions: {
    async login(credentials: { user_name: string; password: string }) {
      try {
        const response = await apiClient.post('/login', credentials)
        this.resp = response
        this.data = response.data
        this.token = response.data.data.token
      } catch (error) {
        throw new Error('登录失败')
      }
    },
  }
})