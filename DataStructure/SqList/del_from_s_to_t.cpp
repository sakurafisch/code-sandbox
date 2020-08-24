#include "SqList.cpp"

// 删除顺序表 L 中给定值 s 与 t 之间的所有元素
bool del_from_s_to_t(SqList &L, ElemType s, ElemType t) {
    if (L.length == 0 || s >= t)
        return false;
    int k = 0;
    for (int i = 0; i < L.length; ++i) 
        if (L.data[i] >= s && L.data[i] <= t)
            ++k;
        else 
            L.data[i - k] = L.data[i];
    L.length -= k;
    return true;
}

// 删除有序表 L 中给定值 s 与 t 之间的所有元素
bool del_from_s_to_t_2(SqList &L, ElemType s, ElemType t) {
    if (s >= t || L.length == 0)
        return false;
    int i, j;

    // 找到首个大于等于 s 的元素下标
    for (i = 0; i < L.length; ++i)
        if (L.data[i] >= s)
            break;
    if (i >= L.length)
        return false;

    // 找到首个大于 t 的元素下标
    for (j = i; j < L.length; ++j) 
        if (L.data[j] > t) 
            break;
    
    while (j < L.length)
        L.data[i++] = L.data[j++];
    L.length = i;
    return true;
}