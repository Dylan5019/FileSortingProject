
package main

import (
    "io"
    "log"
    "os"
    "path/filepath"
    "regexp"
)

var Orginal_Path = ""
var New_Path = ""



func main() {
   

    err := WriteToFile("peach.txt", "I wrote this first\n")
    if err != nil {
        log.Fatal(err)
    }
    

    err2:= WriteToFile("testA.txt", "I wrote this first\n")
    if err != nil {
        log.Fatal(err2)
    }


      err3 := WriteToFile("testB.txt", "I Wrote this Second\n")
    if err2 != nil {
        log.Fatal(err3)
    }

    




var temp []string
temp = readDir()
var pathNames []string
pathNames = delete_empty(temp)



for i:= range pathNames {
    Orginal_Path = pathNames[i]

    match1, _ := regexp.MatchString(`p([a-z]+)ch`, pathNames[i])
    match2, _ := regexp.MatchString(`LICENSE`, pathNames[i])
    if (match1 == true){
      New_Path = "Peach/"+pathNames[i]
    } else if (match2 == true){
      New_Path = "License/"+pathNames[i]
    } else {
       New_Path = "Test/"+pathNames[i]
    }

    sortFile(Orginal_Path, New_Path);
}



}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.WriteString(file, data)
    if err != nil {
        return err
    }
    return file.Sync()
}


func sortFile(or_Path string, nw_path string){

    Orginal_Path := or_Path
    New_Path := nw_path
    e := os.Rename(Orginal_Path, New_Path)
    if e != nil {
      log.Fatal(e)
    }


}


func readDir() []string{
   files, err := filepath.Glob("*")
    if err != nil {
        log.Fatal(err)
    }
    
   // fmt.Println(files) 

  var txtFiles []string



  r, _ := regexp.Compile(`[a-zA-Z0-9]+.txt`)


  for i:= range files {
    txtFiles = append(txtFiles, r.FindString(files[i]))
}
  
 // fmt.Println(txtFiles)

return txtFiles

}


func delete_empty (s []string) []string {
    var r []string
    for _, str := range s {
        if str != "" {
            r = append(r, str)
        }
    }
    return r
}