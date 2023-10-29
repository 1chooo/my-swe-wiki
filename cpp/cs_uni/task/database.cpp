#include "database.hpp"
#include <algorithm>
#include <iostream>

void DataBase::add_student(Student *St)
{
    students.push_back(St);
}

void DataBase::list_student()
{
    for (Student *s : students)
    {
        s->Print();
        std::cout << "\n";
    }
}

void DataBase::GetAllStudentsWithSurname(std::vector<Student *> &foundStudents, const std::string &searchingSurname)
{
    for (Student *s : students)
    {
        if (s->GetSurname().compare(searchingSurname) == 0)
        {
            foundStudents.push_back(s);
        }
    }
}

Student *DataBase::GetStudentViaID(const int &searchingID)
{
    for (Student *s : students)
    {
        if (s->GetID() == searchingID)
        {
            return s;
        }
    }
    return nullptr;
}

void DataBase::ID(int a)
{
    std::vector<int> check;

    for (auto l : students)
    {
        check.push_back(l->GetID());
    }

    std::cout << "\n";

    for (auto k : check)
    {
        if (k == a)
        {
            std::cout << "Istnieje student z podanym ID: " << a << "\n";
            Student *test = GetStudentViaID(a);
            test->Print();
            return;
        }
    }
    std::cout << "Brak ID " << a << " w bazie danych. Obecne ID w bazie danych: " << "\n";

    for (auto k : check)
    {
       std::cout << k << ", ";
    }

}

void DataBase::SortStudentsSurnames(std::vector<Student *> &sortSurnames)
{
    std::cout << "Przed sortowaniem nazwisk: " << "\n";

    for (Student *s : students)
    {
        sortSurnames.push_back(s);
        s->Print();
        std::cout << "\n";
    }
    std::sort(sortSurnames.begin(), sortSurnames.end(), sortSurnamesComparator());

    std::cout << "\n";
    std::cout << "Po sortowaniu nazwisk: " << "\n";

    for (auto i : sortSurnames)
    {
        i->Print();
        std::cout << "\n";
    }
}

void DataBase::SortStudentsID(std::vector<Student *> &sortID)
{
    std::cout << "Przed sortowaniem ID: " << "\n";

    for (Student *s : students)
    {
        sortID.push_back(s);
        s->Print();
        std::cout << "\n";
    }

    std::sort(sortID.begin(), sortID.end(), sortIDComparator());
    std::cout << "\n";
    std::cout << "Po sortowaniu ID: " << "\n";

    for (auto i : sortID)
    {
        i->Print();
        std::cout << "\n";
    }
}

void DataBase::DeleteByIndex(int deleteIndex)
{
    for (int i = 0; i < students.size(); i++)
    {
        if (students[i]->GetIndex() == deleteIndex)
        {
            students.erase(students.begin() + i);
        }
    }
}

