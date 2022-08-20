create table accounts (
  id uuid primary key,
  account_number varchar unique,
  balance float default 0,
  created_at timestamp default current_timestamp
);

create table transactions (
  id uuid primary key,
  from_account uuid references accounts(account_number),
  to_account uuid references accounts(account_number),
  amount float default 0
);
