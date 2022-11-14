CREATE TABLE IF NOT EXISTS `abi`
(
    `id`   binary(12)   NOT NULL,
    `name` varchar(255) NOT NULL COMMENT 'abi名字',
    `data` text         NOT NULL COMMENT 'abi内容',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `abi_event`
(
    `id`     binary(12) NOT NULL,
    `hash`   binary(32) NOT NULL COMMENT '事件签名',
    `data`   text       NOT NULL COMMENT '事件内容',
    `abi_id` binary(12) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_hash` (`hash`),
    KEY `index_abi_id` (`abi_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `abi_function`
(
    `id`     binary(12) NOT NULL,
    `hash`   binary(32) NOT NULL COMMENT '函数签名',
    `data`   text       NOT NULL COMMENT '函数内容',
    `abi_id` binary(12) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_hash` (`hash`),
    KEY `index_abi_id` (`abi_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `block`
(
    `id`              binary(12) NOT NULL,
    `number`          bigint     NOT NULL COMMENT '区块高度',
    `hash`            binary(32) NOT NULL COMMENT '区块hash',
    `status`          int        NOT NULL COMMENT '区块状态：正常、回滚',
    `time`            datetime   NOT NULL COMMENT '出块时间',
    `miner`           binary(20) NOT NULL COMMENT '矿工',
    `chain_id`        int        NOT NULL COMMENT '不同 EVM 链的一个标识',
    `transaction_num` int        NOT NULL COMMENT '交易数量',
    `event_num`       int        NOT NULL COMMENT '事件数量',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_hash` (`hash`),
    KEY `index_number` (`hash`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `event`
(
    `id`               binary(12) NOT NULL,
    `type`             int        NOT NULL COMMENT '事件类型：自生成事件、链上事件',
    `time`             datetime   NOT NULL COMMENT 'event时间',
    `topic0`           binary(32) DEFAULT NULL COMMENT 'topic[0]',
    `topic1`           binary(32) DEFAULT NULL COMMENT 'topic[1]',
    `topic2`           binary(32) DEFAULT NULL COMMENT 'topic[2]',
    `topic3`           binary(32) DEFAULT NULL COMMENT 'topic[3]',
    `data`             blob       NOT NULL COMMENT '事件data',
    `block_id`         binary(12) NOT NULL COMMENT '区块id',
    `block_hash`       binary(32) NOT NULL COMMENT '区块hash',
    `block_num`        bigint     NOT NULL COMMENT '区块高度',
    `chain_id`         int        NOT NULL COMMENT '不同 EVM 链的一个标识',
    `transaction_id`   binary(12) NOT NULL COMMENT '交易id',
    `transaction_hash` binary(32) NOT NULL COMMENT '交易hash',
    `transaction_from` binary(20) NOT NULL COMMENT '发起地址',
    `transaction_to`   binary(20) NOT NULL COMMENT '合约地址/接收人地址',
    `abi_event_id`     binary(12) DEFAULT NULL COMMENT 'abi事件id',
    PRIMARY KEY (`id`),
    KEY `index_block_id` (`block_id`),
    KEY `index_block_hash` (`block_hash`),
    KEY `index_transaction_id` (`transaction_id`),
    KEY `index_transaction_hash` (`transaction_hash`),
    KEY `index_abi_event_id` (`abi_event_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `transaction`
(
    `id`         binary(12)   NOT NULL,
    `hash`       binary(32)   NOT NULL COMMENT '交易hash',
    `to`         binary(20)   NOT NULL COMMENT '发起地址',
    `from`       binary(20)   NOT NULL COMMENT '接收地址',
    `gas`        varchar(255) NOT NULL COMMENT 'gas数量',
    `gas_price`  varchar(255) NOT NULL COMMENT 'gas价格',
    `time`       datetime     NOT NULL COMMENT '交易时间',
    `status`     int          NOT NULL COMMENT '交易状态：成功、失败、回滚',
    `type`       int          NOT NULL COMMENT '交易类型：伦敦协议交易、普通交易',
    `nonce`      bigint       NOT NULL COMMENT '交易nonce',
    `value`      varchar(255) NOT NULL COMMENT '以太数量',
    `data`       blob         NOT NULL COMMENT 'data 用data[:4] 匹配abi函数',
    `block_id`   binary(12)   NOT NULL COMMENT '区块id',
    `block_hash` binary(32)   NOT NULL COMMENT '区块hash',
    `block_num`  bigint       NOT NULL COMMENT '区块高度',
    `chain_id`   int          NOT NULL COMMENT '不同 EVM 链的一个标识',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_hash_block_hash` (`hash`, `block_hash`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;
