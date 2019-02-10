package main

/**
  eg.
    a := []int{8,1,6,3,54,68}
	InsertSort(a)
	fmt.Println(a) //[1 3 6 8 54 68]
 */
func InsertSort(a []int) {
	l := len(a)
	for i := 1 ; i < l; i ++ {
		for j := i ; j > 0 ; j-- {
			if a[j] > a[j-1] {
				break
			}
			a[j] , a[j-1] = a[j-1],a[j]
		}
	}
	return
}
