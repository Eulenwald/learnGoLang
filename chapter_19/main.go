package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
)

type user struct {
	Name    string
	Alter   int
	Sprüche []string
}

func main() {
	fmt.Println("Kapitel 19 Uebungen")

	//uebung01()
	//uebung02()
	//uebung03()
	//uebung04()
	uebung05()
}

/*
	Lektion 198 – Übung 1

Nehmen Sie den folgenden Code als Grundlage: https://go.dev/play/p/wgCYvv88GnY
Packen Sie den []user in JSON und geben Sie dieses aus.
Hilfe: Denken Sie daran, was Sie tun müssen, um eine Variable außerhalb ihres Pakets verfügbar zu
machen
*/
func uebung01() {
	fmt.Println("Uebung 1")

	u1 := user{Name: "Homer", Alter: 36}
	u2 := user{Name: "Marge", Alter: 34}
	u3 := user{Name: "Bart", Alter: 10}
	u4 := user{Name: "Lisa", Alter: 8}
	u5 := user{Name: "Maggie", Alter: 1}

	users := []user{u1, u2, u3, u4, u5}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonUsers))

	fmt.Println(users)
}

/*
	Lektion 200 – Übung 2

Sie erhalten nach dem Funktionsaufruf zum Abschalten eines Antiviren-Programmes folgende
Rückmeldung in JSON:
{"action":"stop","beta":false,"error":
{"code":0},"finished":true,"language":"enu","last_stage":"stopped"
,"package":"AntiVirus","pid":28386,"scripts":
[{"code":0,"message":"","type":"stop"}],"stage":"stopped","status"
:"stop","status_description":"translate from systemd
status","success":true,"username":"","version":"1.5.3-3077"}
Übertragen die die enthaltenden Daten in ein geeignetes Struct in Go und geben Sie anschließend
den Boolean-Wert aus, der den Erfolg anzeigt.
*/
type jsonStructModel struct {
	Action string `json:"action"`
	Beta   bool   `json:"beta"`
	Error  struct {
		Code int `json:"code"`
	} `json:"error"`
	Finished  bool   `json:"finished"`
	Language  string `json:"language"`
	LastStage string `json:"last_stage"`
	Package   string `json:"package"`
	Pid       int    `json:"pid"`
	Scripts   []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"scripts"`
	Stage             string `json:"stage"`
	Status            string `json:"status"`
	StatusDescription string `json:"status_description"`
	Success           bool   `json:"success"`
	Username          string `json:"username"`
	Version           string `json:"version"`
}

func uebung02() {
	fmt.Println("Uebung 2")
	jsonStruct := jsonStructModel{}
	meldung := `{"action":"stop","beta":false,"error":
	{"code":0},"finished":true,"language":"enu","last_stage":"stopped"
	,"package":"AntiVirus","pid":28386,"scripts":
	[{"code":0,"message":"","type":"stop"}],"stage":"stopped","status"
	:"stop","status_description":"translate from systemd status","success":true,"username":"","version":"1.5.3-3077"}`

	err := json.Unmarshal([]byte(meldung), &jsonStruct)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hat wohl geklappt", jsonStruct.Success)
	}
}

/*
	Lektion 202 – Übung 3

Nehmen Sie den folgenden Code als Grundlage: https://go.dev/play/p/EMjjGmbZY4_Z
„Encoden“ sie den Wert vom Typ []user in JSON und senden Sie das Ergebnis an Stdout.
Hilfe: Nutzen Sie json.NewEncoder(os.Stdout).encode(v interface{})
*/
func uebung03() {

	fmt.Println("Uebung 03")

	u1 := user{Name: "Homer", Alter: 36}
	u2 := user{Name: "Marge", Alter: 34}
	u3 := user{Name: "Bart", Alter: 10}
	u4 := user{Name: "Lisa", Alter: 8}
	u5 := user{Name: "Maggie", Alter: 1}

	users := []user{u1, u2, u3, u4, u5}

	json.NewEncoder(os.Stdout).Encode(users)
}

/*
	Lektion 204 – Übung 4

Nehmen Sie den folgenden Code als Grundlage: https://go.dev/play/p/3tBoyBfzyLv
Sortieren die []int und []string.
*/
func uebung04() {
	fmt.Println("Uebung 04")

	i := []int{1, 5, 9, 7, 5, 3, 8, 2, 4, 6, 0, -1, 1024, 23, 42}
	s := []string{"witch", "collapse", "practice", "feed", "shame", "open", "despair", "creek", "road", "again", "ice", "least"}

	fmt.Println("Unsortiert:", i)
	// HIER SOLL IHRE LÖSUNG HIN
	sort.Ints(i)
	fmt.Println("Sortiert:", i)

	fmt.Println("Unsortiert:", s)
	sort.Strings(s)
	// HIER SOLL IHRE LÖSUNG HIN
	fmt.Println("Sortiert:", s)

}

/*
	Sortieren Sie []user nach

Name
Alter
Sortieren Sie jeden []string „Sprüche“ jeden Users aphabetisch.
Geben Sie alles übersichtlich aus.
z.B. ähnlich:
Name, Alter
Spruch1
Spruch2
Spruch3
usw...
*/
func uebung05() {
	fmt.Println("Uebung 05")

	u1 := user{Name: "Homer", Alter: 36, Sprüche: []string{"Doh", "Mmh, Donuts!", "Noch ein Duff!"}}
	u2 := user{Name: "Marge", Alter: 34, Sprüche: []string{"Baaart!", "Alles ist in Ordnung.", "Ich muss zum Frisör."}}
	u3 := user{Name: "Bart", Alter: 10, Sprüche: []string{"Do the Bartman", "Ey, Caramba", "El Barto was here!"}}
	u4 := user{Name: "Lisa", Alter: 8, Sprüche: []string{"Wo ist mein Saxophon?", "Ich bin Vegetarierin", "Welche Uni soll ich wählen?"}}
	u5 := user{Name: "Maggie", Alter: 1, Sprüche: []string{"*Nuckel*", "*Rülps*", "[schweigen]"}}

	fmt.Println("unsorted")
	users := []user{u1, u2, u3, u4, u5}
	fmt.Println(users[1].Sprüche)
	for _, v := range users {
		fmt.Printf("u: %v\n", v.Sprüche)
		sort.Strings(v.Sprüche)
		fmt.Printf("s: %v\n", v.Sprüche)
	}
	
	// Sort by Name
	slices.SortFunc(users, sortByName)
	fmt.Println("Sorted bz Name")
	for _, v := range users {
		printUser(&v)
	}

	// Sort by Age
	slices.SortFunc(users, sortByAge)
	fmt.Println("Sorted bz Age")
	for _, v := range users {
		printUser(&v)
	}
}

func printUser(u *user) {
	fmt.Printf("%s ist %d Jahr(e) alt.\nIhre\\Seine typischen Saetze lauten ", u.Name, u.Alter)
	for _, v := range u.Sprüche {
		fmt.Printf("\"%v\"", v)
	}
	fmt.Println()
}
func sortByName(a, b user) int {
	if n := cmp.Compare(a.Name, b.Name); n != 0 {
		return n
	}
	// If names are equal, order by age
	return cmp.Compare(a.Alter, b.Alter)
}

func sortByAge(a, b user) int {
	return cmp.Compare(a.Alter, b.Alter)
}

