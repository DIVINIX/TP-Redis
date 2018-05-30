package models

import (
    "errors"
    "github.com/mediocregopher/radix.v2/pool"
    "log"
    "strconv"
)

var db *pool.Pool
 
func init() {
    var err error
    db, err = pool.New("tcp", "localhost:6379", 10)
    if err != nil {
        log.Panic(err)
    }
}

var ErrNoNote = errors.New("models: Pas de note")

// Strcture qui correspond a un note
type Note struct {
    Title  string
}

// Retourne le nombre total de notes
func GetNbNotes() (int, error){
    conn, err := db.Get()
    if err != nil {
        return 0,err
    }
    defer db.Put(conn)
 
    nbNotes, err := conn.Cmd("GET", "nbnotes").Int()
 
    if err != nil {
        nbNotes = 0
    }
   
    return nbNotes, nil
}
 
// Méthode pour ajouter une note dans la base redis avec le titre de la note en paramètre
func InsertNote(title string) (string, error) {
    conn, err := db.Get()
    if err != nil {
        return "",err
    }
    defer db.Put(conn)
 
    nbNotes, err := GetNbNotes()
 
    if err != nil {
        nbNotes = 0
    }
 
    nbNotes = nbNotes + 1
 
    err = conn.Cmd("MULTI").Err
    if err != nil {
        return "",err
    }
   
    err = conn.Cmd("HMSET", "notes:"+strconv.Itoa(nbNotes), "title", title).Err
    if err != nil {
        return "",err
    }
 
    err = conn.Cmd("INCR", "nbnotes").Err
    if err != nil {
        return "",err
    }
 
    err = conn.Cmd("EXEC").Err
    if err != nil {
        return "",err
    }
 
    return strconv.Itoa(nbNotes), nil
}
 
// Fonction qui permet de remplir une note
func PopulateNote(reply map[string]string) (*Note, error) {
    n := new(Note)
    n.Title = reply["title"]
    return n, nil
}
 
// Fonctioj qui permet de trouver une note dans la base redis
func FindNote(id string) (*Note, error) {
    reply, err := db.Cmd("HGETALL", "notes:"+id).Map()
    if err != nil {
        return nil, err
    } else if len(reply) == 0 {
        return nil, ErrNoNote
    }
 
    return PopulateNote(reply)
}
 
// Fonction qui permet de supprimer une note dans la base redis
func DeleteNote(id string) (error) {
    conn, err := db.Get()
    if err != nil {
        return err
    }
    defer db.Put(conn)
 
    status, err := conn.Cmd("del", "notes:"+id).Int()
 
    if status == 0 {
        return ErrNoNote
    }
   
    return nil
}