const { Pool } = require('pg')
const config = require('../config.json')
const dbConfig = config.db

const pool = new Pool({
    user: dbConfig.user,
    host: dbConfig.host,
    database: dbConfig.database,
    password: dbConfig.password,
    port: dbConfig.port
})

pool.on('error', (err, client) => {
    console.error('postgresql error: ', err)
    process.exit(-1)
})

function setup() {
    return new Promise((resolve, reject) => {
        pool.connect((err, client) => {
            if (err) {
                reject(err)
            }
            resolve(client)
        })
    }).catch(e => {
        console.error("error", e)
    })
}

module.exports = setup
