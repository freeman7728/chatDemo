/*
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-16 14:22:51
 * @LastEditTime: 2024-09-18 19:25:20
 * @LastEditors: freeman7728
 */
// src/stores/userStore.ts
import { defineStore } from 'pinia'
import apiClient from '@/plugins/axios'
import router from '@/router';
import axios from 'axios';
import iziToast from 'izitoast';
interface UserState {
  userInfo: { name: string } | null;
  token: string;
  data: any;
  resp:any;
}
export const useUserStore = defineStore('userStore', {
  state: ():UserState => ({
    userInfo: null,
    token: "",
    data:null,
    resp:null
  }),
  actions: {
    async login(credentials: { user_name: string; password: string }) {
      try {
        const response = await apiClient.post('/login', credentials)
        this.resp = response
        this.data = response.data
        if(response.data.data){
          this.token = response.data.data.token
        }
      } catch (error) {
        throw new Error('未知错误')
      }
    },
    // 定义函数类型
  async checkAuthorization(): Promise<boolean> {
      try {
        // 发起 GET 请求
        const response = await apiClient.get('/ping');
        return true
        // 处理成功响应
      } catch (error) {
        if (axios.isAxiosError(error)) {
          // 处理 Axios 错误
          // console.error('AxiosError:', {
          //   message: error.message,
          //   response: error.response,
          //   request: error.request,
          // });
          if (error.response?.status === 403) {
            return false
          }
        } else {
          // 处理其他类型的错误
          console.error('Unknown Error:', error);
          return false
        }
        return false
      }
    }
  }
})