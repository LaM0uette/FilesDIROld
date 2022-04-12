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
	return fmt.Sprint(`<magenta>Export Excel...   `)
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
<cyan>Dossiers principal:</> <green>%s</>
<cyan>Requête utilisée:</> <green>%s</>
<cyan>Nombre de Threads:</> <green>%v</>
<cyan>Nombre de Goroutines:</> <green>%v</>

<magenta>#### - RESULTATS :</>
<cyan>Fichiers trouvés:</> <green>%v</>
<cyan>Temps d'exécution de la recherche:</> <green>%v</>
<cyan>Temps d'exécution total:</> <green>%v</>

<magenta>#### - EXPORTS :</>
<cyan>Logs:</> <green>%s</>
<cyan>Dumps:</> <green>%s</>
<cyan>Export Excel:</> <green>%s</>

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
