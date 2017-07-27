const postgresSetup = require('./db')
const redisSetup = require('./redis')

async function initial(app, next) {
    let db, redis
    try {
        [db, redis] = await Promise.all([postgresSetup(), redisSetup()])
    }catch(e) {
        console.error(e)
        process.exit(-1)
    }
    app.db = db
    app.redis = redis
    await next()
    db.release()
    redis.quit()

}

module.exports = initial