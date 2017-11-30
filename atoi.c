/*
*
* Vasu Mahalingam, Nov 2017
*
*/
/*
Submission Details
1043 / 1047 test cases passed.
Status: Wrong Answer
Submitted: 0 minutes ago
Input:
"9223372036854775809"
Output:
-2147483648
Expected:
2147483647

*** Just gave up here :-) *** 

Implement atoi to convert a string to an integer.

Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.

Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.

Update (2015-02-10):
The signature of the C++ function had been updated. If you still see your function signature accepts a const char * argument, please click the reload button  to reset your code definition.

spoilers alert... click to show requirements for atoi.

Requirements for atoi:
The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned. If the correct value is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.
*/
#include<stdlib.h>
#include<stdio.h>
#include<string.h>
#include<limits.h>

int validate(char c) {  
    int d = c - '0';
    return (d < 0 || d > 9)? -1 : 0;
}

int myAtoi(char* str) {
    int sign = 1, i = 0;
    signed long long int value = 0;
    if (str == NULL) {
        return 0;
    }   
    // ignore whitespaces
    while (str[i] == ' ') i++;
    if (str[i] == '-') { 
        sign = -1; 
        i++; 
    } else { 
        sign = 1; 
        if (str[i] == '+') { 
            i++; 
        }
    }
    for (;i<strlen(str);) { 
      if (validate(str[i]) < 0) { // return sign * value;
          break; 
      }   
      value = value * 10 + (str[i] - '0');
      i++;
    }   
    value = value * sign;
    if (value >= INT_MAX) return INT_MAX; 
    else if (value <= INT_MIN) return INT_MIN; 
    return value;
}

