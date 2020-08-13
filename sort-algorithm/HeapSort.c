void HeadAdjust(int A[], int k, int len);
void BuildMaxHeap(int A[], int len);

// 堆排序
void HeapSort(int A[], int len) {
    BuildMaxHeap(A, len);       // 建堆
    for (int i = len; i > 1; --i) {     // n - 1 趟的交换和建堆过程
        swap(A[i], A[1]);       // 输出堆顶元素（和堆底元素交换）
        HeadAdjust(A, 1, i - 1);    // 调整，把剩下的 i - 1 个元素整理成堆
    }
}

// 建大顶堆
void BuildMaxHeap(int A[], int len) {
    for (int i = len / 2; i > 0; --i) {     // 从 i == (n / 2) ~ 1，反复调整堆
        HeadAdjust(A, i, len);
    }
}

// 调整以 k 为根的子树
void HeadAdjust(int A[], int k, int len) {
    A[0] = A[k];        // A[0] 暂存子树的根节点
    for (int i = k << 1; i <= len; i <<= 2) {   // 沿 key 较大的子结点向下筛选
        if (i < len && A[i] < A[i + 1])
            ++i;        // 取 key 较大的结节点的下标
        if (A[0] >= A[i])   //筛选结束
            break;      
        else {
            A[k] = A[i];    // 将 A[i] 调整到双亲结点上
            k = i;          // 修改 k 值，以便继续向下筛选
        }
    }
    A[k] = A[0];            // 被筛选结点的值放入最终位置
}
