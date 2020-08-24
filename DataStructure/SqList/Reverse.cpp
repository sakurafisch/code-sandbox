// 就地逆置顺序表
#include "./SqList.cpp"

void Reverse(SqList &L) {
    ElemType temp;
    for (int i = 0; i < L.length / 2; ++i) {
        temp = L.data[i];
        L.data[i] = L.data[L.length - 1];
        L.data[L.length - i - 1] = temp;
    }
}