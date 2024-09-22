<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-20 08:59:01
 * @LastEditTime: 2024-09-22 13:23:23
 * @LastEditors: freeman7728
-->
<template>
    <div class="friend-item-container" @click="goToChat">
        <!-- 头像 -->
        <div class="avatar">
            <img src="/src/assets/avatar.png" alt="" class="avatar-img">
        </div>
        <!-- 姓名 -->
         <div class="name">
            <div>
                ID为{{ friendDetail.to }}的用户
            </div>
         </div>
    </div>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import { useRelationStore } from '@/stores/relation';
import { computed } from 'vue'
import { ref } from 'vue';
import router from '@/router';
const relationStore = useRelationStore()
const props = defineProps<{
    friendId: number
}>()
const friendDetail = computed(() => {
    return relationStore.list[props.friendId] || {}
}) // 保存当前选中的 friendId
const goToChat = () => {
    try{
        router.push({ 
        name: 'ChatPage', 
        query: { 
            id: relationStore.list[props.friendId].from,
            toId:relationStore.list[props.friendId].to
        } 
    });
    }catch(error){
        console.log(error)
    }
};
</script>

<style>
.friend-item-container{
    width: 100%;
    height: 100px;
    display: flex;
    margin: 0 0 20px 0;
    padding:10px;
}
.avatar{
    width: 25%;
    height: 100%;
}
.name{
    width: 70%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 25px;
}
.avatar-img{
    width: 100%;
    height: 100%;
}
.selected {
    background-color: #e0f7fa;
    border-color: #00796b;
}
</style>