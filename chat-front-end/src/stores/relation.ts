import { defineStore } from 'pinia'
import apiClient from '@/plugins/axios'
import axios from 'axios';

// 定义好友关系项接口
interface RelationItem {
    from: number;
    to: number;
}
// 定义 API 响应数据的结构
interface RelationData {
    list: RelationItem[] | null;
}
interface RelationState{
    status:any;
    msg:string;
    data:any;
    error:string;
    len:number;
    list:RelationItem[];
}
export const useRelationStore = defineStore('relationStore',{
    state:():RelationState =>({
        status:null,
        msg:"",
        data:null,
        error:"",
        len:0,
        list:[]
    }),
    actions:{
        async getList(){
            try{
                const response = await apiClient.get(`/relation/get-relation-list`)
                this.data = response.data
                
                if(this.data.data && this.data.data.list){
                    this.list = this.data.data.list
                }else{
                    this.len = 0
                    this.list = []
                    return
                }
                if(this.list){
                    this.len = this.data.data.list.length
                }
            }catch(error){
                console.log(error)
            }
        }
    }
})