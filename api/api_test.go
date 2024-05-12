package api

import (
	"testing"
)

func TestAddTriple(t *testing.T) {
	t.Run("AddTriple", func(t *testing.T) {
		g := Graph{entities: make(map[string]Triple)}
		triple1 := Triple{"blade_runner", "name", "Blade Runner"}
		AddTriple(&g, triple1)
		if len(g.entities) != 1 {
			t.Errorf("Expected 1, got %d", len(g.entities))
		}
	})
}

func TestAddTripleExisting(t *testing.T) {
	t.Run("AddTriple that already exists", func(t *testing.T) {
		g := Graph{entities: make(map[string]Triple)}
		triple1 := Triple{"blade_runner", "name", "Blade Runner"}
		AddTriple(&g, triple1)
		AddTriple(&g, triple1)
		if len(g.entities) != 1 {
			t.Errorf("Expected 1, got %d", len(g.entities))
		}
	})
}

func TestGetTriples(t *testing.T) {
	t.Run("GetTriples", func(t *testing.T) {
		triple1 := Triple{"blade_runner", "name", "Blade Runner"}
		triple2 := Triple{"blade_runner", "release_date", "June 25 1982"}

		g, err := MakeGraph(triple1)
		if err != nil {
			t.Errorf("Error creating graph")
		}

		AddTriple(g, triple1)
		AddTriple(g, triple2)

		triples := GetTriples(g, "blade_runner")

		if len(triples) != 2 {
			t.Errorf("Expected 2, got %d", len(triples))
		}
	})
}
