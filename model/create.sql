create table if not exists orders
(
    request_id varchar(256) primary key,
    user varchar(42) not null ,
    basic_token_in varchar(66) not null default '0',
    target_token_in varchar(66) not null default '0',
    swap_type int not null,
    result_rate int,
    swap_tx varchar(66),

    pre_basic_token_out varchar(66) not null default '0',
    pre_target_token_out varchar(66) not null default '0',

    final_basic_token_out varchar(66) not null default '0',
    final_target_token_out varchar(66) not null default '0',
    settle_tx varchar(66),

    block bigint,
    index bigint,
    create_time time
);