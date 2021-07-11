package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"strings"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:42586")
	if err != nil {
		log.Fatal(err)
		return
	}
	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
			return
		}
		numsParameter, err := GetParameter(string(line))
		if err != nil {
			log.Fatal(err)
			return
		}
		var result int
		if err = client.Call("Calculator.FormResultOfMultiply", numsParameter, &result); err != nil {
			log.Fatal(err)
			return
		}
		log.Println(numsParameter,result)

		userParameter, err := GetParameterOfUser(string(line))
		if err != nil {
			log.Fatal(err)
			return
		}
		var isValidUser bool
		if err = client.Call("Checker.IsValidUser", userParameter, &isValidUser); err != nil {
			log.Fatal(err)
			return
		}
		log.Println(userParameter, isValidUser)
	}
}
func GetParameterOfUser(line string) (map[string]string, error) {
	line = strings.TrimSpace(line)
	strs := strings.Split(line, " ")
	if len(strs) != 2 {
		err := fmt.Errorf("line(%s) can't be resolve to parameters", line)
		return nil, err
	}
	return map[string]string{
		"email":    strs[0],
		"password": strs[1],
	}, nil
}

func GetParameter(line string) (map[string]int, error) {
	line = strings.TrimSpace(line)
	numStrings := strings.Split(line, " ")
	if len(numStrings) != 2 {
		err := fmt.Errorf("line(%s) can't be resolve to parameters", line)
		return nil, err
	}
	nums := make([]int, 0)
	for i := 0; i < len(numStrings); i++ {
		num, err := strconv.Atoi(numStrings[i])
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		nums = append(nums, num)
	}
	return map[string]int{
		"leftNum":  nums[0],
		"rightNum": nums[1],
	}, nil
}
