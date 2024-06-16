package errors

import "net/http"

func Check_500_Error(w http.ResponseWriter, r *http.Request) {
	//Fonction erreur 500 //Interal server error
	w.WriteHeader(http.StatusInternalServerError)
	http.ServeFile(w, r, "templates/500.html") //Erreur 500 html basic
}
