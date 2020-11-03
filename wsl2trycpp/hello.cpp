#include <iostream>
using namespace std;

class Base {
public:
    int value1;
    int value2;
    Base() {
        value1 = 1;
    }
    Base(int value) : Base() {
        value2 = value;
    }
};

int main() {
    Base *b = new Base(2);
    cout << b->value1 << endl;
    cout << b->value2 << endl;
    delete b;
    return 0;
}
