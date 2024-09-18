/*
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-16 16:00:36
 * @LastEditTime: 2024-09-18 14:46:27
 * @LastEditors: freeman7728
 */
// src/plugins/axios.ts
import axios from 'axios'

const apiClient = axios.create({
  baseURL: 'http://localhost:3000',  // 后端API的基础URL
  timeout: 10000,                      // 请求超时时间
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
    //'Authorization':'yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJteS1pc3N1ZXIiLCJzdWIiOiIxMjMxMjMiLCJleHAiOjE3MjY2NDUyMjYsImlkIjowLCJ1c2VybmFtZSI6IjEyMzEyMyIsInJvbGUiOiJhZG1pbiJ9.CBp0x-7JvoTodSTfJYprAlAswS4elqV6CY11YwvpohY',
  }
})

// 请求拦截器
apiClient.interceptors.request.use(config => {
  // 你可以在这里添加 Authorization Token 或者其他配置
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `${token}`;
  }
  return config
}, error => {
  return Promise.reject(error)
})

// 响应拦截器
apiClient.interceptors.response.use(response => {
  return response
}, error => {
  // 在这里处理 API 错误，可能包括 401 未授权等
  return Promise.reject(error)
})

export default apiClient
