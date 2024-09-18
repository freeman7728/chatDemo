<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-18 12:44:14
 * @LastEditTime: 2024-09-18 13:20:37
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
const store = useUserStore()
console.log(store.token)
</script>
<style>
</style>