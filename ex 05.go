package main

import (
	"fmt"
	"time"
)

type Musica struct {
	Titulo   string
	Artista  string
	Duracao  time.Duration
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func playlistsComTitulo(titulo string, playlists []Playlist) []Playlist {
	var playlistsComTitulo []Playlist

	for _, playlist := range playlists {
		for _, musica := range playlist.Musicas {
			if musica.Titulo == titulo {
				playlistsComTitulo = append(playlistsComTitulo, playlist)
				break
			}
		}
	}

	return playlistsComTitulo
}

func main() {
	playlist1 := Playlist{
		Nome: "Minha Playlist 1",
		Musicas: []Musica{
			{Titulo: "Música 1", Artista: "Artista 1", Duracao: 3 * time.Minute},
			{Titulo: "Música 2", Artista: "Artista 2", Duracao: 4 * time.Minute + 30 * time.Second},
		},
	}

	playlist2 := Playlist{
		Nome: "Minha Playlist 2",
		Musicas: []Musica{
			{Titulo: "Música 2", Artista: "Artista 2", Duracao: 4 * time.Minute + 30 * time.Second},
			{Titulo: "Música 3", Artista: "Artista 3", Duracao: 2 * time.Minute + 45 * time.Second},
		},
	}

	playlist3 := Playlist{
		Nome: "Minha Playlist 3",
		Musicas: []Musica{
			{Titulo: "Música 3", Artista: "Artista 3", Duracao: 2 * time.Minute + 45 * time.Second},
			{Titulo: "Música 4", Artista: "Artista 4", Duracao: 5 * time.Minute},
		},
	}

	tituloDesejado := "Música 2"

	playlists := []Playlist{playlist1, playlist2, playlist3}
	playlistsEncontradas := playlistsComTitulo(tituloDesejado, playlists)

	fmt.Printf("Playlists com a música '%s':\n", tituloDesejado)
	for _, playlist := range playlistsEncontradas {
		fmt.Printf("- %s\n", playlist.Nome)
	}
}
