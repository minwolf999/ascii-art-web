package errors

import "net/http"

func Check_400_Error(w http.ResponseWriter, r *http.Request) {
	//Fonction erreur 400 // Bad request
	w.WriteHeader(http.StatusBadRequest)
	http.ServeFile(w, r, "templates/400.html") //Templace erreur 400 basic
}
