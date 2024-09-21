<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-21 15:53:21
 * @LastEditTime: 2024-09-21 16:11:27
 * @LastEditors: freeman7728
-->
<template>
    <div class="add-container">
        <input type="text" class="input-friendId" v-model="friendId">
        <button class="submit-friendId" @click="submitFriendId">删除好友</button>
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
        const response = await apiClient.post('/relation/del', {
            to: Nid
        });
        const data = response.data
        console.log(data)
        if(data.msg == "检查好友是否存在"){
            iziToast.warning({
                transitionIn:'fadeInDown',
                position:'topCenter',
                title: '检查好友是否存在',
                message: '请检查id',
            });
        }
        else if(data.msg == "关系删除成功"){
            iziToast.success({
                transitionIn:'fadeInDown',
                position:'topCenter',
                title: '关系删除成功',
                message: '关系删除成功',
            });
            relationStore.getList()
        }
        else if(data.msg == "不能删除自己"){
            iziToast.warning({
                transitionIn:'fadeInDown',
                position:'topCenter',
                title: '不能删除自己',
                message: '不能删除自己',
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
    border: lawngreen 3px solid;
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
}
.submit-friendId{
    font-size: 20px;
    width: 100px;
    height: 50px;
    border: 3px red solid;
    border-radius: 10px;
}
</style>