package main

import "testing"

func Test_ValidParentheses(t *testing.T) {

	var tests = []struct {
		name      string
		inputText string
		expected  bool
	}{
		{"Empty String", "", false},
		{"Invalid String with only text", "abc", true},
		{"Invalid String with 1 close braces", "abc}", false},
		{"Valid String 1", "{abc}", true},
		{"Invalid String with two open", "{abc{", false},
		{"Invalid String 1", "{{)}}", false},
		{"Valid String 2", "{a{b+(c*d)/9}+1}", true},
		{"Valid String 3", "(){({})}", true},
		{"InValid String 3", "(){({}})", false},
		{"Valid String 4", "(((((())))))", true},
		{"Invalid String 4", "((((((", false},
		{"Invalid String 5", "))))))((((((", false},
		{"Invalid String 6", "{(}{})}", false},
		{"Valid String 5", "{}()[]{([])}", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := CheckValidParentheses(tt.inputText)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}

}
