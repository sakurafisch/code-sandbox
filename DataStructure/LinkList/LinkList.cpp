#include<iostream>
using namespace std;

typedef int ElemType;

typedef struct LNode
{
    ElemType data;
    struct LNode *next;
}LNode, *LinkList;


LinkList List_HeadInsert(LinkList &L) {
    ; int x;
    L = (LinkList)malloc(sizeof(LNode));
    L->next = NULL;
    while (cin >>x) {
        if (x == 999) break;
        LNode *newnode = (LNode*)malloc(sizeof(LNode));
        newnode->data = x;
        newnode->next = L->next;
        L->next = newnode;
    }
    return L;
}

LinkList HeadInsert(LinkList &L) {
    L = (LinkList)malloc(sizeof(LNode));
    L->next = NULL;
    int x;
    while (cin >>x) {
        if (x == 999) break;
        LNode *newnode = (LNode*)malloc(sizeof(LNode));
        newnode->data = x;
        newnode->next = L->next;
        L->next = newnode;
    }
    return L;
}

LinkList TailInsert(LinkList &L) {
    L = (LinkList)malloc(sizeof(LNode));
    L->next = NULL;
    int x;
    LNode* rear = L;
    while (cin >> x) {
        if (x == 999) break;
        LNode *newnode = (LNode*)malloc(sizeof(LNode));
        newnode->data = x;
        rear->next = newnode;
        rear = newnode;
    }
    return L;
}

LNode *GetElem(LinkList L, int position) {
    if (position == 0) return L;
    if (position < 1) return NULL;
    int count = 1;
    LNode *cur = L->next;
    while (cur) {
        if (count == position)
            break;
        cur = cur->next;
        ++count;
    }
    return cur;
}

LNode *LocateElem(LinkList L, ElemType e) {
    LNode *cur = L->next;
    while(cur) {
        if (cur->data == e)
            break;
        cur = cur->next;
    }
    return cur;
}
