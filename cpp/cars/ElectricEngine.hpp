#pragma once

class ElectricEngine
{
public:
    ElectricEngine(int power, int batteryCapacity);

    int power_;             // in HP 
    int batteryCapacity_;   // in Ah
};

