package main

import "fmt"

func main() {
	var zz string = "################################################################"
	fmt.Println(zz)

	var aa string = "Integrity"
	var ab string = "We need to become very serious about reliability"
	var ac string = "There are two driving forces behind integrity:"
	var ad string = "- Integrity is about every allocution, read and write of memory being accurate, consistent and efficient"
	var ae string = "-- The type system is critical to making sure we have this micro level of integrity"
	var af string = "- Integrity is about every transformation being accurate, consistent and efficient"
	var ag string = "-- Writing less code and error handling is critical to making sure we have this macro level of integrity"
	fmt.Println(aa)
	fmt.Println(ab)
	fmt.Println(ac)
	fmt.Println(ad)
	fmt.Println(ae)
	fmt.Println(af)
	fmt.Println(ag)
	fmt.Println(zz)

	var ba string = "Write less code"
	var bb string = "- There have been studies that have researched the number of bugs you can expect to have in your software"
	var bc string = "- The industry average is around 15 to 50 per 1000 lines of code."
	var bd string = "- One simple way to reduce the number of bugs, and increase the integrity of your software, is to write less code."
	fmt.Println(ba)
	fmt.Println(bb)
	fmt.Println(bc)
	fmt.Println(bd)
	fmt.Println(zz)

	var ca string = "48 critical failures were found in a study looking at a couple hundred bugs in Cassandra, HBase, HDFS, MapReduce, and Redis."
	var cb string = "- 92%: Failures from bad error handling"
	var cc string = "-- 35%: Incorrect handling"
	var cd string = "--- 25%: Simply ignoring an error"
	var ce string = "--- 8%: Catching the wrong exception"
	var cf string = "--- 2%: Incomplere TODOs"
	fmt.Println(ca)
	fmt.Println(cb)
	fmt.Println(cc)
	fmt.Println(cd)
	fmt.Println(ce)
	fmt.Println(cf)
	fmt.Println(zz)

	var da string = "Failure is expected, failure is not an odd case."
	var db string = "Design systems that help you identify failure"
	var dc string = "Design systems that can recover from failure"
	var dd string = "- JBD"
	fmt.Println(da)
	fmt.Println(db)
	fmt.Println(dc)
	fmt.Println(dd)
	fmt.Println(zz)

	var ea string = "Product excellence is the difference between something that only works under conditions,"
	var eb string = "and something that only breaks under certain conditions"
	var ec string = "- Kelsey Hightower"
	fmt.Println(ea)
	fmt.Println(eb)
	fmt.Println(ec)
	fmt.Println(zz)

	var fa string = "Making things easy to do is false economy."
	var fb string = "Focus on making things easy to understand and the rest will follow."
	var fc string = "- Peter Bourgon"
	fmt.Println(fa)
	fmt.Println(fb)
	fmt.Println(fc)
	fmt.Println(zz)
}
