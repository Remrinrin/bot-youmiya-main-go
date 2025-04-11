-- 确保 pg_partman 扩展已安装
CREATE EXTENSION IF NOT EXISTS pg_partman;

-- OneBot 消息主表
CREATE TABLE onebot_messages (
    -- 基础信息
    id BIGSERIAL,                              -- 数据库主键
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- 记录创建时间
    
    -- 消息序列相关
    message_seq BIGINT,                         -- 消息序列号
    real_seq BIGINT,                           -- 真实序列号
    message_id BIGINT,                          -- 消息 ID
    real_id BIGINT,                            -- 真实 ID
    
    -- 事件类型相关
    post_type VARCHAR(32) NOT NULL,             -- 上报类型
    event_type VARCHAR(32),                     -- 事件类型
    sub_type VARCHAR(32),                       -- 子类型
    time BIGINT NOT NULL,                       -- 事件发生的时间戳
    
    -- 消息内容相关
    message_type VARCHAR(32),                   -- 消息类型（私聊/群聊）
    message_format VARCHAR(32),                 -- 消息格式
    message TEXT,                               -- 消息内容
    raw_message TEXT,                           -- 原始消息内容
    font INTEGER,                               -- 字体
    
    -- 参与者信息
    self_id BIGINT NOT NULL,                    -- 机器人 QQ 号
    user_id BIGINT NOT NULL,                    -- 发送者 QQ 号
    group_id BIGINT,                            -- 群号（群消息）
    sender JSONB,                               -- 发送者信息
    
    -- 扩展信息
    extra JSONB,                                -- 额外信息
    
    -- 主键和约束
    PRIMARY KEY (id, created_at),
    CONSTRAINT idx_message_unique UNIQUE (message_id, self_id, created_at)
) PARTITION BY RANGE (created_at);

-- 创建注释
COMMENT ON TABLE onebot_messages IS 'OneBot 消息记录表';

-- 创建 2025 年第一季度分区
CREATE TABLE onebot_messages_y2025q1 PARTITION OF onebot_messages
    FOR VALUES FROM ('2025-01-01') TO ('2025-04-01');

-- 创建当前季度分区（2025年第二季度）
CREATE TABLE onebot_messages_y2025q2 PARTITION OF onebot_messages
    FOR VALUES FROM ('2025-04-01') TO ('2025-07-01');

-- 创建 2025 年第三季度分区
CREATE TABLE onebot_messages_y2025q3 PARTITION OF onebot_messages
    FOR VALUES FROM ('2025-07-01') TO ('2025-10-01');

-- 创建 2025 年第四季度分区
CREATE TABLE onebot_messages_y2025q4 PARTITION OF onebot_messages
    FOR VALUES FROM ('2025-10-01') TO ('2026-01-01');

-- 创建 2026 年第一季度分区
CREATE TABLE onebot_messages_y2026q1 PARTITION OF onebot_messages
    FOR VALUES FROM ('2026-01-01') TO ('2026-04-01');

-- 创建 2026 年第二季度分区
CREATE TABLE onebot_messages_y2026q2 PARTITION OF onebot_messages
    FOR VALUES FROM ('2026-04-01') TO ('2026-07-01');

-- 创建 2026 年第三季度分区
CREATE TABLE onebot_messages_y2026q3 PARTITION OF onebot_messages
    FOR VALUES FROM ('2026-07-01') TO ('2026-10-01');

-- 创建 2026 年第四季度分区
CREATE TABLE onebot_messages_y2026q4 PARTITION OF onebot_messages
    FOR VALUES FROM ('2026-10-01') TO ('2027-01-01');

-- 创建 2027 年第一季度分区
CREATE TABLE onebot_messages_y2027q1 PARTITION OF onebot_messages
    FOR VALUES FROM ('2027-01-01') TO ('2027-04-01');

-- 创建 2027 年第二季度分区
CREATE TABLE onebot_messages_y2027q2 PARTITION OF onebot_messages
    FOR VALUES FROM ('2027-04-01') TO ('2027-07-01');