package main

import {
	"fmt"
	"math"
}

func main() {
	foo := map[string]int{}
  // Add a value in a map with the `=` operator:
  foo["bar"] = 42
  // Here we update the element of `bar`
  foo["bar"] = 73
  // To retrieve a map value, you can use
  baz := foo["bar"]
  // To delete an item from a map, you can use
  delete(foo, "bar")
}

package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    gross:= map[string]int{}
    gross["quarter_of_a_dozen"]=3
    gross["half_of_a_dozen"]=6
    gross["dozen"]=12
    gross["small_gross"]=120
    gross["gross"]=144
    gross["great_gross"]=1728
    return gross
    
	panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bill:=  map[string]int{}
    return bill
	panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    //to check unit in units
    //the value of unit will be stored in the value, if no value- value =0
    value,exists:=units[unit]
    if exists{
        bill[item]+=value
        return true
    }else{
        return false
    }
	panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    itemvalue, itemexists:= bill[item]
    unitvalue, unitexists:= units[unit]
    newQuantity:= itemvalue - unitvalue
    if !itemexists || !unitexists || newQuantity<0{
        return false
    } 
    if bill[item] == units[unit]{
          delete(bill, item)
        return true;
    } else {
		bill[item] -= units[unit];
        return true;
    }
    return true
    
	panic("Please implement the RemoveItem() function")

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    _, itemexists:= bill[item]
    if !itemexists{
        
        return 0, false
    } else {
        return bill[item],true
    }

	panic("Please implement the GetItem() function")
}
