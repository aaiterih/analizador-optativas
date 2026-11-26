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

func TestPorcentajeContinua(t *testing.T) {
	casos := map[string]struct {
		porcentajes map[TipoEvaluacion]float64
		esperado    float64
	}{
		"suma solo las partes de continua": {map[TipoEvaluacion]float64{Teoria: 45, Practicas: 30, Trabajos: 15, Entregables: 10}, 55},
		"solo teoría":                      {map[TipoEvaluacion]float64{Teoria: 100}, 0},
		"sin partes":                       {map[TipoEvaluacion]float64{}, 0},
	}

	for nombre, caso := range casos {
		t.Run(nombre, func(t *testing.T) {
			evaluacion := EvaluacionOptativa{porcentajes: caso.porcentajes}

			if obtenido := evaluacion.porcentajeContinua(); obtenido != caso.esperado {
				t.Errorf("porcentajeContinua() = %v, se esperaba %v", obtenido, caso.esperado)
			}
		})
	}
}

func TestPorcentajeContinuaNoModificaLaEvaluacion(t *testing.T) {
	evaluacion := EvaluacionOptativa{porcentajes: map[TipoEvaluacion]float64{Teoria: 60, Practicas: 40}}

	evaluacion.porcentajeContinua()

	if _, sigue := evaluacion.porcentajes[Teoria]; !sigue {
		t.Error("porcentajeContinua() ha borrado la teoría de la evaluación")
	}
}

func TestExigeNotaMinima(t *testing.T) {
	casos := map[string]struct {
		condiciones map[TipoEvaluacion]CondicionMinima
		esperado    bool
	}{
		"sin condiciones":         {nil, false},
		"con nota mínima":         {map[TipoEvaluacion]CondicionMinima{Teoria: {tipo: NotaMinima, notaMinima: nota(5)}}, true},
		"con aprobar una parte":   {map[TipoEvaluacion]CondicionMinima{Teoria: {tipo: AprobarParte}}, true},
		"con las dos condiciones": {map[TipoEvaluacion]CondicionMinima{Teoria: {tipo: NotaMinima, notaMinima: nota(4)}, Practicas: {tipo: AprobarParte}}, true},
	}

	for nombre, caso := range casos {
		t.Run(nombre, func(t *testing.T) {
			evaluacion := EvaluacionOptativa{condiciones: caso.condiciones}

			if obtenido := evaluacion.exigeNotaMinima(); obtenido != caso.esperado {
				t.Errorf("exigeNotaMinima() = %v, se esperaba %v", obtenido, caso.esperado)
			}
		})
	}
}
