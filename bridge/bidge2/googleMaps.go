package bidge2

type GoogleMap struct {
}

func (g *GoogleMap) DevolverMapa(rua string) string {
	return "Devolvendo mapa com o googleMaps para a rua: " + rua
}
