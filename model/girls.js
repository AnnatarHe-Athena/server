const {
    graphql,
    GraphQLSchema,
    GraphQLObjectType,
    GraphQLString,
    GraphQLID,
    GraphQLInt,
    GraphQLList
} = require('graphql')

const girlType = new GraphQLObjectType({
    name: "Girl",
    description: "girl",
    fields: {
        id: {
            type: GraphQLID
        },
        img: {
            type: GraphQLString
        },
        text: {
            type: GraphQLString
        },
        cate: {
            type: GraphQLInt
        }
    }
})

const girls = {
    type: new GraphQLList(girlType),
    resolve() {
        return []
    }
}

const girl = {
    type: girlType,
    resolve() {
        return { id: 1, img: 'hello', text: 'hello', cate: 1}
    }
}

exports.girls = girls

exports.girl = girl
