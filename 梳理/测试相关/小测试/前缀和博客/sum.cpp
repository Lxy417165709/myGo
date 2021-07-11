# include<bits/stdc++.h>
using namespace std;



// A: 源数组
// length: A的长度 
int* getPrefixSum(int* A,int length){
	int* prefixSum = new int[length+1];	// 注意要 +1，因为prefixSum数组比A数组多一个元素。
	prefixSum[0] = 0;
	for (int i=1;i<=length;i++){
		prefixSum[i] = prefixSum[i-1] + A[i-1];
	} 
	return prefixSum;
} 

// A: 源数组 
// length: 数组A的长度 
// times: 需要查询的次数 
int calculate3(int* A,int length,int times){
	int left, right;
	int* prefixSum = getPrefixSum(A, length); 	// 获取前缀和数组 
	while(times--){
		scanf("%d%d",&left,&right);	// 输入区间
		
		// 计算总和 
		int sum = prefixSum[right+1] - prefixSum[left];
		
		printf("%d\n",sum);	// 输出结果 
	}
}






// A: 源数组
// times: 区间查询次数 
int calculate2(int* A, int times){
	int left, right;
	while(times--){
		scanf("%d%d",&left,&right);	// 输入区间
		
		// 计算总和 
		int sum = 0;
		for (int i=left;i<=right;i++){
			sum += A[i];
		}
		
		printf("%d\n",sum);	// 输出结果 
	}
}


// A: 原数组
// left,right: 区间左右边界 
int calculate(int* A,int left,int right){
	int sum = 0;
	for (int i=left;i<=right;i++){
		sum += A[i];
	}
	return sum;
}






int main() {
	int* A = new int[10];
	for (int i=0;i<10;i++){
		A[i]=i;
	}
	calculate3(A,10,10);
}
