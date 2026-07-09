package storage

import (
	"proyecto/internal/models"
	"proyecto/internal/storage/sqlcdb"
)

// --- Mapeadores de sqlcdb a models ---

func aCita(c sqlcdb.Citum) models.Cita {
	return models.Cita{
		ID:             int(c.ID),
		SolicitanteID:  c.SolicitanteID,
		TecnicoID:      c.TecnicoID,
		Estado:         c.Estado,
		HoraAcordada:   c.HoraAcordada,
		PuntoEncuentro: c.PuntoEncuentro,
	}
}

func aPunto(p sqlcdb.PuntoEncuentro) models.PuntoEncuentro {
	return models.PuntoEncuentro{
		ID:                    int(p.ID),
		NombreLugar:           p.NombreLugar,
		FacultadPerteneciente: p.FacultadPerteneciente,
		DisponibleParaSoporte: p.DisponibleParaSoporte,
	}
}

func aSoporte(s sqlcdb.Soporte) models.Soporte {
	return models.Soporte{
		ID:              int(s.ID),
		CitaID:          int(s.CitaID),
		DispositivoID:   int(s.DispositivoID),
		Solucion:        s.Solucion,
		PiezasCambiadas: s.PiezasCambiadas,
	}
}
