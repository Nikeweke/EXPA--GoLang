// ===========================================================->
// -- Запуск сервера на порт 8080 (нужна библиотека: fmt, net/http)

fmt.Println(http.ListenAndServe(":8080", nil)) // Command start server
// ===========================================================->




// ===========================================================->
// -- Если человек заходит по этому адресу - localhost:8080/static
//    он увидит на страничке что делает фукнция - StaticPage
http.HandleFunc("/static/", StaticPage)
// ===========================================================->




// ===========================================================->
// -- ВЫВОД текста в КОНСОЛЬ (нужна библиотека: fmt)
/*
   Вывод текста в консоль с вставками вместо переменной - %s
   a - переменная в которой html документ, в который вставлены переменные как %s
   GG, Alerk - будут вставляеться по очереди в документ вместо %s, который будет выведен в консоль
*/
fmt.Printf(a, "GG", "Alerk")

fmt.Println("GAGA")
//or
fmt.Fprintf("GAGA")
//or
fmt.Fprint("GAGA")
// ===========================================================->




// ===========================================================->
// -- ВЫВОД текста на страничку в БРАУЗЕРЕ (нужна библиотека: fmt, net/http)
/*
   (wr http.ResponseWriter, req *http.Request) - это обязательные параметры для работы с страницей
    wr - переменная юзаеться при выводе
    req - ....
*/
func Braker(wr http.ResponseWriter, req *http.Request){

  fmt.Fprint(wr,"GAGA") // Если надо вывести просто строку то эта команда

// Это вывод странички, которая в переменной -  t
  fmt.Fprintf(wr, t) // если надо вывести html
}

// -- Хранение страничке html правильно чтобы выводить в БРАУЗЕР
// FILE - html.go
package main

var (
// 1 HTML
      t = `
      <html>
        <head>
          <title>HOME</title>
        </head>
      <body>
        <h2>HALLO</h2>
        <form method="post">
           <input name="msg" type="text">
           <button name="okey" type="submit">OK</button>
        </form>
      </body>
      </html> `

// 2 HTML
     a = `<html>
          <head>
          <title>NOT HOME</title>
          </head>
          <body>
            <h2>%s, (%s)</h2>
          </body>
          </html> `
     )
// ===========================================================->






// ===========================================================->
// -- Сервере как апач для просмотра страничек своих
func main() {

  http.HandleFunc("/static/", StaticPage)

  // START SERVER
  println("\nServer is RUNNING...............................................\n")
  log.Println(" -> Time Started")
  fmt.Println(http.ListenAndServe(":8080", nil))
}

func StaticPage(wr http.ResponseWriter, req *http.Request){
     // Allows
     w.Header().Set("Access-Control-Allow-Origin",  "*")

    /* Allows
      if origin := r.Header().Get("Origin"); origin != "" {
	     w.Header().Set("Access-Control-Allow-Origin", origin)
	     w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	     w.Header().Set("Access-Control-Allow-Headers",  "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	     w.Header().Set("Access-Control-Max-Age", "86400") // 24 hours
	    }
    */

	//  File static page
	http.ServeFile(w, r, r.URL.Path[1:])
}
// ===========================================================->





// ===========================================================->
// -- РЕДИРЕКТ
// Ничего не должно выводиться чтобы эта комманда работала
http.Redirect(wr, req, "/temp/", http.StatusSeeOther)
// ===========================================================->




// ===========================================================->
// -- ВЫВОД времени запуска сервера GO (нужна библиотека: fmt, log)

log.Println(" -> Time Started")
// ===========================================================->





// ===========================================================->
// -- КАК ловить нажатие на кнопку НА ФОРМАХ
// На форме button - это <input type="buton" name="btn" value="enter">
if req.Method == "POST" && req.FormValue("btn") == "enter" {

   // получение данных с полей
   email    := req.FormValue("email")
   password := req.FormValue("password")

  // Проверка поля на пустоту
  if !FilledOut(email) || !FilledOut(password) { fmt.Println("Email or Password is empty")  }

  // Проверка на валидность почты
  if !ValidEmail(email){ fmt.Println("Email is Shitty") }


  fmt.Println("\nEmail ->" + email)
  fmt.Println("Pass ->" + password)
 }
// ===========================================================->




/////////////////////////////////////////////////////////////////////////////////////////////////// #### COOKIES
//================================== Действие: Ставит куки
func SetCookie(wr http.ResponseWriter, Obj, Obj1 string){

  t := 86400*5 // Куки на 5 дней
	cookie := &http.Cookie{ Name: Obj,
                          Value: Obj1,
                          HttpOnly: false,
                          MaxAge: t,
                          Path: "/"  }
	http.SetCookie(wr, cookie)
}
//==================================


//================================== Действие: Удаляет куки
func DelCookie(wr http.ResponseWriter, NameCookie string) {
   t := -86400*100 // Куки на 5 дней
	 cookie := &http.Cookie{ Name: NameCookie,
                           Value: "",
                           MaxAge: t,
                           Path: "/"    }
	 http.SetCookie(wr, cookie)
}
//==================================


//==================================
func ReadCookie(req *http.Request, NameCookie string) string {
	ck, err := req.Cookie(NameCookie)

	// Error
	if err != nil {
	   return ""
	}  else {
	   return string(ck.Value)
	}

	// На протяжении месяца
	// ck.Expires.Month().String()
}
//==================================
/////////////////////////////////////////////////////////////////////////////////////////////////// #### COOKIES




/////////////////////////////////////////////////////////////////////////////////////////////////// #### DATABASE
// ===========================================================->
// -- Подключение к БД (нужна библиотека - 	 "database/sql";  _"github.com/go-sql-driver/mysql")
var db *sql.DB // Var of connect
func Dbini(){ // Function
	 dbs, err:= sql.Open("mysql", "root:@tcp(localhost:3306)/golang")
	 if err != nil{ log.Println(err) }
	 db=dbs
}

/* How to use :
    func main(){
                // Connect to DB
                 Dbini()
                 defer db.Close()
             }*/
// ===========================================================->




// ===========================================================->
// -- ЗАПРОС к БД (нужна библиотека - 	 "database/sql";  _"github.com/go-sql-driver/mysql")
query := fmt.Sprintf("INSERT INTO test(name, surname, middle) VALUES ('%s', '%s', '%s')","gopher", "shit","shit1")
result, err := db.Exec(query)
           // or
result, err := db.Exec("INSERT INTO test(name, surname, middle) VALUES (?, ?, ?)","gopher", "shit","shit1")
if err != nil { fmt.Println(err) }
// ===========================================================->




// ===========================================================->
// -- ВЫВОД с БД (нужна библиотека - 	 "database/sql";  _"github.com/go-sql-driver/mysql")
result, err := db.Query("SELECT name, surname, middle FROM test")
if err != nil { fmt.Println(err) }

fmt.Fprintf(wr, `<html>
              <head></head>
               <body> `)
i := 1
for result.Next(){

  var name string
  var surname string
  var middle string
  err = result.Scan( &name, &surname, &middle)
  if err != nil{ fmt.Println(err) }

  fmt.Fprintf(wr, "Запись №%v -  %s  %s %s <br>", i,name, surname, middle)
   i++
}

fmt.Fprintf(wr, `</body></html>`)
// ===========================================================->





// ===========================================================->
// -- ВЫВОД С БД -  1 ЗАПИСИ
var name string // переменная в которую будет присвоеные данные с БД
// Сначала идет переменная подключения, потом запрос на 1 результат(QueryRow) и присваивание(Scan)
db.QueryRow("SELECT name FROM test WHERE id=?", U).Scan(&name)
fmt.Fprintf(wr, "Name : %s",name)
// ===========================================================->
/////////////////////////////////////////////////////////////////////////////////////////////////// #### DATABASE



/////////////////////////////////////////////////////////////////////////////////////////////////// #### TEMPLATES
// ===========================================================->
// -- ВЫВОД страниц файла типа index.html (нужна библиотека - "html/template")
func Template(wr http.ResponseWriter, req *http.Request){

	templates := template.Must(template.ParseFiles("templates/index.html")) // PREPARE FILE

	templates.ExecuteTemplate(wr, "index.html", nil) // SHOW FILE
	                     // or with catcher of error
	 if err :=  templates.ExecuteTemplate(wr, "index.html", nil)
   err != nil { http.Error(wr, err.Error(), http.StatusInternalServerError) }
}
// ===========================================================->





// ===========================================================->
// -- ВЫВОД страниц файла типа index.html с вставками (нужна библиотека - "html/template")
// -- Это записать в fns.go, интерфейс -> type Mst map[string]interface{}
func Template(wr http.ResponseWriter, req *http.Request){

 var I [5]Mst
     I[0]=Mst{"S1":"ddddd22","S2":"ssssddd32"}
     I[1]=Mst{"S1":"ddddd22","S2":"ssssddd32"}
     I[2]=Mst{"S1":"ddddd33","S2":"ssssddd22"}

		  var R [5]Mst
			    R[0]=Mst{"Fam":"Степа","Names":"Котелков"}
      		R[1]=Mst{"Fam":"Вася","Names":"Прохоров"}
		      R[2]=Mst{"Fam":"Петя","Names":"Степанов"}
          R[3]=Mst{"Fam":"Герда","Names":"Мацареловна"}


	templates := template.Must(template.ParseFiles("templates/index.html")) // PREPARE FILE

	p         :=`<table class='table'> <tr> <td>Head Office</td> </tr> </table>`
  profile   := Mst{"T":p, "Y":"Наименование отчета", "U":I, "NM":R} // Mst - in fns.go

// In file index.html that was PREPARED , inserted profile vars that above
 if err :=  templates.ExecuteTemplate(wr, "index.html", profile)
 err != nil { http.Error(wr, err.Error(), http.StatusInternalServerError) }
}
/* FILE - index.html(folder: templates)
<html>
<head>
  <!-- Latest compiled and minified CSS -->
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">

<!-- Optional theme -->
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap-theme.min.css" integrity="sha384-fLW2N01lMqjakBkx3l/M9EahuwpSfeNvV63J5ezn3uZzapT0u7EYsXMjQV+0En5r" crossorigin="anonymous">

<!-- Latest compiled and minified JavaScript -->
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
</head>
<body>

  <h1>ITS IS TEMPLATE</h1>
      <h1>{{.T}}</h1>
      <h1>{{.Y}}</h1>


<table class="table table-bordered table-striped">
{{range .U}}
   <tr><td>{{.S1}}</td><td> {{.S2}}</td></tr>

{{end}}
</table>

<h1>Собаки и люди</h1>
<table class="table table-bordered table-striped">
{{range .NM}}
   <tr><td>{{.Fam}}</td><td> {{.Names}}</td></tr>
{{end}}
</table>


</body>
</html>*/
// ===========================================================->
/////////////////////////////////////////////////////////////////////////////////////////////////// #### TEMPLATES
