import { defineStore } from "pinia";
import { message } from "ant-design-vue";
import router from "@/router";

export const useStore = defineStore("webook", {
  state: () => {
    return {
      UserInfo: {
        exp: 0,
        user_name: "",
        role: "",
        user_id: "",
        token: "",
      },
    };
  },
  actions: {
    setUserInfo(info) {
      this.$patch({
        UserInfo: info,
      });
      localStorage.setItem("userinfo", JSON.stringify(info));
    },
    loadUserInfo() {
      let userinfoRaw = localStorage.getItem("userinfo");
      if (userinfoRaw == null) {
        return;
      }
      let userinfo = JSON.parse(userinfoRaw);
      //     判断是否过期
      let exp = userinfo.exp;
      let now = new Date().getTime();
      if (exp <= now) {
        message.warn("登录信息已过期");
        router.push({ name: "login" });
      }
      this.setUserInfo(userinfo);
    },
  },
});
