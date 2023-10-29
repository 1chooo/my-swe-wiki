#pragma once
#include <iostream>
#include <string>
#include <vector>

enum class Gender {
    Male,
    Female,
    Other
};

class Student {
    std::string name_;
    std::string surname_;
    std::string address_;
    int ID_;
    Gender gender_;
    int index_;

public:
    std::string GetName() { return name_; }
    std::string GetSurname() { return surname_; }
    std::string GetAddress() { return address_; }
    int GetID() { return ID_; }
    Gender GetGender() { return gender_; }
    std::string GetGenderInString();//
    int GetIndex() { return index_; }

    void SetName(const std::string& name) { name_ = name; }
    void SetSurname(const std::string& surname) { surname_ = surname; }
    void SetAddress(const std::string& address) { address_ = address; }
    void SetID(const int& ID) { ID_ = ID; }
    void SetGender(const Gender& gender) { gender_ = gender; }
    void SetIndex(const int& index) { index_ = index; }

    //ustawiam konstruktor klasy Person
    Student(std::string name, std::string surname, std::string adr, int id, int index, Gender gender)
        : name_(name), surname_(surname), address_(adr), ID_(id), index_(index), gender_(gender) {}

    void Print();
};

