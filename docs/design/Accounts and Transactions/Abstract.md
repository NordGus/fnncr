# Abstract

## Designing financo's core: Accounts and Transactions.

The way Accounts and Transactions are design is like a graph, where Accounts are
the graph's nodes and the Transactions are its edges.

<img src="Account-Abstract.drawio.svg" alt="diagram">

## Accounts
Are containers that can represent a capital store or wallet in the system, debts
held or incurred, available credit or any source of income or expense.

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

## Transactions
Are time series-like records that connect money movements between
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

