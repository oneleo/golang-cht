# 參考
- [Gin Tutorials](https://legacy.gitbook.com/book/chenyitian/gin-tutorials/details)

# 安裝
- git for windows
- golang 
- go get -u github.com/gin-gonic/gin

# Step 1：建置論壇骨架

## 新建「功能表」模板：./templates/menu.html

```html
<!--menu.html-->

<nav class="navbar navbar-default">
  <div class="container">
    <div class="navbar-header">
      <a class="navbar-brand" href="/">
        Home
      </a>
    </div>
  </div>
</nav>
```

## 新建「標頭」模板：./templates/header.html

```html
<!--header.html-->

<!doctype html>
<html>

<head>
    <!--Use the title variable to set the title of the page-->
    <title>{{ .title }}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta charset="UTF-8">

    <!--Use bootstrap to make the application look nice-->
    <!--An overview of Bootstrap: https://bootstrapdocs.com/v3.3.6/docs/getting-started/-->
    <!-- Latest compiled and minified CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css"
        integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">

    <!-- Optional theme -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap-theme.min.css"
        integrity="sha384-fLW2N01lMqjakBkx3l/M9EahuwpSfeNvV63J5ezn3uZzapT0u7EYsXMjQV+0En5r" crossorigin="anonymous">

    <!-- Latest compiled and minified JavaScript -->
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js"
        integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS"
        crossorigin="anonymous"></script>
</head>

<body class="container">
    <!--Embed the menu.html template at this location-->
    {{ template "menu.html" . }}
```

## 新建「頁尾」模板：./templates/footer.html

```html
<!--footer.html-->

  </body>

</html>
```

## 新建「首頁」模板：./templates/index.html

```html
<!--index.html-->

<!--Embed the header.html template at this location-->
{{ template "header.html" .}}

  <h1>Hello Gin!</h1>

<!--Embed the footer.html template at this location-->
{{ template "footer.html" .}}
```

## 新建 Go 程式碼：./main.go

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

  // Set the router as the default one provided by Gin
  router = gin.Default()

  // Process the templates at the start so that they don't have to be loaded
  // from the disk again. This makes serving HTML pages very fast.
  router.LoadHTMLGlob("templates/*")

  // Define the route for the index page and display the index.html template
  // To start with, we'll use an inline route handler. Later on, we'll create
  // standalone functions that will be used as route handlers.
  router.GET("/", func(c *gin.Context) {

    // Call the HTML method of the Context to render a template
    c.HTML(
      // Set the HTTP status to 200 (OK)
      http.StatusOK,
      // Use the index.html template
      "index.html",
      // Pass the data that the page uses (in this case, 'title')
      gin.H{
        "title": "Home Page",
      },
    )
  })

  // Start serving the application
  router.Run()
}
```

## 測試

```powershell
> go build -o app.exe
> ./app.exe
```

- 至 [localhost:8080](http://127.0.0.1:8080) 查看

# Step 2：建置可瀏覽預設文章的論壇

## 新建 Go 程式碼：./routes.go

```go
// routes.go

package main

func initializeRoutes() {

  // Handle the index route
  router.GET("/", showIndexPage)
}
```

## 【修改】 Go 程式碼：./main.go

```go
package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()
}
```

## 新建 Go 程式碼：./models.article.go

```go
// models.article.go

package main

type article struct {
  ID      int    `json:"id"`
  Title   string `json:"title"`
  Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
  article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
  article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
  return articleList
}
```

## 【修改】「首頁」模板：./templates/index.html

```html
<!--index.html-->

<!--Embed the header.html template at this location-->
{{ template "header.html" .}}

  <!--Loop over the `payload` variable, which is the list of articles-->
  {{range .payload }}
    <!--Create the link for the article based on its ID-->
    <a href="/article/view/{{.ID}}">
      <!--Display the title of the article -->
      <h2>{{.Title}}</h2>
    </a>
    <!--Display the content of the article-->
    <p>{{.Content}}</p>
  {{end}}

<!--Embed the footer.html template at this location-->
{{ template "footer.html" .}}
```

## 新建 Go 程式碼：./handlers.article.go

```go
// handlers.article.go

package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
  articles := getAllArticles()

  // Call the HTML method of the Context to render a template
  c.HTML(
    // Set the HTTP status to 200 (OK)
    http.StatusOK,
    // Use the index.html template
    "index.html",
    // Pass the data that the page uses
    gin.H{
      "title":   "Home Page",
      "payload": articles,
    },
  )
}
```

## 測試

```powershell
> go build -o app.exe
> ./app.exe
```

- 至 [localhost:8080](http://127.0.0.1:8080) 查看

# Step 3：建置可查看文章的論壇

## 【修改】 Go 程式碼：./routes.go

```go
// routes.go

package main

func initializeRoutes() {

  // Handle the index route
  router.GET("/", showIndexPage)

  // Handle GET requests at /article/view/some_article_id
  router.GET("/article/view/:article_id", getArticle)
}
```

## 新建「文章」模板：./templates/article.html

```html
<!--article.html-->

<!--Embed the header.html template at this location-->
{{ template "header.html" .}}

<!--Display the title of the article-->
<h1>{{.payload.Title}}</h1>

<!--Display the content of the article-->
<p>{{.payload.Content}}</p>

<!--Embed the footer.html template at this location-->
{{ template "footer.html" .}}
```

## 【修改】 Go 程式碼：./models.article.go

```go
// models.article.go

package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
```

## 【修改】 Go 程式碼：./handlers.article.go

```go
// handlers.article.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the HTML method of the Context to render a template
			c.HTML(
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the index.html template
				"article.html",
				// Pass the data that the page uses
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
```
## 測試

```powershell
> go build -o app.exe
> ./app.exe
```

- 至 [localhost:8080](http://127.0.0.1:8080) 查看

# Step 4：建置可回傳 XML、JSON 格式的論壇

## 【修改】 Go 程式碼：./main.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
```

## 【修改】 Go 程式碼：./handlers.article.go

```go
// handlers.article.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
```

## 測試

```powershell
> go build -o app.exe
> ./app.exe
```

- 至 [localhost:8080](http://127.0.0.1:8080) 查看

## 使用 CMD 測試

```cmd
> curl -X GET -H "Accept: application/json" http://localhost:8080/
> curl -X GET -H "Accept: application/xml" http://localhost:8080/article/view/1
```

# Step 5：建置的論壇

## 新建 Go 程式碼：./models.user.go

```go
// models.user.go

package main

import "errors"

type user struct {
    Username string `json:"username"`
    Password string `json:"-"`
}

var userList = []user{
    user{Username: "user1", Password: "pass1"},
    user{Username: "user2", Password: "pass2"},
    user{Username: "user3", Password: "pass3"},
}

func registerNewUser(username, password string) (*user, error) {
    return nil, errors.New("placeholder error")
}

func isUsernameAvailable(username string) bool {
    return false
}
```

## 新建 Go 程式碼：./handlers.user.go

```go
// handlers.user.go

package main

import "github.com/gin-gonic/gin"

func showRegistrationPage(c *gin.Context) {
}

func register(c *gin.Context) {
}
```

## 【修改】 Go 程式碼：./routes.go

```go
// routes.go

package main

func initializeRoutes() {

    router.GET("/", showIndexPage)

    userRoutes := router.Group("/u")
    {
        userRoutes.GET("/register", showRegistrationPage)

        userRoutes.POST("/register", register)
    }

    router.GET("/article/view/:article_id", getArticle)
}
```

## 新建「註冊使用者」模板：./templates/register.html

```html
<!--register.html-->

<!--Embed the header.html template at this location-->
{{ template "header.html" .}}

<h1>Register</h1>

<div class="panel panel-default col-sm-6">
  <div class="panel-body">
    <!--If there's an error, display the error-->
    {{ if .ErrorTitle}}
    <p class="bg-danger">
      {{.ErrorTitle}}: {{.ErrorMessage}}
    </p>
    {{end}}
    <!--Create a form that POSTs to the `/u/register` route-->
    <form class="form" action="/u/register" method="POST">
      <div class="form-group">
        <label for="username">Username</label>
        <input type="text" class="form-control" id="username" name="username" placeholder="Username">
      </div>
      <div class="form-group">
        <label for="password">Password</label>
        <input type="password" name="password" class="form-control" id="password" placeholder="Password">
      </div>
      <button type="submit" class="btn btn-primary">Register</button>
    </form>
  </div>
</div>  


<!--Embed the footer.html template at this location-->
{{ template "footer.html" .}}
```

## 新建「登入成功」模板：./templates/login-successful.html

```html
<!--login-successful.html-->

<!--Embed the header.html template at this location-->
{{ template "header.html" .}}

<div>
  You have successfully logged in.
</div>

<!--Embed the footer.html template at this location-->
{{ template "footer.html" .}}
```

## 【修改】「功能表」模板：./templates/menu.html

```html
<!--menu.html-->

<nav class="navbar navbar-default">
  <div class="container">
    <div class="navbar-header">
      <a class="navbar-brand" href="/">
        Home
      </a>
    </div>
    <ul class="nav navbar-nav">
        <li><a href="/u/register">Register</a></li>
    </ul>
  </div>
</nav>
```

## 【修改】Go 程式碼：./models.user.go

```go
// models.user.go

package main

import (
    "errors"
    "strings"
)

type user struct {
    Username string `json:"username"`
    Password string `json:"-"`
}

// For this demo, we're storing the user list in memory
// We also have some users predefined.
// In a real application, this list will most likely be fetched
// from a database. Moreover, in production settings, you should
// store passwords securely by salting and hashing them instead
// of using them as we're doing in this demo
var userList = []user{
    user{Username: "user1", Password: "pass1"},
    user{Username: "user2", Password: "pass2"},
    user{Username: "user3", Password: "pass3"},
}

// Register a new user with the given username and password
func registerNewUser(username, password string) (*user, error) {
    if strings.TrimSpace(password) == "" {
        return nil, errors.New("The password can't be empty")
    } else if !isUsernameAvailable(username) {
        return nil, errors.New("The username isn't available")
    }

    u := user{Username: username, Password: password}

    userList = append(userList, u)

    return &u, nil
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
    for _, u := range userList {
        if u.Username == username {
            return false
        }
    }
    return true
}
```

## 【修改】Go 程式碼：./handlers.user.go

```go
// handlers.user.go

package main

import (
    "math/rand"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func generateSessionToken() string {
    // We're using a random 16 character string as the session token
    // This is NOT a secure way of generating session tokens
    // DO NOT USE THIS IN PRODUCTION
    return strconv.FormatInt(rand.Int63(), 16)
}

func showRegistrationPage(c *gin.Context) {
    // Call the render function with the name of the template to render
    render(c, gin.H{
        "title": "Register"}, "register.html")
}

func register(c *gin.Context) {
    // Obtain the POSTed username and password values
    username := c.PostForm("username")
    password := c.PostForm("password")

    if _, err := registerNewUser(username, password); err == nil {
        // If the user is created, set the token in a cookie and log the user in
        token := generateSessionToken()
        c.SetCookie("token", token, 3600, "", "", false, true)
        c.Set("is_logged_in", true)

        render(c, gin.H{
            "title": "Successful registration & Login"}, "login-successful.html")

    } else {
        // If the username/password combination is invalid,
        // show the error message on the login page
        c.HTML(http.StatusBadRequest, "register.html", gin.H{
            "ErrorTitle":   "Registration Failed",
            "ErrorMessage": err.Error()})

    }
}
```

## 【修改】Go 程式碼：./models.user.go

```go
// models.user.go

package main

import (
    "errors"
    "strings"
)

type user struct {
    Username string `json:"username"`
    Password string `json:"-"`
}

// For this demo, we're storing the user list in memory
// We also have some users predefined.
// In a real application, this list will most likely be fetched
// from a database. Moreover, in production settings, you should
// store passwords securely by salting and hashing them instead
// of using them as we're doing in this demo
var userList = []user{
    user{Username: "user1", Password: "pass1"},
    user{Username: "user2", Password: "pass2"},
    user{Username: "user3", Password: "pass3"},
}

func isUserValid(username, password string) bool {
    return false
}

// Register a new user with the given username and password
func registerNewUser(username, password string) (*user, error) {
    if strings.TrimSpace(password) == "" {
        return nil, errors.New("The password can't be empty")
    } else if !isUsernameAvailable(username) {
        return nil, errors.New("The username isn't available")
    }

    u := user{Username: username, Password: password}

    userList = append(userList, u)

    return &u, nil
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
    for _, u := range userList {
        if u.Username == username {
            return false
        }
    }
    return true
}
```

## 【修改】Go 程式碼：./handlers.user.go

```go
// handlers.user.go

package main

import (
    "math/rand"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func showLoginPage(c *gin.Context) {}

func performLogin(c *gin.Context) {}

func logout(c *gin.Context) {}

func generateSessionToken() string {
    // We're using a random 16 character string as the session token
    // This is NOT a secure way of generating session tokens
    // DO NOT USE THIS IN PRODUCTION
    return strconv.FormatInt(rand.Int63(), 16)
}

func showRegistrationPage(c *gin.Context) {
    // Call the render function with the name of the template to render
    render(c, gin.H{
        "title": "Register"}, "register.html")
}

func register(c *gin.Context) {
    // Obtain the POSTed username and password values
    username := c.PostForm("username")
    password := c.PostForm("password")

    if _, err := registerNewUser(username, password); err == nil {
        // If the user is created, set the token in a cookie and log the user in
        token := generateSessionToken()
        c.SetCookie("token", token, 3600, "", "", false, true)
        c.Set("is_logged_in", true)

        render(c, gin.H{
            "title": "Successful registration & Login"}, "login-successful.html")

    } else {
        // If the username/password combination is invalid,
        // show the error message on the login page
        c.HTML(http.StatusBadRequest, "register.html", gin.H{
            "ErrorTitle":   "Registration Failed",
            "ErrorMessage": err.Error()})

    }
}
```

## 【修改】 Go 程式碼：./routes.go

```go
// routes.go

package main

func initializeRoutes() {

    router.GET("/", showIndexPage)

    userRoutes := router.Group("/u")
    {
        userRoutes.GET("/register", showRegistrationPage)

		userRoutes.POST("/register", register)
		
		userRoutes.GET("/login", showLoginPage)
        userRoutes.POST("/login", performLogin)
        userRoutes.GET("/logout", logout)
    }

    router.GET("/article/view/:article_id", getArticle)
}
```

## 新建「登入」模板：./templates/login.html

```html
<!--login.html-->

{{ template "header.html" .}}

<h1>Login</h1>


<div class="panel panel-default col-sm-6">
  <div class="panel-body">

    {{ if .ErrorTitle}}
    <p class="bg-danger">
      {{.ErrorTitle}}: {{.ErrorMessage}}
    </p>
    {{end}}

    <form class="form" action="/u/login" method="POST">
      <div class="form-group">
        <label for="username">Username</label>
        <input type="text" class="form-control" id="username" name="username" placeholder="Username">
      </div>
      <div class="form-group">
        <label for="password">Password</label>
        <input type="password" class="form-control" id="password" name="password" placeholder="Password">
      </div>
      <button type="submit" class="btn btn-primary">Login</button>
    </form>
  </div>
</div>

{{ template "footer.html" .}}
```

## 【修改】「功能表」模板：./templates/menu.html

```html
<!--menu.html-->

<nav class="navbar navbar-default">
  <div class="container">
    <div class="navbar-header">
      <a class="navbar-brand" href="/">
        Home
      </a>
    </div>
    <ul class="nav navbar-nav">
		<li><a href="/u/register">Register</a></li>
        <li><a href="/u/login">Login</a></li>
        <li><a href="/u/logout">Logout</a></li>
    </ul>
  </div>
</nav>
```

## 【修改】Go 程式碼：./models.user.go

```go
// models.user.go

package main

import (
    "errors"
    "strings"
)

type user struct {
    Username string `json:"username"`
    Password string `json:"-"`
}

// For this demo, we're storing the user list in memory
// We also have some users predefined.
// In a real application, this list will most likely be fetched
// from a database. Moreover, in production settings, you should
// store passwords securely by salting and hashing them instead
// of using them as we're doing in this demo
var userList = []user{
    user{Username: "user1", Password: "pass1"},
    user{Username: "user2", Password: "pass2"},
    user{Username: "user3", Password: "pass3"},
}

func isUserValid(username, password string) bool {
    for _, u := range userList {
        if u.Username == username && u.Password == password {
            return true
        }
    }
    return false
}

// Register a new user with the given username and password
func registerNewUser(username, password string) (*user, error) {
    if strings.TrimSpace(password) == "" {
        return nil, errors.New("The password can't be empty")
    } else if !isUsernameAvailable(username) {
        return nil, errors.New("The username isn't available")
    }

    u := user{Username: username, Password: password}

    userList = append(userList, u)

    return &u, nil
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
    for _, u := range userList {
        if u.Username == username {
            return false
        }
    }
    return true
}
```

## 【修改】Go 程式碼：./handlers.user.go

```go
// handlers.user.go

package main

import (
    "math/rand"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func showLoginPage(c *gin.Context) {
    render(c, gin.H{
        "title": "Login",
    }, "login.html")
}

func performLogin(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    if isUserValid(username, password) {
        token := generateSessionToken()
        c.SetCookie("token", token, 3600, "", "", false, true)

        render(c, gin.H{
            "title": "Successful Login"}, "login-successful.html")

    } else {
        c.HTML(http.StatusBadRequest, "login.html", gin.H{
            "ErrorTitle":   "Login Failed",
            "ErrorMessage": "Invalid credentials provided"})
    }
}

func logout(c *gin.Context) {
    c.SetCookie("token", "", -1, "", "", false, true)

    c.Redirect(http.StatusTemporaryRedirect, "/")
}

func generateSessionToken() string {
    // We're using a random 16 character string as the session token
    // This is NOT a secure way of generating session tokens
    // DO NOT USE THIS IN PRODUCTION
    return strconv.FormatInt(rand.Int63(), 16)
}

func showRegistrationPage(c *gin.Context) {
    // Call the render function with the name of the template to render
    render(c, gin.H{
        "title": "Register"}, "register.html")
}

func register(c *gin.Context) {
    // Obtain the POSTed username and password values
    username := c.PostForm("username")
    password := c.PostForm("password")

    if _, err := registerNewUser(username, password); err == nil {
        // If the user is created, set the token in a cookie and log the user in
        token := generateSessionToken()
        c.SetCookie("token", token, 3600, "", "", false, true)
        c.Set("is_logged_in", true)

        render(c, gin.H{
            "title": "Successful registration & Login"}, "login-successful.html")

    } else {
        // If the username/password combination is invalid,
        // show the error message on the login page
        c.HTML(http.StatusBadRequest, "register.html", gin.H{
            "ErrorTitle":   "Registration Failed",
            "ErrorMessage": err.Error()})

    }
}
```

## 【修改】 Go 程式碼：./models.article.go

```go
// models.article.go

package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func createNewArticle(title, content string) (*article, error) {
    return nil, nil
}
```

## 【修改】 Go 程式碼：./handlers.article.go

```go
// handlers.article.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func showArticleCreationPage(c *gin.Context) {}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createArticle(c *gin.Context) {}
```

## 【修改】 Go 程式碼：./routes.go

```go
// routes.go

package main

func initializeRoutes() {

    router.GET("/", showIndexPage)

    userRoutes := router.Group("/u")
    {
        userRoutes.GET("/register", showRegistrationPage)

		userRoutes.POST("/register", register)
		
		userRoutes.GET("/login", showLoginPage)
        userRoutes.POST("/login", performLogin)
        userRoutes.GET("/logout", logout)
    }

    articleRoutes := router.Group("/article")
    {
        // route from Part 1 of the tutorial
        articleRoutes.GET("/view/:article_id", getArticle)

        articleRoutes.GET("/create", showArticleCreationPage)

        articleRoutes.POST("/create", createArticle)
    }
}
```

## 新建「新增文章」模板：./templates/create-article.html

```html
<!--create-article.html-->

{{ template "header.html" .}}

<h1>Create Article</h1>

<div class="panel panel-default col-sm-12">
  <div class="panel-body">
    {{ if .ErrorTitle}}
    <p class="bg-danger">
      {{.ErrorTitle}}: {{.ErrorMessage}}
    </p>
    {{end}}

    <form class="form" action="/article/create" method="POST">
      <div class="form-group">
        <label for="title">Username</label>
        <input type="text" class="form-control" id="title" name="title" placeholder="Title">
      </div>
      <div class="form-group">
        <label for="content">Password</label>
        <textarea name="content" class="form-control" rows="10" id="content" laceholder="Article Content"></textarea>
      </div>
      <button type="submit" class="btn btn-primary">Submit</button>
    </form>
  </div>
</div>

{{ template "footer.html" .}}
```

## 新建「送出文章」模板：./templates/submission-successful.html

```html
<!--submission-successful.html-->

{{ template "header.html" .}}

<div>
  <strong>The article was successfully submitted.</strong>

  <a href="/article/view/{{.payload.ID}}">{{.payload.Title}}</a>
</div>

{{ template "footer.html" .}}
```

## 【修改】「功能表」模板：./templates/menu.html

```html
<!--menu.html-->

<nav class="navbar navbar-default">
  <div class="container">
    <div class="navbar-header">
      <a class="navbar-brand" href="/">
        Home
      </a>
    </div>
    <ul class="nav navbar-nav">
		<li><a href="/article/create">Create Article</a></li>
		<li><a href="/u/register">Register</a></li>
        <li><a href="/u/login">Login</a></li>
        <li><a href="/u/logout">Logout</a></li>
    </ul>
  </div>
</nav>
```

## 【修改】 Go 程式碼：./models.article.go

```go
// models.article.go

package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func createNewArticle(title, content string) (*article, error) {
    a := article{ID: len(articleList) + 1, Title: title, Content: content}

    articleList = append(articleList, a)

    return &a, nil
}
```

## 【修改】 Go 程式碼：./handlers.article.go

```go
// handlers.article.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func showArticleCreationPage(c *gin.Context) {
    render(c, gin.H{
        "title": "Create New Article"}, "create-article.html")
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createArticle(c *gin.Context) {
    title := c.PostForm("title")
    content := c.PostForm("content")

    if a, err := createNewArticle(title, content); err == nil {
        render(c, gin.H{
            "title":   "Submission Successful",
            "payload": a}, "submission-successful.html")
    } else {
        c.AbortWithStatus(http.StatusBadRequest)
    }
}
```

# Step 5：建置的論壇

## 新建 Go 程式碼：./middleware.auth.go

```go
// middleware.auth.go

package main

import "github.com/gin-gonic/gin"

func ensureLoggedIn() gin.HandlerFunc {
    return func(c *gin.Context) {

    }
}

func ensureNotLoggedIn() gin.HandlerFunc {
    return func(c *gin.Context) {

    }
}

func setUserStatus() gin.HandlerFunc {
    return func(c *gin.Context) {

    }
}
```

## 【修改】Go 程式碼：./middleware.auth.go

```go
// middleware.auth.go

package main

import "github.com/gin-gonic/gin"

func ensureLoggedIn() gin.HandlerFunc {
    return func(c *gin.Context) {
        loggedInInterface, _ := c.Get("is_logged_in")
        loggedIn := loggedInInterface.(bool)
        if !loggedIn {
            c.AbortWithStatus(http.StatusUnauthorized)
        }
    }
}

func ensureNotLoggedIn() gin.HandlerFunc {
    return func(c *gin.Context) {
        loggedInInterface, _ := c.Get("is_logged_in")
        loggedIn := loggedInInterface.(bool)
        if loggedIn {
            c.AbortWithStatus(http.StatusUnauthorized)
        }
    }
}

func setUserStatus() gin.HandlerFunc {
    return func(c *gin.Context) {
        if token, err := c.Cookie("token"); err == nil || token != "" {
            c.Set("is_logged_in", true)
        } else {
            c.Set("is_logged_in", false)
        }
    }
}
```

## 【修改】 Go 程式碼：./routes.go

```go
// routes.go

package main

func initializeRoutes() {

    router.Use(setUserStatus())

    router.GET("/", showIndexPage)

    userRoutes := router.Group("/u")
    {
        userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

        userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

        userRoutes.GET("/logout", ensureLoggedIn(), logout)

        userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)

        userRoutes.POST("/register", ensureNotLoggedIn(), register)
    }

    articleRoutes := router.Group("/article")
    {
        articleRoutes.GET("/view/:article_id", getArticle)

        articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)

        articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
    }
}
```

## 【修改】「功能表」模板：./templates/menu.html

```html
<!--menu.html-->

<nav class="navbar navbar-default">
  <div class="container">
    <div class="navbar-header">
      <a class="navbar-brand" href="/">
        Home
      </a>
    </div>
    <ul class="nav navbar-nav">
      {{ if .is_logged_in }}
        <!--Display this link only when the user is logged in-->
        <li><a href="/article/create">Create Article</a></li>
      {{end}}
      {{ if not .is_logged_in }}
        <!--Display this link only when the user is not logged in-->
        <li><a href="/u/register">Register</a></li>
      {{end}}
      {{ if not .is_logged_in }}
        <!--Display this link only when the user is not logged in-->
        <li><a href="/u/login">Login</a></li>
      {{end}}
      {{ if .is_logged_in }}
        <!--Display this link only when the user is logged in-->
        <li><a href="/u/logout">Logout</a></li>
      {{end}}
    </ul>
  </div>
</nav>
```




