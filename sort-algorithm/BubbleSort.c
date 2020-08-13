#define false 0
#define true 1
typedef int bool;


// 冒泡排序
void BubbleSort(int A[], int n) {
    for (int i = 0; i < n; ++i) {
        bool flag = false;
        for (int j = n - 1; j > i; --j)
            if (A[j - 1] > A[j]) {
                swap(A[j - 1], A[j]);
                flag = true;
            }
        if (flag == false) {
            return;     // 本次遍历后没有发生交换，说明已经有序，故排序结束
        }
    }
}