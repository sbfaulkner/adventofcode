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


    ifs >> node;

    std::cout << "Sum: " << node.sum() << std::endl;
    std::cout << "Value: " << node.value() << std::endl;

    ifs.close();
  }

  catch (std::ios_base::failure& e) {
    std::cerr << strerror(errno) << ": " << INPUT << std::endl;
  }

  return 0;
}
