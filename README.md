REST Api

`gin` lib

GET
`/todos`

- res

```
[
    {
        "id": "1",
        "Item": "Clean Room",
        "completed": false
    },
    {
        "id": "2",
        "Item": "Read Book",
        "completed": false
    },
    {
        "id": "3",
        "Item": "Record Video",
        "completed": false
    }
]

```

GET
`/todos/:id` : `id=1`

- res

```
{
    "id": "1",
    "item": "Clean Room",
    "completed": false
}
```

`/todos/:id` : `id=10`

- res

```
{
    "message": "Todo not found"
}
```

POST
`/todos`

- req

```
{
    "id": "1",
    "item": "Make bed",
    "completed": false
}
```

- res

```
{
    "id": "1",
    "item": "Make bed",
    "completed": false
}
```
