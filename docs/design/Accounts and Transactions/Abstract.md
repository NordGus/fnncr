# Abstract

## Designing financo's core: Accounts and Transactions.

The way Accounts and Transactions are designed is like a graph, where Accounts
are the graph's nodes, and the Transactions are its edges.

<img src="Account-Abstract.drawio.svg" alt="diagram">

## Accounts
Accounts are containers that can represent a capital store or wallet in the
system, debts held or incurred, available credit or any source of income or
expense.

| field       | type      | additional                           |
|-------------|-----------|--------------------------------------|
| id          | uuid      | primary key                          |
| parent_id   | uuid      | index                                |
| kind        | text      | index, not nullable                  |
| currency    | text      | index, not nullable                  |
| name        | text      | not nullable                         |
| description | text      |                                      |
| color       | text      | not nullable                         |
| icon        | text      | not nullable                         |
| limit       | integer   | not nullable, default `0`            |
| is_archived | boolean   | index, not nullable, default `false` |
| created_at  | timestamp | not nullable                         |
| updated_at  | timestamp | not nullable                         |
| deleted_at  | timestamp | index                                |

Each account can have a parent account vía the `parent_id`.
This is a design decision made, so the user can define child accounts,
so they can better classify their finances inside the system's limitations.

Each account must have a `kind`, that will be treated as an enum to indicate the
type of account stored in the system.
The values should be defined as `{family}.{type}`.
Where family would be one of the following `system`, `capital`, `debt`,
`external`.
As for the types, each family would have their own:
- `system`
  - `history`
- `capital`
  - `normal`
  - `savings`
- `debt`
  - `loan`
  - `credit`
- `external`
  - `income`
  - `expense`

This should be enforced on a database level.

Each account must have a `currency`, that will be treated as an enum, or at
least be normalized to a defined list of values to indicate the `currency`
stored or handled by the account stored in the system.
This would indicate the system if the user should indicate the amount received
by the target account to store as the transaction's exchange rate and maintain
system coherency.

## Transactions
Transactions are time series-like records that connect money movements between
[Accounts](#accounts).

| field       | type      | additional                                   |
|-------------|-----------|----------------------------------------------|
| id          | uuid      | primary key                                  |
| from_id     | uuid      | foreign key to accounts, index, not nullable |
| to_id       | uuid      | foreign key to accounts, index, not nullable |
| from_amount | integer   | not nullable                                 |
| to_amount   | integer   | not nullable                                 |
| issued_at   | date      | index, not nullable                          |
| executed_at | date      | index                                        |
| created_at  | timestamp | not nullable                                 |
| updated_at  | timestamp | not nullable                                 |
| deleted_at  | timestamp | index                                        |
