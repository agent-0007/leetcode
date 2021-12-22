package hashmaps

import "fmt"

func Test() {
    peopleList := map[string]int {
        "Max" : 24,
        "Bob" : 20,
        "Alice" : 18,
    }
    for name, age := range peopleList {
        fmt.Printf("%s\t%d\n", name, age)
    }
    fmt.Println(peopleList)
    peopleList["Piter"] = 46
    fmt.Println(peopleList)
    delete(peopleList, "Bob")
    fmt.Println(peopleList)
}

func TestMaps() {
    ages := make(map[string]int)
    ages["max"] = 36
    ages["Bob"] = 24
    fmt.Println(ages)
    fmt.Println(ages["Bob"])
}