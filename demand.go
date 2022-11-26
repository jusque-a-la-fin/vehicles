package models

type Demand struct {
	priorities    []string
	min_price     string
	max_price     string
	deviation     string
	manufacturers []string
	condition     string
	bodies        []string
}
