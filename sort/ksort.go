package main

/**
  eg.
    a := []int{8,5,7,16,15,23, 1, 6, 3, 54, 68}
	QuickSort(a,0,len(a) - 1)
	fmt.Println(a) //[1 3 5 6 7 8 15 16 23 54 68]
 */
func QuickSort(a []int, _left, _right int) {
	lena := len(a)
	if _left > _right || _left < 0 || _right > lena - 1{
		return
	}

    temp := a[_left]
    left := _left
    right := _right

    for left != right {
    	for left < right {
    		if a[right] >= temp {
    			right --
			} else {
				a[left] = a[right]
				left ++
				break
			}
		}

    	for left < right {
			if a[left] >= temp {
				a[right] = a[left]
				right --
				break
			} else {
				left ++
			}
		}
	}
	a[right] = temp
	QuickSort(a,_left,left - 1)
	QuickSort(a,right + 1,_right)

	return
}
