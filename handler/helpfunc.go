package handler

import (
	"crypto/rand"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"github.com/figurehighscore/model"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"sort"
)

type ServerStatus struct {
	Alive bool
	Time  time.Time
}

type Message struct {
	Message string
}

func responseJSON(data interface{}, response http.ResponseWriter, request *http.Request) {
	js, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(js)
}

func responseMessage(message string, response http.ResponseWriter, request *http.Request) {
	data := &Message{
		Message: message,
	}
	js, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(js)
}

func ClearString(str string) string {
	str = str[1 : len(str)-1]
	return str
}

func randResID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func NewPlayer(ctx context.Context) (string, error) {
	token := randToken()
	person := &model.Player{
		Token: token,
		Registered: time.Now(),
	}

	key := datastore.NewIncompleteKey(ctx, "Player", nil)

	if _, err := datastore.Put(ctx, key, person); err != nil {
		return "", err
	}

	return token, nil
}

func PlayerExist(token string, ctx context.Context) (bool, error) {
	q := datastore.NewQuery("Player").Filter("Token =", token)
	var player []model.Player
	_, err := q.GetAll(ctx, &player)
	if err!= nil {
		return false, err
	}

	 if len(player) > 0 {
		 return true, nil
	 } else {
		 return false, nil
	 }
}

func FigureExist(id string, ctx context.Context) (bool, error) {
	q := datastore.NewQuery("FigureHighScore").Filter("FigureID =", id)
	var figure []model.FigureHighScore
	_, err := q.GetAll(ctx, &figure)
	if err!= nil {
		return false, err
	}

	if len(figure) > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func SortResults(results []model.Result) []model.Result {
	sort.Slice(results, func(i, j int) bool { return results[i].LapTime > results[j].LapTime })

	return results
}

func MaxResult(results []model.Result) model.Result {
	max := 0
	for i, res := range results {
		if res.LapTime > results[max].LapTime {
			max = i
		}
	}
	return results[max]
}