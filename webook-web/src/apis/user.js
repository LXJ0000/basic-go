import {httpInstance} from "@/utils/http.js";

export const emailLoginApi = (data) => {
// 代理
    return httpInstance.post("/user/login/", data)
}