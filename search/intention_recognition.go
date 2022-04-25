package search

import (
	"dianasdog/database"
	"fmt"
)

//First string stores the tablename, Second string stores the columnname and final value point to the acmatcher
var acmap map[string]map[string]*Matcher

// @title: init
// @description:initialize the acmap
// @param: do not need a param
// @return: do not need a return-value
func init() {
	acmap = make(map[string]map[string]*Matcher)
	tables, err := database.ShowTablesInDict()
	if err != nil {
		fmt.Println(err)
	}
	for _, table := range tables {
		fields, err := database.GetAllField(table)
		if err != nil {
			fmt.Println(err)
		}
		acmap[table] = make(map[string]*Matcher)
		for _, field := range fields {
			acmap[table][field] = NewMatcher()
			dict, err := database.GetAllWord(table, field)
			if err != nil {
				fmt.Println(err)
			}
			acmap[table][field].Build(dict)
		}
	}
}

// @title: IntentionRecognition
// @description: recognize the intent from the query
// @param: query          string    the query posted from the frontend
// @return: intentList    []string  the list of the resources
func IntentionRecognition(query string) []string {
	intentList := []string{}
	for table := range acmap {
		for field := range acmap[table] {
			if field != "garbage" && field != "intent" {
				check := acmap[table][field].Check(query)
				if check {
					intentList = append(intentList, table)
					break
				}
			}
		}
	}
	return intentList
}
