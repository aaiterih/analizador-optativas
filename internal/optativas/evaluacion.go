package optativas

type TipoEvaluacion string

const (
	Teoria        TipoEvaluacion = "teoria"
	Practicas     TipoEvaluacion = "practicas"
	Trabajos      TipoEvaluacion = "trabajos"
	Entregables   TipoEvaluacion = "entregables"
	Participacion TipoEvaluacion = "participacion"
	Examen        TipoEvaluacion = "examen"
)

type ParteEvaluacion struct {
	tipo       TipoEvaluacion
	porcentaje float64
}
