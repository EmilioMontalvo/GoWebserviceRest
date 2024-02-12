package main

import (
	"database/sql"
	"encoding/xml"
	"log"
	"net/http"
)

var connString string = "server=localhost;user id=usr_polimusic;password=usr_polimusic;port=1433;database=BDD_PoliMusic"
var db *sql.DB

func GetSongsFromDB() ([]Song, error) {
	db, err := sql.Open("mssql", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT ID_SONG, SONG_NAME, SONG_PATH, PLAYS FROM TBL_SONG")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []Song

	for rows.Next() {
		var song Song
		err := rows.Scan(&song.ID, &song.Name, &song.Path, &song.Plays)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

// Función para manejar la solicitud de encontrar todas las canciones
func GetSongsHandler(w http.ResponseWriter, r *http.Request) {
	songs, err := GetSongsFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := GetSongsResponse{
		Songs: songs,
	}

	w.Header().Set("Content-Type", "application/xml")
	if err := xml.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Manejadores de solicitud para las operaciones definidas en el esquema XSD
	http.HandleFunc("/getSongs", GetSongsHandler)
	// Agrega más manejadores según sea necesario

	log.Fatal(http.ListenAndServe(":8080", nil))
}
