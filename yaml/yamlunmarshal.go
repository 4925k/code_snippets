package main

import "fmt"

// type pathUrl struct {
// 	path string `yaml:"path"`
// 	url  string `yaml:"url"`
// }

func main() {
	var yaml string = `
- path: dogs
  url: https://www.google.com/search?q=dogs&source=lnms&tbm=isch&sa=X&ved=2ahUKEwi5kJu0pfzzAhVtzDgGHXuLAQwQ_AUoAXoECAEQAw&biw=1536&bih=746&dpr=1.25
- path: /rickroll
  url: https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley
`
	yamlbyte := []byte(yaml)
	// var data []pathUrl
	// yaml.Unmarshal(yamlbyte, &data)
	fmt.Println(yamlbyte)
}
