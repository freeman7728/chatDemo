<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-20 11:23:05
 * @LastEditTime: 2024-09-23 14:20:30
 * @LastEditors: freeman7728
-->
<template>
    <div class="chat-container">
        <!-- 对方名字 -->
        <div class="info-container">
            <div class="info">
                {{ query.toId }}
            </div>
        </div>
        <!-- 聊天区域 -->
        <div class="message-container" ref="containerRef">
            <div
                v-for="(message, index) in messages"
                :key="index"
                :class="message.sender === 1 ? 'send-message' : 'recv-message'"
            >
                {{ message.content }}
            </div>
        </div>
        <!-- 文本区域 -->
            <div class="typing-container">
            <textarea id = "typing" type="text" class="typing-area" v-model="msg">
            </textarea>
            <button class="submit-message" @click="submitMessageHandler">
                发送
            </button>
            </div>
    </div>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { toRef } from 'vue';
import { ref,onBeforeUnmount } from 'vue';
import { onUpdated } from 'vue';
import {bu} from '@/plugins/axios';
const route = useRoute()
let query = toRef(route,'query')
const msg = ref('')
const messages = ref<{ type: number, content: string,sender:number }[]>([]);
const containerRef = ref()
onUpdated(() => {
  containerRef.value.scrollTop = containerRef.value.scrollHeight
})
const submitMessageHandler = ()=>{
    if (socket.value && socket.value.readyState === WebSocket.OPEN) {
        const message = {
            type: 1,
            content: msg.value,
            sender:1
        };
        messages.value.push(message);
        socket.value.send(JSON.stringify(message));
        console.log('发送消息:', msg.value);
    } else {
        console.error('WebSocket 未连接');
    }
}
const socket = ref<WebSocket | null>(null);
    const connectWebSocket = () => {
    const wsurl = `${bu}/ws?uid=${query.value.id}&touid=${query.value.toId}&token=${localStorage.getItem("token")}`
    socket.value = new WebSocket(wsurl);
    socket.value.onopen = () => {
        console.log('WebSocket连接已打开');
    };
    
    socket.value.onmessage = (event) => {
        const messageData = JSON.parse(event.data)
        if (messageData.code != 50004 && messageData.content) {
            const message = {
                type: 1, // Received message type
                content: messageData.content,
                sender:0
            };

            // Add the message to the messages list
            messages.value.push(message);
            console.log('接收到消息:', messageData.content);
        } 
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
onBeforeUnmount(() => {
    if (socket.value) {
        socket.value.close();
    }
});
</script>

<style>
.chat-container{
    width: 100%;
    height: 100%;
    display: flex;
    display: flex;
    flex-direction: column;
}
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: rgb(239, 239, 239);
  border-radius: 2px;
}

::-webkit-scrollbar-thumb {
  background: #bfbfbf;
  border-radius: 10px;
}

::-webkit-scrollbar-thumb:hover {
  background: #bfbfbf;
}

::-webkit-scrollbar-corner {
  background: #bfbfbf;
}
.message-container{
    border-top: 1px rgb(161, 160, 160) solid;
    border-left: 1px rgb(161, 160, 160) solid;
    border-right: 1px rgb(161, 160, 160) solid;
    width: 100%;
    height: 70%;
    display: flex;
    flex-direction: column;
    overflow-y: scroll;
}
.recv-message{
    max-width: 80%;
    min-width: 10%;
    border: 3px white solid;
    margin-left: 3%;
    margin-top: 3%;
    margin-right: auto;
    word-wrap: break-word; /* 自动换行 */
    word-break: break-all; /* 允许在长单词中间换行 */
    overflow-wrap: break-word; /* 防止长字串溢出 */
    border-radius: 10px;
    background-color: white;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 3px;
}
.send-message{
    max-width: 80%;
    min-width: 10%;
    border: 3px rgb(134, 233, 53) solid;
    background-color: rgb(134, 233, 53);
    margin-left: auto;
    margin-right: 3%;
    margin-top: 3%;
    word-wrap: break-word; /* 自动换行 */
    word-break: break-all; /* 允许在长单词中间换行 */
    overflow-wrap: break-word; /* 防止长字串溢出 */
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 3px;
}
.typing-container{
    width: 100%;
    height: 20%;
    display: flex;
    flex-direction: column;
    position: relative;
}
.info-container{
    border-top: 1px rgb(161, 160, 160) solid;
    border-left: 1px rgb(161, 160, 160) solid;
    border-right: 1px rgb(161, 160, 160) solid;
    width: 100%;
    height: 10%;
    display: flex;
    align-items: center;
}
.info{
    margin: 20px;
    font-size: 50px;
}
.typing-area{
    border-top: 1px rgb(161, 160, 160) solid;
    border-left: 1px rgb(161, 160, 160) solid;
    border-right: 1px rgb(161, 160, 160) solid;
    border-bottom: 1px rgb(161, 160, 160) solid;
    width: 100%;
    height: 100%;
    resize: none;
    font-size: large;
    padding: 10px;
    font-size: 20px;
}
.submit-message{
    border: 0px;
    position: absolute;
    width: 90px;
    height: 30px;
    margin: 0 5px 0 0;
    left: calc(100% - 100px);
    color: rgb(35, 196, 21);
    font-weight: 20px;
    font-size: 20px;
    top: calc(100% - 40px);
    border-radius: 5px;
}
button:hover{
    background-color: rgb(133, 131, 131);
    color: black;
}
</style>

