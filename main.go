package main

import (
	//"net/http"
	//"fmt"
	"math/rand" //for randInt
	"time" //for randInt
    "log"
	"fmt"
	
	"github.com/satori/go.uuid" //generates the UUID
	)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func airspacearrival(){
	
}

func main() {
	//begin
	log.SetFlags(0)
	log.Println(" \n")
	log.Println("go-airport v0.0.1 Listening on port 8080")
	//log.Fatal(http.ListenAndServe(":8080", status_api.Handlers()))
	
	type Gvars struct{
		I int
		X int
		Queue int
		WorkCenter int
		CounterW int
		Completed int
		Arrivalrate int
		Completionrate int
		Iterations int
	}
	
	type Airspace struct{
			Active	bool
			Jets	int
			Typet	string	
	}
		
	type Jet struct { //constructor
		
		
		
			Guid uuid.UUID 
			Airline string
			Flight int
			Passengers int
			Maxpassengers int
			Crew int
			Fuel int
			Maxfuel int
			Cargo int
			Maxcargo int
			Model string
	
			}
	
	
	//a
	
	type Tower struct{
		Tower int
		Active bool
		Modifier float32
		Baserate float32
		Workcounter int
	}
	
	type Runway struct{
		Num int
		Active bool
		Jets int  //actually a list! need to fix this.
		Landing int
		Takeoff int
		Heading int
		Length int
	}
	
	type Taxiway struct{
		Runwaynum int
		Active bool
		Jets int //actually an array
		Arrivals int
		Departures int
	}
	
	type Gate struct{
		Num int
		Jets int //actually an array
		Active bool
		Passengers int
		Cargo int
		Fuel float32
	}
	
	gate1 := &Gate{
		Num: 1,
		Jets: 0,
		Passengers: 10,
		Cargo: 1,
		Fuel: 20000,
	}
	
	gate2 := &Gate{
		Num: 2,
		Jets: 0,
		Passengers: 20,
		Cargo: 2,
		Fuel: 15000,
	}
	
	taxiway := &Taxiway{
		Runwaynum: 1,
		Active: true,
		Jets: 0, //actually an array
		Arrivals: 0,
		Departures: 0,
	}
	
	tower := &Tower{
		Tower: 1,
		Active: true,
		Modifier: 1.5,
		Baserate: 3,
		Workcounter: 0,
	}
	
	gvars := &Gvars{
		I: 0,
		X: 0,
		Completed: 0,
		Arrivalrate: 2,
		Iterations: 3500,
	}


	airspace := &Airspace{
		Active: true,
		Jets: 0,
		Typet: "civillian",
	}
	
	runway1 := &Runway{
		Num: 1,
		Active: true,
		Jets: 0,
		Takeoff: 0,
		Landing: 0,
		Length: 10000,
		
		
	}

	jet := &Jet{
		Guid: uuid.NewV4(),
		Airline: "American",
		Flight: random(100, 9999),
		Passengers: random(20, 125),
		Maxpassengers: 125,
		Crew: random(3, 5),
		Fuel: random(20, 99),
		Maxfuel: 100,
		Cargo: random(0, 4),
		Maxcargo: 4,
		Model: "Boeing 777", 
		}


	fmt.Printf("%+v \n", gvars)		
	fmt.Printf("%+v \n", jet)	//now with names
	fmt.Printf("%+v \n", airspace)
	fmt.Printf("%+v \n", runway1)
	fmt.Printf("%+v \n", tower)
	fmt.Printf("%+v \n", taxiway)
	fmt.Printf("%+v \n", gate1)
	fmt.Printf("%+v \n", gate2)


	
	/*func JetTails () {
		log.Printf("Guid", jet.Guid)
		log.Printf("/n Airline", jet.Airline)
		log.Printf("/n Flight", jet.Flight)
		log.Printf("/n Passengers", jet.Passengers)
		log.Printf("Maxpassengers", jet.Maxpassengers)
		log.Printf("Crew", jet.Crew)
		log.Printf("Fuel", jet.Fuel)
		log.Printf("Maxfuel", jet.Maxfuel)
		log.Printf("Cargo", jet.Cargo)
		log.Printf("Maxcargo", jet.Maxcargo)
		log.Printf("Model", jet.Model)
	//}
	valu := jet.Cargo + jet.Maxcargo
	log.Printf("", valu)
	//JetTails()*/
		
}