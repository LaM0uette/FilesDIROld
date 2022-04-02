package build

import (
	"Test/config"
	"fmt"
)

// DrawStart : Display the screen of start application
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

// DrawStartSearch : Display the start of search
func DrawStartSearch() {
	fmt.Printf("**********  DEBUT DE LA RECHERCHE  **********\n")
}

// DrawEndSearch : Display the end of search
func DrawEndSearch(path, saveFolder string, nbrFiles int) {
	fmt.Printf(`**********  FIN DE LA RECHERCHE  **********


            BILAN DES RECHERCHES
Dossier: %s
Fichiers trouvés: %v
Emplacement de sauvegarde: %s


`, path, nbrFiles, saveFolder)
}
