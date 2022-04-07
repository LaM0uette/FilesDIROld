package task

import (
	"FilesDIR/globals"
	"fmt"
	"path/filepath"
	"time"
)

func DrawStart() string {
	return fmt.Sprintf(`
		███████╗██╗██╗     ███████╗██████╗ ██╗██████╗ 
		██╔════╝██║██║     ██╔════╝██╔══██╗██║██╔══██╗
		█████╗  ██║██║     █████╗  ██║  ██║██║██████╔╝
		██╔══╝  ██║██║     ██╔══╝  ██║  ██║██║██╔══██╗
		██║     ██║███████╗███████╗██████╔╝██║██║  ██║
		╚═╝     ╚═╝╚══════╝╚══════╝╚═════╝ ╚═╝╚═╝  ╚═╝
		Version: %s               Auteur: %s


`, globals.Version, globals.Author)
}

func DrawInitSearch() string {
	return fmt.Sprint(`Initialisation de la recherche...
`)
}

func DrawRunSearch() string {
	return fmt.Sprint(`==================   DEBUT DES RECHERCHES   ==================`)
}

func DrawEndSearch() string {
	return fmt.Sprint(`==================   FIN DES RECHERCHES   ==================
`)
}

func DrawWriteExcel() string {
	return fmt.Sprint(`
Sauvegarde du fichier Excel...`)
}

func DrawSaveExcel() string {
	return fmt.Sprint(`
Fichier Excel sauvegardé avec succes.
`)
}

func DrawEnd(s *Sch, timerSearch time.Duration, timerTotal time.Duration) string {
	return fmt.Sprintf(`
==================  BILAN DES RECHERCHES  ==================

INFOS GENERALES:
Dossiers principal: %s
Nombre de Threads: %v
Nombre de Goroutines: %v

RESULTATS:
Fichiers trouvés: %v
Temps d'exécution de la recherche: %v
Temps d'exécution total: %v

EXPORTS:
Logs: %s
Dumps: %s
Export Excel: %s

============================================================
%s
Auteur: %s
Version: %s
`, s.SrcPath, s.PoolSize, s.NbGoroutine, s.NbFiles, timerSearch, timerTotal, filepath.Join(globals.TempPathGen, "logs"), filepath.Join(globals.TempPathGen, "dumps"), s.DstPath, globals.Name, globals.Author, globals.Version)
}
