use wgx_foundation;

alter table user add user_id varchar(32);
alter table user add did varchar(100);

alter table space add cos_folder varchar(100);
alter table space add display_order int default 1;

create table if not exists `contract` (
    `id` varchar(36) not null primary key,
    `contract_address` varchar(64) not null,
    `space_id` binary(16) not null,
    `blockchain_id` varchar(64) not null,
    `contract_type` varchar(255) not null,
    `asset_type` varchar(255) not null,
    `is_shown_in_wallet` boolean not null default false,
    `created_at` timestamp(3) not null default current_timestamp(3),
    `updated_at` timestamp(3) null on update current_timestamp(3),
    unique index `idx_contract_address_blockchain_id` using hash(contract_address, blockchain_id),
    index `idx_space_id` using hash(space_id)
);