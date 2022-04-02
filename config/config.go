package config

type Data struct {
	Name    string
	Version float32
}

func Config() Data {
	c := Data{
		Name:    "FileDIR",
		Version: 0.2,
	}
	return c
}
