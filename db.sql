CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    src INTEGER UNIQUE NOT NULL
);
INSERT INTO categories(name, src) VALUES('所有', 0), ('有颜值', 4), ('美腿控', 3), ('黑丝袜', 7), ('小翘臀', 6), ('大胸妹', 2), ('大杂烩', 5) ON CONFLICT (src) DO NOTHING;
CREATE TABLE IF NOT EXISTS cells (
    id SERIAL PRIMARY KEY,
    img VARCHAR(255) UNIQUE NOT NULL,
    text VARCHAR(255) NOT NULL,
    cate INTEGER REFERENCES categories (id)
);
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(48) UNIQUE NOT NULL,
    name VARCHAR(36) UNIQUE NOT NULL,
    pwd VARCHAR(255) NOT NULL,
    avatar VARCHAR(255) NOT NULL DEFAULT '',
    bio VARCHAR(255) NOT NULL DEFAULT ''
);
-- a user has many collection
CREATE TABLE IF NOT EXISTS collections (
    id SERIAL PRIMARY KEY,
    cell INTEGER REFERENCES cells(id),
    owner INTEGER REFERENCES users(id)
);

ALTER TABLE cells ADD COLUMN createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN createdBy integer REFERENCES users (id),
    ADD COLUMN updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

ALTER TABLE categories ADD COLUMN createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE users ADD COLUMN createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

INSERT INTO users(email, name, pwd, avatar, bio) VALUES('null.dbmeinv@athena-anna.com', 'dbmeinv', 'pwd', 'null', 'dbmeinv'),
('null.weibo@athena-anna.com', 'weibo', 'pwd', 'null', 'weibo'),
('null.wechat@athena-anna.com', 'wechat', 'pwd', 'null', 'wechat'),
('null.gank@athena-anna.com', 'gank', 'pwd', 'null', 'gank');


-- 2017-10-06 添加权限控制字段
-- 2  -> public 共有的，谁都可以看
-- 3  -> 受保护的，只有发布者自己能看
-- 4+ -> 暂未定义
ALTER TABLE cells ADD COLUMN premission SMALLINT NOT NULL DEFAULT 2,
    ADD COLUMN likes BIGINT NOT NULL DEFAULT 0;
INSERT INTO users(email, name, pwd, avatar, bio) VALUES('null.zhihu@athena-anna.com', 'zhihu', 'pwd', 'null', 'zhihu');
-- 123456 测试数据
INSERT INTO users(email, name, pwd, avatar, bio) VALUES('i@annatarhe.com', 'AnnatarHe', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 'null', 'make the world a better place');





















SELECT categories.id, categories.name, categories.src, count(cells.id) AS count FROM categories, cells WHERE categories.id = cells.cate GROUP BY categories.id;





