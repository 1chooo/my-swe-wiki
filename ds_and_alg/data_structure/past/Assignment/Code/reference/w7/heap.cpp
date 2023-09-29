#include <iostream>
#include <string>
#include <vector>

struct Node {
  std::string element;
  int key;
  Node()
      : element(), key() {}
  Node(std::string e, int k)
      : element(e), key(k) {}
  Node(const Node &other)
      : element(other.element), key(other.key) {}
  Node(Node &&other)
      : element(std::move(other.element)), key(other.key) {}
  Node &operator=(const Node &other)
  {
    element = other.element;
    key = other.key;
    return (*this);
  }
  Node &operator=(Node &&other)
  {
    element = std::move(other.element);
    key = other.key;
    return (*this);
  }
};

class Heap {
private:
  std::vector<Node> _heap;
  int find();
  int parent();

public:
  Heap() {}
  Heap(int n) { _heap.reserve(n); }

  void max_insert(Node other)
  {
    _heap.push_back(other);
    int i = _heap.size() - 1;
    while ((i != 0) && (other.key > _heap[(i - 1) / 2].key)) {
      std::swap(_heap[i], _heap[(i - 1) / 2]);
      i = (i - 1) / 2;    // (i-1)/2 is parent
    }
    //for ( const auto &s : _heap ) std::cout << s.key << " ";
    //std::cout << '\n';
  }

  void pop_front()
  {
    int n = _heap.size() - 1;
    Node temp = _heap[n--];
    int parent{}, child{ 1 };
    while (child <= n) {
      if (child < n && _heap[child].key < _heap[child + 1].key)
        ++child;

      _heap[parent] = std::move(_heap[child]);
      parent = child, child = 2 * parent + 1;
    }

    _heap[parent] = temp;
  }

  inline void output()
  {
    std::cout << _heap[0].element << '\n';
  }
};

int main()
{
  int n, key;
  std::cin >> n;

  Heap heap(n);
  std::string str;
  for (; n != 0; --n) {
    std::cin >> str >> key;
    heap.max_insert(Node(str, key));
  }

  std::cout << "First three things to do:" << '\n';

  for (int i{}; i < 3; ++i) {
    heap.output();
    heap.pop_front();
  }

  return 0;
}