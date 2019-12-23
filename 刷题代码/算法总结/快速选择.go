package main


// 选出第k大的数，第k大不存在则返回最大的(这个规则可以自己设定，r==-1时)
// 如果nums有重复元素，那结果是不准确的，所以必须要求nums无重复元素
func fastSelect(nums []int,left int,right int,k int) int{

	base:=nums[left]
	l,r:=left,right
	for ;l<=r ;  {
		for ;l<=r && nums[l]<=base;  {
			l++
		}
		for ;l<=r && nums[r]>=base;  {
			r--
		}
		// l<r,l<=r都可以
		if l<=r{
			nums[l],nums[r] = nums[r],nums[l]
			l++
			r--
		}
	}
	// 表示第k大不存在，在这里设置
	if r==-1 {
		return fastSelect(nums,0, len(nums)-1,1)
	}
	nums[r],nums[left] = nums[left],nums[r]
	if r== len(nums)-k {
		return nums[r]
	}else{
		if r>len(nums)-k {
			return fastSelect(nums,left,r-1,k)
		}else{
			return fastSelect(nums,r+1,right,k)
		}
	}


}
