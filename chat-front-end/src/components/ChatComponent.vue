<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-20 11:23:05
 * @LastEditTime: 2024-09-20 14:14:43
 * @LastEditors: freeman7728
-->
<template>
    <div class="chat-container">
        <RouterView>
            {{ query.toId }}
        </RouterView>
    </div>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { toRef } from 'vue';
const route = useRoute()
let query = toRef(route,'query')
</script>

<style>
.chat-container{
    width: 100%;
    height: 100%;
    display: flex;
    border: 2px yellow solid;
}
</style>

