package display

import (
	"FilesDIR/globals"
	"fmt"
	"path/filepath"
	"time"
)

func DrawStart() string {
	return fmt.Sprintf(`<fg=48,207,37>
		███████╗██╗██╗     ███████╗██████╗ ██╗██████╗ 
		██╔════╝██║██║     ██╔════╝██╔══██╗██║██╔══██╗
		█████╗  ██║██║     █████╗  ██║  ██║██║██████╔╝
		██╔══╝  ██║██║     ██╔══╝  ██║  ██║██║██╔══██╗
		██║     ██║███████╗███████╗██████╔╝██║██║  ██║
		╚═╝     ╚═╝╚══════╝╚══════╝╚═════╝ ╚═╝╚═╝  ╚═╝</>
		<cyan>Version:</> <fg=48,207,37>%s</>              <cyan>Auteur:</> <fg=48,207,37>%s</>


`, globals.Version, globals.Author)
}

func DrawInitSearch() string {
	return fmt.Sprint(`<fg=214,99,144>Initialisation du programme...</>`)
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
	return fmt.Sprint(`<fg=214,99,144>Export Excel...   `)
}

func DrawSaveExcel() string {
	return fmt.Sprint(`<fg=214,99,144>Fichier Excel sauvegardé avec succes.</>`)
}

func DrawEnd(SrcPath, DstPath, ReqFinal string, NbGoroutine, NbFiles, PoolSize int, timerSearch time.Duration, timerTotal time.Duration) string {
	return fmt.Sprintf(`<cyan>
+============================================================+
|                    BILAN DES RECHERCHES                    |                     
+============================================================+
</>
<fg=214,99,144>#### - INFOS GENERALES :</>
<cyan>Dossiers principal:</> <fg=48,207,37>%s</>
<cyan>Requête utilisée:</> <fg=48,207,37>%s</>
<cyan>Nombre de Threads:</> <fg=48,207,37>%v</>
<cyan>Nombre de Goroutines:</> <fg=48,207,37>%v</>

<fg=214,99,144>#### - RESULTATS :</>
<cyan>Fichiers trouvés:</> <fg=48,207,37>%v</>
<cyan>Temps d'exécution de la recherche:</> <fg=48,207,37>%v</>
<cyan>Temps d'exécution total:</> <fg=48,207,37>%v</>

<fg=214,99,144>#### - EXPORTS :</>
<cyan>Logs:</> <fg=48,207,37>%s</>
<cyan>Dumps:</> <fg=48,207,37>%s</>
<cyan>Export Excel:</> <fg=48,207,37>%s</>

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
