# Material 02 - Inheritance

### Lab 05

請你告訴我，對於物件導向，繼承這個機制帶來的好處和壞處。

**My Ans:**
- 好處是：可以實現特異化，並且還能保有原先特性
- 壞處是：需要很有 comman sense 以免寫出無意義的

**Prof Ans:**
- Benefits: reuse, 便利的部分可以擴充推到 subclass 如果可以結合 polymorphism 就可以更好的擴充
- Drawbacks: 用錯了程式碼會更難維護

### Lab 06

Please try to explain why a pet destructor is called to print the message on the screen?

**My Ans:**

`cat` 的 destructor，當你用一個 subclass override 掉 base class 的 method

**Prof Ans:**

`(pet)pussy (troncated)` 會強制轉，強制把 `pussy (cat)` 變成 `pet`，因此會把 `cat` 的記憶體部分切掉，


### Lab 07

Java has no member initialization list. So how can you complete the example in slide 12?

**My Ans:**

用 `super()` 並且只能在 constructor 的第一行


### Lab 11

The following code contains two classes.

```cpp
class person {
    person() ;
    person(int age, int height);
}

class student: public person {
   student(int age, int height, int studentID)   
}
```

(a) How can the subclass student pass the parameters age and height to a base class person in C++?


**My Ans:**

```cpp
student :: student(int age, int height, int studentID) person(age, height)  {
    int _studentID = studentID;
}
```

person 會變成 argument 給 student


### Lab 08

Why constructor and destructor is not inherited? Please explain it in a short sentence.

**My Ans:**

Each subtype of the base class has its own constructor and destructor. -> copy constructors and destructors are not inherited.


### Lab 09

In principle, each subclass has its constructor.   
What is the main task of a subclass constructor? That is, what a subclass constructor should do?

**My Ans:**

因為每個 subclass 都需要建構屬於自己特異化的部分，也就是會建立自己的 memory space


### Lab 13

What is overloading?


**My Ans:**

```cpp
class XXX {
  void dosSomething(int x, string b);
  void dosSomething(int x, string b, char c);
}
```

同名的 `function` 但是 `signature` 不同，如上案例

### Lab 15

請說明繼承的基本原則是什麼?

**My Ans:**

子類別是不是服類別的群集，`Is American a Human?`，滿足並且可以沿用做特異化


### Lab 18

以下的兩個 subclass 投影片告訴你沒有意義. 請你進一步釐清為何沒有意義? 你可以舉例子嗎?

![](../../materials/02_inheritance/imgs/01.png)

**My Ans:**

如果 American, Chinese 沒有跟 Citizen 做出其他的特徵，那就直接用 Human 就好了

如果真的要用，可以加上 American 公民與一般 Citizen 不同的內容，像是健保等等，做出特異化才有必要

