package main

/**
  eg.
    a := []int{8,1,6,3,54,68}
	SelectSort(a)
	fmt.Println(a) //[1 3 6 8 54 68]
 */
func SelectSort(a []int) {
	l := len(a)
	var min int
	for i := 0 ; i < l; i ++ {
		min = i
		for j := i + 1 ; j < l ; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[min] , a[i] = a[i],a[min]
	}
	return
}
