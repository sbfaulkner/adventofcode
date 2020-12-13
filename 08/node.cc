#include "node.h"

Node::Node() {
  num_children = 0;
  num_metadata = 0;
  children = NULL;
  metadata = NULL;
}

int Node::sum() {
  int sum = 0;

  for (int c = 0; c < num_children; c++) {
    sum += children[c].sum();
  }

  sum += metadata_sum();

  return sum;
}

int Node::value() {
  int value = 0;

  if (num_children == 0) {
    value += metadata_sum();
  } else {
    for (int m = 0; m < num_metadata; m++) {
      int c = metadata[m] - 1;

      if (c < num_children) {
        value += children[c].value();
      }
    }
  }

  return value;
}

int Node::metadata_sum() {
  int sum = 0;

  for (int m = 0; m < num_metadata; m++) {
    sum += metadata[m];
  }

  return sum;
}

std::istream& operator>> (std::istream& is, Node& n) {
  is >> n.num_children >> n.num_metadata;

  if (n.children) {
    delete[] n.children;
  }

  if (n.num_children > 0) {
    n.children = new Node[n.num_children];

    for (int c = 0; c < n.num_children; c++) {
      is >> n.children[c];
    }
  } else {
    n.children = NULL;
  }

  if (n.metadata) {
    delete[] n.metadata;
  }

  if (n.num_metadata > 0) {
    n.metadata = new int[n.num_metadata];

    for (int m = 0; m < n.num_metadata; m++) {
      is >> n.metadata[m];
    }
  } else {
    n.metadata = NULL;
  }

  return is;
}

std::ostream& operator<< (std::ostream&os, Node& n) {
  os << "(" << n.num_children << " " << n.num_metadata << ") { [";

  for (int c = 0; c < n.num_children; c++) {
    if (c > 0) os << " ";
    os << n.children[c];
  }

  os << "] [";

  for (int m = 0; m < n.num_metadata; m++) {
    if (m > 0) os << " ";
    os << n.metadata[m];
  }

  return os << "] }";
}
