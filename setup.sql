
create table if not exists Addresses (
    ID               serial primary key,
    CreatedAt        timestamptz not null default (now() at time zone 'UTC'),
    Address          inet not null
);

create or replace procedure InsertAddressIfChanged(p_address inet) as $$
declare
    v_last_address inet;
begin
    select Address into v_last_address from Addresses order by CreatedAt desc limit 1;
    if v_last_address is distinct from p_address then
        insert into Addresses (Address) values (p_address);
    end if;
end;
$$ language plpgsql;
