# include<bits/stdc++.h>
using namespace std;


// 函数作用: 判断 B 中的元素是否存在于 A，并返回结果。 
bool* isLiving1(int* A,int* B,int lengthA,int lengthB){
	bool* hashTable = constructHashTale(A,lengthA,10000); // 以 A 为基础构建 hashTable。 
	bool* C = new bool[lengthB];
	for (int i=0;i<lengthB;i++){
		
		// hashTable[B[i]] == true时:  表示 hashTable 中存在元素 B[i]，即 B[i] 存在于 A 中。 
		// hashTable[B[i]] == false时: 表示 hashTable 中不存在元素 B[i]，即 B[i] 存在于 A 中。 
		C[i] = hashTable[B[i]];
	}
	return C;
}

// A: 源数组
// lengthA: 源数组长度
// maxValue: 源数组的最大取值 
// 函数作用: 以数组 A 为基础构建散列数组 hashTable,并返回。 
bool* constructHashTale(int* A,int lengthA,int maxValue){
	bool* hashTable = new bool[maxValue+1];	// 堆中初始化时，数组元素都为 false 
	for (int i=0;i<lengthA;i++){
		hashTable[A[i]] = true;
	}
	return hashTable;
} 




bool* isLiving(int* A,int* B,int lengthA,int lengthB){
	bool* C = new bool[lengthB];
	for (int i=0;i<lengthB;i++){
		
		// 判断 B[i] 是否存在于数组A 
		C[i] = false;	// 默认不存在 
		for (int t=0;t<lengthA;t++){
			
			// 存在 
			if (A[t]==B[i]) C[i] = true;
		} 
	}
	return C; 
}

int main(){
	int* A = new int[10];
	for (int i=0;i<10;i++){
		A[i]=i;
	}
	
	int* B = new int[10];
	for (int i=0;i<10;i++){
		B[i]=11*i;
	}
	bool* C = isLiving(A,B,10,10);
	for (int i=0;i<10;i++){
		cout<<C[i]<<endl;
	}
	cout<<"--"<<endl;
	bool* D = isLiving(A,B,10,10);
	for (int i=0;i<10;i++){
		cout<<D[i]<<endl;
	}
	
}
