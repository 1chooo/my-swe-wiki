#include <exception>
#include <iostream>
#include <memory>
#include <queue>

template <typename T>
class CBT {
  struct Node {
    std::weak_ptr<Node> _parent;
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
    T key;

    std::shared_ptr<Node> parent()
    {
      auto tmp = _parent.lock();
      if (tmp != nullptr)
        return tmp;
    }

    Node(T _k, std::shared_ptr<Node> _p = nullptr)
        : key(_k), _parent(_p) {}
    Node()
        : key(), _parent(nullptr), left(nullptr), right(nullptr) {}
  };

  std::shared_ptr<Node> root = nullptr;

  std::ostream &_inorder(std::shared_ptr<Node> &node, std::ostream &o);
  std::ostream &_preorder(std::shared_ptr<Node> &node, std::ostream &o);
  std::ostream &_postorder(std::shared_ptr<Node> &node, std::ostream &o);

  size_t _size(std::shared_ptr<Node> &_node) const;

public:
  CBT() = default;
  ~CBT() { destroy(); }

  // non - copyable, non - movable
  CBT(const CBT &) = delete;
  CBT(CBT &&) = delete;
  CBT &operator=(const CBT &) = delete;
  CBT &operator=(CBT &&) = delete;

  size_t size() const { return _size(root); }
  bool isempty() { return (_size(root) == 0) ? true : false; }

  void complete_insert(T k);
  void destroy();

  std::ostream &inorder();
  std::ostream &preorder();
  std::ostream &postorder();
};

template <typename T>
void CBT<T>::destroy()
{
  root = nullptr;
}

template <typename T>
size_t CBT<T>::_size(std::shared_ptr<Node> &_node) const
{
  if (_node == nullptr)
    return 0;    // terminal, return 0
  else
    return _size(_node->left) + _size(_node->right) + 1;    // count + 1
}

template <typename T>
void CBT<T>::complete_insert(T k)
{
  // 如果 Tree 為空，加上 root
  if (root == nullptr) {
    root = std::make_shared<Node>(k);
    return;
  }

  // BFS
  std::queue<std::shared_ptr<Node>> _queue;
  _queue.push(root);
  std::shared_ptr<Node> tmp = nullptr;
  while (!_queue.empty()) {
    tmp = _queue.front();
    _queue.pop();

    if (tmp->left == nullptr) {
      tmp->left = std::make_shared<Node>(k, tmp);
      return;
    }
    else if (tmp->right == nullptr) {
      tmp->right = std::make_shared<Node>(k, tmp);
      return;
    }
    else {
      _queue.push(tmp->left);
      _queue.push(tmp->right);
    }
  }
};

template <typename T>
std::ostream &CBT<T>::preorder()
{
  if (isempty()) {
    throw std::runtime_error("Missing element!\n");
  }
  return _preorder(root, std::cout);
}

template <typename T>
std::ostream &CBT<T>::postorder()
{
  if (isempty()) {
    throw std::runtime_error("Missing element!\n");
  }
  return _postorder(root, std::cout);
}

template <typename T>
std::ostream &CBT<T>::inorder()
{
  if (isempty()) {
    throw std::runtime_error("Missing element!\n");
  }
  return _inorder(root, std::cout);
}

template <typename T>
std::ostream &CBT<T>::_preorder(std::shared_ptr<Node> &i, std::ostream &o)
{
  o << i->key << " ";
  if (i->left != nullptr) {
    _preorder(i->left, o);
  }
  if (i->right != nullptr) {
    _preorder(i->right, o);
  }
  return o;
}

template <typename T>
std::ostream &CBT<T>::_postorder(std::shared_ptr<Node> &i, std::ostream &o)
{
  if (i->left != nullptr) {
    _postorder(i->left, o);
  }
  if (i->right != nullptr) {
    _postorder(i->right, o);
  }
  o << i->key << " ";
  return o;
}

template <typename T>
std::ostream &CBT<T>::_inorder(std::shared_ptr<Node> &i, std::ostream &o)
{
  if (i->left != nullptr) {
    _inorder(i->left, o);
  }
  o << i->key << " ";
  if (i->right != nullptr) {
    _inorder(i->right, o);
  }
  return o;
}

int main()
{
  CBT<int> Tree;
  int i;

  while (std::cin >> i) {
    Tree.complete_insert(i);
  }

  Tree.inorder();
}