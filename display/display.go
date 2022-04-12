package display

import (
	"FilesDIR/globals"
	"fmt"
	"path/filepath"
	"time"
)

func DrawStart() string {
	return fmt.Sprintf(`<fg=%[3]s>
		███████╗██╗██╗     ███████╗██████╗ ██╗██████╗ 
		██╔════╝██║██║     ██╔════╝██╔══██╗██║██╔══██╗
		█████╗  ██║██║     █████╗  ██║  ██║██║██████╔╝
		██╔══╝  ██║██║     ██╔══╝  ██║  ██║██║██╔══██╗
		██║     ██║███████╗███████╗██████╔╝██║██║  ██║
		╚═╝     ╚═╝╚══════╝╚══════╝╚═════╝ ╚═╝╚═╝  ╚═╝</>
		<fg=%[4]s;op=bold;>Version:</> <fg=%[3]s>%[1]s</>              <fg=%[4]s;op=bold;>Auteur:</> <fg=%[3]s>%[2]s</>


`, globals.Version, globals.Author, globals.Th1, globals.Th2)
}

func DrawInitSearch() string {
	return fmt.Sprintf(`<fg=%[1]s>Initialisation du programme...</>`,
		globals.Param)
}

func DrawCheckMinimumPoolSize() string {
	return fmt.Sprintf(`<fg=%[1]s>Poolsize mise à</> <fg=%[2]s;op=bold;>2</> <fg=%[1]s>(ne peut pas être inférieur à</> <fg=%[2]s;op=bold;>2</><fg=%[1]s>)</>`,
		globals.Param, globals.Th1)
}

func DrawSetMaxThread(v any) string {
	return fmt.Sprintf(`<fg=%[1]s>Nombre de threads mis à :</> <fg=%[3]s;op=bold;>%[2]v</>`,
		globals.Param, v, globals.Th1)
}

func DrawRunSearch() string {
	return fmt.Sprintf(`<fg=%[1]s>
+============================================================+
|     **********     DEBUT DES RECHERCHES     **********     |
+============================================================+
</>`, globals.Th2)
}

func DrawFileSearched(num int, file string) string {
	return fmt.Sprintf("\r<bg=%[1]s;fg=255,255,255;op=bold;>[OK]</> <fg=%[2]s;op=bold;>N°</><fg=%[1]s>%[3]v</> <fg=%[2]s;op=bold;>|=|</> <fg=%[1]s>%[4]s</>\n",
		globals.Th1, globals.Th2, num, file)
}

func DrawSearchedFait(num int) string {
	return fmt.Sprintf("\r<fg=%[2]s>Fait:</> <fg=%[1]s>%[3]v</>",
		globals.Th1, globals.Th2, num)
}

func DrawEndSearch() string {
	return fmt.Sprintf(`<fg=%[1]s>
+============================================================+
|     **********      FIN DES RECHERCHES      **********     |
+============================================================+
</>`, globals.Th2)
}

func DrawWriteExcel(i, imax int) string {
	return fmt.Sprintf("\r<fg=%[1]s>Export Excel...</><fg=%[2]s>%[3]v</><fg=%[1]s>/</><fg=%[2]s>%[4]v</>",
		globals.Param, globals.Th1, i, imax)
}

func DrawGenerateExcelSave(imax int) string {
	return fmt.Sprintf("\r<fg=%[1]s>Nombre de lignes sauvegardées :</>  <fg=%[2]s>%[3]v</><fg=%[1]s>/</><fg=%[2]s>%[3]v</>\n",
		globals.Param, globals.Th1, imax)
}

func DrawSetSaveWord(word string) string {
	return fmt.Sprintf("<fg=%[1]s>Nom du fichier de sauvergarde mis par défaut :</> <fg=%[2]s;op=bold;>%[3]v</>",
		globals.Param, globals.Th1, word)
}

func DrawSaveExcel() string {
	return fmt.Sprintf(`<fg=%[1]s>Fichier Excel sauvegardé avec succes.</>`,
		globals.Param)
}

func DrawEnd(SrcPath, DstPath, ReqFinal string, NbGoroutine, NbFiles, PoolSize int, timerSearch time.Duration, timerTotal time.Duration) string {
	return fmt.Sprintf(`<fg=%[3]s>
+============================================================+
|                    BILAN DES RECHERCHES                    |                     
+============================================================+
</>
<fg=%[3]s;op=bold;>#### - INFOS GENERALES :</>
<fg=%[1]s>Dossiers principal:</> <fg=%[2]s>%[4]s</>
<fg=%[1]s>Requête utilisée:</> <fg=%[2]s>%[5]s</>
<fg=%[1]s>Nombre de Threads:</> <fg=%[2]s>%[6]v</>
<fg=%[1]s>Nombre de Goroutines:</> <fg=%[2]s>%[7]v</>

<fg=%[3]s;op=bold;>#### - RESULTATS :</>
<fg=%[1]s>Fichiers trouvés:</> <fg=%[2]s>%[8]v</>
<fg=%[1]s>Temps d'exécution de la recherche:</> <fg=%[2]s>%[9]v</>
<fg=%[1]s>Temps d'exécution total:</> <fg=%[2]s>%[10]v</>

<fg=%[3]s;op=bold;>#### - EXPORTS :</>
<fg=%[1]s>Logs:</> <fg=%[2]s>%[11]s</>
<fg=%[1]s>Dumps:</> <fg=%[2]s>%[12]s</>
<fg=%[1]s>Export Excel:</> <fg=%[2]s>%[13]s</>

<fg=%[3]s>+=========  Auteur:</> <fg=%[2]s>%[14]s</>       <fg=%[3]s>Version:</> <fg=%[2]s>%[15]s</>  <fg=%[3]s>=========+</>
`,
		globals.Param,
		globals.Th1,
		globals.Th2,
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
