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

func DrawRunSearch() {
	fmt.Print(`Début de la recherche...

`)
	time.Sleep(2 * time.Second)
}

func DrawEndSearch() {
	fmt.Print(`==================   FIN DES RECHERCHES   ==================
`)
}

func DrawEnd(s *Sch, timer time.Duration) {
	fmt.Printf(`


==================  BILAN DES RECHERCHES  ==================

INFOS GENERALES:
Dossiers principal: %s
Temps d'exécution: %v

RESULTATS:
Fichiers trouvés: %v

============================================================


`, s.SrcPath, timer, s.NbFiles)
}
