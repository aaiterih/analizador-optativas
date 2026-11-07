package optativas

type EvaluacionOptativa struct {
	porcentajes map[TipoEvaluacion]float64
	condiciones map[TipoEvaluacion]CondicionMinima
}
