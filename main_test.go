package main

import "testing"

func Test_showASCIILogo(t *testing.T) {
	t.Run("Should give error when file not exist", func(t *testing.T) {
		filename := "notExistingFile.txt"
		err := showASCIILogo(filename)
		if err == nil {
			t.Fatalf("expected error to be not nil got %v", err)
		}
	})
	t.Run("Should not give error when file exist", func(t *testing.T) {
		filename := "slash-ascii.txt"
		err := showASCIILogo(filename)
		if err != nil {
			t.Fatalf("expected error to be nil got %v", err)
		}
	})
}
