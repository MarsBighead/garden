package structures

import (
	"fmt"
	"reflect"
)

func dataListVsStruct() {
	pointBaseList := [][2]int{{4, 6}, {}, {-7, 11}, {15, 17}, {14, -8}}
	for _, point := range pointBaseList {
		fmt.Printf(" (%d %d)\n", point[0], point[1])
	}
	pointBaseStruct := []struct{ x, y int }{{4, 6}, {}, {-7, 11}, {15, 17}, {14, -8}}
	for _, point := range pointBaseStruct {
		fmt.Printf(" (x, y) is (%d %d)\n", point.x, point.y)
	}
}

type Person struct {
	Title     string
	Forenames []string
	Surename  string
}

type Author struct {
	Names Person
	Title []string
	Yob   int
}

// Union test union struct
func Union() {
	author1 := Author{
		Person{" Mr ", []string{"Robert", " Louis", " Balfour"}, "Stevenson"},
		[]string{" Kidnapped ", " Treasure Island "},
		1850}
	fmt.Println("Author1 is ", author1)
	author1.Names.Title = ""
	author1.Names.Forenames = []string{" Oscar ", " Fingal ", " O'Flahertie ", " Wills"}
	author1.Names.Surename = " Wilde "
	author1.Title = []string{" The Picture of  Dorian Gray "}
	author1.Yob += 4
	fmt.Println("Modified Author1 is ", author1)
}

type Author2 struct {
	Person
	Title []string
	Yob   int
}

func MixedUnion() {

	author2 := Author2{
		Person{" Mr ", []string{"Robert", " Louis", " Balfour"}, "Stevenson"},
		[]string{" Kidnapped ", " Treasure Island "},
		1850}
	fmt.Println("author2 is ", author2)
	author2.Title = []string{" The Picture of  Dorian Gray "}
	author2.Person.Title = ""
	author2.Forenames = []string{" Oscar ", " Fingal ", " O'Flahertie ", " Wills"}
	author2.Surename = " Wilde "
	author2.Yob += 4
	fmt.Println(author2)
	ts := transferStruct(author2)
	fmt.Println("Author1 ", ts.Surename)

}

func transferStruct(i interface{}) Author2 {
	v := reflect.ValueOf(i)
	fmt.Printf("Elem of interface i is %v\n", reflect.ValueOf(&v).Elem())
	fmt.Printf("Kind of interface i is %v\n", v.Kind())
	fmt.Printf("Type of interface i is %v\n", v.Type())
	if value, ok := i.(Author); ok {
		fmt.Printf("Interface %v is struct Author2.\n", value)
	} else {
		fmt.Printf("Interface %v is not struct Author.\n", i)
	}
	val := i.(Author2)
	return val
}

type Tasks struct {
	slice []string
	Count
}
type Count struct{ X int }

func (tasks *Tasks) Add(task string) {
	tasks.slice = append(tasks.slice, task)
}
