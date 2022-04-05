package build

import (
	"Test/config"
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


`, config.Config().Version, config.Config().Author)
}

func DrawStartSearch() {
	fmt.Printf("**********  DEBUT DE LA RECHERCHE  **********\n")
}

func DrawEndSearch(time time.Duration, path, reqUse, saveFolder string, nbrFolder, nbrFiles int) {
	fmt.Printf(`**********  FIN DE LA RECHERCHE  **********


BILAN DES RECHERCHES :
La recherche à mis : %s
------------------------------------------------------------------

Dossier principal de recherche: %s
Requête utilisée: %s

Nombre de dossiers parents: %v
Fichiers trouvés: %v
Emplacement de sauvegarde: %s


`, time, path, reqUse, nbrFolder, nbrFiles, saveFolder)
}
