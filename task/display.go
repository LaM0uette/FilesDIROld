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

type DrawStruct struct {
	Pattern string
}

var DrawInitSearch = DrawStruct{
	Pattern: fmt.Sprintf(`Initialisation de la recherche...

`),
}

var DrawRunSearch = DrawStruct{
	Pattern: fmt.Sprintf(`==================   DEBUT DES RECHERCHES   ==================
`),
}

var DrawEndSearch = DrawStruct{
	Pattern: fmt.Sprintf(`==================   FIN DES RECHERCHES   ==================


`),
}

var DrawWriteExcel = DrawStruct{
	Pattern: fmt.Sprintf(`Sauvegarde du fichier Excel...
`),
}

var DrawSaveExcel = DrawStruct{
	Pattern: fmt.Sprintf(`Fichier Excel sauvegardé avec succes.
`),
}

func (d *DrawStruct) Draw() {
	fmt.Print(d.Pattern)
}
