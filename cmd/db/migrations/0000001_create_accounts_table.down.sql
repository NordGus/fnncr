DROP INDEX index_on_accounts_account_parent;
DROP INDEX index_on_accounts_account_kind;
DROP INDEX index_on_accounts_account_currency;
DROP INDEX index_on_accounts_account_archived;
DROP INDEX index_on_accounts_account_deleted;

DROP TABLE public.accounts;

DROP EXTENSION "uuid-ossp";
