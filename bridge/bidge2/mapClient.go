package bidge2

import "fmt"

type MapClient struct {
	Mapa Mapa
}

func (c *MapClient) ImprimirMapa() {
	fmt.Println(c.Mapa.DevolverMapa("Primeiro Ministro - Pilarzinho"))
}

func (c *MapClient) SetMapType(m Mapa) {
	c.Mapa = m
}
