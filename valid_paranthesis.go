
// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

var closed = map[string]string{
	"}": "{",
	")": "(",
	"]": "[",
}

func isValid(s string) bool {
	stack := []string{}
	for _, i := range s {
		if j, ok := closed[string(i)]; ok {
            length := len(stack)
			if length > 0 && j == stack[length-1] {
				stack = stack[:length-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, string(i))
		}
	}
    if len(stack) > 0 {
        return false 
    } else {
        return true
    }
}


