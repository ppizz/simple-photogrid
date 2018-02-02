/* ==================================================================
GALLERY.GO 
ppizz 2017 V0.1 photogrid
Serveur WEB demo:  simple photogrid ppizz
--Description:
Web server en GO avec Gorilla/mux router
Les données relatives aux photos sont stockées dans une base SQLite
Les accès à la base sont encapsulés dans un package GO: catalog
la fonction Init ouvre la base demo "jpg.db" qui se trouve dans "./static/PHOTO/"
la fonction GetPhotoDB return les enregistrements la table pĥoto
les données sont retournées en reponse a la requete POST JSON 
--Prerequis:
Linux: Ubuntu, Debian...
IE>=11, Firefox>=25, Chrome>=30
Golang >= 1.6 + package gorilla/mux
--Developpement:
installer le package catalog avec: "go install github.com/ppizz/catalog"
compilation avec "go build gallery.go", lancer le serveur web avec "./gallery"
=====================================================================*/

package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	Cat "github.com/ppizz/catalog"
)

// CONST
const VERSION = "Photogrid V0.1"

// TYPE
type typRequestDir struct {
	NomDir string `json:"NomDir"`
}

// Glob VAR
var _DirPhoto = "./static/PHOTO/"
var _DirName = "SD1"
var _NbPhoto int

func main() {
	log.Println(VERSION)
	fmt.Println("Dossierphoto: ", _DirPhoto + _DirName)
	
	// initialisation catalogue
	log.Println("Ouverture catalogue photo")
	log.Println(Cat.GetVersion())
	Cat.Init(_DirPhoto)

	// Mise en place des routes sur le port 8080
	router := mux.NewRouter()
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    router.HandleFunc("/", HandlerIndex)
	router.HandleFunc("/getphoto", HandlerGetPhoto).Methods("POST")

	log.Println("Ouvrir Simple-Photogrid en HTTP port 8080")
	http.ListenAndServe(":8080", router)
}

// -------------------------------------------------------------
// Handler pour index.html et ses fichiers static
func  HandlerIndex(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./static/index.html")
}

// -------------------------------------------------------------
// Handler pour l'affichage des vignettes photo
// Renvoie en JSOn le nom et les parametres jpeg des fichiers du Dir
// le client (navigateur) demande les photos du sous-dossier: r.NomDir 
// r.NomDir n'est pas utilisé dans cette demo => sous dossier par defaut "SD1"

func HandlerGetPhoto(res http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var err error
	var r typRequestDir
	err = decoder.Decode(&r)
	if err != nil {
		panic(err)
	}
	res.Header().Set("Content-Type", "application/json")
	// _DirName = r.NomDir
	_NbPhoto = Cat.GetPhotoDB(_DirName)
    if _NbPhoto > 0 {
		outgoingJSON, error := json.Marshal(Cat.Tab[0:_NbPhoto])
	    if error != nil {
		     log.Println(error.Error())
		      http.Error(res, error.Error(), http.StatusInternalServerError)
		     return
		}
		fmt.Fprint(res, string(outgoingJSON))	
	}
}

