package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func ModelComputerClub(inputFile string) string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("can't open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanLines)

	var numtables int

	if scanner.Scan() {
		n, _ := fmt.Sscanf(scanner.Text(), "%d", &numtables)
		if n == 0 {
			return scanner.Text()
		}
	} else {
		log.Fatal("Number of tables not provided")
	}

	var openingHours, openingMinutes, closingHours, closingMinutes, openingTime, closingTime int
	if scanner.Scan() {
		re, err := regexp.Compile("^([0-1][0-9]|2[0-3]):[0-5][0-9] ([0-1][0-9]|2[0-3]):[0-5][0-9]$")
		if err != nil {
			log.Fatal(err)
		}
		if !re.MatchString(scanner.Text()) {
			return scanner.Text()
		}

		fmt.Sscanf(scanner.Text(), "%d:%d %d:%d", &openingHours, &openingMinutes, &closingHours, &closingMinutes)
		openingTime = openingHours*60 + openingMinutes
		closingTime = closingHours*60 + closingMinutes
	} else {
		log.Fatal("failed to scan working hours", scanner.Text())
	}

	var cost int
	if scanner.Scan() {
		n, _ := fmt.Sscanf(scanner.Text(), "%d", &cost)
		if n == 0 {
			return scanner.Text()
		}
	} else {
		log.Fatal("failed to scan cost", scanner.Text())
	}

	clients := make(map[string]int)
	queue := ClientQueue{make([]string, 0, 1), 0}
	computers := make([]Computer, numtables)
	freeTables := numtables

	for ind := range computers {
		computers[ind] = Computer{
			isOccupied: false,
			user:       "",
			since:      0,
			cost:       cost,
		}
	}

	var result strings.Builder
	result.WriteString(fmt.Sprintf("%02d:%02d\n", openingHours, openingMinutes))

	var hours, minutes, eventId int
	var body string
	for scanner.Scan() {

		re, err := regexp.Compile("^([0-1][0-9]|2[0-3]):[0-5][0-9] ([13-4] [a-z1-9_-]+|2 [a-z1-9_-]+ [0-9]+)$")
		if err != nil {
			log.Fatal(err)
		}

		if !re.MatchString(scanner.Text()) {
			return scanner.Text()
		}

		fmt.Sscanf(scanner.Text(), "%d:%d %d %s", &hours, &minutes, &eventId, &body)

		result.WriteString(scanner.Text())
		result.WriteString("\n")
		now := hours*60 + minutes

		switch eventId {
		case 1:
			if _, ok := clients[body]; ok {
				result.WriteString(fmt.Sprintf("%02d:%02d 13 YouShallNotPass\n", hours, minutes))
				continue
			}

			if now < openingTime || now > closingTime { //TODO: implement for working hours like 06:00-01:00
				result.WriteString(fmt.Sprintf("%02d:%02d 13 NotOpenYet\n", hours, minutes))
				continue
			}

			clients[body] = 0
			queue.Push(body)
		case 2:
			var user string
			var pc int
			fmt.Sscanf(scanner.Text(), "%d:%d %d %s %d", &hours, &minutes, &eventId, &user, &pc)

			if pc < 1 || pc > numtables {
				log.Fatal("pc out of tables range", scanner.Text())
			}

			if computers[pc-1].isOccupied {
				result.WriteString(fmt.Sprintf("%02d:%02d 13 PlaceIsBusy\n", hours, minutes))
				continue
			}

			if _, ok := clients[user]; !ok {
				result.WriteString(fmt.Sprintf("%02d:%02d 13 ClientUnknown\n", hours, minutes))
				continue
			}

			if clients[user] != 0 {
				computers[clients[user]-1].Free(now)
			} else {
				queue.Remove(user)
			}
			clients[user] = pc
			computers[pc-1].Occupy(user, now)
			freeTables--
		case 3:
			if freeTables > 0 {
				result.WriteString(fmt.Sprintf("%02d:%02d 13 ICanWaitNoLonger!\n", hours, minutes))
				continue
			}

			if queue.size > numtables {
				queue.Remove(body)
				delete(clients, body)
				result.WriteString(fmt.Sprintf("%02d:%02d 11 %s", hours, minutes, body))
				continue
			}
		case 4:
			if _, ok := clients[body]; !ok {
				result.WriteString(fmt.Sprintf("%02d:%02d 13 ClientUnknown\n", hours, minutes))
				continue
			}

			freeTables++

			if clients[body] != 0 {
				computers[clients[body]-1].Free(now)
				var client string
				if !queue.IsEmpty() {
					client, err = queue.Pop()
					computers[clients[body]-1].Occupy(client, now)
					clients[client] = clients[body]
					result.WriteString(fmt.Sprintf("%02d:%02d 12 %s %d\n", hours, minutes, client, clients[body]))
					freeTables--
				}
				if err != nil {
					log.Fatal(err)
				}
			}

			delete(clients, body)
		default:
			log.Fatal("event id not found", scanner.Text())
		}
	}

	i := 0
	keys := make([]string, len(clients))
	for k := range clients {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	for _, user := range keys {
		freeTables++

		if clients[user] != 0 {
			computers[clients[user]-1].Free(closingTime)
		}

		result.WriteString(fmt.Sprintf("%02d:%02d 11 %s \n", closingHours, closingMinutes, user))
		delete(clients, body)
	}

	result.WriteString(fmt.Sprintf("%02d:%02d\n", closingHours, closingMinutes))

	for ind, pc := range computers {
		hours := pc.totaltime / 60
		minutes := pc.totaltime % 60
		result.WriteString(fmt.Sprintf("%d %d %02d:%02d\n", ind+1, pc.revenue, hours, minutes))
	}
	return result.String()
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Error: no file provided")
	}

	result := ModelComputerClub(os.Args[1])

	fmt.Print(result)
}
