package main

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
    var s =[]int {2,6,9}
    return s
	panic("Please implement the FavoriteCards function")
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    if index >= len(slice)  || index < 0 {
        return -1
    } else{
        return slice[index]
    }
	panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
        if index < len(slice)  && index >= 0 {
            slice[index]=value
        } else{
            slice = append(slice, value)
        }
       return slice
	panic("Please implement the SetItem function")
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    var newSlice = append(values, slice...)
    return newSlice
	panic("Please implement the PrependItems function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    if index < len(slice)  && index >= 0 {
       return append(slice[:index],slice[index+1:]...)
    } else{
        return slice
    }
	panic("Please implement the RemoveItem function")
}
func monkeyCount1(n int) []int {
    // Your code here, happy coding!
    var a = []int {}
   
   // write a for loop that pushes value from 1 to n
   var i = 1
   for i=1;i<=n;i++{
     a[i]=i
   }
   return a
 }
