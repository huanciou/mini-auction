local result = redis.call("HGETALL", KEYS[1])

return result