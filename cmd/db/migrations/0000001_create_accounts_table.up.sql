CREATE EXTENSION "uuid-ossp";

CREATE TABLE public.accounts (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    parent_id uuid,
    kind VARCHAR NOT NULL,
    currency VARCHAR NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    color VARCHAR NOT NULL,
    icon VARCHAR NOT NULL,
    credit BIGINT NOT NULL DEFAULT 0,
    is_archived BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX index_on_accounts_account_parent ON public.accounts (parent_id);
CREATE INDEX index_on_accounts_account_kind ON public.accounts (kind);
CREATE INDEX index_on_accounts_account_currency ON public.accounts (currency);
CREATE INDEX index_on_accounts_account_archived ON public.accounts (is_archived);
CREATE INDEX index_on_accounts_account_deleted ON public.accounts (deleted_at);