package main
import "fmt"
func main () {

	superhero := map[string]map[string]string{

		"Superman" :map[string]string{
			"RealName" :"Clark Kent",
			"City" : "MetropoLis",

		},

		"Batman" : map[string] string {
			"RealName" : "Bruce wayne",
			"City" :"Gotham City",

		},

		
		}
		if temp, hero := superhero["Superman"]; hero {
			fmt.Println(temp["RealName"], temp["City"])
	}

}

