package main

import (
	"bufio"
	models "challenge/models"
	print "challenge/print"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var users []models.User
var organizations []models.Organization
var tickets []models.Ticket

func readFile() {
	jsonFileUser, errUser := os.Open("./data/users.json")
	if errUser != nil {
		fmt.Println(errUser)
	}
	byteValueUser, _ := ioutil.ReadAll(jsonFileUser)
	json.Unmarshal(byteValueUser, &users)
	jsonFileOrganization, errOrganization := os.Open("./data/organizations.json")
	if errOrganization != nil {
		fmt.Println(errOrganization)
	}
	byteValueOrganization, _ := ioutil.ReadAll(jsonFileOrganization)
	json.Unmarshal(byteValueOrganization, &organizations)
	jsonFileTicket, errTicket := os.Open("./data/tickets.json")
	if errTicket != nil {
		fmt.Println(errTicket)
	}
	byteValueTicket, _ := ioutil.ReadAll(jsonFileTicket)
	json.Unmarshal(byteValueTicket, &tickets)
	defer jsonFileUser.Close()
	defer jsonFileOrganization.Close()
	defer jsonFileTicket.Close()
}

func inputSearch() (string, string) {
	var term string
	var value string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter search term \n")
	if scanner.Scan() {
		term = scanner.Text()
	}
	fmt.Printf("Enter search value \n")
	if scanner.Scan() {
		value = scanner.Text()
	}
	return term, value
}

func printOrganizationAndTicketFromUser(i int) {
	for k := 0; k < len(organizations); k++ {
		if organizations[k].Id == users[i].OrganizationId {
			fmt.Printf("%30s ", "organization_name")
			fmt.Printf("\t\t\t")
			fmt.Printf("%s ", organizations[k].Name)
			fmt.Printf("\n")
		}
	}
	t := 0
	for k := 0; k < len(tickets); k++ {
		if tickets[k].SubmitterId == users[i].Id {
			fmt.Printf("%30s ", "ticket_"+strconv.Itoa(t))
			fmt.Printf("\t\t\t")
			fmt.Printf("%s ", tickets[k].Subject)
			fmt.Printf("\n")
			t = t + 1
		}
	}
}

func printUserAndTicketFromOrganization(i int) {
	for k := 0; k < len(users); k++ {
		if users[k].OrganizationId == organizations[i].Id {
			fmt.Printf("%30s ", "user_name")
			fmt.Printf("\t\t\t")
			fmt.Printf("%s ", users[k].Name)
			fmt.Printf("\n")
		}
	}
	t := 0
	for k := 0; k < len(tickets); k++ {
		if tickets[k].OrganizationId == organizations[i].Id {
			fmt.Printf("%30s ", "ticket_"+strconv.Itoa(t))
			fmt.Printf("\t\t\t")
			fmt.Printf("%s ", tickets[k].Subject)
			fmt.Printf("\n")
			t = t + 1
		}
	}
}

func printUserAndOrganizationFormTicket(i int) {
	for k := 0; k < len(organizations); k++ {
		if organizations[k].Id == tickets[i].OrganizationId {
			fmt.Printf("%30s ", "organization_name")
			fmt.Printf("\t\t\t")
			fmt.Printf("%s ", organizations[k].Name)
			fmt.Printf("\n")
		}
	}
	t := 0
	for k := 0; k < len(users); k++ {
		if users[k].Id == tickets[i].SubmitterId {
			fmt.Printf("%30s ", "user_"+strconv.Itoa(t))
			fmt.Printf("\t\t\t")
			fmt.Printf("%s ", users[k].Name)
			fmt.Printf("\n")
			t = t + 1
		}
	}
}

func searchUserMany(term string, value string) string {

	fmt.Println("")
	fmt.Println("Search user for " + term + " with a value of " + value)
	term = getFieldName(term, "json", models.User{})
	if term == "" {
		fmt.Println("Term not exists at user")
		return fmt.Sprintf("Term not exists at user")
	}
	var usersSearch []models.User
	var compareValue int
	var compareBool bool
	var errCompare error

	r := reflect.ValueOf(users[0])
	f := reflect.Indirect(r).FieldByName(term)
	if f.String() == "<int Value>" {
		compareValue, errCompare = strconv.Atoi(value)
		if errCompare != nil {
			fmt.Println("Wrong value type")
			return fmt.Sprintf("Wrong value type")
		}
	} else if f.String() == "<bool Value>" {
		compareBool, errCompare = strconv.ParseBool(value)
		if errCompare != nil {
			fmt.Println("Wrong value type")
			return fmt.Sprintf("Wrong value type")
		}
	}
	lenUsers := len(users)
	var loop float64
	loop = float64(lenUsers / 100)
	loopInt := int(math.Ceil(loop))
	if loopInt*100 < len(users) {
		loopInt += 1
	}
	var wg sync.WaitGroup

	for i := 0; i < loopInt; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if i < (loopInt - 1) {
				for j := 100 * i; j < 100*i+99; j++ {
					r := reflect.ValueOf(users[j])
					f := reflect.Indirect(r).FieldByName(term)
					if f.String() == "<int Value>" {
						if int(f.Int()) == compareValue {
							usersSearch = append(usersSearch, users[j])
							print.PrintlnUser(users[j])
							printOrganizationAndTicketFromUser(j)
						}
					} else if f.String() == "<bool Value>" {
						if bool(f.Bool()) == compareBool {
							usersSearch = append(usersSearch, users[j])
							print.PrintlnUser(users[j])
							printOrganizationAndTicketFromUser(j)
						}
					} else {
						s := f.Interface().(string)
						if s == value {
							usersSearch = append(usersSearch, users[j])
							print.PrintlnUser(users[j])
							printOrganizationAndTicketFromUser(j)
						}
					}
				}
			} else {
				for j := 100 * i; j < len(users); j++ {
					r := reflect.ValueOf(users[j])
					f := reflect.Indirect(r).FieldByName(term)
					if f.String() == "<int Value>" {

						if int(f.Int()) == compareValue {
							usersSearch = append(usersSearch, users[j])
							print.PrintlnUser(users[j])
							printOrganizationAndTicketFromUser(j)
						}
					} else if f.String() == "<bool Value>" {
						if bool(f.Bool()) == compareBool {
							usersSearch = append(usersSearch, users[j])
							print.PrintlnUser(users[j])
							printOrganizationAndTicketFromUser(j)
						}
					} else {
						s := f.Interface().(string)
						if s == value {
							usersSearch = append(usersSearch, users[j])
							print.PrintlnUser(users[j])
							printOrganizationAndTicketFromUser(j)
						}
					}
				}
			}
		}()
		wg.Wait()

	}

	if len(usersSearch) == 0 {
		fmt.Println("No results found")
		return fmt.Sprintf("No results found")
	} else {
		return fmt.Sprintf("Have results")
	}
}

func searchOrganizationMany(term string, value string) string {

	fmt.Println("")
	fmt.Println("Search organization for " + term + " with a value of " + value)
	term = getFieldName(term, "json", models.Organization{})
	if term == "" {
		fmt.Println("Term not exists at organization")
		return fmt.Sprintf("Term not exists at organization")
	}
	var organizationsSearch []models.Organization
	var compareValue int
	var compareBool bool
	var errCompare error

	r := reflect.ValueOf(organizations[0])
	f := reflect.Indirect(r).FieldByName(term)
	if f.String() == "<int Value>" {
		compareValue, errCompare = strconv.Atoi(value)
		if errCompare != nil {
			fmt.Println("Wrong value type")
			return fmt.Sprintf("Wrong value type")
		}
	} else if f.String() == "<bool Value>" {
		compareBool, errCompare = strconv.ParseBool(value)
		if errCompare != nil {
			fmt.Println("Wrong value type")
			return fmt.Sprintf("Wrong value type")
		}
	}
	lenOrganizations := len(organizations)
	var loop float64
	loop = float64(lenOrganizations / 100)
	loopInt := int(math.Ceil(loop))
	if loopInt*100 < len(organizations) {
		loopInt += 1
	}
	var wg sync.WaitGroup

	for i := 0; i < loopInt; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if i < (loopInt - 1) {
				for j := 100 * i; j < 100*i+99; j++ {
					r := reflect.ValueOf(organizations[j])
					f := reflect.Indirect(r).FieldByName(term)
					if f.String() == "<int Value>" {
						if int(f.Int()) == compareValue {
							organizationsSearch = append(organizationsSearch, organizations[j])
							print.PrintlnOrganization(organizations[j])
							printUserAndTicketFromOrganization(j)
						}
					} else if f.String() == "<bool Value>" {
						if bool(f.Bool()) == compareBool {
							organizationsSearch = append(organizationsSearch, organizations[j])
							print.PrintlnOrganization(organizations[j])
							printUserAndTicketFromOrganization(j)
						}
					} else {
						s := f.Interface().(string)
						if s == value {
							organizationsSearch = append(organizationsSearch, organizations[j])
							print.PrintlnOrganization(organizations[j])
							printUserAndTicketFromOrganization(j)
						}
					}
				}
			} else {
				for j := 100 * i; j < len(organizations); j++ {
					r := reflect.ValueOf(organizations[j])
					f := reflect.Indirect(r).FieldByName(term)
					if f.String() == "<int Value>" {
						if int(f.Int()) == compareValue {
							organizationsSearch = append(organizationsSearch, organizations[j])
							print.PrintlnOrganization(organizations[j])
							printUserAndTicketFromOrganization(j)
						}
					} else if f.String() == "<bool Value>" {
						if bool(f.Bool()) == compareBool {
							organizationsSearch = append(organizationsSearch, organizations[j])
							print.PrintlnOrganization(organizations[j])
							printUserAndTicketFromOrganization(j)
						}
					} else {
						s := f.Interface().(string)
						if s == value {
							organizationsSearch = append(organizationsSearch, organizations[j])
							print.PrintlnOrganization(organizations[j])
							printUserAndTicketFromOrganization(j)
						}
					}
				}
			}
		}()
		wg.Wait()

	}

	if len(organizationsSearch) == 0 {
		fmt.Println("No results found")
		return fmt.Sprintf("No results found")
	} else {
		return fmt.Sprintf("Have results")
	}
}

func searchTicketMany(term string, value string) string {

	fmt.Println("")
	fmt.Println("Search ticket for " + term + " with a value of " + value)
	term = getFieldName(term, "json", models.Ticket{})
	if term == "" {
		fmt.Println("Term not exists at ticket")
		return fmt.Sprintf("Term not exists at ticket")
	}
	var ticketsSearch []models.Ticket
	var compareValue int
	var compareBool bool
	var errCompare error

	r := reflect.ValueOf(tickets[0])
	f := reflect.Indirect(r).FieldByName(term)
	if f.String() == "<int Value>" {
		compareValue, errCompare = strconv.Atoi(value)
		if errCompare != nil {
			fmt.Println("Wrong value type")
			return fmt.Sprintf("Wrong value type")
		}
	} else if f.String() == "<bool Value>" {
		compareBool, errCompare = strconv.ParseBool(value)
		if errCompare != nil {
			fmt.Println("Wrong value type")
			return fmt.Sprintf("Wrong value type")
		}
	}
	lenTickets := len(tickets)

	var loop float64
	loop = float64(lenTickets / 100)
	loopInt := int(math.Ceil(loop))

	if loopInt*100 < len(tickets) {

		loopInt += 1
	}
	var wg sync.WaitGroup

	for i := 0; i < loopInt; i++ {
		wg.Add(1)
		go func() {

			defer wg.Done()
			if i < (loopInt - 1) {

				for j := 100 * i; j < 100*i+99; j++ {
					r := reflect.ValueOf(tickets[j])
					f := reflect.Indirect(r).FieldByName(term)
					if f.String() == "<int Value>" {
						if int(f.Int()) == compareValue {
							ticketsSearch = append(ticketsSearch, tickets[j])
							print.PrintlnTicket(tickets[j])
							printUserAndOrganizationFormTicket(j)
						}
					} else if f.String() == "<bool Value>" {
						if bool(f.Bool()) == compareBool {
							ticketsSearch = append(ticketsSearch, tickets[j])
							print.PrintlnTicket(tickets[j])
							printUserAndOrganizationFormTicket(j)
						}
					} else {
						s := f.Interface().(string)
						if s == value {
							ticketsSearch = append(ticketsSearch, tickets[j])
							print.PrintlnTicket(tickets[j])
							printUserAndOrganizationFormTicket(j)
						}
					}
				}
			} else {
				for j := 100 * i; j < len(tickets); j++ {
					r := reflect.ValueOf(tickets[j])
					f := reflect.Indirect(r).FieldByName(term)
					if f.String() == "<int Value>" {
						if int(f.Int()) == compareValue {
							ticketsSearch = append(ticketsSearch, tickets[j])
							print.PrintlnTicket(tickets[j])
							printUserAndOrganizationFormTicket(j)
						}
					} else if f.String() == "<bool Value>" {
						if bool(f.Bool()) == compareBool {
							ticketsSearch = append(ticketsSearch, tickets[j])
							print.PrintlnTicket(tickets[j])
							printUserAndOrganizationFormTicket(j)
						}
					} else {
						s := f.Interface().(string)
						if s == value {
							ticketsSearch = append(ticketsSearch, tickets[j])
							print.PrintlnTicket(tickets[j])
							printUserAndOrganizationFormTicket(j)
						}
					}
				}
			}
		}()
		wg.Wait()

	}

	if len(ticketsSearch) == 0 {
		fmt.Println("No results found")
		return fmt.Sprintf("No results found")
	} else {
		return fmt.Sprintf("Have results")
	}
}

func searchUser(term string, value string) string {
	//term, value:=inputSearch()
	fmt.Println("")
	fmt.Println("Search user for " + term + " with a value of " + value)
	term = getFieldName(term, "json", models.User{})
	if term == "" {
		fmt.Println("Term not exists at user")
		return fmt.Sprintf("Term not exists at user")
	}
	var usersSearch []models.User
	var compareValue int
	var compareBool bool
	var errCompare error
	for i := 0; i < len(users); i++ {
		r := reflect.ValueOf(users[i])

		f := reflect.Indirect(r).FieldByName(term)

		if f.String() == "<int Value>" {
			if i == 0 {
				compareValue, errCompare = strconv.Atoi(value)
				if errCompare != nil {
					fmt.Println("Wrong value type")
					return fmt.Sprintf("Wrong value type")
				}
			}
			if int(f.Int()) == compareValue {
				usersSearch = append(usersSearch, users[i])
				print.PrintlnUser(users[i])
				printOrganizationAndTicketFromUser(i)

			}
		} else if f.String() == "<bool Value>" {
			if i == 0 {
				compareBool, errCompare = strconv.ParseBool(value)
				if errCompare != nil {
					fmt.Println("Wrong value type")
					return fmt.Sprintf("Wrong value type")
				}
			}
			if bool(f.Bool()) == compareBool {
				usersSearch = append(usersSearch, users[i])
				print.PrintlnUser(users[i])
				printOrganizationAndTicketFromUser(i)
			}
		} else {
			s := f.Interface().(string)
			if s == value {
				usersSearch = append(usersSearch, users[i])
				print.PrintlnUser(users[i])
				printOrganizationAndTicketFromUser(i)
			}
		}
	}
	if len(usersSearch) == 0 {
		fmt.Println("No results found")
		return fmt.Sprintf("No results found")
	} else {
		return fmt.Sprintf("Have results")
	}
}

func getFieldName(tag, key string, s interface{}) (fieldname string) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		panic("bad type")
	}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get(key), ",")[0] // use split to ignore tag "options"
		if v == tag {
			return f.Name
		}
	}
	return ""
}

func searchOrganization(term string, value string) string {
	//term, value:=inputSearch()
	fmt.Println("")
	fmt.Println("Search organization for " + term + " with a value of " + value)
	term = getFieldName(term, "json", models.Organization{})
	if term == "" {
		fmt.Println("Term not exists at organization")
		return fmt.Sprintf("Term not exists at organization")
	}
	var organizationSearch []models.Organization
	var compareValue int
	var compareBool bool
	var errCompare error
	for i := 0; i < len(organizations); i++ {
		r := reflect.ValueOf(organizations[i])

		f := reflect.Indirect(r).FieldByName(term)

		if f.String() == "<int Value>" {
			if i == 0 {
				compareValue, errCompare = strconv.Atoi(value)
				if errCompare != nil {
					fmt.Println("Wrong value type")
					return fmt.Sprintf("Wrong value type")
				}
			}
			if int(f.Int()) == compareValue {
				organizationSearch = append(organizationSearch, organizations[i])
				print.PrintlnOrganization(organizations[i])
				printUserAndTicketFromOrganization(i)

			}
		} else if f.String() == "<bool Value>" {
			if i == 0 {
				compareBool, errCompare = strconv.ParseBool(value)
				if errCompare != nil {
					fmt.Println("Wrong value type")
					return fmt.Sprintf("Wrong value type")
				}
			}
			if bool(f.Bool()) == compareBool {
				organizationSearch = append(organizationSearch, organizations[i])
				print.PrintlnOrganization(organizations[i])
				printUserAndTicketFromOrganization(i)
			}
		} else {
			s := f.Interface().(string)
			if s == value {
				organizationSearch = append(organizationSearch, organizations[i])
				print.PrintlnOrganization(organizations[i])
				printUserAndTicketFromOrganization(i)
			}
		}
	}
	if len(organizationSearch) == 0 {
		fmt.Println("No results found")
		return fmt.Sprintf("No results found")
	} else {
		return fmt.Sprintf("Have results")
	}

}

func searchTickets(term string, value string) string {
	//term, value:=inputSearch()
	fmt.Println("")
	fmt.Println("Search ticket for " + term + " with a value of " + value)
	term = getFieldName(term, "json", models.Ticket{})
	if term == "" {
		fmt.Println("Term not exists at ticket")
		return fmt.Sprintf("Term not exists at ticket")
	}
	var ticketSearch []models.Ticket
	var compareValue int
	var compareBool bool
	var errCompare error
	for i := 0; i < len(tickets); i++ {
		r := reflect.ValueOf(tickets[i])

		f := reflect.Indirect(r).FieldByName(term)

		if f.String() == "<int Value>" {
			if i == 0 {
				compareValue, errCompare = strconv.Atoi(value)
				if errCompare != nil {
					fmt.Println("Wrong value type")
					return fmt.Sprintf("Wrong value type")
				}
			}
			if int(f.Int()) == compareValue {
				ticketSearch = append(ticketSearch, tickets[i])
				print.PrintlnTicket(tickets[i])
				printUserAndOrganizationFormTicket(i)
			}
		} else if f.String() == "<bool Value>" {
			if i == 0 {
				compareBool, errCompare = strconv.ParseBool(value)
				if errCompare != nil {
					fmt.Println("Wrong value type")
					return fmt.Sprintf("Wrong value type")
				}
			}
			if bool(f.Bool()) == compareBool {
				ticketSearch = append(ticketSearch, tickets[i])
				print.PrintlnTicket(tickets[i])
				printUserAndOrganizationFormTicket(i)
			}
		} else {
			s := f.Interface().(string)
			if s == value {
				ticketSearch = append(ticketSearch, tickets[i])
				print.PrintlnTicket(tickets[i])
				printUserAndOrganizationFormTicket(i)
			}
		}
	}

	if len(ticketSearch) == 0 {
		fmt.Println("No results found")
		return fmt.Sprintf("No results found")
	} else {
		return fmt.Sprintf("Have results")
	}

}

func viewListSearchable() {
	var user models.User
	var ticket models.Ticket
	var organization models.Organization
	u := reflect.ValueOf(&user).Elem()
	typeOfU := u.Type()
	fmt.Println("-------------------------------------------------")
	fmt.Println("Search Users with")
	for i := 0; i < u.NumField(); i++ {

		fmt.Println(typeOfU.Field(i).Name)
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println("Search Tickets With")
	t := reflect.ValueOf(&ticket).Elem()
	typeOfT := t.Type()
	for i := 0; i < t.NumField(); i++ {

		fmt.Println(typeOfT.Field(i).Name)
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println("Search Organization With")
	o := reflect.ValueOf(&organization).Elem()
	typeOfO := o.Type()
	for i := 0; i < o.NumField(); i++ {

		fmt.Println(typeOfO.Field(i).Name)
	}

}

func search() {
	var j int
	fmt.Println("Select 1) Users or 2) Tickets or 3) Organizations")
	_, err := fmt.Scanf("%d", &j)
	if err != nil {
		fmt.Println("Error Input", err)
	}
	switch j {
	case 1:
		searchUserMany(inputSearch())
		break
	case 3:
		searchOrganizationMany(inputSearch())
		break
	case 2:
		searchTicketMany(inputSearch())
		break
	default:
		break
	}

}

func main() {
	var i int
	readFile()
	for {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Select Search Options:")
		fmt.Println("* Press 1 to search")
		fmt.Println("* Press 2 to view a list of searchable fields")
		fmt.Println("* Type 'quit' to exit")
		fmt.Println("")
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Println("Error Input", err)
			break
		}
		switch i {
		case 1:
			search()
			break
		case 2:
			viewListSearchable()
			break
		default:
			break
		}

		fmt.Println("----------------------------------------------------")
	}
}
