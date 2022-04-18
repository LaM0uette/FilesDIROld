package pkg

import "time"

func (s *Search) CompileFichesAppuis() {
	s.DrawSep("PARAMETRES")

	s.DrawParam("INITIALISATION DE LA COMPILATION EN COURS")

	s.Timer.CompileStart = time.Now()

	s.Timer.CompileEnd = time.Since(s.Timer.CompileStart)
}
