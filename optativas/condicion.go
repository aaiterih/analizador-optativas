package optativas

type TipoCondicion string

const (
	NotaMinima   TipoCondicion = "nota_minima"
	AprobarParte TipoCondicion = "aprobar_parte"
)

type CondicionMinima struct {
	parte      TipoEvaluacion
	tipo       TipoCondicion
	notaMinima *float64
}
