package main

import "fmt"
import "math"
func avg(xs []float64) float64{
    total:= 0.0
    for _, value:= range xs {
        fmt.Println(value)
        total+= value
    }
    return total / float64(len(xs))
}

type Circle struct {
    x, y, radius float64
}

func (c *Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func main(){
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)

    output:= input *2
    fmt.Println(output)

    i:= 1
    for i <= 10 {
        if i % 2 == 0 {
            fmt.Println(i, "even")
        } else {
            fmt.Println(i, "odd")
        }
        i += 1
    }

    fmt.Println("Arrays");
    ar := []float64 { 12.9, 19.0, 17.2, 10.3, 23, }
    fmt.Println("Average:", avg(ar))
    fmt.Println("Slices of:", ar);
    slice1 := ar[3:5]
    slice2 := make([]float64,2)
    copy(slice2, slice1)
    slice3:= append(slice2, 8,9)
    fmt.Println(slice1,slice2, slice3)
    fmt.Println("Maps")
    elements := map[string]map[string]string {
		"H": {"name": "Hydrogen", "state": "gas", },
		"He": {"name": "Helium", "state": "gas", },
        "Li": {"name": "Lithium", "state": "solid", },
        "Be": {"name": "Beryllium", "state": "solid", },
        "B": {"name": "Boron", "state": "solid", },
        "C": {"name": "Carbon", "state": "solid", },
        "N": {"name": "Nitrogen", "state": "gas", },
        "O": {"name": "Oxygen", "state": "gas", },
        "F": {"name": "Fluorine", "state": "gas", },
        "Ne": {"name": "Neon", "state": "gas", },
    }
	fmt.Println("Map:", elements, "\nEnter a key:")
	var sin string
	fmt.Scanf("%s", &sin);
	if result, ok := elements[sin]; ok {
		fmt.Println(result["name"], result["state"]);
	} else {
		fmt.Println("Invalid key");
    }

    fmt.Println("Structs")
    c:= Circle{0,0,10.1}
    fmt.Println(c.area())

}
