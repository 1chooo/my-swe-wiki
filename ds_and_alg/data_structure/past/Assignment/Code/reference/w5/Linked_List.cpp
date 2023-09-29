// come from
// https://solarianprogrammer.com/2019/02/22/cpp-17-implementing-singly-linked-list-smart-pointers/
#include <iostream>
#include <memory>
#include <sstream>
#include <string>

namespace ds {
  struct Node {
    int data;
    std::unique_ptr<Node> next;
    Node()
        : data{}, next{ nullptr } {}
    Node(int data)
        : data{ data }, next{ nullptr } {}
  };

  class Linked_List {
  public:
    Linked_List()
        : head{ nullptr } {}
    Linked_List(const Linked_List &other)
    {    // copy constructor
      Node *current = other.head.get();    // 現在複製到的 node
      Node *new_phead{ nullptr };
      std::unique_ptr<Node> new_head{ nullptr };
      while (current != nullptr) {
        auto temp{ std::make_unique<Node>(current->data) };    // 複製 node

        if (new_head != nullptr) {    // 不是第一次
          new_phead->next = std::move(temp);    // 下一個元素 = temp
          new_phead = new_phead->next.get();    // 指向下一個
        }
        else {    // 第一次
          new_head = std::move(temp);    // 新 List 的頭
          new_phead = new_head.get();    // 指向頭
        }

        current = current->next.get();
      }
      head = std::move(new_head);    // 新 List 的頭
    }
    Linked_List(Linked_List &&other) { head = std::move(other.head); }
    ~Linked_List() { clean(); }
    void push_front(int data)
    {
      auto temp{ std::make_unique<Node>(data) };
      if (head != nullptr) {
        temp->next = std::move(head);    // 把現在的頭接到輸入的 Node 的後面
        head = std::move(temp);    // 把輸入的 Node 改成頭
      }
      else {
        head = std::move(temp);    // 直接把輸入的 Node 當頭
      }
    }
    void push_back(int data)
    {
      auto temp{ std::make_unique<Node>(data) };
      if (head != nullptr) {
        Node *current = head.get();
        while (true) {
          if (current->next == nullptr)    // 找到最後一個
            break;

          current = current->next.get();    // 換下一個
        }

        current->next = std::move(temp);    // 找到最後一個後 push 進去
        current = nullptr;
      }
      else {
        head = std::move(temp);    // 如果原本就空的那就直接 push
      }
    }
    void pop_front()
    {
      if (head == nullptr)
        return;

      std::unique_ptr<Node> temp = std::move(head);    // 把 temp 指到 head
      head = std::move(
        temp->next);    // 把 head 指到剛剛的下一項，也就是說 head 往後了一格
    }
    void pop_back()
    {
      if (head == nullptr)
        return;

      Node *current = head.get();
      while (true) {
        if (current->next->next.get() ==
            nullptr)    // 下一個如果是最後一個了，那就要把 next 拿掉
          break;

        current = current->next.get();
      }

      std::unique_ptr<Node> temp{ std::move(current->next) };    // 直接拿掉
    }
    void insert(int index, int data)
    {
      if (index == 0) {    // insert 0 代表直接插入頭部
        push_front(data);
        return;
      }

      int cnt{};
      auto temp{ std::make_unique<Node>(data) };
      Node *current = head.get();
      while (true) {
        if (current != nullptr) {    // 如果 List 非空
          if (cnt + 1 ==
              index) {    // 如果下一個是要找的元素，就插在現在這個的後方
            if (current->next == nullptr)
              push_back(data);
            else {
              temp->next = std::move(current->next);
              current->next = std::move(temp);
            }

            current = nullptr;
            return;
          }
        }
        else
          return;

        if (current->next ==
            nullptr) {    // 如果下一個是空的，那代表 index 不存在
          return;
        }

        current = current->next.get();
        ++cnt;
      }
    }
    void erase(int index)
    {
      if (index == 0) {
        pop_front();    // erase 0 代表把頭那個元素拿掉
        return;
      }

      int cnt{};
      Node *current = head.get();
      while (true) {
        if (current != nullptr) {    // 如果 List 非空
          if (current->next ==
              nullptr) {    // 如果下一個是空的，那代表 index 不存在
            return;
          }
          if (cnt + 1 == index) {
            if (current->next->next == nullptr)
              pop_back();
            else
              current->next = std::move(
                current->next
                  ->next);    // 如果下一個是要找的，那就直接把後面那個元素接過來

            current = nullptr;
            return;
          }
        }
        else
          return;

        current = current->next.get();
        ++cnt;
      }
    }
    void clean()
    {
      while (head != nullptr) {
        head = std::move(head->next);    // 逐一刪掉 Node，直接刪的話會因為
          // destructor 造成 stackoverflow
      }
    }
    void reverse()
    {
      Linked_List temp_list;
      Node *current = head.get();
      while (current != nullptr) {
        temp_list.push_front(current->data);
        current = current->next.get();
      }
      clean();
      head = std::move(temp_list.head);
    }
    friend std::ostream &operator<<(std::ostream &os, const Linked_List &list)
    {
      Node *head = list.head.get();
      while (head != nullptr) {
        os << head->data << "-->";    // 輸出目前 head 的 data
        head = head->next.get();    // head 往後一格
      }
      os << "null";
      return os;
    }

  private:
    std::unique_ptr<Node> head;
  };
}    // namespace ds

int main()
{
  ds::Linked_List list;
  std::string input, compact{};
  std::stringstream ss;
  int index{}, data{};
  getline(std::cin, input);
  ss << input;

  while (true) {
    ss >> compact;
    if (compact == "addBack") {
      ss >> data;
      list.push_back(data);
    }
    else if (compact == "addFront") {
      ss >> data;
      list.push_front(data);
    }
    else if (compact == "addIndex") {
      ss >> index >> data;
      list.insert(index, data);
    }
    else if (compact == "deleteIndex") {
      ss >> index;
      list.erase(index);
    }
    else if (compact == "exit") {
      std::cout << list;
      return 0;
    }
  }
}