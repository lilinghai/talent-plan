package main

import "sync"

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	//sort.Slice(src, func(i, j int) bool { return src[i] < src[j] })
	//MergeSortSequence(src,0,len(src)-1)
	MergeSortCurrency(src,0,len(src)-1)
	//QuickSortCurrency(src,0,len(src)-1)
}

var max int=1<<10

//并行归并排序
func MergeSortCurrency(array []int64, l int,r int){
	if l >= r {
		return
	}
	if r-l<=max{
		MergeSortSequence(array,l,r)
	}else {
		m := (l + r) / 2
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			MergeSortCurrency(array, l, m)
		}()

		go func() {
			defer wg.Done()
			MergeSortCurrency(array, m+1, r)
		}()

		wg.Wait()
		merge(array, l, m, r);
	}
}

//串行归并排序
func MergeSortSequence(array []int64, l int,r int) {

	//如果只有一个元素，那就不用排序了
	if l >= r {
		return
	}
	//取中间的数，进行拆分
	m:= (l + r) / 2

	//左边的数不断进行拆分
	MergeSortSequence(array, l, m)

	//右边的数不断进行拆分
	MergeSortSequence(array, m + 1, r)

	//合并
	merge(array, l, m, r);

}

func merge(array []int64, l int,m int,r int) {
	leftArr:=make([]int64,m-l+1)
	rightArr:=make([]int64,r-m)
	var li,ri int
	for k:=l;k<r+1;k++{
		if k<=m{
			leftArr[li]=array[k]
			li++
		}else{
			rightArr[ri]=array[k]
			ri++
		}
	}
	var i,j int

	//比较这两个数组的值，哪个小，就往数组上放
	//2,6,10,1,7,8
	for i<len(leftArr)&&j<len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			array[l] = leftArr[i]
			i++
		} else {
			array[l] = rightArr[j]
			j++
		}
		l++
	}

	//如果左边的数组还没比较完，右边的数都已经完了，那么将左边的数抄到大数组中(剩下的都是大数字)
	for i<len(leftArr) {
		array[l]=leftArr[i]
		i++
		l++
	}
	//如果右边的数组还没比较完，左边的数都已经完了，那么将右边的数抄到大数组中(剩下的都是大数字)
	for j<len(rightArr) {
		array[l]=rightArr[j]
		j++
		l++
	}
}

func QuickSortCurrency(arr[]int64,start int ,end int){
	if start>=end{
		return
	}
	//获取分区点
	if end-start<=max{
		QuickSortSequence(arr,start,end)
	}else {
		p := partition(arr, start, end)
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			QuickSortSequence(arr, start, p-1)

		}()

		go func() {
			defer wg.Done()
			QuickSortSequence(arr, p+1, end)
		}()

		wg.Wait()
	}
}

func QuickSortSequence(arr[]int64,start int ,end int){
	if start>=end{
		return
	}
	//获取分区点
	p:=partition(arr,start,end)
	//fmt.Println(arr,p)
	QuickSortSequence(arr,start,p-1)
	QuickSortSequence(arr,p+1,end)
}

func partition(arr[]int64,start int,end int)int{
	base:=arr[end]
	//双指针，j指向第一个大于左边的元素
	j:=start
	for i:=start;i<end;i++{
		if arr[i]<base{
			tmp:=arr[i]
			arr[i]=arr[j]
			arr[j]=tmp
			j++
		}
	}
	arr[end]=arr[j]
	arr[j]=base
	return j
}