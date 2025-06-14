package handler

import (
	"fmt"
	"net/http"
)

func HtmlRendering(w http.ResponseWriter, r *http.Request) {

	var name string = "Kendrick"

	htmlContent := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Go HTML Rendering</title>
	</head>
	<body>
		<h1>Hello, %s! This is a HTML document rendered with Go!</h1>
	</body>
	</html>
	`, name)
	fmt.Fprint(w, htmlContent)
}

// https://vercel.com/0neand0nlykens-projects/speedrun-fast-hands-on-webdev/DMWAdd9prGmxx3k8ADqYFZrmVMN4
