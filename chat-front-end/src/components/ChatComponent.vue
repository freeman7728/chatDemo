<!--
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-20 11:23:05
 * @LastEditTime: 2024-09-20 22:16:27
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

            </button>
            </div>
    </div>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { toRef } from 'vue';
import { ref,onBeforeUnmount } from 'vue';
import { onUpdated } from 'vue';
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
    const wsurl = `http://127.0.0.1:3000/ws?uid=${query.value.id}&touid=${query.value.toId}&token=${localStorage.getItem("token")}`
    socket.value = new WebSocket(wsurl);
    socket.value.onopen = () => {
        console.log('WebSocket连接已打开');
    };
    
    socket.value.onmessage = (event) => {
        // if(event.data.data.code != 50004){
            
        // }
        const messageData = JSON.parse(event.data);
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
    border: 2px yellow solid;
    display: flex;
    flex-direction: column;
}
.message-container{
    width: 100%;
    height: 70%;
    /* max-height: 70%; */
    /* height: 535px; */
    border: 2px purple solid;
    display: flex;
    flex-direction: column;
    overflow-y: scroll;
}
.recv-message{
    max-width: 80%;
    min-width: 10%;
    border: 3px blue solid;
    margin-left: 3%;
    margin-top: 3%;
    margin-right: auto;
    word-wrap: break-word; /* 自动换行 */
    word-break: break-all; /* 允许在长单词中间换行 */
    overflow-wrap: break-word; /* 防止长字串溢出 */
    border-radius: 10px;
}
.send-message{
    max-width: 80%;
    min-width: 10%;
    border: 3px greenyellow solid;
    margin-left: auto;
    margin-right: 3%;
    margin-top: 3%;
    word-wrap: break-word; /* 自动换行 */
    word-break: break-all; /* 允许在长单词中间换行 */
    overflow-wrap: break-word; /* 防止长字串溢出 */
    border-radius: 10px;
}
.typing-container{
    width: 100%;
    height: 20%;
    border: 2px purple solid;
    display: flex;
    flex-direction: column;
}
.info-container{
    width: 100%;
    height: 10%;
    border: 2px purple solid;
    display: flex;
    align-items: center;
}
.info{
    margin: 20px;
    font-size: 50px;
}
.typing-area{
    margin: 5px;
    width: calc(100% - 10px);
    height: calc(75% - 10px);
    resize: none;
    font-size: large;
}
.submit-message{
    position: relative;
    width: 90px;
    height: 30px;
    margin: 0 5px 0 0;
    left: calc(100% - 95px);
}
</style>

