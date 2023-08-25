package bidge2

type MapLink struct {
}

func (m *MapLink) DevolverMapa(rua string) string {
	return "Devolvendo mapa com MapLink para a rua: " + rua
}
