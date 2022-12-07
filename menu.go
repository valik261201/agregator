package main

type Foods struct {
	Foods []Food `json:"foods"`
}

type Food struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Preparation_time  int    `json:"preparation_time"`
	Complexity        int    `json:"complexity"`
	Cooking_apparatus string `json:"cooking_apparatus"`
}
