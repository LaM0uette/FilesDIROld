package task

import (
	"FilesDIR/globals"
	"fmt"
	"path/filepath"
	"time"
)

func DrawStart() {
	fmt.Printf(`
		███████╗██╗██╗     ███████╗██████╗ ██╗██████╗ 
		██╔════╝██║██║     ██╔════╝██╔══██╗██║██╔══██╗
		█████╗  ██║██║     █████╗  ██║  ██║██║██████╔╝
		██╔══╝  ██║██║     ██╔══╝  ██║  ██║██║██╔══██╗
		██║     ██║███████╗███████╗██████╔╝██║██║  ██║
		╚═╝     ╚═╝╚══════╝╚══════╝╚═════╝ ╚═╝╚═╝  ╚═╝
		Version: %s               Auteur: %s


`, globals.Version, globals.Author)
	time.Sleep(200 * time.Millisecond)
}

func DrawSetupSearch() {
	fmt.Print(`Initialisation de la recherche...

`)
	time.Sleep(1 * time.Second)
}

func DrawRunSearch() {
	fmt.Print(`==================   DEBUT DES RECHERCHES   ==================

`)
	time.Sleep(1 * time.Second)
}

func DrawEndSearch() {
	fmt.Print(`==================   FIN DES RECHERCHES   ==================


`)
}

func DrawWriteExcel() {
	fmt.Print(`Sauvegarde du fichier Excel...
`)
}

func DrawSaveExcel() {
	fmt.Print(`Fichier Excel sauvegardé avec succes.
`)
}

func DrawEnd(s *Sch, timer time.Duration) {
	fmt.Printf(`

==================  BILAN DES RECHERCHES  ==================

INFOS GENERALES:
Dossiers principal: %s
Nombre de Threads: %v
Nombre de Goroutines: %v

RESULTATS:
Fichiers trouvés: %v
Temps d'exécution: %v

EXPORTS:
Logs: %s
Dumps: %s
Export Excel: %s

============================================================
%s
Auteur: %s
Version: %s


`, s.SrcPath, s.PoolSize, s.NbGoroutine, s.NbFiles, timer, filepath.Join(globals.TempPathGen, "logs"), filepath.Join(globals.TempPathGen, "dumps"), s.DstPath, globals.Name, globals.Version, globals.Author)
}
