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
	return fmt.Sprint(`<yellow>Initialisation du programme...</><green>`)
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
	return fmt.Sprint(`<magenta>Sauvegarde du fichier Excel...</>  <green>`)
}

func DrawSaveExcel() string {
	return fmt.Sprint(`<magenta>Fichier Excel sauvegardé avec succes.</>`)
}

func DrawEnd(SrcPath, DstPath, ReqFinal string, NbGoroutine, NbFiles, PoolSize int, timerSearch time.Duration, timerTotal time.Duration) string {
	return fmt.Sprintf(`<cyan>
+============================================================+
|                    BILAN DES RECHERCHES                    |                     
+============================================================+
</>
<magenta>#### - INFOS GENERALES :</>
Dossiers principal: <green>%s</>
Requête utilisée: <green>%s</>
Nombre de Threads: <green>%v</>
Nombre de Goroutines: <green>%v</>

<magenta>#### - RESULTATS :</>
Fichiers trouvés: <green>%v</>
Temps d'exécution de la recherche: <green>%v</>
Temps d'exécution total: <green>%v</>

<magenta>#### - EXPORTS :<magenta/>
Logs: <green>%s</>
Dumps: <green>%s</>
Export Excel: <green>%s</>

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
