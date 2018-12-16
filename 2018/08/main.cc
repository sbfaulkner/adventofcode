#include <iostream>
#include <fstream>
#include "node.h"

#define INPUT "input"

int main() {
  std::ifstream ifs;

  ifs.exceptions(ifs.exceptions() | std::ios::failbit);

  try {
    ifs.open(INPUT);

    Node node;

    int sum = 0;

    while (ifs >> node) {
      sum += node.sum();
    }

    std::cout << "Sum: " << sum << std::endl;

    ifs.close();
  }

  catch (std::ios_base::failure& e) {
    std::cerr << strerror(errno) << ": " << INPUT << std::endl;
  }

  return 0;
}
