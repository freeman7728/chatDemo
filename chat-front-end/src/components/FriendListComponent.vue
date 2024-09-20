<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-20 08:59:01
 * @LastEditTime: 2024-09-20 17:15:09
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
                {{ friendDetail }}
            </div>
         </div>
    </div>
</template>

<script lang="ts" setup>
    import { onMounted } from 'vue'
    import { useRelationStore } from '@/stores/relation';
    import { computed } from 'vue'
    import router from '@/router';
    const relationStore = useRelationStore()
    const props = defineProps<{
        friendId: number
    }>()
    const friendDetail = computed(() => {
        return relationStore.list[props.friendId] || {}
    })
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
    border: 3px blue solid;
    display: flex;
    margin: 0px 0px 5px 0px;
}
.avatar{
    width: 30%;
    height: 100%;
    border: 3px green solid;
}
.name{
    width: 70%;
    height: 100%;
    border: 3px green solid;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 25px;
}
.avatar-img{
    width: 100%;
    height: 100%;
}
</style>