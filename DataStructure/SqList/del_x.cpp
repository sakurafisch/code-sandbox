#include "./SqList.cpp"

// 删除顺序表 L 中所有值为 x 的元素

void del_x_1(SqList &L, ElemType x) {
    int k = 0; // 记录值不等于 x 的元素
    for (int i = 0; i < L.length; ++i) 
        if (L.data[i] != x) {
            L.data[k] = L.data[i];
            ++k;
        }
    L.length = k;
}

void del_x_2(SqList &L, ElemType x) {
    int k = 0, i = 0;
    while (i < L.length) {
        if (L.data[i] == x)
            ++k;
        else 
            L.data[i - k] = L.data[i];
        ++i;
    }
    L.length = L.length - k;
}
