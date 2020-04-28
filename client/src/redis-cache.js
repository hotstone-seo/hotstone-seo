import bluebird from 'bluebird'
import redis from 'redis'
import fetch from 'node-fetch'

bluebird.promisifyAll(redis.RedisClient.prototype)

// TODO: handling ReplyError: Ready check failed: NOAUTH Authentication required.
export default class RedisCache {
    constructor(opts) {
        this.redis = redis.createClient(opts)
        this.redis.on("error", function(err) {
            console.error("redis-cache error: ", err)
        });
    }
    async match(req) {
        const res = await this.redis.getAsync(req.url)
        if (res) {
            const parsed = JSON.parse(res)
            return new fetch.Response(parsed.body, {
                url: req.url,
                headers: parsed.headers,
                status: 200
            })
        }

    }
    async put(req, res) {
        const body = await res.text()
        await this.redis.setAsync(req.url, JSON.stringify({
            body: body,
            headers: res.headers.raw()
        }))

        return new fetch.Response(body, {
            url: req.url,
            headers: res.headers,
            status: res.status
        })
    }
    'delete'(req) {
        return this.redis.unlinkAsync(req.url)
    }
}

