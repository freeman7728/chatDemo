<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-21 13:49:40
 * @LastEditTime: 2024-09-22 15:05:17
 * @LastEditors: freeman7728
-->
<template>
    <div class="add-container">
        <input type="text" class="input-friendId" v-model="friendId" placeholder="请输入对方ID">
        <button class="submit-friendId" @click="submitFriendId">添加好友</button>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const friendId = ref('')
import apiClient from '@/plugins/axios';
import iziToast from 'izitoast';
import { useRelationStore } from '@/stores/relation';
const relationStore = useRelationStore()
const submitFriendId = async () => {
    try {
        // 调用异步函数来发送 friendId
        await sendFriendId(friendId.value); // friendId 是你要发送的数据
    } catch (error) {
        console.error("Failed to send friend ID:", error);
    }
}
const sendFriendId = async (id: string) => {
    try {
        // 模拟发送请求，比如使用 fetch 或 axios 发送请求
        const Nid: number = Number(id);
        const response = await apiClient.post('/relation/create', {
            to: Nid
        });
        const data = response.data
        console.log(data)
        if(data.msg == "对方不存在"){
            iziToast.warning({
                transitionIn:'fadeInDown',
                position:'topCenter',
                title: '对方不存在',
                message: '请检查id',
            });
        }
        else if(data.msg == "不能重复添加好友"){
            iziToast.warning({
                transitionIn:'fadeInDown',
                position:'topCenter',
                title: '不能重复添加好友',
                message: '好友已存在',
            });
        }
        else if(data.msg == "关系建立成功"){
            iziToast.success({
                transitionIn:'fadeInDown',
                position:'topCenter',
                title: '关系建立成功',
                message: '关系建立成功',
            });
            relationStore.getList()
        }
        else if(data.msg == "不能添加自己"){
            iziToast.warning({
                transitionIn:'fadeInDown',
                position:'topCenter',
                title: '不能添加自己',
                message: '不能添加自己',
            });
            relationStore.getList()
        }
        else{
            iziToast.warning({
                transitionIn:'fadeInDown',
                position:'topCenter',
                title: '未知错误',
                message: '未知错误',
            });
        }
    } catch (error) {
        iziToast.warning({
            transitionIn:'fadeInDown',
            position:'topCenter',
            title: '未知错误',
            message: '未知错误',
        });
    }
}
</script>

<style>
.add-container{
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
}
.input-friendId{
    width: 300px;
    height: 50px;
    margin: 20px;
    font-size: 30px;
    padding: 10px;
}
.submit-friendId{
    font-size: 20px;
    width: 100px;
    height: 50px;
    border-radius: 10px;
}
</style>

