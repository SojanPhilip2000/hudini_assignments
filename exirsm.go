
func monkeyCount(n int) []int {
	// Your code here, happy coding!
	var a = []int {}
   
   // write a for loop that pushes value from 1 to n
   var i = 1
   for i=1;i<=n;i++{
	 a = append(a, i)
   }
   return a
 }

 
func FindMultiples(integer, limit int) []int {
	// Your code here
	//first creeate an empty slice
	var a = []int {}
	var i=1
	for i=1;i<=(limit/integer);i++{
	  var intgr = integer * i
	  a= append(a,intgr)
	}
	return a
  }


func PowersOfTwo(n int) []uint64 {
  // your code goes here
  a := []uint64 {}
  for i:=0;i<=n;i++{
    intgr := math.Pow(2,float64(i))
    a= append(a,uint64(intgr))
  }
  return a
}


func Points(games []string) int {
  // your code here!
  points:=0
  for i:= 0; i< len(games); i++{
    x := games[i][0]
    y := games[i][2]
    
    if x > y {
      points = points + 3
    } else if x==y{
      points = points + 1
    } else {
      continue
    }
  }
  return points
  
}

func GetSum(a, b int) int {
	if a>b{
	  temp:=a
	  a=b
	  b=temp
	}
	sum:= 0
	for i:=a;i<=b;i++{
	  sum = sum + i
	}
	return sum
  }

  import "strings"

func FindShort(s string) int {
  // your code
  count:=1000
  word:= strings.Split(s," ")
  for i:= 0; i<len(word);i++{
    if count>len(word[i]){
      count=len(word[i])
    }
  }
  return count
}

func LeastLarger(a []int, i int) int {
	least:=1000
	index := -1
	for j:=0;j<len(a); j++{
	  if a[j]>a[i]{
		if a[j]<least{
		  least = a[j]
		  index = j
		}
	  }
	}
	return index
  }

  func OddCount(n int) int{
	//your code here
	count:=0
	for i:= 0; i<n;i++{
	  if i>0{
		if i%2==1{
		  count=count+1
		}
	  }
	}
	return count
  }

  func ShuffleAnimals() []string {
	s :=[]string {"ant","beaver","cat","dog","elephant","fox","giraffe","hedgehog"}
	rand.Shuffle(len(s), func(i,j int){
		   s[i], s[j] = s[j], s[i]
   })
	   return s
	   panic("Please implement the ShuffleAnimals function")
   }

   package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    sum:=0
    i:=0
    for i=0; i<len(birdsPerDay);i++{
        sum=sum+birdsPerDay[i]
    }
    return sum
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    sum:=0
    i:=0
    w1:=(week-1)*7
    w2:=(week*7)-1
    for i=w1; i<=w2; i++{
        sum=sum+birdsPerDay[i]
    }
    return sum
	panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    i:=0
    for i=0; i<len(birdsPerDay);i+=2{
         birdsPerDay[i]=birdsPerDay[i]+1
    }
    return birdsPerDay
    panic("Please implement the FixBirdCountLog() function")
}


package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgprptime int)  int{
    if avgprptime ==0{
        avgprptime =2
    } 
    
    layersNum:= len(layers)
    estime := layersNum*avgprptime
    return estime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodleCount int, sauceCount float64){
    i:=0
    noodleCount =0
    sauceCount = 0
    for i=0;i<len(layers);i++{
        if layers[i] == "sauce"{
            sauceCount+=0.2
        }
        if layers[i] == "noodles"{
            noodleCount+=50
        } 
    }
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendslist, mylist []string) {
    friendlast:= len(friendslist)-1
    mylast:= len(mylist)-1
    mylist[mylast]=friendslist[friendlast]
    }



// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, number int) []float64{
    i:=0
    a:=[]float64{}
    for i=0; i<len(quantities);i++{
        a=append(a,quantities[i])
        a[i]= quantities[i]/2*float64(number)
    }
    return a
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.


package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    intrest:=0.0
    if balance<0{
        intrest= 3.213
    } else if balance<1000 {
        intrest= 0.5
    } else if balance<5000 {
        intrest= 1.621
    } else {
        intrest= 2.475
    }
    return float32(intrest)
	panic("Please implement the InterestRate function")
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    return float64(InterestRate(balance)/100)*balance
	panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    return Interest(balance)+balance
	panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    year:=0
    for balance< targetBalance{
        balance=AnnualBalanceUpdate(balance)
        year=year+1
    }
    return year
	panic("Please implement the YearsBeforeDesiredBalance function")
}
