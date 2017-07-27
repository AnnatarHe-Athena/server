const redis = require('redis')
const config = require('../config.json')
const redisConfig = config.redis

function setup() {
    return new Promise((resolve, reject) => {
        const client = redis.createClient({
            host: redisConfig.host,
            port: redisConfig.port
        })

        client.on('error', err => {
            console.error("redis error: ", err)
            client.end(true)
            process.exit(-1)
        })

        resolve(client)
    })
}

module.exports = setup