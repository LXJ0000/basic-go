1. 本地缓存替换 Redis 
   - 定义 CodeCache 接口，目前只有 CodeRedisCache
   - 提供基于本地缓存实现的 CodeLocalCache
   - 保证**单机**并发安全
2. 注册、登陆后缓存用户信息
3. 长短 Token 、退出登录 [ok]
4. 微信登陆
5. 高可用短信服务平台
6. 布隆过滤器
7. 日志模块
8. 配置模块
9. 装饰器