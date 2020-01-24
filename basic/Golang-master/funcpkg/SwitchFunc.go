package funcpkg

import "fmt"

// SwitchFunc switch文のサンプルです
func SwitchFunc(skill string) {

	switch skill {
	case "PHP":
		fmt.Println("My Skill is PHP")
	case "Golang":
		fmt.Println("My Skill is Golang")
	case "Java":
		fmt.Println("My Skill is Java")
	case "Ruby":
		fmt.Println("My Skill is Ruby")
		fallthrough
	case "Rails":
		fmt.Println("My Skill is Ruby on Rails")
	case "JavaScript", "HTML", "CSS":
		fmt.Println("My Skill is Web")
	default:
		fmt.Println("My Skill is Nothing")
	}
}