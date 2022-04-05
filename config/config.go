package config

type Data struct {
	Name    string
	Author  string
	Version float32
}

func Config() Data {
	c := Data{
		Name:    "FileDIR",
		Author:  "LaM0uette",
		Version: 0.4,
	}
	return c
}
