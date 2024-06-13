package main

import "testing"

func Test_isPalindrome(t *testing.T) {

	if is := isPalindrome(121); !is {
		t.Log(is)
		t.Error("Expected", true)
	}

	if is := isPalindrome(-121); is {
		t.Error("Expected", false)
	}

	if is := isPalindrome(10); is {
		t.Error("Expected", false)
	}

}
