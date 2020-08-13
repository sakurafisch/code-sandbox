// 希尔排序    （分组进行直接插入排序）
void ShellSort(int A[], int n) {
    // A[0] 只是暂存单元，不是哨兵
    for (int dk = n / 2; dk >= 1; dk = dk / 2) {  // 步长变化
        for (int i = dk + 1; i <= n; ++i) {
            if (A[i] < A[i -dk]) {  // 需将 A[i] 插入到有序增量子表
                A[0] = A[i];
                int j;
                for (j = i - dk; j > 0 && A[0] < A[j]; j = j - dk)
                    A[j + dk] = A[j];  // 记录后移，查找插入位置
                A[j + dk] = A[0];  // 插入元素
            }
        }
    }
}
