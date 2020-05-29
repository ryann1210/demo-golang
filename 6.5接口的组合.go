package main

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com", map[string]string{
		"name":   "ryan",
		"course": "golang",
	})
}

type PosterAndGetter interface {
	Poster
	Get(host string)
}

func session(s PosterAndGetter) {
	//s.Post()
	//s.Get()
}

func main() {

}
