//...
import "github.com/flaviocopes/gitometer/api/handlers"
//...
http.HandleFunc("/api/index", handlers.Index)
http.HandleFunc("/api/repo/", handlers.Repo)
//...