## API Endpoint : http://localhost:8080


#### Get all accounts

```
  GET localhost:8080/accounts
```

#### Get account

```
  GET localhost:8080/accounts/:id
```

#### create account

```
  POST localhost:8080/accounts
```

| json body | value     | 
| :-------- | :------- |
| `owner` | `user1` |
| `currency` | `USD` |



#### Transfer money

```
  POST localhost:8080/transfers
```

| json body | value     | 
| :-------- | :------- |
| `from_account_id` | `user1` |
| `to_account_id` | `user2` |
| `amount` | `1` |
| `currency` | `USD` |
