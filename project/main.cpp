// 期末作業3.cpp : 此檔案包含 'main' 函式。程式會於該處開始執行及結束執行。
//

#include <fcntl.h>
#include <sys/uio.h>
#include <iostream>
#include <random>
#include <time.h>
#include <conio.h>
#include <windows.h>
#include <iomanip>
using namespace std;
#define cardsnumber 52 //表總牌數
#define suitnumber 13  //表單種花色牌數
#define players 4      //玩家應該數量

bool playerstatus[4] = { 0, 0, 0, 0 };  //0表電腦玩家,1表人類玩家
bool playercanplay[4] = { 1, 1, 1, 1 }; //0表此輪不能玩(因為已經過牌),1表可以
int cardstable[cardsnumber] = { 0 };    //每張牌的分配狀態(0-51)(0-4,0為廢牌區)(梅花、方塊、紅心、黑桃)
int handstatus[players + 1] = { 0 };    //每位玩家的持牌數量(0-4,0為廢牌區)
int playernow = 0, winner = 0;        //目前玩家編號(輪轉順序：從小到大)
int player1[suitnumber];              //人類手牌情況
int counter = 0;                   //過牌次數
int ctr1 = 0;           //紀錄按左右鍵的次數
int idx, idx2;          //分別記錄電腦與人類選牌數
int check2;             //確認是否按下Enter
int memo;               //記下電腦最近出的牌

int memory[6] = { 0 };                //記下電腦最近的出牌內容
int inputcards[6] = { 0 };              //出牌內容
int suitnum = 0, suitvalue = -1;      //出牌牌型，出牌大小(價值)(-1表第一位)
int presuitnum = 0, presuitvalue = 0; //前一次出牌牌型，出牌大小(價值)(-1表第一位)
int selected[5];                      //選擇出牌的位置
void gotoxy(int xpos, int ypos)       //移動座標
{
    COORD scrn;
    HANDLE hOuput = GetStdHandle(STD_OUTPUT_HANDLE);
    scrn.X = xpos; scrn.Y = ypos;
    SetConsoleCursorPosition(hOuput, scrn);
}
void SetColor(int color = 7)          
{
    HANDLE hConsole;
    hConsole = GetStdHandle(STD_OUTPUT_HANDLE);
    SetConsoleTextAttribute(hConsole, color);
}
void sorting(int x[], int n);
void opening();
void printmemorystatus(int n);
void printhandstatus(int n);                         //輸出所有牌的狀態
bool checkhavingcards(int numberofcard, int player); //輸出是否擁有這些牌
int checklegalcards(int numberofcard);               //輸出牌型
bool checklargercards(int numberofcard);             //檢查是否較大
int algfindingcards(int player);                     //電腦玩家決定出牌
void display(int selector, int num1, int num2)
{
    gotoxy(10, 15);
    printhandstatus(0);
    gotoxy(10, 25);
    cout << num1;
    gotoxy(10, 27);
    cout << playernow;
    gotoxy(selector, 17);
    wcout << "^";
    for (int i = 0; i < num1; i++)
    {
        if (selected[i] != -1)
        {
            gotoxy(selected[i], 13);
            wcout << "*";
        }
    }
    if (num1 == 0)
    {
        gotoxy(10, 4);
        wcout << "Pass";
    }
    else
    {
        gotoxy(10, 4);
        printhandstatus(num1);
    }
    if (num2 == 0)
    {
        gotoxy(10, 7);
        wcout << "Pass";
    }
    else
    {
        gotoxy(10, 7);
        printmemorystatus(num2);
    }
}

int main()
{
    int selector_pos = 12;//游標的位子
    std::default_random_engine generator(time(NULL));            //建立亂樹種子
    std::uniform_int_distribution<int> distribution(1, players); //亂數分布

    for (int i = 0; i < cardsnumber; i++)
    {
        int cardplayer = distribution(generator); //將每張牌分給每位玩家
        if (handstatus[cardplayer] < cardsnumber / players)
        {
            if (cardplayer == 1)
                player1[handstatus[cardplayer]] = i;
            cardstable[i] = cardplayer;
            handstatus[cardplayer] += 1; //紀錄玩家持牌數量
            if (i == 2)
            {
                playernow = cardplayer; //紀錄梅花3分配給誰
            }
        }
        else
        {
            i -= 1; //若持牌數量太多，重新分配此張牌
        }
    }
    sorting(player1, 13);

    for (int i = 0; i < 1; i++)
    {
        playerstatus[i] = 1;
    }

    opening();
    if (char a = _getch() == 's')
        system("cls");
    system("color F0");
    gotoxy(10, 15);
    printhandstatus(0);

    if (playernow == 1)
    {
        gotoxy(10,13);
        wcout << "You First.";
    }
    else
    {
        gotoxy(10,13);
        wcout << "Computer First.";
    }

    while (true) //若有贏家就結束遊戲
    {
        if (winner == 1)
        {
            gotoxy(10, 15); wcout << "You Win";
            break;
        }
        else if (winner == 2 || winner == 3 || winner == 4)
        {
            gotoxy(10, 15); wcout << "You Lose";
            break;
        }
        if (handstatus[playernow] == 0) //確保playernow在範圍內，且playernow的手牌不為0
        {
            playernow += 1;
            playernow = (playernow - 1) % players + 1;
        }

        if (playerstatus[playernow - 1] == 1) //若playernow是人類玩家
        {
            for (int i = 0;i < 13;i++)
            {
                if (cardstable[player1[i]] == 0)
                {
                    player1[i] = -2;
                }
            }
            sorting(player1, 13);
            int check = 0; //紀錄牌型是否正確與合法
            for (int i = 0; i < 6; i++)
            {
                inputcards[i] = 0; //清空出牌佔態
            }
            for (int i = 0; i < 5; i++)
            {
                selected[i] = -1;
            }
            int selector_pos = 12;//游標的位子
            idx = 0; //idx表輸入牌數
            ctr1 = 13 - handstatus[1];
            while (char c = _getch())
            {
                check2 = 0;
                switch (c)
                {
                case 'a':
                {
                    if (selector_pos > 12)
                    {
                        selector_pos -= 6;
                        ctr1--;
                    }
                    else
                    {
                        selector_pos = 90 - 6 * (13 - handstatus[1] + 1);
                        ctr1 = 12;
                    }
                    break;
                }

                case 'd':
                {
                    //if(pos)
                    if (selector_pos < (90 - 6 * (13 - handstatus[1] + 1)))
                    {
                        selector_pos += 6;
                        ctr1++;
                    }
                    else
                    {
                        selector_pos = 12;
                        ctr1 = 13 - handstatus[1];
                    }
                    break;
                }

                case ' ':
                {
                    if (idx < 5)
                    {
                        selected[idx] = selector_pos;
                        inputcards[idx] = player1[ctr1];
                        idx++;
                    }
                    break;
                }
                case '\b':
                {
                    if (idx > 0)
                    {
                        selected[idx] = -1;
                        inputcards[idx] = 0;
                        idx--;
                    }
                    break;
                }
                case '\r':
                {
                    check2 = 1;
                    if (idx == 0) //若是不出牌(過牌)
                    {
                        if (presuitnum == 0 && presuitvalue == 0)
                        {
                            break;//do nothing
                        }
                        else
                        {
                            playercanplay[playernow - 1] = 0;
                            check = 1;
                            playernow += 1;
                            playernow = (playernow - 1) % players + 1;
                            counter += 1; //紀錄過牌人數
                            if (counter == players - 1)
                            {
                                suitnum = 0; //清空判斷之佔態
                                suitvalue = 0;
                                presuitnum = 0;
                                presuitvalue = 0;
                                counter = 0;
                                for (int k = 0; k < players; k++)
                                {
                                    playercanplay[k] = 1;
                                }
                                gotoxy(12, 20);
                                wcout << "This is new round.\n\n";
                                system("pause");
                            }
                            break;
                        }
                    }
                    else
                    {
                        counter = 0;
                        check = checkhavingcards(idx, playernow); //檢查是否有牌
                        gotoxy(12, 25);
                        cout << playernow;
                        if (check == 0)
                        {
                            gotoxy(12, 20);
                            wcout << "You don't have some of the cards.\n";
                            gotoxy(12, 22);
                            wcout << "Please input again to follow the rule.\n";
                            idx = 0;
                            for (int i = 0; i < 5; i++)
                            {
                                selected[i] = -1;
                            }
                            system("pause");
                            break;
                        }
                        else
                        {
                            for (int i = 0; i < idx - 1; i++)
                            {
                                for (int j = i + 1; j < idx; j++)
                                {
                                    if (inputcards[i] > inputcards[j]) //排序
                                    {
                                        swap(inputcards[i], inputcards[j]);
                                    }
                                    else if (inputcards[i] == inputcards[j]) //確定無重複
                                    {
                                        check = 0;
                                    }
                                }
                            }
                            if (check == 0)
                            {
                                gotoxy(12, 20);
                                wcout << "Input repeat.\n";
                                gotoxy(12, 22);
                                wcout << "Please input again to follow the rule.\n";
                                idx = 0;
                                for (int i = 0; i < 5; i++)
                                {
                                    selected[i] = -1;
                                }
                                system("pause");
                                break;
                            }
                            else
                            {
                                check = check * checklegalcards(idx); //檢查是否合法
                                if (check == 0)
                                {
                                    gotoxy(12, 20);
                                    wcout << "Input not legal.\n";
                                    gotoxy(12, 22);
                                    wcout << "Please input again to follow the rule.\n";
                                    idx = 0;
                                    for (int i = 0; i < 5; i++)
                                    {
                                        selected[i] = -1;
                                    }
                                    system("pause");
                                    break;
                                }
                                else
                                {
                                    check = check * checklargercards(idx); //檢查是否比較大
                                    if (check == 0)
                                    {
                                        gotoxy(12, 20);
                                        wcout << "Input not larger.\n";
                                        gotoxy(12, 22);
                                        wcout << "Please input again to follow the rule.\n";
                                        idx = 0;
                                        for (int i = 0; i < 5; i++)
                                        {
                                            selected[i] = -1;
                                        }
                                        system("pause");
                                        break;
                                    }
                                    else
                                    {
                                        handstatus[playernow] -= idx;   //減少手牌數量
                                        if (handstatus[playernow] == 0) //檢查是否有人贏了
                                        {
                                            winner = playernow;
                                            gotoxy(12, 20);
                                            wcout << "The winner is : " << winner << endl;
                                            system("pause");
                                            break;
                                        }
                                        for (int i = 0; i < idx; i++)
                                        {
                                            cardstable[inputcards[i]] = 0; //設定出牌狀態
                                        }
                                        presuitnum = suitnum;
                                        presuitvalue = suitvalue; //設定出牌牌型與大小
                                        playernow += 1;
                                        playernow = (playernow - 1) % players + 1; //設定下一位玩家
                                        break;
                                    }
                                }
                            }
                        }
                    }
                    break;
                }
                }
                system("cls");
                display(selector_pos, idx, memo);
                if (check2 == 1 && check != 0)
                    break;
            }
        }
        else
        {
            {
                int idx2 = 0;
                if (playercanplay[playernow - 1] == 0)
                {
                    playernow += 1;
                    playernow = (playernow - 1) % players + 1;
                    counter += 1; //紀錄過牌人數
                    if (counter == players - 1)
                    {
                        suitnum = 0; //清空判斷之佔態
                        suitvalue = 0;
                        presuitnum = 0;
                        presuitvalue = 0;
                        counter = 0;
                        for (int k = 0; k < players; k++)
                        {
                            playercanplay[k] = 1;
                        }
                        gotoxy(12, 20);
                        wcout << "This is new round.\n\n";
                        system("pause");
                    }
                }
                else
                {
                    for (int i = 0; i < 6; i++)
                    {
                        inputcards[i] = 0; //清空出牌佔態
                    }
                    idx2 = algfindingcards(playernow);
                    if (idx2 == 0)
                    {
                        gotoxy(12, 20);
                        wcout << "Player " << playernow << " check.\n\n";
                        system("pause");
                        playercanplay[playernow - 1] = 0;
                        playernow += 1;
                        playernow = (playernow - 1) % players + 1; //設定下一位玩家
                        counter += 1;                              //紀錄過牌人數
                        if (counter == players - 1)
                        {
                            suitnum = 0; //清空判斷之佔態
                            suitvalue = 0;
                            presuitnum = 0;
                            presuitvalue = 0;
                            counter = 0;
                            for (int k = 0; k < players; k++)
                            {
                                playercanplay[k] = 1;
                            }
                            gotoxy(12, 20);
                            wcout << "This is new round.\n\n";
                            system("pause");
                        }
                    }
                    else
                    {
                        memo = 0;
                        counter = 0;
                        handstatus[playernow] -= idx2;   //減少手牌數量
                        if (handstatus[playernow] == 0) //檢查是否有人贏了
                        {
                            winner = playernow;
                            printf("The winner is : %d\n", winner);
                            break;
                        }
                        for (int i = 0; i < idx2; i++)
                        {
                            cardstable[inputcards[i]] = 0; //設定出牌狀態
                        }
                        for (int i = 0; i < idx2; i++)
                        {
                            memory[i] = inputcards[i]; //設定出牌狀態
                            memo++;
                        }
                        presuitnum = suitnum;
                        presuitvalue = suitvalue; //設定出牌牌型與大小
                        playernow += 1;
                        playernow = (playernow - 1) % players + 1; //設定下一位玩家
                    }
                }
            }
        }
        system("cls");
        display(selector_pos, idx, memo);
        gotoxy(10, 17);
    }
    system("pause");
    return 0;
}

//輸出是否擁有這些牌
bool checkhavingcards(int numberofcard, int player)
{
    bool checking;
    if (suitvalue == -1) //若此輪為第一輪
    {
        checking = 0;
    }
    else
    {
        checking = 1;
    }
    for (int i = 0; i < numberofcard; i++)
    {
        if (cardstable[inputcards[i]] != player)
        {
            return 0;
        }
        if (suitvalue == -1 && checking == 0 && inputcards[i] == 2)
        {
            checking = 1; //檢查必須出梅花3
        }
    }
    if (checking == 0 && suitvalue == -1)
    {
        return 0;
    }
    return 1;
}

//輸出牌型(0表不存在,1表單張,2表對子,3表順子,4表葫蘆,5表鐵支,6表同花順)
int checklegalcards(int numberofcard)
{
    if (numberofcard == 1)
    {
        suitnum = 1;
        suitvalue = inputcards[0];
        return 1; //單張
    }
    else if (numberofcard == 2)
    {
        if (inputcards[0] % suitnumber != inputcards[1] % suitnumber)
        {
            suitnum = 0;
            suitvalue = 0;
            return 0;
        }
        else
        {
            suitnum = 2;
            suitvalue = inputcards[1];
            return 2; //對子
        }
    }
    else if (numberofcard == 5)
    {
        int tmp[5], tmpsuit[5];
        for (int i = 0; i < numberofcard; i++) //前處裡，取出資料
        {
            tmp[i] = inputcards[i] % suitnumber;
            tmpsuit[i] = inputcards[i] / suitnumber;
        }
        for (int i = 0; i < numberofcard - 1; i++) //將取出的資料排序
        {
            for (int j = i + 1; j < numberofcard; j++)
            {
                if (tmp[i] > tmp[j])
                {
                    int tmptmp = tmp[i];
                    tmp[i] = tmp[j];
                    tmp[j] = tmptmp;
                    tmptmp = tmpsuit[i];
                    tmpsuit[i] = tmpsuit[j];
                    tmpsuit[j] = tmptmp;
                }
            }
        }
        int tmpcounter = 0, idx;               //前變數表計數
        for (int i = 0; i < numberofcard; i++) //取出特徵值
        {
            if (tmp[i] == tmp[2])
            {
                tmpcounter += 1;
            }
            else
            {
                idx = i;
            }
        }
        if (tmpcounter == 4)
        {
            suitnum = 5;
            suitvalue = tmpsuit[2] * suitnumber + tmp[2];
            return 5; //鐵支
        }
        else if (tmpcounter == 3)
        {
            if (idx == 1)
            {
                if (tmp[0] == tmp[1])
                {
                    suitnum = 4;
                    suitvalue = tmpsuit[2] * suitnumber + tmp[2];
                    return 4; //葫蘆
                }
            }
            else if (idx == 4)
            {
                if (tmp[3] == tmp[4])
                {
                    suitnum = 4;
                    suitvalue = tmpsuit[2] * suitnumber + tmp[2];
                    return 4; //葫蘆
                }
            }
        }
        bool tmpcheck1 = 1;
        for (int i = 2; i < numberofcard; i++)
        {
            if (tmp[i] - tmp[1] != i - 1) //判斷順子結構
            {
                suitnum = 0;
                suitvalue = 0;
                return 0;
            }
            if (tmpsuit[0] != tmpsuit[i] && tmpcheck1 == 1) //判斷同花順
            {
                tmpcheck1 = 0;
            }
        }
        if (tmpsuit[0] != tmpsuit[1] && tmpcheck1 == 1) //判斷同花順
        {
            tmpcheck1 = 0;
        }
        if (tmp[1] - tmp[0] == 1) //其他
        {
            if (tmp[0] == 1)
            {
                suitvalue = tmpsuit[0] * suitnumber + tmp[0];
            }
            else
            {
                suitvalue = tmpsuit[4] * suitnumber + tmp[4];
            }
            if (tmpcheck1 == 0)
            {
                suitnum = 3;
                return 3; //順子
            }
            else
            {
                suitnum = 6;
                return 6; //同花順
            }
        }
        else if (tmp[0] == 0 && tmp[1] == 9) //10 J Q K A
        {
            if (tmpcheck1 == 0)
            {
                suitnum = 3;
                suitvalue = tmpsuit[0] * suitnumber + tmp[0];
                return 3; //順子
            }
            else
            {
                suitnum = 6;
                suitvalue = tmpsuit[0] * suitnumber + tmp[0];
                return 6; //同花順
            }
        }
        else
        {
            suitnum = 0;
            suitvalue = 0;
            return 0;
        }
    }
    else //不合法的出牌數量
    {
        return 0;
    }
}

//檢查是否較大
bool checklargercards(int numberofcard)
{
    if (suitnum > 2)
    {
        if (suitnum < presuitnum)
        {
            return 0;
        }
        else if (suitnum == presuitnum)
        {
            if (suitvalue == 3) //順子
            {
                int tmp = suitvalue % suitnumber;
                int pretmp = presuitvalue % suitnumber;
                if (tmp < 2)
                {
                    tmp += suitnumber; //校正大小
                }
                if (pretmp < 2)
                {
                    pretmp += suitnumber; //校正大小
                }
                if (tmp > pretmp)
                {
                    return 1;
                }
                else if (tmp == pretmp)
                {
                    if (suitvalue > presuitvalue)
                    {
                        return 1;
                    }
                    else
                    {
                        return 0;
                    }
                }
                else
                {
                    return 0;
                }
            }
            else if (suitvalue == 4 || suitvalue == 5) //葫蘆、鐵支
            {
                int tmp = suitvalue % suitnumber;
                int pretmp = presuitvalue % suitnumber;
                if (tmp < 2)
                {
                    tmp += suitnumber; //校正大小
                }
                if (pretmp < 2)
                {
                    pretmp += suitnumber; //校正大小
                }
                if (tmp > pretmp)
                {
                    return 1;
                }
                else
                {
                    return 0;
                }
            }
            else //同花順
            {
                int tmp = suitvalue % suitnumber;
                int pretmp = presuitvalue % suitnumber;
                if (tmp < 2)
                {
                    tmp += suitnumber; //校正大小
                }
                if (pretmp < 2)
                {
                    pretmp += suitnumber; //校正大小
                }
                if (tmp > pretmp)
                {
                    return 1;
                }
                else if (tmp == pretmp)
                {
                    if (suitvalue > presuitvalue)
                    {
                        return 1;
                    }
                    else
                    {
                        return 0;
                    }
                }
                else
                {
                    return 0;
                }
            }
        }
        else if (suitnum > 4 && presuitnum <= 4) //絕對勝利
        {
            return 1;
        }
        else if (presuitnum == 0)
        {
            return 1;
        }
        else
        {
            return 0; //牌型不符
        }
    }
    else
    {
        if (suitnum == presuitnum)
        {
            int tmp = suitvalue % suitnumber;
            int pretmp = presuitvalue % suitnumber;
            if (tmp < 2)
            {
                tmp += suitnumber; //校正大小
            }
            if (pretmp < 2)
            {
                pretmp += suitnumber; //校正大小
            }
            if (tmp > pretmp)
            {
                return 1;
            }
            else if (tmp == pretmp)
            {
                if (suitvalue > presuitvalue)
                {
                    return 1;
                }
                else
                {
                    return 0;
                }
            }
            else
            {
                return 0;
            }
        }
        else if (presuitnum == 0)
        {
            return 1;
        }
        else
        {
            return 0;
        }
    }
}

//電腦玩家決定出牌，return idx(出牌數量)
int algfindingcards(int player)
{
    int tmpcheck;
    if (presuitnum == 1) //前玩家出單張
    {
        int tmp1[13], tmpsuit1[13], idxnum = 0;
        for (int i = 0; i < cardsnumber; i++)
        {
            if (cardstable[i] != player)
            {
                continue;
            }
            tmp1[idxnum] = i % suitnumber;
            if (tmp1[idxnum] < 2)
            {
                tmp1[idxnum] += suitnumber;
            }
            tmpsuit1[idxnum] = i / suitnumber;
            idxnum += 1;
        }
        for (int i = 0; i < idxnum - 1; i++) //將取出的資料排序
        {
            for (int j = 0; j < idxnum - 1 - i; j++)
            {
                if (tmp1[j] > tmp1[j + 1])
                {
                    int tmptmp = tmp1[j];
                    tmp1[j] = tmp1[j + 1];
                    tmp1[j + 1] = tmptmp;
                    tmptmp = tmpsuit1[j];
                    tmpsuit1[j] = tmpsuit1[j + 1];
                    tmpsuit1[j + 1] = tmptmp;
                }
            }
        }
        for (int i = 0; i < idxnum; i++)
        {
            if (tmp1[i] >= suitnumber)
            {
                tmp1[i] -= suitnumber;
            }
            inputcards[0] = tmpsuit1[i] * suitnumber + tmp1[i]; //找出最小的牌又能出的牌
            tmpcheck = checklegalcards(1);
            tmpcheck = tmpcheck * checklargercards(1);
            if (tmpcheck != 0)
            {
                return 1;
            }
        }
        return 0;
    }
    else if (presuitnum == 2) //前玩家出對子
    {
        for (int i = 0; i < cardsnumber - 1; i++)
        {
            if (cardstable[i] != player)
            {
                continue;
            }
            for (int j = i + 1; j < cardsnumber; j++)
            {
                if (cardstable[j] != player)
                {
                    continue;
                }
                inputcards[0] = i;
                inputcards[1] = j;
                tmpcheck = checklegalcards(2);
                tmpcheck = tmpcheck * checklargercards(2);
                if (tmpcheck != 0)
                {
                    return 2;
                }
            }
        }
        return 0;
    }
    else if (presuitnum >= 3) //前玩家出順子
    {
        for (int a = 0; a < cardsnumber - 4; a++)
        {
            if (cardstable[a] != player)
            {
                continue;
            }
            for (int b = a + 1; b < cardsnumber - 3; b++)
            {
                if (cardstable[b] != player)
                {
                    continue;
                }
                for (int c = b + 1; c < cardsnumber - 2; c++)
                {
                    if (cardstable[c] != player)
                    {
                        continue;
                    }
                    for (int d = c + 1; d < cardsnumber - 1; d++)
                    {
                        if (cardstable[d] != player)
                        {
                            continue;
                        }
                        for (int e = d + 1; e < cardsnumber; e++)
                        {
                            if (cardstable[e] != player)
                            {
                                continue;
                            }
                            inputcards[0] = a;
                            inputcards[1] = b;
                            inputcards[2] = c;
                            inputcards[3] = d;
                            inputcards[4] = e;
                            tmpcheck = checklegalcards(5);
                            tmpcheck = tmpcheck * checklargercards(5);
                            if (tmpcheck != 0)
                            {
                                return 5;
                            }
                        }
                    }
                }
            }
        }
        return 0;
    }
    else if (suitvalue == -1) //當電腦要先出時
    {
        //可以出五張時
        for (int a = 0; a < cardsnumber - 3; a++)
        {
            if (cardstable[a] != player || a == 2)
            {
                continue;
            }
            for (int b = a + 1; b < cardsnumber - 2; b++)
            {
                if (cardstable[b] != player || b == 2)
                {
                    continue;
                }
                for (int c = b + 1; c < cardsnumber - 1; c++)
                {
                    if (cardstable[c] != player || c == 2)
                    {
                        continue;
                    }
                    for (int d = c + 1; d < cardsnumber; d++)
                    {
                        if (cardstable[d] != player || d == 2)
                        {
                            continue;
                        }
                        inputcards[0] = 2; //梅花3
                        inputcards[1] = a;
                        inputcards[2] = b;
                        inputcards[3] = c;
                        inputcards[4] = d;
                        tmpcheck = checklegalcards(5);
                        tmpcheck = tmpcheck * checklargercards(5);
                        if (tmpcheck != 0)
                        {
                            return 5;
                        }
                    }
                }
            }
        }
        //可以出對子時
        for (int a = 0; a < cardsnumber; a++)
        {
            if (cardstable[a] != player || a == 2)
            {
                continue;
            }
            inputcards[0] = 2; //梅花3
            inputcards[1] = a;
            tmpcheck = checklegalcards(2);
            tmpcheck = tmpcheck * checklargercards(2);
            if (tmpcheck != 0)
            {
                return 2;
            }
        }
        //只能出單張時
        inputcards[0] = 2; //梅花3
        tmpcheck = checklegalcards(1);
        tmpcheck = tmpcheck * checklargercards(1);
        if (tmpcheck != 0)
        {
            return 1;
        }
        return 0;
    }
    else if (presuitnum == 0 && presuitvalue == 0) //場上沒有牌，前三家都過
    {
        //可以出五張時
        for (int a = 0; a < cardsnumber - 4; a++)
        {
            if (cardstable[a] != player)
            {
                continue;
            }
            for (int b = a + 1; b < cardsnumber - 3; b++)
            {
                if (cardstable[b] != player)
                {
                    continue;
                }
                for (int c = b + 1; c < cardsnumber - 2; c++)
                {
                    if (cardstable[c] != player)
                    {
                        continue;
                    }
                    for (int d = c + 1; d < cardsnumber - 1; d++)
                    {
                        if (cardstable[d] != player)
                        {
                            continue;
                        }
                        for (int e = d + 1; e < cardsnumber; e++)
                        {
                            if (cardstable[e] != player)
                            {
                                continue;
                            }
                            inputcards[0] = a;
                            inputcards[1] = b;
                            inputcards[2] = c;
                            inputcards[3] = d;
                            inputcards[4] = e;
                            tmpcheck = checklegalcards(5);
                            tmpcheck = tmpcheck * checklargercards(5);
                            if (tmpcheck != 0)
                            {
                                return 5;
                            }
                        }
                    }
                }
            }
        }
        //可以出對子時
        for (int i = 0; i < cardsnumber - 1; i++)
        {
            if (cardstable[i] != player)
            {
                continue;
            }
            for (int j = i + 1; j < cardsnumber; j++)
            {
                if (cardstable[j] != player)
                {
                    continue;
                }
                inputcards[0] = i;
                inputcards[1] = j;
                tmpcheck = checklegalcards(2);
                tmpcheck = tmpcheck * checklargercards(2);
                if (tmpcheck != 0)
                {
                    return 2;
                }
            }
        }
        //只能出單張時
        int tmp0[13], tmpsuit0[13], idxnum = 0;
        for (int i = 0; i < cardsnumber; i++)
        {
            if (cardstable[i] != player)
            {
                continue;
            }
            tmp0[idxnum] = i % suitnumber;
            if (tmp0[idxnum] < 2)
            {
                tmp0[idxnum] += suitnumber;
            }
            tmpsuit0[idxnum] = i / suitnumber;
            idxnum += 1;
        }
        for (int i = 0; i < idxnum - 1; i++) //將取出的資料排序
        {
            for (int j = 0; j < idxnum - 1 - i; j++)
            {
                if (tmp0[j] > tmp0[j + 1])
                {
                    int tmptmp = tmp0[j];
                    tmp0[j] = tmp0[j + 1];
                    tmp0[j + 1] = tmptmp;
                    tmptmp = tmpsuit0[j];
                    tmpsuit0[j] = tmpsuit0[j + 1];
                    tmpsuit0[j + 1] = tmptmp;
                }
            }
        }
        if (tmp0[0] >= suitnumber)
        {
            tmp0[0] -= suitnumber;
        }
        inputcards[0] = tmpsuit0[0] * suitnumber + tmp0[0];
        tmpcheck = checklegalcards(1);
        tmpcheck = tmpcheck * checklargercards(1);
        if (tmpcheck != 0)
        {
            return 1;
        }
        return 0;
    }
    return 0;
}

//輸出所有牌的狀態
void printhandstatus(int num)
{
    if (num == 0)
    {
        for (int j = 0; j < 13; j++)
        {
            _setmode(_fileno(stdout), _O_U16TEXT);
            if (player1[j] == -2)
            {
                //do nithing
            }
            else if (player1[j] / suitnumber == 0)
            {
                std::wcout << left << setw(3) << L"\u2663";
            }
            else if (player1[j] / suitnumber == 1)
            {
                std::wcout << left << setw(3) << L"\u2662";
            }
            else if (player1[j] / suitnumber == 2)
            {
                std::wcout << left << setw(3) << L"\u2661";
            }
            else if (player1[j] / suitnumber == 3)
            {
                std::wcout << left << setw(3) << L"\u2660";
            }
            _setmode(_fileno(stdout), _O_TEXT);
            if (player1[j] % suitnumber < 9 && player1[j] % suitnumber > 0)
            {
                cout << left << setw(2) << player1[j] % suitnumber + 1;
            }
            else if (player1[j] % suitnumber == 9)
            {
                cout << left << setw(2) << player1[j] % suitnumber + 1;
            }
            else if (player1[j] % suitnumber == 0)
            {
                cout << left << setw(2) << "A";
            }
            else if (player1[j] % suitnumber == 10)
            {
                cout << left << setw(2) << "J";
            }
            else if (player1[j] % suitnumber == 11)
            {
                cout << left << setw(2) << "Q";
            }
            else if (player1[j] % suitnumber == 12)
            {
                cout << left << setw(2) << "K";
            }
        }
        printf("\n");
    }
    else if (num < 0)
    {
        num = -num;
        printf("player %d : ", num);
        for (int j = 0; j < cardsnumber; j++)
        {
            if (cardstable[j] == num)
            {
                _setmode(_fileno(stdout), _O_U16TEXT);
                if (j / suitnumber == 0)
                {
                    std::wcout << L"\u2663";
                }
                else if (j / suitnumber == 1)
                {
                    std::wcout << L"\u2662";
                }
                else if (j / suitnumber == 2)
                {
                    std::wcout << L"\u2661";
                }
                else if (j / suitnumber == 3)
                {
                    std::wcout << L"\u2660";
                }
                _setmode(_fileno(stdout), _O_TEXT);
                if (j % suitnumber < 9 && j % suitnumber > 0)
                {
                    printf(" %d  ", j % suitnumber + 1);
                }
                else if (j % suitnumber == 9)
                {
                    printf("%d  ", j % suitnumber + 1);
                }
                else if (j % suitnumber == 0)
                {
                    printf(" A  ");
                }
                else if (j % suitnumber == 10)
                {
                    printf(" J  ");
                }
                else if (j % suitnumber == 11)
                {
                    printf(" Q  ");
                }
                else if (j % suitnumber == 12)
                {
                    printf(" K  ");
                }
            }
        }
        printf("\n");
    }
    else
    {
        printf("player plays : ");
        for (int i = 0; i < num; i++)
        {
            int j = inputcards[i];
            _setmode(_fileno(stdout), _O_U16TEXT);
            if (j / suitnumber == 0)
            {
                std::wcout << L"\u2663";
            }
            else if (j / suitnumber == 1)
            {
                std::wcout << L"\u2662";
            }
            else if (j / suitnumber == 2)
            {
                std::wcout << L"\u2661";
            }
            else
            {
                std::wcout << L"\u2660";
            }
            _setmode(_fileno(stdout), _O_TEXT);
            if (j % suitnumber < 9 && j % suitnumber > 0)
            {
                printf(" %d  ", j % suitnumber + 1);
            }
            else if (j % suitnumber == 9)
            {
                printf("%d  ", j % suitnumber + 1);
            }
            else if (j % suitnumber == 0)
            {
                printf(" A  ");
            }
            else if (j % suitnumber == 10)
            {
                printf(" J  ");
            }
            else if (j % suitnumber == 11)
            {
                printf(" Q  ");
            }
            else if (j % suitnumber == 12)
            {
                printf(" K  ");
            }
        }
        printf("\n\n");
    }
    return;
}

void printmemorystatus(int num)
{
    if (num == 0)
    {
        //do nothing
    }
    else
    {
        printf("COM plays : ");
        for (int i = 0; i < num; i++)
        {
            int j = memory[i];
            _setmode(_fileno(stdout), _O_U16TEXT);
            if (j / suitnumber == 0)
            {
                std::wcout << L"\u2663";
            }
            else if (j / suitnumber == 1)
            {
                std::wcout << L"\u2662";
            }
            else if (j / suitnumber == 2)
            {
                std::wcout << L"\u2661";
            }
            else
            {
                std::wcout << L"\u2660";
            }
            _setmode(_fileno(stdout), _O_TEXT);
            if (j % suitnumber < 9 && j % suitnumber > 0)
            {
                printf(" %d  ", j % suitnumber + 1);
            }
            else if (j % suitnumber == 9)
            {
                printf("%d  ", j % suitnumber + 1);
            }
            else if (j % suitnumber == 0)
            {
                printf(" A  ");
            }
            else if (j % suitnumber == 10)
            {
                printf(" J  ");
            }
            else if (j % suitnumber == 11)
            {
                printf(" Q  ");
            }
            else if (j % suitnumber == 12)
            {
                printf(" K  ");
            }
        }
        printf("\n\n");
    }
    return;
}

void sorting(int x[], int n)
{

    for (int i = n - 1; i > 0; i--)
    {
        for (int j = 0; j <= i - 1; j++)
        {
            if (x[j] %13> x[j + 1]%13)
            {
                swap(x[j], x[j+1]);
            }
        }
    }
    for (int i = n - 1; i > 0; i--)
    {
        for (int j = 0; j <= i - 1; j++)
        {
            if (x[j] % 13 == x[j + 1] % 13 && x[j] > x[j + 1])
            {
                swap(x[j], x[j + 1]);
            }
        }
    }
}


void opening()
{
    wcout << "         o   o   o   o     o   o   o   o       o          o    o   o   o   o        o  o   o   o   " << endl;Sleep(200);SetColor(11);
    wcout << "         o           o     o           o       o        o      o                    o          o   " << endl;Sleep(200);SetColor(12);
    wcout << "         o           o     o           o       o       o       o                    o          o   " << endl;Sleep(200);SetColor(13);
    wcout << "         o           o     o           o       o     o         o                    o          o   " << endl;Sleep(200);SetColor(14);
    wcout << "         o           o     o           o       o   o           o                    o          o   " << endl;Sleep(200);SetColor(15);
    wcout << "         o   o   o   o     o           o       o o             o   o   o   o        o  o   o   o   " << endl;Sleep(200);SetColor(11);
    wcout << "         o                 o           o       o o             o                    o  o           " << endl;Sleep(200);SetColor(12);
    wcout << "         o                 o           o       o   o           o                    o    o         " << endl;Sleep(200);SetColor(13);
    wcout << "         o                 o           o       o     o         o                    o      o       " << endl;Sleep(200);SetColor(14);
    wcout << "         o                 o           o       o       o       o                    o        o     " << endl;Sleep(200);SetColor(15);
    wcout << "         o                 o           o       o         o     o                    o          o   " << endl;Sleep(200);SetColor(12);
    wcout << "         o                 o   o   o   o       o           o   o   o   o   o        o            o " << endl;Sleep(200);SetColor(13);
    SetColor();
    wcout << "                                       --press s to start--                                    " << endl;
}

// 執行程式: Ctrl + F5 或 [偵錯] > [啟動但不偵錯] 功能表
// 偵錯程式: F5 或 [偵錯] > [啟動偵錯] 功能表

// 開始使用的提示: 
//   1. 使用 [方案總管] 視窗，新增/管理檔案
//   2. 使用 [Team Explorer] 視窗，連線到原始檔控制
//   3. 使用 [輸出] 視窗，參閱組建輸出與其他訊息
//   4. 使用 [錯誤清單] 視窗，檢視錯誤
//   5. 前往 [專案] > [新增項目]，建立新的程式碼檔案，或是前往 [專案] > [新增現有項目]，將現有程式碼檔案新增至專案
//   6. 之後要再次開啟此專案時，請前往 [檔案] > [開啟] > [專案]，然後選取 .sln 檔案