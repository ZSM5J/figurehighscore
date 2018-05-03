# Figures

## GET /api/figures
Returns a list of all figures in JSON format.

**http response**

```json
[
    {
        "figureID": "square",
        "distance": 58
    },
    {
        "figureID": "test2",
        "distance": 55
    }
]
```

## GET /api/figures/{id}
Returns all results in this figure(id=figureID)  in JSON format.

**http request**

```
https://figurehighscore.appspot.com/api/figures/test2
```

**http response**

```json
[
    {
        "resID": "11a76f1b78e8dea07002cf8e08821f90",
        "figureID": "test2",
        "lapTime": 555,
        "username": "Petr",
        "token": "a47a27512c35da573f0fb82b106087edcb4e951f3f74fed05e581b6aadc4c318"
    }
]
```

## POST /api/figures/new
Create new figure.

**http request**

```json
{
 "figureID": "triangle",
 "distance": 80
}

```

**http response**

```json
{
    "Message": "New Figure is added."
}
```

## DELETE /api/figures/{id}
Delete choosen figure.

**http request**

```
https://figurehighscore.appspot.com/api/figures/Triangle
```

**http response**

```json
{
    "Message": "Figure is deleted."
}
```


# Results

## GET /api/results
Returns a list of all results in JSON format.

**http response**

```json
[
    {
        "resID": "24ae24e962b84bbfb591a3de2d111b76",
        "figureID": "square",
        "lapTime": 33,
        "username": "Kevin",
        "token": "a9225ab22963febdec2eca46d3187426f342ad3b93d78fff284b34e82cd75d2b"
    },
    {
        "resID": "5f71c8bf0469fec046656f947325be8a",
        "figureID": "test2",
        "lapTime": 11,
        "username": "Tony",
        "token": "b68fc96c7f626ba463941f4b552b8122c9f2a162f04540f5ef4f767243c25287"
    },
]
```

## POST /api/results/new
Create new result. If token is empty server will generate new token and be sended in responce.

**http request**

```json
{
 "figureID": "Triangle",
 "username": "John",
 "lapTime": 80,
 "token": ""
}

```

**http response**

```json
{
    "Message": "1d9fa62427cdd943d0beea02b1301a35671eccdf9a1657988771e63e5edfb900"
}
```

## DELETE /api/results/{id}
Delete choosen result.

**http request**

```
https://figurehighscore.appspot.com/api/results/24ae24e962b84bbfb591a3de2d111b76
```

**http response**

```json
{
    "Message": "Result is deleted."
}
```

# PLayers

## GET /api/players
Returns a list of all players in JSON format.

**http response**

```json
[
    {
        "token": "a9225ab22963febdec2eca46d3187426f342ad3b93d78fff284b34e82cd75d2b",
        "registred": "2018-05-03T15:54:50.33823Z"
    },
    {
        "token": "78f249f77e74802008f453cdcce5308b9f28297108bc100aee068d95f4854403",
        "registred": "2018-05-02T23:37:57.272601Z"
    }       
]
```

## GET /api/players/{id}
Returns this player's results in JSON format.

**http request**

```
https://figurehighscore.appspot.com/api/players/78f249f77e74802008f453cdcce5308b9f28297108bc100aee068d95f4854403
```

**http response**

```json
[
    {
        "resID": "24ae24e962b84bbfb591a3de2d111b76",
        "figureID": "square",
        "lapTime": 33,
        "username": "Kevin",
        "token": "a9225ab22963febdec2eca46d3187426f342ad3b93d78fff284b34e82cd75d2b"
    },
    {
        "resID": "f44b88ca41ba704780f2c7386344db0d",
        "figureID": "square",
        "lapTime": 222,
        "username": "Kevin",
        "token": "a9225ab22963febdec2eca46d3187426f342ad3b93d78fff284b34e82cd75d2b"
    }      
]
```

## DELETE /api/players/{id}
Delete choosen player.

**http request**

```
https://figurehighscore.appspot.com/api/players/78f249f77e74802008f453cdcce5308b9f28297108bc100aee068d95f4854403
```

**http response**

```json
{
    "Message": "Player is deleted."
}
```






