int Partition(int A[], int low, int high);

// 快速排序
void QuickSort(int A[], int low, int high) {
    if (low < high) {
        int pivotpos - Partition(A, low, high);
        QuickSort(A, low, pivotpos - 1);
        QuickSort(A, pivotpos + 1, high);
    }
}

// Partition() 用于划分子表
int Partition(int A[], int low, int high) {
    int pivot = A[low];             // 将当前表中第一个元素设为枢轴
    while (low < high) {            // 循环跳出条件
        while (low < high && A[high] >= pivot) --high;
        A[low] = A[high];           // 比枢轴小的元素移到左端
        while (low < high && A[low] <= pivot) ++low;
        A[high] = A[low];           // 比枢轴大的元素移到右端
    }
    A[low] = pivot;                 // 枢轴元素存放到最终位置
    return low;                     // 返回存放枢轴的最终位置
}
