<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-18 12:44:14
 * @LastEditTime: 2024-09-18 19:25:36
 * @LastEditors: freeman7728
-->
<template>
    <div>
        <h1>
            这是home页面
        </h1>
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
</style>