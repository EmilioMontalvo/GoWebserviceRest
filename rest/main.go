package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

var connString string = "server=localhost;user id=usr_polimusic;password=usr_polimusic;port=1433;database=BDD_PoliMusic"
var db *sql.DB

func main() {

	router := gin.Default()
	router.GET("/Song", getSongs)
	router.GET("/Song/:id", getSongByID)
	router.POST("/Song", insertSong)
	router.PUT("/Song/:id", updateSong)
	router.DELETE("/Song/:id", deleteSong)

	router.Run("localhost:8080")
}

func getSongs(c *gin.Context) {
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las canciones"})
		return
	}
	defer db.Close()

	// Realizar la consulta
	rows, err := db.Query("SELECT [ID_SONG], [SONG_NAME], [SONG_PATH], [PLAYS] FROM [TBL_SONG]")
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las canciones"})
		return
	}
	defer rows.Close()

	// Iterar sobre los resultados y construir una lista de canciones
	var songs []Song
	for rows.Next() {
		var song Song
		err := rows.Scan(&song.Id, &song.Name, &song.Path, &song.Plays)

		if err != nil {
			fmt.Println("Error al escanear fila:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las canciones"})
			return
		}
		songs = append(songs, song)
	}

	c.JSON(http.StatusOK, songs)
}

func getSongByID(c *gin.Context) {
	id := c.Param("id")

	// Abrir la conexión a la base de datos
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar a la base de datos"})
		return
	}
	defer db.Close()

	query := "SELECT ID_SONG, SONG_NAME, SONG_PATH, PLAYS FROM TBL_SONG WHERE ID_SONG = @id"

	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al preparar la consulta"})
		return
	}
	defer stmt.Close()

	var song Song

	// Ejecución de la consulta con el ID como parámetro
	err = stmt.QueryRow(sql.Named("id", id)).Scan(&song.Id, &song.Name, &song.Path, &song.Plays)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la canción"})
		return
	}

	// Devolver la canción encontrada en formato JSON
	c.JSON(http.StatusOK, song)
}

func insertSong(c *gin.Context) {
	// Parsear los datos de la canción desde el cuerpo de la solicitud
	var song Song
	if err := c.BindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de canción inválidos"})
		return
	}

	// Abrir la conexión a la base de datos
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar a la base de datos"})
		return
	}
	defer db.Close()

	// Preparar la consulta para insertar la canción
	query := "INSERT INTO TBL_SONG (SONG_NAME, SONG_PATH, PLAYS) VALUES (@name, @path, @plays)"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al preparar la consulta"})
		return
	}
	defer stmt.Close()

	// Ejecutar la consulta para insertar la canción
	result, err := stmt.Exec(sql.Named("name", song.Name), sql.Named("path", song.Path), sql.Named("plays", song.Plays))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar la canción"})
		return
	}

	// Obtener el número de filas afectadas por la inserción
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el número de filas afectadas"})
		return
	}

	// Devolver la cantidad de filas afectadas en formato JSON
	c.JSON(http.StatusOK, gin.H{"rows_affected": rowsAffected})
}

func updateSong(c *gin.Context) {
	id := c.Param("id")

	// Parsear los datos de la canción desde el cuerpo de la solicitud
	var song Song
	if err := c.BindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de canción inválidos"})
		return
	}

	// Abrir la conexión a la base de datos
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar a la base de datos"})
		return
	}
	defer db.Close()

	// Preparar la consulta para actualizar la canción
	query := "UPDATE TBL_SONG SET SONG_NAME = @name, SONG_PATH = @path, PLAYS = @plays WHERE ID_SONG = @id"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al preparar la consulta"})
		return
	}
	defer stmt.Close()

	// Ejecutar la consulta para actualizar la canción
	result, err := stmt.Exec(sql.Named("name", song.Name), sql.Named("path", song.Path), sql.Named("plays", song.Plays), sql.Named("id", id))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la canción"})
		return
	}

	// Obtener el número de filas afectadas por la actualización
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el número de filas afectadas"})
		return
	}

	// Devolver la cantidad de filas afectadas en formato JSON
	c.JSON(http.StatusOK, gin.H{"rows_affected": rowsAffected})
}

func deleteSong(c *gin.Context) {
	id := c.Param("id")

	// Abrir la conexión a la base de datos
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar a la base de datos"})
		return
	}
	defer db.Close()

	// Preparar la consulta para eliminar la canción por ID
	query := "DELETE FROM TBL_SONG WHERE ID_SONG = @id"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al preparar la consulta"})
		return
	}
	defer stmt.Close()

	// Ejecutar la consulta para eliminar la canción
	result, err := stmt.Exec(sql.Named("id", id))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la canción"})
		return
	}

	// Obtener el número de filas afectadas por la eliminación
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el número de filas afectadas"})
		return
	}

	// Verificar si se eliminó la canción correctamente
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "La canción no fue encontrada"})
		return
	}

	// La canción fue eliminada correctamente
	c.JSON(http.StatusOK, gin.H{"message": "Canción eliminada correctamente"})
}
