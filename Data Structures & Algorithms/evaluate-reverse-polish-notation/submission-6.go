func evalRPN(tokens []string) int {

	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		
		if token == "+" || 
			token == "-" || 
			token == "*" || 
			token == "/" {

			n := len(stack) - 1

			result := arithmeticOperation(
				stack[n-1],
				stack[n],
				token,
			)

			stack[n-1] = result
			stack = stack[:n]
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}


func arithmeticOperation(x int, y int, operator string) int {
	result := 0
	switch {
		case operator == "+":
			return x + y
		case operator == "-":
			return x - y
		case operator == "*":
			return x * y
		case operator == "/":
			return x / y
	}

	return result

}
