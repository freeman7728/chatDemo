/* eslint-disable @typescript-eslint/no-unused-vars */
/*
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-16 14:22:51
 * @LastEditTime: 2024-09-22 14:54:03
 * @LastEditors: freeman7728
 */
import { defineStore } from 'pinia'
import apiClient from '@/plugins/axios'
import axios from 'axios';
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
    async register(credentials: { user_name: string; password: string }) {
      try {
        const response = await apiClient.post('/register', credentials)
        this.resp = response
        this.data = response.data
      } catch (error) {
        throw new Error('未知错误')
      }
    },
    // 定义函数类型
  async checkAuthorization(): Promise<boolean> {
      try {
        // 发起 GET 请求
        const response = await apiClient.get('/ping');
        localStorage.setItem("id",response.data.data)
        return true
        // 处理成功响应
      } catch (error) {
        if (axios.isAxiosError(error)) {
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