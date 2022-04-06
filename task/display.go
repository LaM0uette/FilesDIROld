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

func DrawEnd(s *Sch, timer time.Duration) {
	fmt.Printf(`
==================  BILAN DES RECHERCHES  ==================

Dossiers principal: %s
Temps d'exécution: %v

============================================================
Fichiers trouvés: %v


`, s.SrcPath, timer, s.NbFiles)
}
