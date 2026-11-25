package optativas

const umbralEvaluacionContinua = 50.0

type EvaluacionOptativa struct {
	porcentajes map[TipoEvaluacion]float64
	condiciones map[TipoEvaluacion]CondicionMinima
}

func superaUmbralContinua(porcentajeContinua float64) bool {
	return porcentajeContinua > umbralEvaluacionContinua
}