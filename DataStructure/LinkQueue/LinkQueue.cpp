#include<iostream>

using namespace std;

typedef int ElemType;

typedef struct LinkQueue
{
    LinkNode *front, *rear;
};

typedef struct LinkNode
{
    ElemType data;
    struct LinkNode *next;
};

void InitQueue(LinkQueue &Q) {
    Q.front = Q.rear = (LinkNode*)malloc(sizeof(LinkNode));
    Q.front->next = NULL;
}

bool IsEmpty(LinkQueue Q) {
    if (Q.front == Q.rear)
        return true;
    else
        return false;
}

void EnQueue(LinkQueue &Q, ElemType x) {
    LinkNode *newnode = (LinkNode*)malloc(sizeof(LinkNode));
    newnode->data = x; newnode->next = NULL;
    Q.rear->next = newnode;
    Q.rear = newnode;
}

bool DeQueue(LinkQueue &Q, ElemType &x) {
    if (Q.front == Q.rear)
        return false;
    LinkNode *p = Q.front->next;
    x = p->data;
    Q.front->next = p->next;
    if (Q.rear == p) {
        Q.rear = Q.front;
    }
    free(p);
    return true;
}

