# figurehighscore
api for game



## Models

FigureHighscore {
	Int figureHighscoreId // Internaly for server
	String figureId // Used by client
	[Result] results
	Int distance
}

Result {
	Int resultId
	Int lapTime
	String visualUserName
	->Person person
}

Person {
	Int personId
	String token
}


## Endpoints:

#fetch
/highscore/ [get]
body {
	string figureId
}

result:
body {
	FigureHighscore figureHighscore
}

#send
/hightscore/ [post]
body {
	string figureId
	int lapTime
	string token //optional
}

result:
body {
	string token
}

// if there is no person with this token, or if token is null then create a new token
// if there is no matching figureId -> error message

## Homepage

#Login

#Firstpage:
- List all figureHighscores
- create a new FigureHighscore

#Highscorepage
- list all results (Top 50)
- delete a result
- click on username -> Personpage

#Personpage
- list all results






