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
}

func DrawRunSearch() {
	fmt.Print(`
==================  DEBUT DES RECHERCHES  ==================
`)
}

func DrawEndSearch() {
	fmt.Print(`
==================   FIN DES RECHERCHES   ==================
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
