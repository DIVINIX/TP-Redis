package models
curl -i http://localhost:12345/..
curl -i -L -d "note" http://localhost:12345/..
import (
    "errors"
    // Import the Radix.v2 pool package, NOT the redis package.
    "github.com/mediocregopher/radix.v2/pool"
    "log"
    "strconv"
)
 
// Declare a global db variable to store the Redis connection pool.
var db *pool.Pool
 
func init() {
    var err error
    // Establish a pool of 10 connections to the Redis server listening on
    // port 6379 of the local machine.
    db, err = pool.New("tcp", "localhost:6379", 10)
    if err != nil {
        log.Panic(err)
    }
}
 
// Create a new error message and store it as a constant. We'll use this
// error later if the FindNote() function fails to find an Note with a
// specific id.
var ErrNoNote = errors.New("models: no note found")
 
type Note struct {
    Title  string
}
 
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
 
func PopulateNote(reply map[string]string) (*Note, error) {
    n := new(Note)
    n.Title = reply["title"]
    return n, nil
}
 
func FindNote(id string) (*Note, error) {
    reply, err := db.Cmd("HGETALL", "notes:"+id).Map()
    if err != nil {
        return nil, err
    } else if len(reply) == 0 {
        return nil, ErrNoNote
    }
 
    return PopulateNote(reply)
}
 
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