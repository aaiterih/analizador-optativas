package optativas

import "testing"

func TestEsContinua(t *testing.T) {
	casos := map[TipoEvaluacion]bool{
		Practicas:     true,
		Trabajos:      true,
		Entregables:   true,
		Participacion: true,
		Teoria:        false,
		"desconocido": false,
	}

	for tipo, esperado := range casos {
		if obtenido := tipo.esContinua(); obtenido != esperado {
			t.Errorf("%v.esContinua() = %v, se esperaba %v", tipo, obtenido, esperado)
		}
	}
}