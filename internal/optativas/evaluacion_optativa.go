package optativas

import (
	"errors"
	"maps"
	"math"
	"slices"
)

const (
	umbralEvaluacionContinua = 50.0
	toleranciaSuma           = 1e-9
)

var (
	ErrPorcentajeFueraDeRango = errors.New("hay un porcentaje fuera del rango de 0 a 100")
	ErrPorcentajesNoSuman100  = errors.New("los porcentajes no suman 100")
)

type EvaluacionOptativa struct {
	porcentajes map[TipoEvaluacion]float64
	condiciones map[TipoEvaluacion]CondicionMinima
}

func NuevaEvaluacionOptativa(porcentajes map[TipoEvaluacion]float64, condiciones map[TipoEvaluacion]CondicionMinima) (EvaluacionOptativa, error) {
	if err := errors.Join(validarPorcentajes(porcentajes), validarCondiciones(condiciones)); err != nil {
		return EvaluacionOptativa{}, err
	}

	return EvaluacionOptativa{porcentajes: porcentajes, condiciones: condiciones}, nil
}

func validarPorcentajes(porcentajes map[TipoEvaluacion]float64) error {
	valores := slices.Collect(maps.Values(porcentajes))

	return errors.Join(
		errorSi(slices.ContainsFunc(valores, fueraDeRango), ErrPorcentajeFueraDeRango),
		errorSi(!sumaCien(valores), ErrPorcentajesNoSuman100),
	)
}

func fueraDeRango(porcentaje float64) bool {
	return porcentaje < 0 || porcentaje > 100
}

func sumaCien(valores []float64) bool {
	return math.Abs(suma(valores)-100) < toleranciaSuma
}

func suma(valores []float64) float64 {
	total := 0.0
	for _, valor := range valores {
		total += valor
	}
	return total
}

func errorSi(condicion bool, err error) error {
	if condicion {
		return err
	}
	return nil
}

func superaUmbralContinua(porcentajeContinua float64) bool {
	return porcentajeContinua > umbralEvaluacionContinua
}

func (e EvaluacionOptativa) porcentajeContinua() float64 {
	continua := maps.Clone(e.porcentajes)
	maps.DeleteFunc(continua, func(tipo TipoEvaluacion, _ float64) bool {
		return !tipo.esContinua()
	})

	return suma(slices.Collect(maps.Values(continua)))
}
