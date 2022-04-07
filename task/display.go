package task

import (
	"FilesDIR/globals"
	"fmt"
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
		Version: %v               Auteur: %s


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

func DrawSaveExcel(s *Sch) {
	fmt.Printf(`Fichier Excel sauvergardé: %s
`, s.DstPath)
}

func DrawEnd(s *Sch, timer time.Duration) {
	fmt.Printf(`

==================  BILAN DES RECHERCHES  ==================

INFOS GENERALES:
Dossiers principal: %s
Nombre de Threads: %v
Nombre de Goroutines: %v
Temps d'exécution: %v

RESULTATS:
Fichiers trouvés: %v

============================================================


`, s.SrcPath, s.PoolSize, s.NbGoroutine, timer, s.NbFiles)
}
