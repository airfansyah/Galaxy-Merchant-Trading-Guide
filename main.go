package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "regexp"
  "strconv"
)

var m = map[string]int{
"I" : 1,
"V" : 5,
"X" : 10,
"L" : 50,
"C" : 100,
"D" : 500,
"M" : 1000,
}

var mCredits = map[string][]string{}
var mQuestion = map[string][]string{}
var def = map[string]string{}
var value string

func isSymbolValid(a string)bool{
  if m[a] != 0 {
    return true
  }
  return false
}

func countRoman(a []string)int{
  var x []int
  var ttl int
  ttl = 0
  for _, a := range a {
    x = append(x, m[a])
  }
  length := len(x)
  for i:=0; i < length; i++ {
    if i > 0 {
      if x[i] > x[i-1] {
        ttl += x[i] - x[i-1] - x[i-1]
        continue
      }else{
        ttl += x[i]
      }
    }else {
      ttl += x[i] //10+30+1+1
    }
  }
  return ttl
}

func countRomanCredits(a []string, b []string, c string)float64{
  var x []int
  var y []int
  var ttl float64

  for _, a := range a {
     x = append(x, m[a])
   }
  length := len(x)
  for i:=0; i < length; i++ {
    if i > 0 {
      if x[i] > x[i-1] {
        ttl += float64(x[i]) - float64(x[i-1]) - float64(x[i-1])
        continue
      }else {
        ttl += float64(x[i])
      }
    }else {
      ttl += float64(x[i])
    }
  }
  var ttl2 float64
  ttl2=0
  for _, b := range b {
    y = append(y, m[b])
  }
  lengthb := len(y)
  for i:=0; i < lengthb; i++ {
    if i > 0 {
      if float64(y[i]) > float64(y[i-1]) {
        ttl2 += float64(y[i]) - float64(y[i-1]) - float64(y[i-1])
        continue
      }else {
        ttl2 += float64(y[i])
      }
    }else {
      ttl2 += float64(y[i])
    }
  }
  cc, _ := strconv.Atoi(c)
  d := float64(ttl2)*(float64(cc)/float64(ttl))
  return d
}


func handlingFile(a string)[]string{
  readFile, err := os.Open(a)
  if err != nil {
    fmt.Println("error : ", err)
  }
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

  var textPerLine []string
  for fileScanner.Scan() {
    textPerLine = append(textPerLine, fileScanner.Text())
  }
  readFile.Close()
  return textPerLine
}

func main() {
  file := handlingFile("input.txt")


  for _, text := range file {
    if strings.Contains(text, "is") && !strings.Contains(text, "how") && !strings.Contains(text, "Credits") { //define
      split := strings.Split(text," is ")
      if split[0] != "how much" && split[0] != "how many" {
        def[split[0]] = split[1]
      }
    }
    if strings.Contains(text, "is") && strings.Contains(text, "Credits") && !strings.Contains(text, "how"){ //define Credits
      var ra []string
      regex := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
      val := regex.FindAllString(text, -1)
      value = val[0] // value static 34

      split := strings.Split(text," is ") // glob glob Silver 34 Credits
      subSplit := strings.Split(split[0], " ") //[glob glogb Silver]
      ai := len(subSplit)-1 //2
      unit := subSplit[ai:] //Silver for Key

      superSplit := strings.Split(split[0], " ") //[glob glogb Silver]
      superSplit = superSplit[:len(superSplit)-1] // glob glob
      superSplit = append(superSplit,value) //[glob glob 34]
      //debug
      // fmt.Println("sequence DEFINE UNIT -",unit[0])
      // fmt.Println("    parameter DEFINE -", superSplit)
      // fmt.Println("               value -", value)
      for _, a := range superSplit {
        ra = append(ra,a)
        mCredits[unit[0]] = ra
      }
      //debug
      // fmt.Println("                maps -",mCredits)
      // fmt.Println("-----------------------")

    }
    if strings.Contains(text, "?") && strings.Contains(text, "how") && strings.Contains(text, "is"){      //question and answer
        var roman []string
        var roman2 []string
        split := strings.Split(text," is ")

        if split[0] == "how much" {
          var str string
          subSplit := strings.Split(split[1], " ")
          subSplit = subSplit[:len(subSplit)-1] //[pish tegj glob glob]

          for _,a := range subSplit {
            if a != "?" {
              roman = append(roman,def[a])
            }
          }
          str = strings.Join(subSplit," ")
          fmt.Println(str,"is",countRoman(roman))
        }

      if split[0] == "how many Credits" {
        var qr []string
        split := strings.Split(text," is ")
        subSplit := strings.Split(split[1]," ")
        subSplit = subSplit[:len(subSplit)-1] //[prok pish Silver] question
        superSplit := subSplit[:len(subSplit)-1] // [prok pish]
        ai := len(subSplit)-1
        unit := subSplit[ai:] //Silver
        //debug
        // fmt.Println("--------------------------")
        // fmt.Println("sequence QUESTION unit -",unit)
        // fmt.Println("    parameter QUESTION -", superSplit)
        for _, a:= range superSplit {
            if a!="" {
              roman2 = append(roman2,def[a])
              qr = append(qr,a)
              mQuestion[unit[0]] = qr
            }
        }
        //debug
        // fmt.Println("                  maps -", mQuestion)
        // fmt.Println("--------------------------")
        var str []string
        var rmValQ string
        var rmValD string
        var valS string
        for tt := range mQuestion[unit[0]] {
        //  str += mQuestion[unit[0]][tt]+" "
          str = append(str, mQuestion[unit[0]][tt])
          rmValD += def[mCredits[unit[0]][tt]]
          rmValQ += def[mQuestion[unit[0]][tt]]
          valS = mCredits[unit[0]][len(mCredits[unit[0]])-1]
        }
        //debug
        // fmt.Println("      DEF",rmValD)
        // fmt.Println(" QUESTION",rmValQ)
        // fmt.Println("    VALUE",valS)
        uu := strings.Split(rmValD,"")
        ii := strings.Split(rmValQ,"")
        strs := strings.Join(str," ")
        fmt.Println(strs,"is",countRomanCredits(uu, ii, valS),"Credits")
      }
    }
    if !strings.Contains(text, "is") {
      fmt.Println("I have no idea what you are talking about")
    }
  }
}
