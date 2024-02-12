package main

import "encoding/xml"

type Song struct {
	XMLName xml.Name `xml:"Song"`
	ID      int      `xml:"ID"`
	Name    string   `xml:"Name"`
	Path    string   `xml:"Path"`
	Plays   int      `xml:"Plays"`
}

// Estructura de solicitud para encontrar todas las canciones
type GetSongsRequest struct{}

// Estructura de respuesta para encontrar todas las canciones
type GetSongsResponse struct {
	XMLName xml.Name `xml:"getSongsResponse"`
	Songs   []Song   `xml:"songs"`
}

// Estructura de solicitud para encontrar una canción por ID
type GetSongByIdRequest struct {
	ID int `xml:"id"`
}

// Estructura de respuesta para encontrar una canción por ID
type GetSongByIdResponse struct {
	XMLName xml.Name `xml:"getSongByIdResponse"`
	Song    Song     `xml:"song"`
}

// Estructura de solicitud para guardar una canción
type PostSongRequest struct {
	Song Song `xml:"song"`
}

// Estructura de respuesta para guardar una canción
type PostSongResponse struct {
	XMLName xml.Name `xml:"postSongResponse"`
	Song    Song     `xml:"song"`
}

// Estructura de solicitud para eliminar una canción
type DeleteSongRequest struct {
	ID int `xml:"id"`
}

// Estructura de respuesta para eliminar una canción
type DeleteSongResponse struct {
	XMLName xml.Name `xml:"deleteSongResponse"`
	Message string   `xml:"message"`
}

// Estructura de solicitud para actualizar una canción
type UpdateSongRequest struct {
	Song Song `xml:"song"`
}

// Estructura de respuesta para actualizar una canción
type UpdateSongResponse struct {
	XMLName xml.Name `xml:"updateSongResponse"`
	Message string   `xml:"message"`
}
