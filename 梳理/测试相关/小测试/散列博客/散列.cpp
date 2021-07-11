# include<bits/stdc++.h>
using namespace std;


// ��������: �ж� B �е�Ԫ���Ƿ������ A�������ؽ���� 
bool* isLiving1(int* A,int* B,int lengthA,int lengthB){
	bool* hashTable = constructHashTale(A,lengthA,10000); // �� A Ϊ�������� hashTable�� 
	bool* C = new bool[lengthB];
	for (int i=0;i<lengthB;i++){
		
		// hashTable[B[i]] == trueʱ:  ��ʾ hashTable �д���Ԫ�� B[i]���� B[i] ������ A �С� 
		// hashTable[B[i]] == falseʱ: ��ʾ hashTable �в�����Ԫ�� B[i]���� B[i] ������ A �С� 
		C[i] = hashTable[B[i]];
	}
	return C;
}

// A: Դ����
// lengthA: Դ���鳤��
// maxValue: Դ��������ȡֵ 
// ��������: ������ A Ϊ��������ɢ������ hashTable,�����ء� 
bool* constructHashTale(int* A,int lengthA,int maxValue){
	bool* hashTable = new bool[maxValue+1];	// ���г�ʼ��ʱ������Ԫ�ض�Ϊ false 
	for (int i=0;i<lengthA;i++){
		hashTable[A[i]] = true;
	}
	return hashTable;
} 




bool* isLiving(int* A,int* B,int lengthA,int lengthB){
	bool* C = new bool[lengthB];
	for (int i=0;i<lengthB;i++){
		
		// �ж� B[i] �Ƿ����������A 
		C[i] = false;	// Ĭ�ϲ����� 
		for (int t=0;t<lengthA;t++){
			
			// ���� 
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
