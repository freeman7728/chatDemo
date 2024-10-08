<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-18 12:44:14
 * @LastEditTime: 2024-10-08 11:14:30
 * @LastEditors: freeman7728
-->
<template>
    <div class="home-container">
        <!-- 添加功能按钮 -->
        <div class="tabar">
          <div class="local-id">
            本机id
          </div>
          <div class="local-id-container">
            {{ localId }}
          </div>
          <div class="logout">
            <button class="logout-btn" @click="logout">
              登出
            </button>
          </div>
          <div class="add-friend">
            <button class="add-friend-btn" @click="delFriend">
              删除<br>好友
            </button>
          </div>
          <div class="add-friend">
            <button class="add-friend-btn" @click="addNewFriend">
              添加<br>好友
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
import { onBeforeMount } from 'vue';
import { ref } from 'vue';
import { bu } from '@/plugins/axios';
const store = useUserStore()
const relationStore = useRelationStore()
let localId = ref<string | null>('')
relationStore.getList()
const logout = () => {
  localStorage.removeItem("token")
  localStorage.removeItem("id")
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
onBeforeMount(async () => {
  try {
    const isAuthorized = await store.checkAuthorization();
    localId.value = localStorage.getItem("id")
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

const socket = ref<WebSocket | null>(null);
    const connectWebSocket = () => {
    const wsurl = `${bu}/userWebsocket?token=${localStorage.getItem('token')}`
    socket.value = new WebSocket(wsurl);
    socket.value.onopen = () => {
        console.log('UserWebSocket连接已打开');
    };
    
    socket.value.onmessage = (event) => {
        const messageData = JSON.parse(event.data)
        console.log('接收到消息:', messageData.uid);
        relationStore.getList()
        // console.log('接收到消息:', message);
        // 这里可以更新消息列表
    };

    socket.value.onclose = () => {
        console.log('WebSocket连接已关闭');
    };

    socket.value.onerror = (error) => {
        console.error('WebSocket发生错误:', error);
    };
};
connectWebSocket()
</script>

<style>
.home-container{
  width: 1300px;
  height: 800px;
  border: black 1px solid;
  display: flex; 
  justify-content: flex-start;
  background-color: rgb(175, 174, 174);
  /* background-color:#2c2c2c */

}
.tabar{
  width: calc(13% - 40px); /* 考虑左右 margin */
  height: calc(100% - 40px); /* 考虑上下 margin */
  border: 1px solid gray;
  margin: 20px 0px 20px 20px;
  display: flex;
  flex-direction: column;
  background-color: rgb(48, 49, 51);
}
.local-id{
  margin-top: 10%;
  margin-left: auto;
  margin-right: auto;
  font-size: 30px;
  font-weight: 5px;
  color: aliceblue;
}
.local-id-container{
  width: 100%;
  margin-top: 10%;
  margin-left: auto;
  margin-right: auto;
  font-size: 30px;
  font-weight: 5px;
  padding: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: aliceblue;
}
.add-friend{
  width: 100%;
  height: 120px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.add-friend-btn{
  font-size: 30px;
  margin: 5px;
  height: calc(100% - 10px);
  width: calc(100% - 10px);
  border-radius: 10px;
  border: 0px;
  background-color: rgb(48, 49, 51);
  color: rgb(123, 123, 124);
}
.logout{
  width: 100%;
  height: 120px;
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
  border-radius: 10px;
  border: 0px;
  background-color: rgb(48, 49, 51);
  color: rgb(123, 123, 124);
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
  margin: 20px 0px 20px 0px;
  position: relative;
  background-color: rgb(175, 179, 182);
  overflow-y: auto;
}
.chat{
  width: calc(63% - 40px); /* 考虑左右 margin */
  height: calc(100% - 40px); /* 考虑上下 margin */
  margin: 20px 0px 20px 0px;
}
</style>