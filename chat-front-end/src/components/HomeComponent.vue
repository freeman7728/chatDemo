<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-18 12:44:14
 * @LastEditTime: 2024-09-18 20:23:56
 * @LastEditors: freeman7728
-->
<template>
    <div class="home-container">
        <!-- 添加功能按钮 -->
        <div class="tabar">
          
        </div>
        <!-- 渲染该用户的好友列表 -->
        <div class="friends-list">

        </div>
        
        <div class="chat">

        </div>
    </div>
</template>

<script lang="ts" setup>
import { useUserStore } from '@/stores/user'
import { onMounted } from 'vue';
import router from '@/router';
import iziToast from 'izitoast';
const store = useUserStore()
onMounted(async () => {
  try {
    console.log()
    const isAuthorized = await store.checkAuthorization();
    if(!isAuthorized){
      console.log('Unauthorized. Please log in.');
            // 处理 403 错误，例如重定向到登录页面
      iziToast.warning({
        transitionIn:'fadeInDown',
        position:'topCenter',
        title: '登录状态异常',
        message: '请重新登录',
        });
      router.push('/login');
    }
  } catch (error) {
    // 处理授权检查失败的情况，例如重定向到错误页面
    router.push('/error');
  }
});
</script>

<style>
.home-container{
  width: 1300px;
  height: 800px;
  border: 3px,solid;
  border-color: aqua;
  display: flex; 
  justify-content: flex-start;
}
.tabar{
  width: calc(13% - 40px); /* 考虑左右 margin */
  height: calc(100% - 40px); /* 考虑上下 margin */
  border: 3px solid aqua;
  margin: 20px 0px 20px 20px;
}
.friends-list{
  width: calc(30% - 40px); /* 考虑左右 margin */
  height: calc(100% - 40px); /* 考虑上下 margin */
  border: 3px solid aqua;
  margin: 20px 0px 20px 0px;
}
.chat{
  width: calc(63% - 40px); /* 考虑左右 margin */
  height: calc(100% - 40px); /* 考虑上下 margin */
  border: 3px solid aqua;
  margin: 20px 0px 20px 0px;
}
</style>