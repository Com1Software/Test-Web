package main

import (
	"net"
	"net/http"
	"os/exec"
	"runtime"

	"fmt"
	"os"
)

//----------------------------------------------------------------
func main() {
	fmt.Println("Go Web Server")
	switch {
	//-------------------------------------------------------------
	case len(os.Args) == 2:

		fmt.Println("Not")

		//-------------------------------------------------------------
	default:

		fmt.Println("Server running....")
		fmt.Println("Listening on port 8080")
		fmt.Println("")
		//------------------------------------------------ Home Page
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			xdata := InitPage()
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ About Page
		http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
			xdata := AboutPage()
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ PlayGround Page
		http.HandleFunc("/playground", func(w http.ResponseWriter, r *http.Request) {
			xdata := PlayGroundPage()
			fmt.Fprint(w, xdata)
		})

		//------------------------------------------------- Static Handler

		fs := http.FileServer(http.Dir("static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))

		//------------------------------------------------- Start Server
		Openbrowser("http://localhost:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}
}

// Openbrowser : Opens default web browser to specified url
func Openbrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "linux":
		cmd = "chromium-browser"
		args = []string{""}

	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func InitPage() string {
	//---------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	xip := "localhost"
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Web Server Start Page</title>"
	xdata = xdata + "  <link REL='StyleSheet' TYPE='text/css' HREF='static/css/style.css'>"
	//------------------------------------------------------------------------
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body>"
	xdata = xdata + "<p>Go Web Server.</p>"
	//---------
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			xip = fmt.Sprintf("%s", ipv4)
		}
	}
	xip = "localhost"
	xdata = xdata + "<p> Host Port IP : " + xip + "</p>"

	xdata = xdata + "  <A HREF='http://" + xip + ":8080/about'> [ About ] </A>  "

	xdata = xdata + "  <A HREF='http://" + xip + ":8080/playground'> [ PlayGround ] </A>  "

	xdata = xdata + "<BR><BR>Go Web Server"
	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata
}

//----------------------------------------------------------------
func AboutPage() string {
	//---------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>About Page</title>"
	xdata = xdata + "  <link REL='StyleSheet' TYPE='text/css' HREF='static/css/style.css'>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: lightblue;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: white;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p {"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<p>Go Web Server</p>"
	xdata = xdata + "<div id='txt'></div>"
	//---------
	xdata = xdata + "Go Web Server"
	//------------------------------------------------------------------------
	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

//----------------------------------------------------------------
//----------------------------------------------------------------
func PlayGroundPage() string {
	//---------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>The PlayGround</title>"
	xdata = xdata + "  <link REL='StyleSheet' TYPE='text/css' HREF='static/css/style.css'>"
	//------------------------------------------------------------------------
	xdata = xdata + "<script>"
	//------------------------------------------------------------------------
	xdata = xdata + "function startTime() {"
	xdata = xdata + "  var today = new Date();"
	xdata = xdata + "  var d = today.getDay();"
	xdata = xdata + "  var h = today.getHours();"
	xdata = xdata + "  var m = today.getMinutes();"
	xdata = xdata + "  var s = today.getSeconds();"
	xdata = xdata + "  var ampm = h >= 12 ? 'pm' : 'am';"
	xdata = xdata + "  var mo = today.getMonth();"
	xdata = xdata + "  var dm = today.getDate();"
	xdata = xdata + "  var yr = today.getFullYear();"
	xdata = xdata + "  m = checkTimeMS(m);"
	xdata = xdata + "  s = checkTimeMS(s);"
	xdata = xdata + "  h = checkTimeH(h);"
	//------------------------------------------------------------------------
	xdata = xdata + "  switch (d) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       day = 'Sunday';"
	xdata = xdata + "    break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "    day = 'Monday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "        day = 'Tuesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "        day = 'Wednesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "        day = 'Thursday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "        day = 'Friday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "        day = 'Saturday';"
	xdata = xdata + "}"
	//------------------------------------------------------------------------------------
	xdata = xdata + "  switch (mo) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       month = 'January';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "       month = 'Febuary';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "       month = 'March';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "       month = 'April';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "       month = 'May';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "       month = 'June';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "       month = 'July';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 7:"
	xdata = xdata + "       month = 'August';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 8:"
	xdata = xdata + "       month = 'September';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 9:"
	xdata = xdata + "       month = 'October';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 10:"
	xdata = xdata + "       month = 'November';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 11:"
	xdata = xdata + "       month = 'December';"
	xdata = xdata + "       break;"
	xdata = xdata + "}"
	//  -------------------------------------------------------------------
	xdata = xdata + "  document.getElementById('txt').innerHTML = ' - '+h + ':' + m + ':' + s+' '+ampm+' - '+day+', '+month+' '+dm+', '+yr;"
	xdata = xdata + "  var t = setTimeout(startTime, 500);"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeMS(i) {"
	xdata = xdata + "  if (i < 10) {i = '0' + i};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeH(i) {"
	xdata = xdata + "  if (i > 12) {i = i -12};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	xdata = xdata + "</script>"
	//------------------------------------------------------------------------
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<p>The PlayGround.</p>"
	xdata = xdata + "<div id='txt'></div>"
	xdata = xdata + "<p id='demo1'></p>"
	xdata = xdata + "<button onclick='myFunction1()'>Click me 1</button> - "
	xdata = xdata + "<button onclick='myFunction2()'>Click me 2</button>"
	xdata = xdata + "<p id='demo2'></p>"
	//---------
	xdata = xdata + "Go Web Server"
	//------------------------------------------------------------------------
	xdata = xdata + "<script>"
	//----------
	xdata = xdata + " function myFunction1() {"
	xdata = xdata + "document.getElementById('demo1').innerHTML = 'Hello<BR> World';"
	xdata = xdata + "document.getElementById('demo2').innerHTML = 'Hello<BR> World';"
	xdata = xdata + "  }"
	//----------
	xdata = xdata + " function myFunction2() {"
	xdata = xdata + "document.getElementById('demo1').innerHTML = 'Hello World 2 Upper screen stuff';"
	xdata = xdata + "document.getElementById('demo2').innerHTML = 'Hello World 2 Lower screen stuff';"
	xdata = xdata + "  }"
	//----------
	xdata = xdata + "</script>"
	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

func DateTimeDisplay(xdata string) string {
	//------------------------------------------------------------------------
	xdata = xdata + "<script>"
	xdata = xdata + "function startTime() {"
	xdata = xdata + "  var today = new Date();"
	xdata = xdata + "  var d = today.getDay();"
	xdata = xdata + "  var h = today.getHours();"
	xdata = xdata + "  var m = today.getMinutes();"
	xdata = xdata + "  var s = today.getSeconds();"
	xdata = xdata + "  var ampm = h >= 12 ? 'pm' : 'am';"
	xdata = xdata + "  var mo = today.getMonth();"
	xdata = xdata + "  var dm = today.getDate();"
	xdata = xdata + "  var yr = today.getFullYear();"
	xdata = xdata + "  m = checkTimeMS(m);"
	xdata = xdata + "  s = checkTimeMS(s);"
	xdata = xdata + "  h = checkTimeH(h);"
	//------------------------------------------------------------------------
	xdata = xdata + "  switch (d) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       day = 'Sunday';"
	xdata = xdata + "    break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "    day = 'Monday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "        day = 'Tuesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "        day = 'Wednesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "        day = 'Thursday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "        day = 'Friday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "        day = 'Saturday';"
	xdata = xdata + "}"
	//------------------------------------------------------------------------------------
	xdata = xdata + "  switch (mo) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       month = 'January';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "       month = 'Febuary';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "       month = 'March';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "       month = 'April';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "       month = 'May';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "       month = 'June';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "       month = 'July';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 7:"
	xdata = xdata + "       month = 'August';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 8:"
	xdata = xdata + "       month = 'September';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 9:"
	xdata = xdata + "       month = 'October';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 10:"
	xdata = xdata + "       month = 'November';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 11:"
	xdata = xdata + "       month = 'December';"
	xdata = xdata + "       break;"
	xdata = xdata + "}"
	//  -------------------------------------------------------------------
	xdata = xdata + "  document.getElementById('txt').innerHTML = ' - '+h + ':' + m + ':' + s+' '+ampm+' - '+day+', '+month+' '+dm+', '+yr;"
	xdata = xdata + "  var t = setTimeout(startTime, 500);"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeMS(i) {"
	xdata = xdata + "  if (i < 10) {i = '0' + i};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeH(i) {"
	xdata = xdata + "  if (i > 12) {i = i -12};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	xdata = xdata + "</script>"
	return xdata

}