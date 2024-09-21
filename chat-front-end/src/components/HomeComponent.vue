<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-18 12:44:14
 * @LastEditTime: 2024-09-21 16:14:13
 * @LastEditors: freeman7728
-->
<template>
    <div class="home-container">
        <!-- 添加功能按钮 -->
        <div class="tabar">
          <div class="logout">
            <button class="logout-btn" @click="logout">
              登出
            </button>
          </div>
          <div class="add-friend">
            <button class="add-friend-btn" @click="delFriend">
              删除好友
            </button>
          </div>
          <div class="add-friend">
            <button class="add-friend-btn" @click="addNewFriend">
              添加好友
            </button>
          </div>
        </div>
        <!-- 渲染该用户的好友列表 -->
        <div class="friends-list">
          <div v-if="relationStore.list.length === 0" class="if-null">
            好友列表为空
          </div>
          <!-- <FriendListComponent></FriendListComponent> -->
            <FriendListComponent 
                v-for="(friend,idx) in relationStore.list" 
                :key=idx
                :friendId="idx"
            />
        </div>
        
        <div class="chat">
          <RouterView :key="$route.fullPath">
            
          </RouterView>
        </div>
    </div>
</template>

<script lang="ts" setup>
import FriendListComponent from './FriendListComponent.vue';
import ChatComponent from './ChatComponent.vue';
import { useUserStore } from '@/stores/user'
import { onMounted } from 'vue';
import router from '@/router';
import iziToast from 'izitoast';
import { useRelationStore } from '@/stores/relation';
const store = useUserStore()
const relationStore = useRelationStore()
relationStore.getList()
const logout = () => {
  localStorage.removeItem("token")
  router.push('/login');
  iziToast.success({
      transitionIn:'fadeInDown',
      position:'topCenter',
      title: '登出成功',
      message: '已退出登录',
  });
}
const addNewFriend = () => {
  router.push('/home/add-friend')
}
const delFriend = () => {
  router.push('/home/del-friend')
}
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
  display: flex;
  flex-direction: column;
}
.add-friend{
  width: 100%;
  height: 120px;
  border: 3px red solid;
  display: flex;
  justify-content: center;
  align-items: center;
}
.add-friend-btn{
  font-size: 30px;
  margin: 5px;
  height: calc(100% - 10px);
  width: calc(100% - 10px);
  border: 3px orange solid;
  border-radius: 10px;
}
.logout{
  width: 100%;
  height: 120px;
  border: 3px red solid;
  margin-top: auto;
  display: flex;
  justify-content: center;
  align-items: center;
}
.logout-btn{
  font-size: 30px;
  margin: 5px;
  height: calc(100% - 10px);
  width: calc(100% - 10px);
  border: 3px orange solid;
  border-radius: 10px;
}
.if-null{
  position: absolute; /* 绝对定位 */
  top: 50%; /* 距离顶部 50% */
  left: 50%; /* 距离左侧 50% */
  transform: translate(-50%, -50%); /* 使用 transform 移动到中心 */
  font-size: 80px;
  font-weight: 20px;
}
.friends-list{
  width: calc(30% - 40px); /* 考虑左右 margin */
  height: calc(100% - 40px); /* 考虑上下 margin */
  border: 3px solid aqua;
  margin: 20px 0px 20px 0px;
  position: relative;
}
.chat{
  width: calc(63% - 40px); /* 考虑左右 margin */
  height: calc(100% - 40px); /* 考虑上下 margin */
  border: 3px solid aqua;
  margin: 20px 0px 20px 0px;
}
</style>