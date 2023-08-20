CREATE TABLE `follows` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `block_number` bigint(20) unsigned DEFAULT NULL,
  `block_hash` varchar(191) DEFAULT NULL,
  `sub_type` varchar(191) DEFAULT NULL,
  `burn` bigint(20) unsigned DEFAULT NULL,
  `user` varchar(191) DEFAULT NULL,
  `followee` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_follows_burn` (`burn`),
  KEY `idx_follows_user` (`user`),
  KEY `idx_follows_deleted_at` (`deleted_at`),
  KEY `idx_follows_block_number` (`block_number`),
  KEY `idx_follows_block_hash` (`block_hash`),
  KEY `idx_follows_sub_type` (`sub_type`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

CREATE TABLE `likes` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `block_number` bigint(20) unsigned DEFAULT NULL,
  `block_hash` varchar(191) DEFAULT NULL,
  `sub_type` varchar(191) DEFAULT NULL,
  `burn` bigint(20) unsigned DEFAULT NULL,
  `tweet` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_likes_tweet` (`tweet`),
  KEY `idx_likes_deleted_at` (`deleted_at`),
  KEY `idx_likes_block_number` (`block_number`),
  KEY `idx_likes_block_hash` (`block_hash`),
  KEY `idx_likes_sub_type` (`sub_type`),
  KEY `idx_likes_burn` (`burn`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

CREATE TABLE `pub_keys` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `block_number` bigint(20) unsigned DEFAULT NULL,
  `block_hash` varchar(191) DEFAULT NULL,
  `user` varchar(191) DEFAULT NULL,
  `pub_key` longblob,
  PRIMARY KEY (`id`),
  KEY `idx_pub_keys_deleted_at` (`deleted_at`),
  KEY `idx_pub_keys_block_number` (`block_number`),
  KEY `idx_pub_keys_block_hash` (`block_hash`),
  KEY `idx_pub_keys_user` (`user`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

CREATE TABLE `send_messages` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `block_number` bigint(20) unsigned DEFAULT NULL,
  `block_hash` varchar(191) DEFAULT NULL,
  `sub_type` varchar(191) DEFAULT NULL,
  `burn` bigint(20) unsigned DEFAULT NULL,
  `to` varchar(191) DEFAULT NULL,
  `message` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_send_messages_block_hash` (`block_hash`),
  KEY `idx_send_messages_sub_type` (`sub_type`),
  KEY `idx_send_messages_burn` (`burn`),
  KEY `idx_send_messages_to` (`to`),
  KEY `idx_send_messages_deleted_at` (`deleted_at`),
  KEY `idx_send_messages_block_number` (`block_number`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

CREATE TABLE `tweets` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `block_number` bigint(20) unsigned DEFAULT NULL,
  `block_hash` varchar(191) DEFAULT NULL,
  `sub_type` varchar(191) DEFAULT NULL,
  `burn` bigint(20) unsigned DEFAULT NULL,
  `tweet` longtext,
  `retweet_of` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_tweets_sub_type` (`sub_type`),
  KEY `idx_tweets_burn` (`burn`),
  KEY `idx_tweets_deleted_at` (`deleted_at`),
  KEY `idx_tweets_block_number` (`block_number`),
  KEY `idx_tweets_block_hash` (`block_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1
