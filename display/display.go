package display

import (
	"FilesDIR/globals"
	"fmt"
	"path/filepath"
	"time"
)

func DrawStart() string {
	return fmt.Sprintf(`<lightCyan>
			███████╗██╗██╗     ███████╗██████╗ ██╗██████╗ 
			██╔════╝██║██║     ██╔════╝██╔══██╗██║██╔══██╗
			█████╗  ██║██║     █████╗  ██║  ██║██║██████╔╝
			██╔══╝  ██║██║     ██╔══╝  ██║  ██║██║██╔══██╗
			██║     ██║███████╗███████╗██████╔╝██║██║  ██║
			╚═╝     ╚═╝╚══════╝╚══════╝╚═════╝ ╚═╝╚═╝  ╚═╝</>
			<cyan>Version:</> <yellow>%s</>              <cyan>Auteur:</> <yellow>%s</>


`, globals.Version, globals.Author)
}

func DrawInitSearch() string {
	return fmt.Sprint(`<yellow>Initialisation du programme...</>`)
}

func DrawRunSearch() string {
	return fmt.Sprint(`<cyan>
+============================================================+
|                    DEBUT DES RECHERCHES                    |
+============================================================+
</>`)
}

func DrawEndSearch() string {
	return fmt.Sprint(`<cyan>
+============================================================+
|                     FIN DES RECHERCHES                     |                      
+============================================================+
</>`)
}

func DrawWriteExcel() string {
	return fmt.Sprint(`Sauvegarde du fichier Excel...  `)
}

func DrawSaveExcel() string {
	return fmt.Sprint(`Fichier Excel sauvegardé avec succes.`)
}

func DrawEnd(SrcPath, DstPath, ReqFinal string, NbGoroutine, NbFiles, PoolSize int, timerSearch time.Duration, timerTotal time.Duration) string {
	return fmt.Sprintf(`<cyan>
+============================================================+
|                    BILAN DES RECHERCHES                    |                     
+============================================================+
</>
<magenta>#### - INFOS GENERALES :</>
Dossiers principal: <yellow>%s</>
Requête utilisée: <yellow>%s</>
Nombre de Threads: <yellow>%v</>
Nombre de Goroutines: <yellow>%v</>

<magenta>#### - RESULTATS :</>
Fichiers trouvés: <yellow>%v</>
Temps d'exécution de la recherche: <yellow>%v</>
Temps d'exécution total: <yellow>%v</>

<magenta>#### - EXPORTS :</>
Logs: <yellow>%s</>
Dumps: <yellow>%s</>
Export Excel: <yellow>%s</>

<cyan>+=========  Auteur:</> <yellow>%s</>       <cyan>Version:</> <yellow>%s</>  <cyan>=========+</>

`,
		SrcPath,
		ReqFinal,
		PoolSize,
		NbGoroutine,

		NbFiles,
		timerSearch,
		timerTotal,

		filepath.Join(globals.TempPathGen, "logs"),
		filepath.Join(globals.TempPathGen, "dumps"),
		DstPath,

		globals.Author, globals.Version)
}

func DrawInitCompiler() string {
	return fmt.Sprint(`Initialisation de la compilation...`)
}

func DrawRunCompiler() string {
	return fmt.Sprint(`
+============================================================+
|                   DEBUT DES COMPILATIONS                   |
+============================================================+
`)
}

func DrawEndCompiler() string {
	return fmt.Sprint(`
+============================================================+
|                    FIN DES COMPILATIONS                    |
+============================================================+
`)
}
