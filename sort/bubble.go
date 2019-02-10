package main

/**
  eg.
    a := []int{8,1,6,3,54,68}
	BubbleSort(a)
	fmt.Println(a) //[1 3 6 8 54 68]
 */
func BubbleSort(a []int) {
	l := len(a)
	for i := 0 ; i < l; i ++ {
		for j := 0 ; j < l - i -1 ; j++ {
			if a[j] > a[j+1] {
				a[j] , a[j+1] = a[j+1],a[j]
			}
		}
	}
	return
}
