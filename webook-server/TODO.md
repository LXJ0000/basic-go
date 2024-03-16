1. 本地缓存替换 Redis 
   - 定义 CodeCache 接口，目前只有 CodeRedisCache
   - 提供基于本地缓存实现的 CodeLocalCache
   - 保证**单机**并发安全
2. 注册、登陆后缓存用户信息