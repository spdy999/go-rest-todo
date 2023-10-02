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
    "ID": "1",
    "Item": "Make bed",
    "Completed": false
}
```
