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

func imprimirPlaylistInfo(p Playlist) {
	fmt.Printf("Playlist: %s\n", p.Nome)

	tempoTotal := calcularTempoTotal(p)
	fmt.Printf("Tempo total da playlist: %s\n", tempoTotal)

	fmt.Println("Músicas:")
	for _, musica := range p.Musicas {
		fmt.Printf("- Título: %s, Artista: %s, Duração: %s\n", musica.Titulo, musica.Artista, musica.Duracao)
	}
}

func calcularTempoTotal(p Playlist) time.Duration {
	var tempoTotal time.Duration

	for _, musica := range p.Musicas {
		tempoTotal += musica.Duracao
	}

	return tempoTotal
}

func main() {
	playlistExemplo := Playlist{
		Nome: "Minha Playlist",
		Musicas: []Musica{
			{Titulo: "Música 1", Artista: "Artista 1", Duracao: 3 * time.Minute},
			{Titulo: "Música 2", Artista: "Artista 2", Duracao: 4 * time.Minute + 30 * time.Second},
			{Titulo: "Música 3", Artista: "Artista 3", Duracao: 2 * time.Minute + 45 * time.Second},
		},
	}

	imprimirPlaylistInfo(playlistExemplo)
}
