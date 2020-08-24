#include "./SqList.cpp"

// 删除顺序表 L 中最小值元素结点，并通过引用型参数 value 返回其值
bool Del_Min(SqList &L, ElemType &value) {
    if (L.length == 0) 
        return false;
    value = L.data[0];
    int pos = 0;
    for (int i = 1; i < L.length; ++i) 
        if (L.data[i] < value) {
            value = L.data[i];
            pos = i;
        }
    L.data[pos] = L.data[L.length - 1];
    --L.length;
    return true;
}