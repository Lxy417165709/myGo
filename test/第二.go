package main



import(

	"fmt"
	"regexp"
	"strings"
)

func solve(str string) string{

	// 处理EUR情况
	digitsRegexp := regexp.MustCompile(`^EUR (\d+).(\d+),(\d+)$`)
	arr := digitsRegexp.FindStringSubmatch(str)
	if len(arr)==4 {
		arr[3] = strings.TrimRight(arr[3],"0")
		if arr[3]==""{
			return arr[1] + arr[2]
		}else{
			return arr[1] + arr[2] + "." + arr[3]
		}
	}

	// 处理€情况
	digitsRegexp = regexp.MustCompile(`^€ (\d+),(\d+)$`)
	arr = digitsRegexp.FindStringSubmatch(str)
	if len(arr)==3 {
		arr[2] = strings.TrimRight(arr[2],"0")
		if arr[2]==""{
			return arr[1]
		}else{
			return arr[1] + "." + arr[2]
		}
	}
	// 处理￥情况
	digitsRegexp = regexp.MustCompile(`^￥(\d+).(\d+)$`)
	arr = digitsRegexp.FindStringSubmatch(str)
	if len(arr)==3 {
		arr[2] = strings.TrimRight(arr[2],"0")
		if arr[2]==""{
			return arr[1]
		}else{
			return arr[1] + "." + arr[2]
		}
	}

	// 处理CNY 情况
	digitsRegexp = regexp.MustCompile(`^CNY (\d+),(\d+)$`)
	arr = digitsRegexp.FindStringSubmatch(str)
	if len(arr)==3 {
		return arr[1] + arr[2]
	}

	// 非法情况
	return "0"
}




func main(){



	fmt.Println(solve("EUR 1.,05"))
	fmt.Println(solve("€"))
	fmt.Println(solve("￥409.50"))
	fmt.Println(solve("CNY 1,000"))
	fmt.Println(solve("abs"))
}