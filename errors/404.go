package errors

import "net/http"

func Check_404_Error(w http.ResponseWriter, r *http.Request) {
	//Fonction erreur 404 // Page NOT FOUND
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "templates/404.html") //Templace erreur 404 basic
}
