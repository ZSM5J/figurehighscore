package handler


import (
	"net/http"
	"time"
	"google.golang.org/appengine"
	"github.com/figurehighscore/model"
	"google.golang.org/appengine/datastore"
	_"fmt"
	"io/ioutil"
	"github.com/Jeffail/gabs"
	"strconv"
	"github.com/gorilla/mux"
	"strings"
	"github.com/figurehighscore/config"
)



//StatusHandler is used for check server state
var StatusHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var status ServerStatus
	status.Alive = true
	status.Time = time.Now()
	responseJSON(status, response, request)
})

var NewFigureHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)
	body, _ := ioutil.ReadAll(request.Body)
	json, _ := gabs.ParseJSON(body)
	dist, err := strconv.Atoi(json.Path("distance").String())
	if err != nil {
		responseMessage("Distance isn't int.", response, request)
		return
	}

	newFigure := &model.FigureHighScore{
		FigureID: strings.ToLower(ClearString(json.Path("figureID").String())),
		Distance: dist,
	}
	key := datastore.NewIncompleteKey(ctx, "FigureHighScore", nil)

	if _, err := datastore.Put(ctx, key, newFigure); err != nil {
		responseMessage("Can't save to cloud datastore", response, request)
		return
	}

	responseMessage("New Figure is added.", response, request)
})

var GetFigureListHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)
	q := datastore.NewQuery("FigureHighScore").Order("FigureID")
	var figures []model.FigureHighScore
	_, err := q.GetAll(ctx, &figures)
	if err!= nil {
		responseMessage("Can't access cloud datastore", response, request)
		return
	}
	responseJSON(figures, response, request)
})

var GetResultListHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)
	q := datastore.NewQuery("Result").Order("FigureID")
	var results []model.Result
	_, err := q.GetAll(ctx, &results)
	if err!= nil {
		responseMessage("Can't access cloud datastore", response, request)
		return
	}
	responseJSON(results, response, request)
})

var GetLastResultHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)
	q := datastore.NewQuery("Result").Order("-Created").Limit(25)
	var results []model.Result
	_, err := q.GetAll(ctx, &results)
	if err!= nil {
		responseMessage("Can't access cloud datastore", response, request)
		return
	}
	responseJSON(results, response, request)
})

var GetPlayerListHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)
	q := datastore.NewQuery("Player")
	var players []model.Player
	_, err := q.GetAll(ctx, &players)
	if err!= nil {
		http.Error(response, "Can't access to cloud datastore", http.StatusForbidden)
		return
	}
	responseJSON(players, response, request)
})

var GetResultByFigureHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)

	id := strings.ToLower(mux.Vars(request)["id"])
	q := datastore.NewQuery("Result").Filter("FigureID =", id).Filter("Trashed =", false).Limit(50)
	var results []model.Result
	_, err := q.GetAll(ctx, &results)
	if err!= nil {
		http.Error(response, "Can't access to cloud datastore", http.StatusForbidden)
		return
	}

	results = SortResults(results)
	responseJSON(results, response, request)
})

var GetTrashResultByFigureHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)

	id := strings.ToLower(mux.Vars(request)["id"])
	q := datastore.NewQuery("Result").Filter("FigureID =", id).Filter("Trashed =", true)
	var results []model.Result
	_, err := q.GetAll(ctx, &results)
	if err!= nil {
		http.Error(response, "Can't access to cloud datastore", http.StatusForbidden)
		return
	}

	results = SortResults(results)
	responseJSON(results, response, request)
})

var GetMaxResultHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)

	id := strings.ToLower(mux.Vars(request)["id"])
	q := datastore.NewQuery("Result").Filter("FigureID =", id)
	var results []model.Result
	_, err := q.GetAll(ctx, &results)
	if err!= nil {
		http.Error(response, "Can't access to cloud datastore", http.StatusForbidden)
		return
	}

	result := MaxResult(results)
	responseJSON(result, response, request)
})

var GetResultByPlayerHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)

	token := mux.Vars(request)["token"]
	q := datastore.NewQuery("Result").Filter("Token =", token)
	var results []model.Result
	_, err := q.GetAll(ctx, &results)
	if err!= nil {
		http.Error(response, "Can't access to cloud datastore", http.StatusForbidden)
		return
	}
	responseJSON(results, response, request)
})

var DeleteResultHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)

	id := mux.Vars(request)["id"]
	q := datastore.NewQuery("Result").Filter("ResID =", id)
	var results []model.Result
	keys, err := q.GetAll(ctx, &results)
	if err!= nil {
		http.Error(response, "Can't access to cloud datastore", http.StatusForbidden)
		return
	}

	for _, k := range keys {
		datastore.Delete(ctx, k)
	}

	responseMessage("Result is deleted.", response, request)
})

var TrashResultHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)

	id := mux.Vars(request)["id"]
	q := datastore.NewQuery("Result").Filter("ResID =", id)
	var results []model.Result
	keys, err := q.GetAll(ctx, &results)
	if err!= nil || len(results) == 0 {
		http.Error(response, "Can't access to cloud datastore", http.StatusForbidden)
		return
	}

	newResult := &model.Result{
		ResID:   results[0].ResID,
		FigureID: results[0].FigureID,
		LapTime: results[0].LapTime,
		Username: results[0].Username,
		Token: results[0].Token,
		Created: results[0].Created,
		Trashed: true,
	}


	for _, k := range keys {
		if _, err := datastore.Put(ctx, k, newResult); err != nil {
			http.Error(response, "Can't save to cloud datastore", http.StatusForbidden)
			return
		}
	}

	responseMessage("Result is trashed.", response, request)
})

var UndoTrashResultHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)

	id := mux.Vars(request)["id"]
	q := datastore.NewQuery("Result").Filter("ResID =", id)
	var results []model.Result
	keys, err := q.GetAll(ctx, &results)
	if err!= nil || len(results) == 0 {
		http.Error(response, "Can't access to cloud datastore", http.StatusForbidden)
		return
	}

	newResult := &model.Result{
		ResID:   results[0].ResID,
		FigureID: results[0].FigureID,
		LapTime: results[0].LapTime,
		Username: results[0].Username,
		Token: results[0].Token,
		Created: results[0].Created,
		Trashed: false,
	}


	for _, k := range keys {
		if _, err := datastore.Put(ctx, k, newResult); err != nil {
			http.Error(response, "Can't save to cloud datastore", http.StatusForbidden)
			return
		}
	}

	responseMessage("Trash is undo.", response, request)
})

var DeleteFigureHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)
	body, _ := ioutil.ReadAll(request.Body)
	json, _ := gabs.ParseJSON(body)
	username := ClearString(json.Path("username").String())
	password := ClearString(json.Path("password").String())
	if username != config.Config.Admin.User || password != config.Config.Admin.Pass {
		http.Error(response, "no access", http.StatusForbidden)
		return
	}

	id := mux.Vars(request)["id"]
	q := datastore.NewQuery("FigureHighScore").Filter("FigureID =", id)
	var figures []model.FigureHighScore
	keys, err := q.GetAll(ctx, &figures)
	if err!= nil {
		http.Error(response, "Can't access to cloud datastore", http.StatusForbidden)
		return
	}

	for _, k := range keys {
		datastore.Delete(ctx, k)
	}

	responseMessage("Figure is deleted.", response, request)
})

var DeletePlayerHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)

	token := mux.Vars(request)["token"]
	q := datastore.NewQuery("Player").Filter("Token =", token)
	var players []model.Player
	keys, err := q.GetAll(ctx, &players)
	if err!= nil {
		http.Error(response, "Can't access to cloud datastore", http.StatusForbidden)
		return
	}

	for _, k := range keys {
		datastore.Delete(ctx, k)
	}

	responseMessage("Player is deleted.", response, request)
})

var NewResultHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	ctx := appengine.NewContext(request)
	body, _ := ioutil.ReadAll(request.Body)
	json, _ := gabs.ParseJSON(body)
	figureID := ClearString(json.Path("figureID").String())
	username := ClearString(json.Path("username").String())
	lapTime, err := strconv.Atoi(json.Path("lapTime").String())
	if err != nil {
		responseMessage("LapTime isn't int.", response, request)
		return
	}
	token := ClearString(json.Path("token").String())

	//check figure is exist
	figureExist, err := FigureExist(figureID, ctx)
	if err != nil {
		http.Error(response, "Can't access to cloud datastore", http.StatusForbidden)
		return
	}

	if !figureExist {
		http.Error(response, "Figure with this ID doesn't exist", http.StatusForbidden)
		return
	}
	message := "New Result is added."
	//check token is legal or generate new if token is empty
	if token == "" {
		token , err = NewPlayer(ctx)
		if err != nil {
			responseMessage("Can't create new token for user", response, request)
			return
		}
		message = token
	} else {
		exist, err := PlayerExist(token, ctx)
		if err != nil {
			http.Error(response, "Can't access to cloud datastore", http.StatusForbidden)
			return
		}
		if !exist {
			http.Error(response, "Fake token.", http.StatusForbidden)
			return
		}
	}

	newResult := &model.Result{
		ResID:   randResID(),
		FigureID: figureID,
		LapTime: lapTime,
		Username: username,
		Token: token,
		Trashed: false,
		Created: time.Now(),
	}


	key := datastore.NewIncompleteKey(ctx, "Result", nil)

	if _, err := datastore.Put(ctx, key, newResult); err != nil {
		http.Error(response, "Can't save to cloud datastore", http.StatusForbidden)
		return
	}

	responseMessage(message, response, request)
 })


var LoginHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	json, _ := gabs.ParseJSON(body)
	username := ClearString(json.Path("username").String())
	password := ClearString(json.Path("password").String())
	if username == config.Config.Admin.User && password == config.Config.Admin.Pass {
		responseMessage("login", response, request)
		return
	}

	responseMessage("bad login or password", response, request)
})

