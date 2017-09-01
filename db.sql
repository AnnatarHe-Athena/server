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