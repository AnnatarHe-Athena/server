const Koa = require('koa')
const Router = require('koa-router')
const logger = require('koa-logger')
const graphqlHTTP = require('koa-graphql')
const init = require('./initial/index')
const schema = require('./model/index')

const app = new Koa()
const router = new Router()

app.use(logger())
app.use(init)

router.get('/hello', ctx => {
    ctx.body = "world"
})

router.all('/graphql', async ctx => {
    // 异步代码，必须等着，不然就直接走 404 了
    await graphqlHTTP({
        schema,
        graphiql: true
    })(ctx)
    .catch(e => {
        console.error(e)
    })
})

app.use(router.routes()).use(router.allowedMethods())

app.on('error', e => {
    console.error(e)
})

app.listen(9000, () => {
    console.log('app is runing')
})