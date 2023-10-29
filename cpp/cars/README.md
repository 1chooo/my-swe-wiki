# Object-oriented C++

https://github.com/coders-school/Cars

UWAGA: Praca domowa jest na gałęzi [diamond](https://github.com/coders-school/Cars/tree/diamond). Kod tam to mój kod po lekcjach wideo i jest on inny niż ten na gałęzi main.

## Zadanie

1. <!-- .element: class="fragment fade-in" --> Zaproponuj właściwą abstrakcję (interfejs)
2. <!-- .element: class="fragment fade-in" --> Zastosuj dziedziczenie
3. <!-- .element: class="fragment fade-in" --> Napraw enkapsulację
4. <!-- .element: class="fragment fade-in" --> Użyj polimorfizmu, aby za pomocą wskaźnika na klasę bazową reprezentować dowolny typ samochodu

## Pytania

1. <!-- .element: class="fragment fade-in" --> Jak trzymać silniki? Przez wartość, referencję, czy wskaźnik?
2. <!-- .element: class="fragment fade-in" --> Czy jest problem diamentowy?
3. <!-- .element: class="fragment fade-in" --> Czy są wycieki pamięci?
4. <!-- .element: class="fragment fade-in" --> Czy kod jest testowalny?



---

Task
Propose the proper abstraction (interface)
Use inheritance
Fix encapsulation
Use polymorphism to represent any type of car using a pointer to the base class
Questions
How to store engines? By value, reference, or pointer?
Is there a diamond problem?
Are there memory leaks?
Is the code testable?
Homework
(5 XP) Create the InvalidGear exception. It should be thrown when someone tries to change gear improperly, such as from 5 to R. It should inherit from std::logic_error.
(10 XP) Write unit tests for this code. Test throwing the above exception in particular. Configure CMake appropriately.
(0 XP) Fix the interface to be easy to use correctly and hard to use incorrectly (e.g. accelerate(-999)). No points, because it's hard to automate such a general task and everyone will come up with something different here.
Read the article SOLID - Good Practices in Object-Oriented Programming.