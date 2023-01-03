# transaction-backend

- [transaction API](#transaction-api)
  - [Initial Setup](#initial-setup)
  - [Run Project](#run-project)
  - [Create Account](#create-account)
  - [Get Account](#get-account)
  - [Create Transaction](#create-transaction)
  - [Get Transaction](#get-transaction)

 
 
## Initial Setup
- Clone the project first

Install migration
- `$ brew install golang-migrate`


## Run Project
Install migration
- `$ make dockerstart`

In a new terminal Run the following to create tables
- `$ make migrateupall`


## Create Account
Endpoint: POST http://localhost:3000/api/v1/accounts

Creates a new account.
### Request
```json
{
    "documentNumber": "1234"
}
```

### Response
On success, returns a 200 status code and the newly created account object:
```json
{
    "id": 1,
    "documentNumber": "1234",
    "createdAt": "2022-12-31T20:33:16.719287+00:00"
}
```

## Get Account
Endpoint: GET http://localhost:3000/api/v1/accounts/:id

Gets an account by its ID.

#### Parameters
- id: the ID of the account



### Response
On success, returns a 200 status code and the newly created account object:
```json
{
    "id": 1,
    "documentNumber": "1234",
    "createdAt": "2022-12-31T20:33:16.719287+00:00"
}
```



## Create Transaction
Endpoint: POST http://localhost:3000/api/v1/transactions

Creates a new transaction.
### Request
```json
{
    "accountID": 1,
    "operationType": 1,
    "amount": "100.0"
}
```

### Response
On success, returns a 200 status code and the newly created account object:
```json
{
    "id": 1,
    "accountID": 1,
    "operationType": 1,
    "amount": "100.0",
    "eventDate": "2023-01-03T01:01:30.017127Z"
}
```

## Get Transaction
Endpoint: GET http://localhost:3000/api/v1/transactions/:id

Gets a transaction by its ID.

#### Parameters
- id: the ID of the transaction



### Response
On success, returns a 200 status code and the newly created account object:
```json
{
    "id": 1,
    "accountID": 1,
    "operationType": 1,
    "amount": "100.0",
    "eventDate": "2023-01-03T01:01:30.017127Z"
}
```
