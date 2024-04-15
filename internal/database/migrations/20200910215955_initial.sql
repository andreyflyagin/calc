-- +goose Up
create table currency
(
    symbol    varchar(50) PRIMARY KEY CHECK (upper(symbol) = symbol),
    is_active boolean default true
);

insert into currency
values ('EUR'),
       ('USD'),
       ('CNY'),
       ('USDT'),
       ('USDC'),
       ('ETH');

-- +goose Down
drop table currency;
