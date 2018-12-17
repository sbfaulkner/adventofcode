#include <iostream>

class Node {
  public:
    int num_children;
    int num_metadata;

    Node *children;
    int* metadata;

    Node();

    int sum();
    int value();

  private:
    int metadata_sum();
};

std::istream& operator>> (std::istream& is, Node& n);
std::ostream& operator<< (std::ostream&os, Node& n);