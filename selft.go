package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)
func main() {
//   var i int
//   fmt.Print("Type a number: ")
//   fmt.Scan(&i)
//   fmt.Println("Your number is:", i)


// //r2:= rect{height:20, width:10}
// fmt.Printf("width: %d, height: %d\n",r2.width,r2.height)
// //r2.sqaureit()
// fmt.Printf("width: %d, height: %d\n",r2.width,r2.height)

var text string
count:=0
wordscount:= make(map[string]int)
fmt.Println("Enter the text: ");
reader := bufio.NewReader(os.Stdin)
text, _ = reader.ReadString('\n')
words := strings.Split(strings.TrimSpace(text), " ");
///viswa
for _, word:= range words{
  if wordscount[word]==0{
    wordscount[word] =1
  }else{
    wordscount[word]+= 1
  }
}
for word, count:= range wordscount{
  fmt.Printf("Word: %s, Count: %d",word,count)
}
//tharun
for _, word:= range words{
  wordscount[word]+=1
}
fmt.Println(wordscount)

//erics
// for i:=0;i<len(words);i++{
//   _, exists:= wordscount[words[i]]
//   if !exists{
//     for j:=i+1;j<len(words); j++{
//       if words[i]==words[j]{
//         count++
//       }
//     }
//   }
//   wordscount[words[i]] = count
// }
// fmt.Print(wordscount)
}



// type rect struct{
// 	width,height int
// }

// func (r rect)sqaureit(){
// 	r.height = r.width
// }