package main
import(
	"fmt"
)

type State struct {
    name      string
    population int
    capital   string
}

//Maps
func main () {
	
	// Declaring maps
	map1 := map[int]string{}
	map2 := make(map[int]string)

	// Declaring a map with specific size
	map3 := make(map[int]string, 4096)

	// Attributing in the declaration
	capitals := map[string]string {
    "GO": "Goiânia",
    "PB": "João Pessoa",
    "PR": "Curitiba"}
	
	// Attributing after declaration
	capitals["RN"] = "Natal"
	capitals["AM"] = "Manaus"
	capitals["SE"] = "Aracaju"
	capitals["RJ"] = "Rio de Janeiro"

	fmt.Println(len(capitals), map1, map2, map3, capitals, '\n')

	states := make(map[string]State, 6)

	states["GO"] = State{"Goiás", 6434052, "Goiânia"}
	states["PB"] = State{"Paraíba", 3914418, "João Pessoa"}
	states["PR"] = State{"Paraná", 10997462, "Curitiba"}
	states["RN"] = State{"Rio Grande do Norte", 3373960, "Natal"}
	states["AM"] = State{"Amazonas", 3807923, "Manaus"}
	states["SE"] = State{"Sergipe", 2228489, "Aracaju"}

	fmt.Println(states)


	// lookup
	sergipe := states["SE"]
	fmt.Println(sergipe)

	// The second value returned by the lookup operation is a bool 
	// that will receive the value true if the key is present on the map, or false otherwise.
	saoPaulo, found := states["SP"]
	if found {
    	fmt.Println(saoPaulo)
	} else {
		fmt.Println("SP not found")
	}

	// Update a value
	ages := map[string]int{
    	"John":    37,
    	"Richard": 26}

	ages["Richard"] = 27
	fmt.Println(ages["Richard"])

	// Remove a value
	delete(ages, "Richard")
	fmt.Println(ages)

	// Iterating
	for initials, state := range states {
    	fmt.Printf("%s (%s) has %d inhabitant.\n",
		state.name, initials, state.population)
	}

}