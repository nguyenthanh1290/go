# Valid Parenthesis String

https://leetcode.com/problems/valid-parenthesis-string
(https://leetcode.com/problems/valid-parentheses/)

https://leetcode.com/submissions/detail/325900553/?from=/explore/featured/card/30-day-leetcoding-challenge/530/week-3/3301/

## Description

Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
An empty string is also valid.

Example 1:
```
Input: "()"
Output: True
```

Example 2:
```
Input: "(*)"
Output: True
```

Example 3:
```
Input: "(*))"
Output: True
```

### Note:
The string size will be in the range [1, 100].