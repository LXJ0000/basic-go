# Webook Server
## 1. Gin Session 存储的实现
单机单实例部署，`memstore`；多实例部署，`redis`，确保不同实例都能访问到 `session`
## 2. 刷新登陆状态
- 固定间隔时间刷新
- 长短token
## 3. 参数如何设定
性能测试
## 4. JWT JSON WEB TOKEN
Header Payload Signature 三部分组成
## 5. JWT 优缺点
优点：不依赖第三方存储、适合在分布式环境下使用、性能好（没有 Redis 访问之类）
缺点：对加密依赖较大，比 Session 任意泄露、不要再 JWT 中存放敏感信息
## 6. 保护系统
1. 正常用户会不会搞崩你的系统？
2. 如果有人攻击你的系统，如何解决？
### 限流
1. 如何标识对象？限流对象可以用 IP，APP 端可以考虑使用设备序列号
2. 限流阈值？限制的阈值不是很小，就可以解决用一个 IP 多个用户的问题

- 为限流添加对应的监控和警告
- 对不需要登陆就可以访问的接口限流
- 对核心业务接口限流

为 Gin 插件库添加限流插件
- 单机限流
  - 令牌桶
  - 漏桶
  - 滑动窗口
  - 固定窗口
- 基于 Redis 限流
- 基于 Redis IP 限流
## 7. Kubernetes 容器编排平台
- Pod 实例
- Service 服务
- Deployment 管理
### 安装
[参考文档](https://www.qikqiak.com/post/deploy-k8s-on-win-use-wsl2/)
[参考视频](https://www.bilibili.com/video/BV1Ru41137s2/?spm_id_from=333.1007.top_right_bar_window_history.content.click&vd_source=2cb41caee9551fbf13c606149026e31c)
```bash
kubectl apply -f k8s-webook-deployment.yaml
kubectl apply -f k8s-webook-service.yaml
kubectl get deployments
kubectl get pods
kubectl get services
```