package optativas

import (
	"errors"
	"testing"
)

func nota(valor float64) *float64 {
	return &valor
}

func TestNuevaEvaluacionOptativaCondiciones(t *testing.T) {
	casos := map[string]struct {
		condicion CondicionMinima
		esperado  error
	}{
		"nota mínima de 5":           {CondicionMinima{tipo: NotaMinima, notaMinima: nota(5)}, nil},
		"aprobar una parte sin nota": {CondicionMinima{tipo: AprobarParte}, nil},
		"nota mínima sin nota":       {CondicionMinima{tipo: NotaMinima}, ErrNotaMinimaInvalida},
		"nota mínima de 12":          {CondicionMinima{tipo: NotaMinima, notaMinima: nota(12)}, ErrNotaMinimaInvalida},
		"nota mínima negativa":       {CondicionMinima{tipo: NotaMinima, notaMinima: nota(-1)}, ErrNotaMinimaInvalida},
	}

	for nombre, caso := range casos {
		t.Run(nombre, func(t *testing.T) {
			condiciones := map[TipoEvaluacion]CondicionMinima{Teoria: caso.condicion}

			_, err := NuevaEvaluacionOptativa(map[TipoEvaluacion]float64{Teoria: 100}, condiciones)

			if !errors.Is(err, caso.esperado) {
				t.Errorf("se esperaba %v, se obtuvo %v", caso.esperado, err)
			}
		})
	}
}