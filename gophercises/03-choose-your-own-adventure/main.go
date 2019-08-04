package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/nguyenthanh1290/gophercises/03-choose-your-own-adventure/story"
)

var templates = template.Must(template.ParseFiles("index.html", "story.html"))
var stories map[string]story.Story
var storyNames []string

func main() {
	var err error
	stories, err = loadStories("./story/")
	if err != nil {
		log.Fatal(err)
	}

	storyNames, err = listStories(stories)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/story", storyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadStories(dirname string) (map[string]story.Story, error) {
	stories := make(map[string]story.Story)

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".json") {
			st, _ := loadStory(dirname + f.Name())
			stories[f.Name()] = st
		}
	}

	return stories, nil
}

func loadStory(name string) (story.Story, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	st, err := story.New(data)
	if err != nil {
		return nil, err
	}

	return st, nil
}

func listStories(stories map[string]story.Story) ([]string, error) {
	if len(stories) < 1 {
		return nil, errors.New("Empty stories")
	}

	names := make([]string, len(stories))
	i := 0
	for k := range stories {
		names[i] = strings.TrimRight(k, ".json")
		i++
	}

	return names, nil
}

// rootHandler handles all requests to website root /
// where the user can choose a story to begin with
func rootHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", storyNames)
}

// storyHandler handles requests matching /story?name=[story name]&arc=[next arc]
// Ex: /story?name=gopher&arc=new-york
func storyHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	name := query.Get("name")
	if name == "" {
		http.Redirect(w, r, "/", http.StatusNotFound)
	}
	name += ".json"

	arc := query.Get("arc")
	if arc == "" {
		arc = story.BeginArc
	}

	templates.ExecuteTemplate(w, "story.html", stories[name][arc])
}
