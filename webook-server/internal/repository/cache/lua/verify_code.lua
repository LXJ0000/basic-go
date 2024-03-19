local key = KEYS[1]
local expectedCode = ARGV[1]
local cntKey = key..":cnt"

local code = redis.call("get", key)
local cnt = tonumber(redis.call("get", cntKey))

if cnt == nil or cnt <= 0 then
    return -1 -- 用户输入错误，没有机会
end
if expectedCode == code then
    redis.call("del", cntKey)
    return 0 -- 用户输入正确
else
    redis.call("decr", cntKey)
    return -2 -- 用户输入错误，还有机会
end