local bid = tonumber(ARGV[1])
local bidder = ARGV[2]
local key = 'product:1'

local current_price = tonumber(redis.call('HGET', key, 'Price'))

if bid > current_price then
    redis.call('HSET', key, 'Price', bid)
    redis.call('HSET', key, 'Bidder', bidder)
    return bid
else
    return current_price
end
