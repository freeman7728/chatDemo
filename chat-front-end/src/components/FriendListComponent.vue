<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-20 08:59:01
 * @LastEditTime: 2024-09-20 10:26:28
 * @LastEditors: freeman7728
-->
<template>
    <div class="container">
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
    const relationStore = useRelationStore()
    console.log(relationStore.list)
    const props = defineProps<{
    friendId: number
    }>()
    const friendDetail = computed(() => {
        return relationStore.list[props.friendId].to || {}
    })
    // onMounted(() => {
    //     relationStore.getList()
    // })
 
</script>

<style>
.container{
    width: 100%;
    height: 100px;
    border: 3px blue solid;
    display: flex;
    margin: 0px 0px 20px 0px;
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
}
.avatar-img{
    width: 100%;
    height: 100%;
}
</style>