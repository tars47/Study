package main

import (
	"fmt"
	//"math/rand"
)

func init() {
	fmt.Printf("I am in init function \n")
}

func main() {

	/////////////////////////////////////////////////////////////////////////////

	// x := rand.Intn(251)
	// fmt.Printf("the value of x is %v.\n", x)
	// if x >= 0 && x <= 100 {
	// 	fmt.Printf("between 0 and 100.\n")
	// } else if x >= 101 && x <= 200 {
	// 	fmt.Printf("between 101 and 200.\n")
	// } else if x >= 201 && x <= 250 {
	// 	fmt.Printf("between 201 and 250.\n")

	// }

	////////////////////////////////////////////////////////////////////////////////

	// switch {
	// case x >= 0 && x <= 100:
	// 	fmt.Printf("switch - between 0 and 100.\n")
	// case x >= 101 && x <= 200:
	// 	fmt.Printf("switch - between 101 and 200.\n")
	// case x >= 201 && x <= 250:
	// 	fmt.Printf("switch - between 201 and 250.\n")
	// }

	/////////////////////////////////////////////////////////////////////////////

	// for i := 0; i < 100; i++ {
	// 	fmt.Printf("iteration number %v\n", i)
	// }

	/////////////////////////////////////////////////////////////////////////////

	// arr := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)",
	// 	"Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)",
	// 	"Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
	// 	"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
	// 	"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar",
	// 	"Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
	// 	"Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)",
	// 	"Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough",
	// 	"Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different",
	// 	"Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie",
	// 	"Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

	// fmt.Printf("%v \n %T", len(arr), arr)

	/////////////////////////////////////////////////////////////////////////////

	// xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)",
	// 	"Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)",
	// 	"Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
	// 	"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
	// 	"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar",
	// 	"Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
	// 	"Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)",
	// 	"Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough",
	// 	"Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different",
	// 	"Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie",
	// 	"Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

	// for i, v := range xs {
	// 	fmt.Printf("%v - %v \n", i, v)
	// }

	/////////////////////////////////////////////////////////////////////////////

	/*
			    Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.
		        key                 value
		        `bond_james`        `shaken, not stirred`, `martinis`, `fast cars`
		        `moneypenny_jenny`  `intelligence`, `literature`, `computer science`
		        `no_dr`             `cats`, `ice cream`, `sunsets`

	*/
	// mp := map[string][]string{
	// 	`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
	// 	`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
	// 	`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	// }

	// fmt.Println(mp)
	// fmt.Printf("%#v \n\n", mp)

	// for k, v := range mp {
	// 	fmt.Printf("%v \t %v \n", k, v)
	// }

	// delete(mp, "no_dr")

	// for k, v := range mp {
	// 	fmt.Printf("%v \t %v \n", k, v)
	// }

	// ggwfq := map[string]int{}

	// gg := []string{"in", "my", "younger", "and", "more", "vulnerable", "years", "my", "father", "gave", "me", "some",
	// 	"advice", "that", "i’ve", "been", "turning", "over", "in", "my", "mind", "ever", "since.", "whenever",
	// 	"you", "feel", "like", "criticizing", "anyone,", "he", "told", "me,", "just", "remember", "that", "all",
	// 	"the", "people", "in", "this", "world", "haven’t", "had", "the", "advantages", "that", "you’ve",
	// 	"had.", "he", "didn’t", "say", "any", "more,", "but", "we’ve", "always", "been", "unusually",
	// 	"communicative", "in", "a", "reserved", "way,", "and", "i", "understood", "that", "he", "meant",
	// 	"a", "great", "deal", "more", "than", "that.", "in", "consequence,", "i’m", "inclined", "to",
	// 	"reserve", "all", "judgements,", "a", "habit", "that", "has", "opened", "up", "many", "curious",
	// 	"natures", "to", "me", "and", "also", "made", "me", "the", "victim", "of", "not", "a", "few",
	// 	"veteran", "bores.", "the", "abnormal", "mind", "is", "quick", "to", "detect", "and", "attach",
	// 	"itself", "to", "this", "quality", "when", "it", "appears", "in", "a", "normal", "person,", "and", "so",
	// 	"it", "came", "about", "that", "in", "college", "i", "was", "unjustly", "accused", "of", "being", "a",
	// 	"politician,", "because", "i", "was", "privy", "to", "the", "secret", "griefs", "of", "wild,", "unknown",
	// 	"men.", "most", "of", "the", "confidences", "were", "unsought—frequently", "i", "have",
	// 	"feigned", "sleep,", "preoccupation,", "or", "a", "hostile", "levity", "when", "i", "realized", "by",
	// 	"some", "unmistakable", "sign", "that", "an", "intimate", "revelation", "was", "quivering", "on",
	// 	"the", "horizon;", "for", "the", "intimate", "revelations", "of", "young", "men,", "or", "at", "least",
	// 	"the", "terms", "in", "which", "they", "express", "them,", "are", "usually", "plagiaristic", "and",
	// 	"marred", "by", "obvious", "suppressions.", "reserving", "judgements", "is", "a", "matter", "of",
	// 	"infinite", "hope.", "i", "am", "still", "a", "little", "afraid", "of", "missing", "something", "if", "i",
	// 	"forget", "that,", "as", "my", "father", "snobbishly", "suggested,", "and", "i", "snobbishly",
	// 	"repeat,", "a", "sense", "of", "the", "fundamental", "decencies", "is", "parcelled", "out",
	// 	"unequally", "at", "birth.", "and,", "after", "boasting", "this", "way", "of", "my", "tolerance,", "i",
	// 	"come", "to", "the", "admission", "that", "it", "has", "a", "limit.", "conduct", "may", "be",
	// 	"founded", "on", "the", "hard", "rock", "or", "the", "wet", "marshes,", "but", "after", "a", "certain",
	// 	"point", "i", "don’t", "care", "what", "it’s", "founded", "on.", "when", "i", "came", "back", "from",
	// 	"the", "east", "last", "autumn", "i", "felt", "that", "i", "wanted", "the", "world", "to", "be", "in",
	// 	"uniform", "and", "at", "a", "sort", "of", "moral", "attention", "forever;", "i", "wanted", "no",
	// 	"more", "riotous", "excursions", "with", "privileged", "glimpses", "into", "the", "human", "heart.",
	// 	"only", "gatsby,", "the", "man", "who", "gives", "his", "name", "to", "this", "book,", "was",
	// 	"exempt", "from", "my", "reaction—gatsby,", "who", "represented", "everything", "for", "which",
	// 	"i", "have", "an", "unaffected", "scorn.", "if", "personality", "is", "an", "unbroken", "series", "of",
	// 	"successful", "gestures,", "then", "there", "was", "something", "gorgeous", "about", "him,",
	// 	"some", "heightened", "sensitivity", "to", "the", "promises", "of", "life,", "as", "if", "he", "were",
	// 	"related", "to", "one", "of", "those", "intricate", "machines", "that", "register", "earthquakes",
	// 	"ten", "thousand", "miles", "away.", "this", "responsiveness", "had", "nothing", "to", "do", "with",
	// 	"that", "flabby", "impressionability", "which", "is", "dignified", "under", "the", "name", "of", "the",
	// 	"creative", "temperament—it", "was", "an", "extraordinary", "gift", "for", "hope,", "a", "romantic",
	// 	"readiness", "such", "as", "i", "have", "never", "found", "in", "any", "other", "person", "and",
	// 	"which", "it", "is", "not", "likely", "i", "shall", "ever", "find", "again.", "no—gatsby", "turned", "out",
	// 	"all", "right", "at", "the", "end;", "it", "is", "what", "preyed", "on", "gatsby,", "what", "foul", "dust",
	// 	"floated", "in", "the", "wake", "of", "his", "dreams", "that", "temporarily", "closed", "out", "my",
	// 	"interest", "in", "the", "abortive", "sorrows", "and", "short-winded", "elations", "of", "men.", "my",
	// 	"family", "have", "been", "prominent,", "well-to-do", "people", "in", "this", "middle", "western",
	// 	"city", "for", "three", "generations.", "the", "carraways", "are", "something", "of", "a", "clan,",
	// 	"and", "we", "have", "a", "tradition", "that", "we’re", "descended", "from", "the", "dukes", "of",
	// 	"buccleuch,", "but", "the", "actual", "founder", "of", "my", "line", "was", "my", "grandfather’s", "brother,", "who", "came", "here", "in", "fifty-one,", "sent", "a", "substitute", "to", "the", "civil",
	// 	"war,", "and", "started", "the", "wholesale", "hardware", "business", "that", "my", "father",
	// 	"carries", "on", "today.", "i", "never", "saw", "this", "great-uncle,", "but", "i’m", "supposed", "to",
	// 	"look", "like", "him—with", "special", "reference", "to", "the", "rather", "hard-boiled", "painting",
	// 	"that", "hangs", "in", "father’s", "office.", "i", "graduated", "from", "new", "haven", "in", "1915,",
	// 	"just", "a", "quarter", "of", "a", "century", "after", "my", "father,", "and", "a", "little", "later", "i",
	// 	"participated", "in", "that", "delayed", "teutonic", "migration", "known", "as", "the", "great",
	// 	"war.", "i", "enjoyed", "the", "counter-raid", "so", "thoroughly", "that", "i", "came", "back",
	// 	"restless.", "instead", "of", "being", "the", "warm", "centre", "of", "the", "world,", "the", "middle",
	// 	"west", "now", "seemed", "like", "the", "ragged", "edge", "of", "the", "universe—so", "i",
	// 	"decided", "to", "go", "east", "and", "learn", "the", "bond", "business.", "everybody", "i", "knew",
	// 	"was", "in", "the", "bond", "business,", "so", "i", "supposed", "it", "could", "support", "one",
	// 	"more", "single", "man.", "all", "my", "aunts", "and", "uncles", "talked", "it", "over", "as", "if",
	// 	"they", "were", "choosing", "a", "prep", "school", "for", "me,", "and", "finally", "said,",
	// 	"why—ye-es,", "with", "very", "grave,", "hesitant", "faces.", "father", "agreed", "to", "finance",
	// 	"me", "for", "a", "year,", "and", "after", "various", "delays", "i", "came", "east,", "permanently,",
	// 	"i", "thought,", "in", "the", "spring", "of", "twenty-two.", "the", "practical", "thing", "was", "to",
	// 	"find", "rooms", "in", "the", "city,", "but", "it", "was", "a", "warm", "season,", "and", "i", "had",
	// 	"just", "left", "a", "country", "of", "wide", "lawns", "and", "friendly", "trees,", "so", "when", "a",
	// 	"young", "man", "at", "the", "office", "suggested", "that", "we", "take", "a", "house", "together",
	// 	"in", "a", "commuting", "town,", "it", "sounded", "like", "a", "great", "idea.", "he", "found", "the",
	// 	"house,", "a", "weather-beaten", "cardboard", "bungalow", "at", "eighty", "a", "month,", "but",
	// 	"at", "the", "last", "minute", "the", "firm", "ordered", "him", "to", "washington,", "and", "i", "went",
	// 	"out", "to", "the", "country", "alone.", "i", "had", "a", "dog—at", "least", "i", "had", "him", "for", "a",
	// 	"few", "days", "until", "he", "ran", "away—and", "an", "old", "dodge", "and", "a", "finnish",
	// 	"woman,", "who", "made", "my", "bed", "and", "cooked", "breakfast", "and", "muttered",
	// 	"finnish", "wisdom", "to", "herself", "over", "the", "electric", "stove.", "it", "was", "lonely", "for",
	// 	"a", "day", "or", "so", "until", "one", "morning", "some", "man,", "more", "recently", "arrived",
	// 	"than", "i,", "stopped", "me", "on", "the", "road.", "how", "do", "you", "get", "to", "west", "egg",
	// 	"village?", "he", "asked", "helplessly.", "i", "told", "him.", "and", "as", "i", "walked", "on", "i",
	// 	"was", "lonely", "no", "longer.", "i", "was", "a", "guide,", "a", "pathfinder,", "an", "original",
	// 	"settler.", "he", "had", "casually", "conferred", "on", "me", "the", "freedom", "of", "the",
	// 	"neighbourhood.", "and", "so", "with", "the", "sunshine", "and", "the", "great", "bursts", "of",
	// 	"leaves", "growing", "on", "the", "trees,", "just", "as", "things", "grow", "in", "fast", "movies,", "i",
	// 	"had", "that", "familiar", "conviction", "that", "life", "was", "beginning", "over", "again", "with",
	// 	"the", "summer.", "there", "was", "so", "much", "to", "read,", "for", "one", "thing,", "and", "so",
	// 	"much", "fine", "health", "to", "be", "pulled", "down", "out", "of", "the", "young", "breath-giving",
	// 	"air.", "i", "bought", "a", "dozen", "volumes", "on", "banking", "and", "credit", "and",
	// 	"investment", "securities,", "and", "they", "stood", "on", "my", "shelf", "in", "red"}

	// for _, v := range gg {
	// 	ggwfq[v]++
	// }

	// //fmt.Println(ggwfq)
	// max := 0
	// key := ""
	// for k, v := range ggwfq {
	// 	if v > max {
	// 		max = v
	// 		key = k
	// 	}
	// }
	// fmt.Println(max)
	// fmt.Printf("the most frequently occured word is '%v' with count value of %v", key, ggwfq[key])

	////////////////////////////////////////////////////////////////////////////////

	/*
			Create your own type “person” which will have an underlying type of “struct” so that it can store
		the following data:
		● first name
		● last name
		● favorite ice cream flavors
		Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
		which stores the favorite flavors.

	*/

	// type Person struct {
	// 	firstName                 string
	// 	lastName                  string
	// 	favouriteIceCreamFlavours []string
	// }

	// p1 := Person{
	// 	firstName:                 "sudeep",
	// 	lastName:                  "cd",
	// 	favouriteIceCreamFlavours: []string{"venilla", "butterscotch"},
	// }

	// p2 := Person{
	// 	firstName:                 "sandhya",
	// 	lastName:                  "fp",
	// 	favouriteIceCreamFlavours: []string{"strawberry", "chocolate"},
	// }

	// for _, v := range p1.favouriteIceCreamFlavours {
	// 	fmt.Printf("%v %v likes %v \n", p1.firstName, p1.lastName, v)
	// }
	// for _, v := range p2.favouriteIceCreamFlavours {
	// 	fmt.Printf("%v %v likes %v \n", p2.firstName, p2.lastName, v)
	// }

	/*
			Take the code from the previous exercise, then store the VALUES of type person in a map with
		the KEY of last name. Access each value in the map. Print out the values, ranging over the
		slice.

	*/

	// mp := map[string]Person{
	// 	p1.lastName: p1,
	// 	p2.lastName: p2,
	// }

	// for _, v := range mp {
	// 	for _, v2 := range v.favouriteIceCreamFlavours {
	// 		fmt.Printf("%v %v likes %v \n", v.firstName, v.lastName, v2)
	// 	}
	// }

	/*
		Create a type engine struct, and include this field
		○ electric bool
		● Create a type vehicle struct, and include these fields
		■ engine
		■ make
		■ model
		■ doors
		■ color
		● Create two VALUES of TYPE vehicle
		○ use a composite literal
		● Print out each of these values.
		● Print out a single field from each of these values.
	*/

	// type engine struct {
	// 	electric bool
	// }
	// type vehicle struct {
	// 	engine
	// 	make  string
	// 	model string
	// 	doors int
	// 	color string
	// }

	// v1 := vehicle{
	// 	engine: engine{electric: false},
	// 	make:   "tata",
	// 	model:  "punch",
	// 	doors:  4,
	// 	color:  "camo",
	// }

	// v2 := vehicle{
	// 	engine: engine{electric: true},
	// 	make:   "kia",
	// 	model:  "sonnet",
	// 	doors:  4,
	// 	color:  "black",
	// }

	// fmt.Println(v1)
	// fmt.Println(v2)

	// fmt.Println(v1.electric, v1.make, v1.model)
	// fmt.Println(v2.electric, v2.make, v2.model)

	// fmt.Println(v1.engine.electric, v1.make, v1.model)
	// fmt.Println(v2.engine.electric, v2.make, v2.model)

	/*
			Create and use an anonymous struct with these fields:
		● first string
		● friends map[string]int
		● favDrinks []string
		Print things.
	*/

	// as := struct {
	// 	first     string
	// 	friends   map[string]int
	// 	favDrinks []string
	// }{
	// 	first: "sudeep",
	// 	friends: map[string]int{
	// 		"Jenny": 27,
	// 		"Q":     87,
	// 		"Ian":   47,
	// 	},
	// 	favDrinks: []string{
	// 		"Martini",
	// 		"Water",
	// 	},
	// }

	// for k, v := range as.friends {
	// 	fmt.Println(as.first, "- friends - ", k, v)
	// }

	// for _, v := range as.favDrinks {
	// 	fmt.Println(as.first, "- drinks - ", v)
	// }

	// sa := secretAgent{
	// 	person: person{first: "sudeep"},
	// 	ltk:    false,
	// }

	// sa.speak()
	//speak(sa)

	// val := 43
	// ptr := &val
	// fmt.Println(*ptr)

	// p1 := person{"jerry"}
	// ptr := &p1
	// fmt.Printf("Address of struct = %+v: %p\n", p1, ptr)
	// speak(p1)
	// fmt.Println(p1)

	// xs := []int{1, 2, 3, 4}
	// fmt.Printf("Address of slice = %p: %p\n", xs, &xs)
	// modify(xs)
	// fmt.Println(xs)

	////////////////////////////////////////////////////////////////////////////

	// u1 := user{
	// 	First: "James",
	// 	Age:   32,
	// }

	// u2 := user{
	// 	First: "Moneypenny",
	// 	Age:   27,
	// }

	// u3 := user{
	// 	First: "M",
	// 	Age:   54,
	// }

	// users := []user{u1, u2, u3}

	// fmt.Println(users)
	// b, err := json.Marshal(users)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// os.Stdout.Write(b)

	// s := `[{"First":"James",
	//         "Last":"Bond",
	// 		"Age":32,
	// 		"Sayings":["Shaken, not stirred",
	// 		           "Youth is no guarantee of innovation",
	// 				   "In his majesty's royal service"]},
	// 		{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	// fmt.Println(s)

	// var su []user
	// err := json.Unmarshal([]byte(s), &su)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// fmt.Println(su)

}

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// func modify(x []int) {
// 	fmt.Printf("Address of slice in func = %p: %p\n", x, &x)
// 	x[0] = 0
// }

// type human interface {
// 	speak()
// }

// type person struct {
// 	first string
// }

// type secretAgent struct {
// 	person
// 	ltk bool
// }

// func speak(p person) {
// 	fmt.Printf("Address of struct = %+v: %p\n", p, &p)
// 	p.first = "rock"
// 	fmt.Println("I am", p.first)
// }

// func speak(h human) {
// 	h.speak()
// }
