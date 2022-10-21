SELECT count(*) FROM INFORMATION_SCHEMA.COLUMNS WHERE table_schema = 'spire' AND table_name = 'registered_entries' AND column_name = 'store_sv
id';
ALTER TABLE `registered_entries` ADD `store_svid` boolean;

UPDATE `migrations` SET `code_version` = '1.0.4', `updated_at` = '2022-10-14 08:25:11.298163', `version` = 16;
COMMIT

SELECT count(*) FROM INFORMATION_SCHEMA.TABLES WHERE table_schema = 'spire' AND table_name = 'federated_trust_domains'; -- table won;t be present at this point
CREATE TABLE `federated_trust_domains` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`trust_domain
` varchar(255) NOT NULL,`bundle_endpoint_url` varchar(255),`bundle_endpoint_profile` varchar(255),`endpoint_spiffe_id` varchar(255),`implicit` boolean , PRIMARY KEY (`id`)); -- create table

SELECT count(*) FROM INFORMATION_SCHEMA.STATISTICS WHERE table_schema = 'spire' AND table_name = 'federated_trust_domains' AND index_name = 'u
ix_federated_trust_domains_trust_domain'; -- index won't be there as table is newly created
CREATE UNIQUE INDEX uix_federated_trust_domains_trust_domain ON `federated_trust_domains`(trust_domain); -- index has been created

UPDATE `migrations` SET `code_version` = '1.0.4', `updated_at` = '2022-10-14 08:25:11.776918', `version` = 17;
COMMIT

