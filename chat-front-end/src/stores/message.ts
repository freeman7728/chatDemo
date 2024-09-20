/*
 * @Description: 
 * @author: freeman7728
 * @Date: 2024-09-20 21:31:29
 * @LastEditTime: 2024-09-20 21:36:26
 * @LastEditors: freeman7728
 */
import { defineStore } from 'pinia'

interface MessageState{
    sender:string,
    content:string,
    sendTime:string,
}

export const useMessageStore = defineStore('messageStore',{
    state:():MessageState =>({
        sender:"",
        content:"",
        sendTime:"",
    }),
    actions :{
    }
})