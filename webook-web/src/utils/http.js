import axios from "axios";
import {useStore} from "@/stores/store";

export const httpInstance = axios.create({
    baseURL:"http://localhost:8080",
    timeout:5000,
    headers:{
        "Content-Type": "application/json"
    }
})

httpInstance.interceptors.request.use(request => {
// 一般用于添加用户token
    const store = useStore()
    const token = store.UserInfo.token
    if (token) {
        request.headers.Authorization = `Bearer ${token}`
    }
    return request
}) // todo e => Promise.reject(e)

httpInstance.interceptors.response.use(response => {
    return response.data
})