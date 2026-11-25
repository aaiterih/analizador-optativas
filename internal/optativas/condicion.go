package optativas

import (
	"errors"
	"maps"
	"slices"
)

type TipoCondicion string

const (
	NotaMinima   TipoCondicion = "nota_minima"
	AprobarParte TipoCondicion = "aprobar_parte"
)

var ErrNotaMinimaInvalida = errors.New("una nota mínima tiene que tener una nota entre 0 y 10")

type CondicionMinima struct {
	tipo       TipoCondicion
	notaMinima *float64
}

func validarCondiciones(condiciones map[TipoEvaluacion]CondicionMinima) error {
	invalidas := slices.ContainsFunc(slices.Collect(maps.Values(condiciones)), CondicionMinima.esInvalida)
	return errorSi(invalidas, ErrNotaMinimaInvalida)
}

func (c CondicionMinima) esInvalida() bool {
	return c.tipo == NotaMinima && !notaEnRango(c.notaMinima)
}

func notaEnRango(nota *float64) bool {
	return nota != nil && *nota >= 0 && *nota <= 10
}