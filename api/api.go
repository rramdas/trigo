apackage api

import (
	"github.com/rs/zerolog/log"
)

// Triple is a subject-predicate-object tuple
type Triple struct {
	Subject   string
	Predicate string
	Object    string
}

// Graph is a collection of triples
type Graph struct {
	entities map[string]Triple // map of subject to triple
}

// MakeGraph creates a new graph with a unique ID and at least one non-empty triple
func MakeGraph(t Triple) (*Graph, error) {
	g := &Graph{make(map[string]Triple)}
	g.entities[t.Subject] = t
	return g, nil
}

func AddTriple(g *Graph, t Triple) {
	if _, ok := g.entities[t.Subject]; ok {
		return
	}
	g.entities[t.Subject] = t
	log.Info().Msgf("Added triple: %s %s %s", t.Subject, t.Predicate, t.Object)

}

func GetTriples(g *Graph, s string) []Triple {

	println("Search string: " + s)
	triples := make([]Triple, 0)
	for _, triple := range g.entities {
		println("XXX Triple: " + triple.Subject + " " + triple.Predicate + " " + triple.Object)
		if triple.Subject == s {
			println("found")
			triples = append(triples, triple)
		}
	}
	return triples
}

func deleteTriple(g *Graph, s string) {
	delete(g.entities, s)
}

func copyTriple(g *Graph, s string) Triple {
	return g.entities[s]
}

func main() {

	triple1 := Triple{"blade_runner", "name", "Blade Runner"}
	triple2 := Triple{"blade_runner", "release_date", "June 25 1982"}

	triple3 := Triple{"ridle_scott", "name", "Ridley Scott"}
	triple4 := Triple{"blade_runner", "directed_by", "ridley_scott"}

	println("Subject: " + triple1.Subject)
	println("Predicate: " + triple2.Predicate)
	println("Object: " + triple3.Object)
	println("Subject:" + triple4.Subject)

	g := Graph{make(map[string]Triple)}
	AddTriple(&g, triple1)
	AddTriple(&g, triple2)
	AddTriple(&g, triple3)
	AddTriple(&g, triple4)

	println("Graph: ")
	for key, value := range g.entities {
		println("Key: " + key)
		println("s: " + value.Subject + " p: " + value.Predicate + " o: " + value.Object)
		println("args ...Type")
	}

	triples := GetTriples(&g, "blade_runner")
	println("Triples: ")
	for _, triple := range triples {
		println("s: " + triple.Subject + " p: " + triple.Predicate + " o: " + triple.Object)
	}

}
