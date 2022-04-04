package build

import (
	"Test/config"
	"fmt"
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

func DrawEndSearch(path, saveFolder string, nbrFiles int) {
	fmt.Printf(`**********  FIN DE LA RECHERCHE  **********


            BILAN DES RECHERCHES
Dossier: %s
Fichiers trouvés: %v
Emplacement de sauvegarde: %s


`, path, nbrFiles, saveFolder)
}
