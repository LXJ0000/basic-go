# Webook Server
## 1. Gin Session 存储的实现
单机单实例部署，`memstore`；多实例部署，`redis`，确保不同实例都能访问到 `session`
## 2. 刷新登陆状态
- 固定间隔时间刷新
- 长短token
## 3. 参数如何设定
性能测试