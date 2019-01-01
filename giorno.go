package giorno

import (
    "fmt"
    "time"
    "strconv"
    "net/http"
)

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/fai", fai)
    http.HandleFunc("/robots.txt", robots)
}

func root(w http.ResponseWriter, r *http.Request) {

fmt.Fprintf(w,mioForm)
}

const shortForm = "2006-01-02"

func fai(w http.ResponseWriter, r *http.Request) {
sett := []string {"lunedì","martedì","mercoledì","giovedì","venerdì","sabato","domenica"}
week := []string {"Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"}
a := 0
b := 0
c := 0

xa, e1 := strconv.Atoi(r.FormValue("bda"))
if e1 != nil {
	fmt.Fprintf(w, "<html><body>ERR. Anno. <a href=/>correggi</a></body></html>")
	return
}
xm, e2 := strconv.Atoi(r.FormValue("bdm"))
if e2 != nil {
	fmt.Fprintf(w, "<html><body>ERR. Mese. <a href=/>correggi</a></body></html>")
	return
}
xg, e3 := strconv.Atoi(r.FormValue("bdg"))
if e3 != nil {
	fmt.Fprintf(w, "<html><body>ERR. Giorno. <a href=/>correggi</a></body></html>")
	return
}
fmt.Fprintf(w,"<html><body>")
x := fmt.Sprintf("%d-%02d-%02d", xa, xm, xg)
fmt.Fprintf(w,"Giorno di nascita: %s\n",x)
k, err := fmt.Sscanf(x,"%d-%d-%d",&a,&b,&c)
if err != nil || k != 3 {
  fmt.Fprintf(w, "ERRORE: <a href=/>riprova</a></body></html>")
  return
}
t, e4 := time.Parse(shortForm, x)
if e4 != nil {
  fmt.Fprintf(w, "ERRORE: <a href=/>riprova</a></body></html>")
  return
}
settday := "boh"
for j := 0; j < len(week); j++ {
	if week[j] == t.Weekday().String() {
		settday = sett[j]
		break
	}
}
fmt.Fprintf(w, "<p>Giorno della settimana: %s", settday)
fmt.Fprintf(w, "<p>Anni in cui il compleanno cade di %s<p>", settday)
m := t.Month()
g := t.Day()
we := t.Weekday()

for y := t.Year(); y <= 2099; y++ {
	z := time.Date(y, m, g, 0, 0, 0, 0, time.UTC)
	if z.Weekday() == we {
		fmt.Fprintf(w, "%d-%02d-%02d età %d<br>", y, m, g, y-t.Year())
	}
}

fmt.Fprintf(w,"<p>Grazie per aver usato questa app.<br><a href=/>altra data</a></body></html>")
}

func robots(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "User-agent: *\nDisallow: /\n")
}

const mioForm = `
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body>
<h1>Giorno di Nascita</h1>
Questa app calcola il giorno della settimana corrispondente 
al giorno di nascita indicato.
<p>Di seguito sono visualizzati gli anni (dalla nascita al 2099)<br>
in cui il compleanno cade nello stesso giorno della settimana.
<p>Date impossibili generano risultati imprevedibili.
<form method="POST" action="/fai">
Data di nascita:<br>
Anno: <input name="bda" type="number" min=1900 max=2099 autofocus></input><br>
Mese: <input name=bdm type=number min=1 max=12></input><br>
Giorno: <input name=bdg type=number min=1 max=31></input><br>
<button type=submit>Ok</button>
</form>
___<br><br>
Questo sito non utilizza alcun cookie<br>
e non conserva nessun dato.
</body></html>
`
