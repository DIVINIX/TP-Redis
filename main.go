package main
 
import (
    "github.com/mediocregopher/radix.v2/pool"
    "github.com/julienschmidt/httprouter"
    "log"
    "strconv"
    "net/http"
    "fmt"
    "tp1/models"
    "io/ioutil"
)

// Routage
func main() {
    router := httprouter.New()
    router.POST("/notes", addNote)
    router.GET("/delete/:id", deleteNote)
    router.GET("/notes", getAllNotes)
    router.GET("/notes/:id", showNote)
 
    log.Fatal(http.ListenAndServe(":12345", router))
}
 
// Fonction pour ajouter une note
func addNote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
 
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        http.Error(w, http.StatusText(405), 405)
        return
    }
   
    title, HTTPerr := ioutil.ReadAll(r.Body);
   
    if HTTPerr != nil {
        fmt.Fprintf(w, "%s", HTTPerr)
    }
 
    returnValue, err := models.InsertNote(string(title[:]))
    if err != nil {
        log.Fatal(err)
        return
    }
   
     fmt.Fprintf(w, "ID de la nouvelle note :"+returnValue+"\n")
}

// Fonction pour afficher une note via son ID
func showNote(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
 
    if r.Method != "GET" {
        w.Header().Set("Allow", "GET")
        http.Error(w, http.StatusText(405), 405)
        return
    }
 
    id := p.ByName("id")
    if id == "" {
        http.Error(w, http.StatusText(400), 400)
        return
    }
 
    if _, err := strconv.Atoi(id); err != nil {
        http.Error(w, http.StatusText(400), 400)
         return
    }
 
    note, err := models.FindNote(id)
    if err == models.ErrNoNote {
        http.NotFound(w, r)
    fmt.Fprintf(w, "Pas de note avec cet ID \n")
        return
    } else if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
 
    fmt.Fprintf(w, "Note n°"+id+" %s"+"\n", note.Title)
}
 
//Fonction pour afficher toute les notes
func getAllNotes(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
 
    //Lister les notes
    if r.Method == "GET" {
 
        nbNotes, err := models.GetNbNotes()
    if err != nil {
                log.Fatal(err)
            }
 
        //Pour chaque note entree
        for i:=1;i <= nbNotes;i++ {
        idNote := strconv.Itoa(i)
        note, err := models.FindNote(idNote)
            if err == nil {
                fmt.Fprint(w,"Note n°"+strconv.Itoa(i)+" "+note.Title)
                fmt.Fprint(w,"\n")
            }
        }
    }
}
 
// Fonction pour supprimer une note via son ID
func deleteNote(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    if r.Method != "GET" {
        w.Header().Set("Allow", "GET")
        http.Error(w, http.StatusText(405), 405)
        return
    }
 
    id := p.ByName("id")
    if id == "" {
        http.Error(w, http.StatusText(400), 400)
        return
    }
 
    if _, err := strconv.Atoi(id); err != nil {
        http.Error(w, http.StatusText(400), 400)
         return
    }
 
    err := models.DeleteNote(id)
    if err == models.ErrNoNote {
        http.NotFound(w, r)
    fmt.Fprintf(w, "Pas de note avec cet ID \n")
        return
    } else if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
 
    fmt.Fprintf(w, "Note n°"+id+" supprimée \n", )
}