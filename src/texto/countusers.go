package main

import "fmt"

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	listUsers := []string{"Tim", "Haydn", "John", "Haydn"}
	fmt.Println(listUsers)
	var countedUsers []string
	var score = map[string]int{}
	for range listUsers {
		count := 0

		for i := range listUsers {
			userName := listUsers[i]
			// for i := range listUsers {
			if listUsers[i] == userName && contains(countedUsers, userName) == false {
				fmt.Println("we have a match", userName, ":", listUsers[i])
				count++
				countedUsers = append(countedUsers, userName)
				score[userName] = count

			} else {
				// count++
				// score[userName] = count
				//continue
			}
			fmt.Println(userName, " counted: ", count)
		}
		//
	}

	fmt.Println(score)
}
