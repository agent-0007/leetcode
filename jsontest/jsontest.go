package jsontest

import (
    "encoding/json"
    "fmt"
    "net/http"
)
var JsonUrl = "https://jsonplaceholder.typicode.com/users"
/*
{
    id: 1,
    name: "Leanne Graham",
    username: "Bret",
    email: "Sincere@april.biz",
    address: {
        street: "Kulas Light",
        suite: "Apt. 556",
        city: "Gwenborough",
        zipcode: "92998-3874",
        geo: {
            lat: "-37.3159",
            lng: "81.1496"
        }
    },
    phone: "1-770-736-8031 x56442",
    website: "hildegard.org",
    company: {
        name: "Romaguera-Crona",
        catchPhrase: "Multi-layered client-server neural-net",
        bs: "harness real-time e-markets"
    }
},
*/
type JsonAddressGeo struct {
    Lat string `json:"lat"`
    Lng string `json:"lng"`
}

type JsonAddress struct {
    Street string `json:"street"`
    Suite string `json:"suite"`
    City string `json:"city"`
    Zipcode string `json:"zipcode"`
    Geo JsonAddressGeo `json:"geo"`
}

type JsonCompany struct {
    Name string `json:"name"`
    CatchPhrase string `json:"catch_phrase"`
    Bs string `json:"bs"`
}

type JsonItem struct {
    Id   int    `json:"id,omitempty"`
    Name string `json:"name"`
    Username string `json:"username"`
    Email string `json:"email"`
    Address JsonAddress `json:"address"`
    Phone string `json:"phone"`
    Website string `json:"website"`
    Company JsonCompany `json:"company"`
}


func GetJson(uri string) (*[]JsonItem, error){
    resp, err := http.Get(JsonUrl)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("Error while request: %s", resp.Status)
    }
    var result []JsonItem
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        return nil, err
    }
    resp.Body.Close()
    return &result,nil
}
func JsonTest() {
    // тут будет указатель на результат
    result, err := GetJson(JsonUrl)
    if err != nil {
        fmt.Println("Some shit happends")
    }

    for k,v := range *result {
        fmt.Println(k)
        fmt.Println(v)
        fmt.Printf("Lat: %s, Lon: %s", v.Address.Geo.Lat, v.Address.Geo.Lng)
    }
}

func JsonTestSimple() {
    var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
    type Animal struct {
        Name  string
        Order string
    }
    var animals []Animal
    err := json.Unmarshal(jsonBlob, &animals)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("%+v", animals)
}
