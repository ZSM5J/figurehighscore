# Figures

## GET /api/figures
Returns a list of all figures in JSON format.

**http response**

```json
[
    {
        "FigureID": "test",
        "Distance": 12356
    },
    {
        "FigureID": "test2",
        "Distance": 55
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
        "ResID": "11a76f1b78e8dea07002cf8e08821f90",
        "FigureID": "test2",
        "LapTime": 555,
        "Username": "Petr",
        "Token": "a47a27512c35da573f0fb82b106087edcb4e951f3f74fed05e581b6aadc4c318"
    }
]
```

## POST /api/figures/new
Create new figure.

**http request**

```json
{
 "FigureID": "Triangle",
 "Distance": 80
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


