#include "student.hpp"

void Student::Print() {
    std::cout << "Imie: " << name_ << " Nazwisko: " << surname_ << " PESEL: " << ID_ << " Płeć: " << GetGenderInString() << " index: " << index_ << " adres: " << address_;
}


std::string Student::GetGenderInString() {
    switch (gender_) {
    case Gender::Male:
        return "Chłopak";
        break;
    case Gender::Female:
        return "Dziewczyna";
        break;
    case Gender::Other:
        return "Other";
        break;
    default:
        return "";
        break;
    }
}