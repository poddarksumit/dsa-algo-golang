package main

import (
	"testing"
)

func Test_ReversePrefixUsingStack(t *testing.T) {

	var tests = []struct {
		testName     string
		word         string
		ch           rune
		expectedWord string
	}{
		{"Empty string", "", 'd', ""},
		{"No ch found", "a", 'd', "a"},
		{"Ch found", "a", 'a', "a"},
		{"Ch found 1", "aa", 'a', "aa"},
		{"No ch found 1", "aa", 'b', "aa"},
		{"Ch found 2", "ab", 'b', "ba"},
		{"Ch found 3", "aaaaad", 'd', "daaaaa"},
		{"Ch found 4", "dabcd", 'd', "ddcba"},
		{"Ch found 5", "abcdd", 'd', "dcbad"},
		{"Ch found 6", "abcd", 'd', "dcba"},
		{"Ch found 7", "ab c  i de ia op", 'd', "d i  c bae ia op"},
		{"Ch found 6", "abcdert", 'd', "dcbaert"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans := reversePrefixUsingStack(tt.word, tt.ch)
			if ans != tt.expectedWord {
				t.Errorf("got %v, want %v", ans, tt.expectedWord)
			}
		})

	}
}

func Test_ReversePrefixUsingString(t *testing.T) {

	var tests = []struct {
		testName     string
		word         string
		ch           rune
		expectedWord string
	}{
		{"Empty string", "", 'd', ""},
		{"No ch found", "a", 'd', "a"},
		{"Ch found", "a", 'a', "a"},
		{"Ch found 1", "aa", 'a', "aa"},
		{"No ch found 1", "aa", 'b', "aa"},
		{"Ch found 2", "ab", 'b', "ba"},
		{"Ch found 3", "aaaaad", 'd', "daaaaa"},
		{"Ch found 4", "dabcd", 'd', "ddcba"},
		{"Ch found 5", "abcdd", 'd', "dcbad"},
		{"Ch found 6", "abcd", 'd', "dcba"},
		{"Ch found 7", "ab c  i de ia op", 'd', "d i  c bae ia op"},
		{"Ch found 6", "abcdert", 'd', "dcbaert"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ans := reversePrefixUsingString(tt.word, tt.ch)
			if ans != tt.expectedWord {
				t.Errorf("got %v, want %v", ans, tt.expectedWord)
			}
		})

	}
}

func Test_ReversePrefix(t *testing.T) {

	var tests = []struct {
		name         string
		word         string
		ch           byte
		expectedWord string
	}{
		{"Empty string", "", 'd', ""},
		{"Valid string", "abcdefd", 'd', "dcbaefd"},
		{"No ch found", "a", 'd', "a"},
		{"Ch found", "a", 'a', "a"},
		{"Ch found 1", "aa", 'a', "aa"},
		{"No ch found 1", "aa", 'b', "aa"},
		{"Ch found 2", "ab", 'b', "ba"},
		{"Ch found 3", "aaaaad", 'd', "daaaaa"},
		{"Ch found 4", "dabcd", 'd', "dabcd"},
		{"Ch found 5", "abcdd", 'd', "dcbad"},
		{"Ch found 6", "abcd", 'd', "dcba"},
		{"Ch found 7", "ab c  i de ia opd", 'd', "d i  c bae ia opd"},
		{"Ch found 6", "abcdert", 'd', "dcbaert"},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			var ans = reversePrefix(tt.word, tt.ch)

			if ans != tt.expectedWord {
				t.Errorf("got %v, want %v", ans, tt.expectedWord)
			}
		})
	}
}
