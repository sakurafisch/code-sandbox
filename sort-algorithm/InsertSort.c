// 直接插入排序
void InsertSort(int A[], int n) {
    for (int i = 0; i <= n; ++i) {
        if (A[i] < A[i - 1]) {  // 若 A[i] < 前驱，将 A[i] 插入到有序表
            A[0] = A[i];  // 复制为哨兵，A[0] 不存放元素
            int j;
            for (j = i - 1; A[0] < A[j]; --j) {  // 从后往前查找待插入位置
                A[j + 1] = A[0];  // 向后挪位
            }
            A[j + 1] = A[0];  // 插入元素
        }
    }
}

// 折半插入排序
void BinaryInsertSort(int A[], int n) {
    for (int i = 2; i <= n; ++i) {
        A[0] = A[i];  // A[0] 为哨兵
        int low = 1, high = i - 1;  // 设置折半查找的范围
        while (low <= high) {
            int mid = (low + high) / 2;
            if (A[mid] > A[0]) high = mid - 1;
            else low = mid + 1;
        }
        for (int j = i - 1; j >= high + 1; --j)
            A[j + 1] = A[j];  // 向后挪位
        A[high + 1] = A[0];  // 插入元素
    }
}