package optativas

type CodigoAsignatura string

type OfertaOptativas struct {
	cursoAcademico string
	evaluaciones   map[CodigoAsignatura]EvaluacionOptativa
}
