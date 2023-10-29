#include <stdio.h>
#include <math.h>
#include <Windows.h>

float f(float x, float y, float z) {
    float a = x * x + 9.0f / 4.0f * y * y + z * z - 1;
    return a * a * a - x * x * z * z * z - 9.0f / 80.0f * y * y * z * z * z;
}

float h(float x, float z) {
    for (float y = 1.0f; y >= 0.0f; y -= 0.001f) {
        if (f(x, y, z) <= 0.0f)
            return y;
    }
    return 0.0f;
}

void renderScene() {
    HANDLE console = GetStdHandle(STD_OUTPUT_HANDLE);
    const int width = 80;
    const int height = 25;
    TCHAR buffer[height][width + 1];

    const TCHAR ramp[] = _T(".:-=+*#%@");

    for (int sy = 0; sy < height; sy++) {
        float z = 1.3f - sy * 0.1f;
        TCHAR* p = buffer[sy];

        for (int sx = 0; sx < width; sx++) {
            float x = -1.5f + sx * 0.05f;
            float y = h(x, z);

            if (y > 0.0f) {
                int index = (int)(y * 5.0f);
                *p++ = ramp[index];
            }
            else {
                *p++ = ' ';
            }
        }
        *p = '\0';
    }

    COORD cursorPos = { 0, 0 };
    SetConsoleCursorPosition(console, cursorPos);

    for (int sy = 0; sy < height; sy++) {
        WriteConsole(console, buffer[sy], width, NULL, NULL);
        printf("\n");
    }
}

int main() {
    for (float t = 0.0f;; t += 0.1f) {
        renderScene();
        Sleep(33);
    }

    return 0;
}
