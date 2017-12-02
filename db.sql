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
    ADD COLUMN updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

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


-- 2017-10-12 添加 from_url 字段，用来证明从哪里添加的
--                               用户 id 是多少
ALTER TABLE cells ADD COLUMN from_url VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN from_id VARCHAR(255) NOT NULL DEFAULT '';


-- 2017-10-15 version check for mobile platform
CREATE TABLE IF NOT EXISTS versions(
    id SERIAL PRIMARY KEY,
    platform VARCHAR(32) NOT NULL DEFAULT '',
    version INTEGER NOT NULL DEFAULT 0,
    published_by VARCHAR(32) NOT NULL DEFAULT '',
    link VARCHAR(255) NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    title VARCHAR(32) NOT NULL DEFAULT ''
);



-- 2017-11-21  添加新的三种类型
INSERT INTO categories(name, src) VALUES('知乎', 31), ('微博', 32), ('豆瓣', 33) ON CONFLICT (src) DO NOTHING;

INSERT INTO users(email, name, pwd, avatar, bio) VALUES('null.douban@athena-anna.com', 'douban', 'pwd', 'null', 'douban');
ALTER TABLE cells ADD COLUMN content TEXT NOT NULL DEFAULT '';

INSERT INTO categories(name, src) VALUES('废弃', 51) ON CONFLICT (src) DO NOTHING;

-- 2017-11-30

ALTER TABLE cells ADD COLUMN md5 VARCHAR(255) NOT NULL DEFAULT '';
-- 用户级别 0-19 超级管理员 20-39 一般管理员 40-59 高级用户 60-79 一般用户 80-99 受限用户
ALTER TABLE users ADD COLUMN role SMALLINT NOT NULL DEFAULT '';

-- 2017-11-21 tags
CREATE TABLE IF NOT EXISTS tags(
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL DEFAULT '',
    desc VARCHAR(255) NOT NULL DEFAULT '',
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- begin
-- tags 和 girls 多对多的关系，用来渐渐替换掉 categories 
-- 暂时并未启用，有大块时间的时候，写个工具，做一套数据迁移
CREATE TABLE NOT EXISTS tags_girls(
    id SERIAL PRIMARY KEY,
    tag_id INTEGER REFERENCES tags(id),
    cell_id INTEGER REFERENCES cells(id)
)

--- end
















# SELECT categories.id, categories.name, categories.src, count(cells.id) AS count FROM categories, cells WHERE categories.id = cells.cate GROUP BY categories.id;





