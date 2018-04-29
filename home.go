package main

import (

  "net/http"
  "fmt"
"strconv"
"strings"
"encoding/json"
    "path"
    "html/template"


)

type Data struct {
    Numara string
}
type WebData struct {
     Numaraa  string

    Error string
    Fiyat int
}
var wd *WebData = &WebData{
        Numaraa:"",
        Fiyat:  0,
        Error:"",
}


// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {


    fp := path.Join("", "home.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
//wd := WebData{}
    if err := tmpl.Execute(w, &wd); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

// AJAX Request Handler
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
    //parse request to struct
    //fp := path.Join("templates", "ajax-json.html")
    //tmpl, _ := template.ParseFiles(fp)
   //tmpl, _ := template.ParseFiles("templates/ajax-json.html")

    var d Data
    err := json.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }


    wd.Numaraa=d.Numara
fmt.Printf("Numara:%s\n", wd.Numaraa)




num:=strings.Replace(wd.Numaraa, " ", "", -1)
number, err := strconv.ParseInt(num, 10, 64)
if err != nil || 908508853999 <= number ||   number<=908508850000 {
  wd.Error="Lütfen Geçerli Numara Giriniz !!"
  wd.Fiyat=0
}
if err == nil && 908508853999 >= number &&   number>=908508850000 {


if son4ayni(num){
wd.Fiyat=1000+1000*0.18
} else if onPrefixSimetrik(num) {
wd.Fiyat=1000+1000*0.18

} else if son3ayni(num)  {
wd.Fiyat=500+500*0.18
}else if   num[8]==num[10] && num[9]==num[11] {
wd.Fiyat=500+500*0.18

} else if   num[8]==num[9] && num[10]==num[11] {
wd.Fiyat=250+250*0.18

} else if   num[8]==num[10]  {
wd.Fiyat=100+100*0.18
} else if   ardisik(num){
wd.Fiyat=100+100*0.18

} else if   num[9]=='1'  && num[10]=='9'  {
wd.Fiyat=50+50*0.18
}

  wd.Error="Girilen Numaranın Fiyatı + KDV : "
}

tmpl, _ := template.ParseFiles("index.html")

tmpl.Execute(w, &wd)
fmt.Printf("%s\n", wd.Error)
if wd.Fiyat >0{
fmt.Printf("%d\n", wd.Fiyat)}

}

func main() {

    http.HandleFunc("/", defaultHandler)
      http.HandleFunc("/ajax", ajaxHandler)
      err := http.ListenAndServe(":8080", nil)
      if err != nil {
      }
}
func ardisik(num string)bool {
  if strings.Contains(num, "0123") || strings.Contains(num, "1234") ||strings.Contains(num, "2345") ||strings.Contains(num, "3456"){
    return true
  }else {
    return false
  }
}
func son4ayni(num string)bool {
if num[11]==num[10] && num[10]==num[9] && num[9]==num[8]{
  return true
  }else {
    return false
  }
}
func onPrefixSimetrik(num string)bool {
if  num[5]==num[9] && num[6]==num[10] && num[7]==num[11]{
  return true
  }else {
    return false
  }
}
func son3ayni(num string)bool {
if  num[9]==num[10] && num[10]==num[11] {
  return true
  }else {
    return false
  }
}
