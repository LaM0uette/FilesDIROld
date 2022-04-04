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

func DrawEndSearch(path, saveFolder string, nbrFolder, nbrFiles int) {
	fmt.Printf(`**********  FIN DE LA RECHERCHE  **********


            BILAN DES RECHERCHES
Dossier: %s
Nbr dossier parents: %v
Fichiers trouvés: %v
Emplacement de sauvegarde: %s


`, path, nbrFolder, nbrFiles, saveFolder)
}
