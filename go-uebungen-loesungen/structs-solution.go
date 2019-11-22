package main

type studierender struct {
	name           string
	vorname        string
	matrikelnummer int
	studiengang    string
}

func main() {

	s := studierender{"Müller", "Peter", 1234567, "Informatik"}
	p := &s
	p.studiengang = "Wirtschaftsinformatik"

}
