package main

// fatalIfError calls log.Fatal with a context message if err is not null
func fatalIfError(err error, msg string) {
	if err != nil {
		log.Fatal("error ", msg, ": ", err)
	}
}

// fatalIfEmptyString calls log.Fatal if s is empty, noting that the label is
// empty. This is useful with environment variables:
//
//    foo := os.Getenv(envFoo)
//    fatalIfEmptyString(foo, envFoo)
func fatalIfEmptyString(s, label string) {
	if s == "" {
		log.Fatalf("%s cannot be empty", label)
	}
}
