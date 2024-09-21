<template>
    <div class="login-form">
      <h2>注册账号</h2>
      <div class="img-container">
          <img src="/src/assets/Go-Logo_Blue.png" alt="">
      </div>
      <div>
        <label for="username">账号:</label>
        <input 
          v-model="username" 
          type="text"
          id="username" 
          placeholder="请输入账号或邮箱"
        />
      </div>
      <div>
        <label for="password">密码:</label>
        <input 
          v-model="password" 
          type="password" 
          id="password" 
          placeholder="请输入密码"
        />
      </div>
      <div>
        <button @click="handleRegister">确认注册</button>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue'
  import { useUserStore } from '@/stores/user'
  import iziToast from 'izitoast';
  import { useRouter } from 'vue-router';
  
  // 使用 ref 管理表单数据
  const username = ref('')
  const password = ref('')
  const router = useRouter()
  
  
  // 获取 userStore 以便发起登录请求
  const userStore = useUserStore()
  
  const handleRegister = async () => {
    if (username.value && password.value) {
      try {
        // 调用 Pinia 中的 login 方法
        await userStore.register({
          user_name: username.value,
          password: password.value
        })
        if(userStore.data.status == 200){
          iziToast.success({
          position:'topCenter',
          transitionIn:'fadeInDown',
          title: '注册成功',
          message: 'Welcome',
          });
          setTimeout(()=>{
            router.push('/login')
          },2000)
        }else if (userStore.data.status == 500){
          iziToast.warning({
            transitionIn:'fadeInDown',
            position:'topCenter',
            title: '账号重复',
            message: '请更改账号',
          });
        }
      } catch (error) {
        console.log(error)
        iziToast.warning({
          transitionIn:'fadeInDown',
          position:'topCenter',
          title: '未知错误',
          message: '',
          });
      }
    }else{
      iziToast.warning({
        transitionIn:'fadeInDown',
        position:'topCenter',
        title: '账号和密码不能为空',
        message: '请输入账号和密码',
      });
    }
  }
  </script>
  
  <style scoped>
  
  .img-container{
    z-index: 999;
    position: absolute;
    top: -200px;
    left:450px;
    width: 500px;
    height: 500px;
    rotate: 30deg;
  }
  .img-container img {
    width: 100%; /* 图片将继承容器的宽度 */
    height: 100%; /* 图片将继承容器的高度 */
    object-fit: cover; /* 确保图片不会失真 */
    opacity:0.5;
  }
  .under-area-item{
    margin: 50px;
  }
  .under-area-item a{
    font-size: 15px;
    color: #19b6e6;
  }
  .under-area{
    display: flex;
    justify-content:center;
    align-items:center;
    gap: 50px;
  }
  .login-form {
    width: 800px;
    height: 500px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    box-shadow: 0 5px 20px rgba(0, 0, 0, 0.7);
    backdrop-filter: blur(13px) brightness(163%);
    background-color: rgba(255, 255, 255, 0.54);
    border-radius: 10px;
  }
  
  .login-form h2 {
    text-align: center;
    margin-bottom: 20px;
    font-weight: bold;
    font-size: 40px;
  }
  
  .login-form label {
    display: block;
    margin-bottom: 5px;
    font-weight: bold;
    font-size: 20px;
  }
  
  .login-form input {
    width: 100%;
    padding: 20px;
    margin-bottom: 15px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 20px;
  }
  
  .login-form button {
    width: 100%;
    padding: 10px;
    background-color: #19b6e6;
    color: white;
    border: none;
    border-radius: 4px;
    font-size: 20px;
    cursor: pointer;
  }
  
  .login-form button:hover {
    background-color: #00CC33;
  }
  
  </style>
  