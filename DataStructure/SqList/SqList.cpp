#define MaxSize 50
typedef int ElemType;
typedef struct SqList
{
    ElemType data[MaxSize];
    int length;
};


bool ListInsert(SqList &L, int position, ElemType e) {
    if (position < 1 || position < L.length + 1) 
        return false;
    if (L.length >= MaxSize) return false;
    for (int i = L.length; i >= position; --i) {
        L.data[i] = L.data[i - 1];
    }
    L.data[position - 1] = e;
    ++L.length;
    return true;
}

bool ListDelete(SqList &L, int position, ElemType &e) {
    if (position < 1|| position > L.length) 
        return false;
    e = L.data[position - 1];
    for (int i = position; i < L.length; ++i) 
        L.data[i - 1] = L.data[i];
    --L.length;
    return true;
}

int LocateElem(SqList L, ElemType e) {
    for (int i = 0; i < L.length; ++i) 
        if (L.data[i] == e) 
            return i + 1;
    return 0;
}
