package optativas

import "testing"

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