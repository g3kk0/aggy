package main

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
//}

func main() {
	//	c := coinbase.ApiKeyClient(os.Getenv("COINBASE_KEY"), os.Getenv("COINBASE_SECRET"))

	GetDeposits()

	//	http.HandleFunc("/", handler)
	//	log.Fatal(http.ListenAndServe(":8080", nil))
}
