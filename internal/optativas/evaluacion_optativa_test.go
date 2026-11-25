package optativas

import (
	"errors"
	"testing"
)

func TestSuperaUmbralContinua(t *testing.T) {
	casos := map[float64]bool{
		51:  true,
		100: true,
		50:  false,
		0:   false,
	}

	for porcentaje, esperado := range casos {
		if obtenido := superaUmbralContinua(porcentaje); obtenido != esperado {
			t.Errorf("superaUmbralContinua(%v) = %v, se esperaba %v", porcentaje, obtenido, esperado)
		}
	}
}

func TestNuevaEvaluacionOptativaPorcentajes(t *testing.T) {
	casos := map[string]struct {
		porcentajes map[TipoEvaluacion]float64
		esperado    error
	}{
		"suman 100":                {map[TipoEvaluacion]float64{Teoria: 60, Practicas: 40}, nil},
		"decimales que suman 100":  {map[TipoEvaluacion]float64{Teoria: 33.3, Practicas: 33.3, Trabajos: 33.4}, nil},
		"porcentaje negativo":      {map[TipoEvaluacion]float64{Teoria: 110, Practicas: -10}, ErrPorcentajeFueraDeRango},
		"porcentaje mayor que 100": {map[TipoEvaluacion]float64{Teoria: 150}, ErrPorcentajeFueraDeRango},
		"no suman 100":             {map[TipoEvaluacion]float64{Teoria: 50, Practicas: 30}, ErrPorcentajesNoSuman100},
		"sin ninguna parte":        {map[TipoEvaluacion]float64{}, ErrPorcentajesNoSuman100},
	}

	for nombre, caso := range casos {
		t.Run(nombre, func(t *testing.T) {
			_, err := NuevaEvaluacionOptativa(caso.porcentajes, nil)

			if !errors.Is(err, caso.esperado) {
				t.Errorf("se esperaba %v, se obtuvo %v", caso.esperado, err)
			}
		})
	}
}
