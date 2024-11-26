#include <iostream>

int main() {
    int x = 3;

    int &y = x;

    std::cout << &x << std::endl; // Print the address of x
    std::cout << &y << std::endl; // Print the address of y (which is the same as x)

    y = 4; // Assign 4 to y, which modifies x as well

    std::cout << x << std::endl; // Print the value of x

    return 0;
}