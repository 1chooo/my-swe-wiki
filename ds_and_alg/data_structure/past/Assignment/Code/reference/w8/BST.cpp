#include <iostream>
#include <memory>

template <typename T>
class BST {
  struct Node {
    Node *left;
    Node *right;
    Node *parent;
    T key;

    Node(T k, Node *p = nullptr)
        : key(k), parent(p), left(nullptr), right(nullptr) {}

    Node()
        : key(), parent(), left(nullptr), right(nullptr) {}
  };

  Node *root;

  auto _min(Node *_node) const -> decltype(root);
  auto _max(Node *_node) const -> decltype(root);
  auto _next(Node *_node) const -> decltype(root);
  auto _prev(Node *_node) const -> decltype(root);

  size_t _size(Node *_node) const;
  bool _find_sum(Node *_node, T target, T &cnt, bool &flag) const;

  void _inorder(Node *_node) const;
  void _preorder(Node *_node) const;
  void _postorder(Node *_node) const;

public:
  BST()
      : root(nullptr) {}
  ~BST() = default;
  size_t size() const { return _size(root); }
  bool empty() const { return root == nullptr; }
  void insert(T k);
  bool find_sum(T target) const;
  auto find(T k) const -> decltype(root);
  void remove(T k);
  auto remove(Node *_root, T k) -> decltype(root);

  void inorder() const;
  void preorder() const;
  void postorder() const;
};

template <typename T>
void BST<T>::preorder() const
{
  if (empty())
    throw std::runtime_error("Missing element!\n");

  _preorder(root);
}

template <typename T>
void BST<T>::postorder() const
{
  if (empty()) {
    throw std::runtime_error("Missing element!\n");
  }
  _postorder(root);
}

template <typename T>
void BST<T>::inorder() const
{
  if (empty()) {
    throw std::runtime_error("Missing element!\n");
  }
  _inorder(root);
}

template <typename T>
void BST<T>::_preorder(Node *_node) const
{
  std::cout << _node->key << " ";
  if (_node->left)
    _preorder(_node->left);

  if (_node->right)
    _preorder(_node->right);
}

template <typename T>
void BST<T>::_postorder(Node *_node) const
{
  if (_node->left)
    _postorder(_node->left);

  if (_node->right)
    _postorder(_node->right);

  std::cout << _node->key << " ";
}

template <typename T>
void BST<T>::_inorder(Node *_node) const
{
  if (_node->left)
    _inorder(_node->left);

  std::cout << _node->key << " ";
  if (_node->right)
    _inorder(_node->right);
}

template <typename T>
auto BST<T>::_min(Node *_node) const -> decltype(root)
{
  Node *tmp = _node;
  while (tmp && tmp->left)
    tmp = tmp->left;

  return tmp;
}

template <typename T>
auto BST<T>::_max(Node *_node) const -> decltype(root)
{
  Node *tmp = _node;
  while (tmp && tmp->right)
    tmp = tmp->right;

  return tmp;
}

template <typename T>
auto BST<T>::_next(Node *_node) const -> decltype(root)
{
  if (_node->right)
    return _min(_node->right);

  Node *tmp = _node->parent;
  while (tmp && tmp->parent) {
    if (tmp->parent->right == tmp && tmp->parent->key > tmp->key)
      return tmp->parent;

    tmp = tmp->parent;
  }

  return nullptr;
}

template <typename T>
auto BST<T>::_prev(Node *_node) const -> decltype(root)
{
  if (_node->left)
    return _max(_node->left);

  Node *tmp = _node->parent;
  while (tmp && tmp->parent) {
    if (tmp->parent->left == tmp && tmp->parent->key < tmp->key)
      return tmp->parent;

    tmp = tmp->parent;
  }

  return nullptr;
}

template <typename T>
size_t BST<T>::_size(Node *_node) const
{
  if (_node)
    return _size(_node->left) + _size(_node->right);
  else
    return 0;
}

template <typename T>
void BST<T>::insert(T k)
{
  Node *tmp = root;
  Node *new_parent = nullptr;

  while (tmp) {
    new_parent = tmp;
    if (k < tmp->key)
      tmp = tmp->left;
    else if (k > tmp->key)
      tmp = tmp->right;
    else
      return;
  }

  Node *new_node = new Node(k, new_parent);
  if (new_parent == nullptr)
    root = new_node;
  else if (k < new_parent->key)
    new_parent->left = new_node;
  else
    new_parent->right = new_node;
}

template <typename T>
auto BST<T>::find(T k) const -> decltype(root)
{
  Node *tmp = root;
  while (tmp) {
    if (k < tmp->key)
      tmp = tmp->left;
    else if (k > tmp->key)
      tmp = tmp->right;
    else
      return tmp;
  }

  return nullptr;
}

template <typename T>
auto BST<T>::remove(Node *_root, T k) -> decltype(root)
{
  if (_root) {
    if (k < _root->key)
      _root->left = remove(_root->left, k);
    else if (k > _root->key)
      _root->right = remove(_root->right, k);
    else {
      if (!_root->left && !_root->right)
        return nullptr;
      else if (!_root->left || !_root->right)
        return _root->left ? _root->left : _root->right;

      Node *tmp = _prev(_root);
      _root->key = tmp->key;
      _root->left = remove(_root->left, tmp->key);
    }
  }
  return _root;
}

template <typename T>
void BST<T>::remove(T k)
{
  remove(this->root, k);
}

template <typename T>
bool BST<T>::find_sum(T target) const
{
  if (empty())
    return false;

  int cnt{};
  bool flag = false;
  return _find_sum(root, target, cnt, flag);
}

template <typename T>
bool BST<T>::_find_sum(Node *_node, T target, T &cnt, bool &flag) const
{
  cnt += _node->key;
  if (cnt == target) {
    flag = true;
    return flag;
  }

  if (cnt > target) {
    cnt -= _node->key;
  }
  else {
    if (_node->left)
      _find_sum(_node->left, target, cnt, flag);

    if (_node->right)
      _find_sum(_node->right, target, cnt, flag);

    cnt -= _node->key;
  }

  return flag;
}

int main()
{
  BST<int> tree;
  int cnt{}, num, target;
  std::cin >> cnt;

  for (int i{}; i < cnt; ++i) {
    std::cin >> num;
    tree.insert(num);
  }

  std::cin >> target;
  if (tree.find_sum(target))
    std::cout << "There exit at least one path in binary search tree.\n";
  else
    std::cout << "There have no path in binary search tree.\n";


  return 0;
}