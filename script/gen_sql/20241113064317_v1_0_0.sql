-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS account
(
    id          BIGINT PRIMARY KEY COMMENT '账户ID',
    mobile      VARCHAR(20)  NOT NULL COMMENT '手机号',
    email       VARCHAR(100) NOT NULL COMMENT '电子邮件',
    password    VARCHAR(64)  NOT NULL COMMENT '密码',
    salt        VARCHAR(64)  NOT NULL COMMENT '密码盐',
    number      VARCHAR(15)  NOT NULL DEFAULT '' COMMENT '号码',
    is_deleted  TINYINT      NOT NULL DEFAULT FALSE COMMENT '是否删除',
    create_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX account_mobile_idx (mobile) COMMENT '手机号索引',
    INDEX account_email_idx (email) COMMENT '电子邮件索引',
    INDEX create_time_idx (create_time) COMMENT '创建时间索引',
    INDEX update_time_idx (update_time) COMMENT '更新时间索引'
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user
(
    id               BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID',
    account_id       BIGINT       NOT NULL COMMENT '账户ID',
    mobile           VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '手机号',
    email            VARCHAR(50)  NOT NULL DEFAULT '' COMMENT '电子邮件',
    name             VARCHAR(50)  NOT NULL DEFAULT '' COMMENT '用户名',
    avatar           VARCHAR(255) NOT NULL DEFAULT '' COMMENT '头像',
    background_image VARCHAR(255) NOT NULL DEFAULT '' COMMENT '背景图片',
    signature        VARCHAR(255) NOT NULL DEFAULT '' COMMENT '签名',
    is_deleted       TINYINT      NOT NULL DEFAULT FALSE COMMENT '是否删除',
    create_time      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE INDEX idx_users_account_id (account_id) USING BTREE COMMENT '账户ID唯一索引'
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS template
(
    id          BIGINT PRIMARY KEY COMMENT '模板ID',
    title       VARCHAR(255) NOT NULL COMMENT '标题',
    content     TEXT         NOT NULL COMMENT '内容',
    is_deleted  TINYINT      NOT NULL DEFAULT FALSE COMMENT '是否删除',
    create_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX create_time_idx (create_time) COMMENT '创建时间索引',
    INDEX update_time_idx (update_time) COMMENT '更新时间索引',
    INDEX title_idx (title) COMMENT '标题索引'
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS file
(
    id          BIGINT PRIMARY KEY COMMENT '文件ID',
    domain_name VARCHAR(100) NOT NULL COMMENT '域名',
    biz_name    VARCHAR(100) NOT NULL COMMENT '业务名称',
    hash        VARCHAR(255) NOT NULL COMMENT '文件哈希',
    file_size   BIGINT       NOT NULL DEFAULT 0 COMMENT '文件大小',
    file_type   VARCHAR(255) NOT NULL COMMENT '文件类型',
    uploaded    TINYINT      NOT NULL DEFAULT FALSE COMMENT '是否上传',
    is_deleted  TINYINT      NOT NULL DEFAULT FALSE COMMENT '是否删除',
    create_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX create_time_idx (create_time) COMMENT '创建时间索引',
    INDEX update_time_idx (update_time) COMMENT '更新时间索引',
    INDEX hash_idx (hash) COMMENT '文件哈希索引'
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS video
(
    id            BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '视频ID',
    user_id       BIGINT       NOT NULL COMMENT '用户ID',
    title         VARCHAR(50)  NOT NULL DEFAULT '' COMMENT '标题',
    description   VARCHAR(255) NOT NULL DEFAULT '' COMMENT '描述',
    video_url     VARCHAR(255) NOT NULL DEFAULT '' COMMENT '视频URL',
    cover_url     VARCHAR(255) NOT NULL DEFAULT '' COMMENT '封面URL',
    like_count    BIGINT       NOT NULL DEFAULT 0 COMMENT '点赞数',
    collect_count BIGINT       NOT NULL DEFAULT 0 COMMENT '收藏数',
    forward_count BIGINT       NOT NULL DEFAULT 0 COMMENT '转发数',
    comment_count BIGINT       NOT NULL DEFAULT 0 COMMENT '评论数',
    is_deleted    TINYINT      NOT NULL DEFAULT FALSE COMMENT '是否删除',
    create_time   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS comment
(
    id             BIGINT PRIMARY KEY COMMENT '评论ID',
    video_id       BIGINT       NOT NULL COMMENT '视频ID',
    user_id        BIGINT       NOT NULL COMMENT '发表评论的用户ID',
    parent_id      BIGINT                DEFAULT NULL COMMENT '父评论ID',
    to_user_id     BIGINT                DEFAULT NULL COMMENT '评论所回复的用户ID',
    content        VARCHAR(512) NOT NULL COMMENT '评论内容',
    first_comments JSON         NOT NULL COMMENT '最开始的几条子评论',
    is_deleted     TINYINT      NOT NULL DEFAULT FALSE COMMENT '是否删除',
    create_time    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX video_id_idx (video_id, is_deleted) COMMENT '视频ID索引',
    INDEX user_id_idx (user_id, is_deleted) COMMENT '用户ID索引',
    INDEX create_time_idx (create_time) COMMENT '创建时间索引',
    INDEX update_time_idx (update_time) COMMENT '更新时间索引'
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS collection
(
    id          BIGINT PRIMARY KEY COMMENT '收藏夹ID',
    user_id     BIGINT       NOT NULL COMMENT '用户ID',
    title       VARCHAR(255) NOT NULL COMMENT '标题',
    description TEXT         NOT NULL COMMENT '描述',
    is_deleted  TINYINT      NOT NULL DEFAULT FALSE COMMENT '是否删除',
    create_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX user_id_idx (user_id, is_deleted) COMMENT '用户ID索引',
    INDEX create_time_idx (create_time) COMMENT '创建时间索引',
    INDEX update_time_idx (update_time) COMMENT '更新时间索引'
) COMMENT '收藏夹';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS follow
(
    id             BIGINT PRIMARY KEY COMMENT '关注ID',
    user_id        BIGINT    NOT NULL COMMENT '用户ID',
    target_user_id BIGINT    NOT NULL COMMENT '被关注的用户ID',
    is_deleted     TINYINT   NOT NULL DEFAULT FALSE COMMENT '是否删除',
    create_time    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX user_id_idx (
                       user_id,
                       target_user_id,
                       is_deleted
        ) COMMENT '用户ID索引'
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS interaction
(
    id          BIGINT PRIMARY KEY,
    biz_id      BIGINT             DEFAULT NULL,
    biz         varchar(255)       DEFAULT NULL,
    read_cnt    BIGINT    NOT NULL DEFAULT 0,
    like_cnt    BIGINT    NOT NULL DEFAULT 0,
    collect_cnt BIGINT    NOT NULL DEFAULT 0,
    is_deleted  TINYINT   NOT NULL DEFAULT FALSE COMMENT '是否删除',
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY idx_interaction_bizID_biz (biz_id, biz)
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_collect
(
    id            BIGINT unsigned NOT NULL AUTO_INCREMENT,
    user_id       BIGINT                   DEFAULT NULL,
    biz_id        BIGINT                   DEFAULT NULL,
    biz           VARCHAR(255)             DEFAULT NULL,
    collection_id BIGINT                   DEFAULT NULL,
    status        TINYINT(1)      NOT NULL DEFAULT 0 COMMENT '0=不收藏 1=收藏',
    is_deleted    TINYINT         NOT NULL DEFAULT FALSE COMMENT '是否删除',
    create_time   TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time   TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY idx_userCollect_userID_bizID_biz (user_id, biz_id, biz),
    KEY idx_user_collect_collection_id (collection_id)
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_like
(
    id          BIGINT unsigned NOT NULL AUTO_INCREMENT,
    user_id     BIGINT                   DEFAULT NULL,
    biz_id      BIGINT                   DEFAULT NULL,
    biz         VARCHAR(255)             DEFAULT NULL,
    status      TINYINT(1)      NOT NULL DEFAULT 0 COMMENT '0=不点赞 1=点赞',
    is_deleted  TINYINT         NOT NULL DEFAULT FALSE COMMENT '是否删除',
    create_time TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY idx_userLike_userID_bizID_biz (user_id, biz_id, biz)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS user;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS template;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS file;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS video;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS comment;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS collection;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS follow;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE `interaction`;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE `user_collect`;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE `user_like`;
-- +goose StatementEnd