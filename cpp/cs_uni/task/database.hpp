#pragma once
#include <vector>
#include <string>
#include "student.hpp"

class DataBase
{
    std::vector<Student *> students;

public:
    void add_student(Student *St);

    void list_student();

    void GetAllStudentsWithSurname(std::vector<Student *> &foundStudents, const std::string &searchingSurname); 

    Student *GetStudentViaID(const int &searchingID); 

    void SortStudentsSurnames(std::vector<Student *> &sortSurnames);

    void SortStudentsID(std::vector<Student *> &sortID);

    void DeleteByIndex(int);

    void ID(int);
};
struct sortSurnamesComparator
{
    inline bool operator()(Student *s1, Student *s2)
    {
        return (s1->GetSurname() < s2->GetSurname());
    }
};
struct sortIDComparator
{
    inline bool operator()(Student *s1, Student *s2)
    {
        return (s1->GetID() < s2->GetID());
    }
};
