package main

import (
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locationPage, err := cfg.client.GetLocationAreaPage(cfg.next)
	if err != nil {
		return err
	}
	cfg.previous = locationPage.Previous
	cfg.next = locationPage.Next
	for _, location := range locationPage.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previous == nil {
		return fmt.Errorf("you are on the first page")
	}
	locationPage, err := cfg.client.GetLocationAreaPage(cfg.previous)
	if err != nil {
		return err
	}
	cfg.previous = locationPage.Previous
	cfg.next = locationPage.Next
	for _, location := range locationPage.Results {
		fmt.Println(location.Name)
	}
	return nil
}
