package optativas

type TipoEvaluacion string

const (
	Teoria        TipoEvaluacion = "teoria"
	Practicas     TipoEvaluacion = "practicas"
	Trabajos      TipoEvaluacion = "trabajos"
	Entregables   TipoEvaluacion = "entregables"
	Participacion TipoEvaluacion = "participacion"
)

var tiposEvaluacionContinua = map[TipoEvaluacion]bool{
	Practicas:     true,
	Trabajos:      true,
	Entregables:   true,
	Participacion: true,
}

func (t TipoEvaluacion) esContinua() bool {
	return tiposEvaluacionContinua[t]
}