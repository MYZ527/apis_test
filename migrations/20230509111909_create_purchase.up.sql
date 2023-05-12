CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE purchases (
    Purchase_ID uuid default uuid_generate_v4() not null primary KEY,
    Applicant VARCHAR not null,
    Company_name VARCHAR not null,
    Requisition_department VARCHAR not null,
    Product VARCHAR not null,
    Requisition_quantity int not null,
    Price int not null,
    Position text not null,
    Is_Deleted bool default false not null,
    Created_At timestamp default now() not null,
    Created_By VARCHAR not null,
    Updated_At timestamp null ,
    Updated_By VARCHAR null
);

create unique index purchases_purchase_id_uindex
    on purchases (Purchase_ID);

create or replace function position_code()
returns text as
$$
declare
old_position text := (select position from purchases order by position desc limit 1);
    id_number char(4) := '0001';
    new_position text;
    num integer;

begin
    if old_position is null then
        new_position := 'P'||id_number;
return new_position;

end if;

        num := cast(right(old_position, 4) as integer) + 1;
        id_number :=
        case
            when num < 10 then '000'||num
            when num < 100 then '00'||num
            when num < 1000 then '0'||num
            when num < 10000 then cast(num as text)
end;

    new_position := 'P'||id_number;
return new_position;
end;
$$
language plpgsql;