# include<bits/stdc++.h>
using namespace std;



// A: Դ����
// length: A�ĳ��� 
int* getPrefixSum(int* A,int length){
	int* prefixSum = new int[length+1];	// ע��Ҫ +1����ΪprefixSum�����A�����һ��Ԫ�ء�
	prefixSum[0] = 0;
	for (int i=1;i<=length;i++){
		prefixSum[i] = prefixSum[i-1] + A[i-1];
	} 
	return prefixSum;
} 

// A: Դ���� 
// length: ����A�ĳ��� 
// times: ��Ҫ��ѯ�Ĵ��� 
int calculate3(int* A,int length,int times){
	int left, right;
	int* prefixSum = getPrefixSum(A, length); 	// ��ȡǰ׺������ 
	while(times--){
		scanf("%d%d",&left,&right);	// ��������
		
		// �����ܺ� 
		int sum = prefixSum[right+1] - prefixSum[left];
		
		printf("%d\n",sum);	// ������ 
	}
}






// A: Դ����
// times: �����ѯ���� 
int calculate2(int* A, int times){
	int left, right;
	while(times--){
		scanf("%d%d",&left,&right);	// ��������
		
		// �����ܺ� 
		int sum = 0;
		for (int i=left;i<=right;i++){
			sum += A[i];
		}
		
		printf("%d\n",sum);	// ������ 
	}
}


// A: ԭ����
// left,right: �������ұ߽� 
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
