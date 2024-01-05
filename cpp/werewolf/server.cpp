//Example code: A simple server side code, which echos back the received message.
//Handle multiple socket connections with select and fd_set on Linux 
#include <iostream>
#include <stdio.h>
#include <cstring>
#include <cstdlib> 
#include <errno.h> 
#include <unistd.h>   //close 
#include <arpa/inet.h>    //close 
#include <sys/types.h> 
#include <sys/socket.h> 
#include <netinet/in.h> 
#include <sys/time.h> //FD_SET, FD_ISSET, FD_ZERO macros
#include <ctime> // for random
    
#define TRUE   1 
#define FALSE  0 
#define PORT 8888

const int playersCount = 6;
    
int main(int argc , char *argv[])  
{  
    int opt = TRUE;  
    int master_socket , addrlen , new_socket , client_socket[30] , 
          max_clients = 30 , activity, i , valread , sd, connection = 0;  
    int max_sd;  
    struct sockaddr_in address;  
        
    char buffer[1025];  //data buffer of 1K 
        
    //set of socket descriptors 
    fd_set readfds;  
        
    //a message 
    char *message = "歡迎來到天黑請閉眼\r\n等待其他玩家連線...";  
    
    //initialise all client_socket[] to 0 so not checked 
    for (i = 0; i < max_clients; i++)  
    {  
        client_socket[i] = 0;  
    }  
        
    //create a master socket 
    if( (master_socket = socket(AF_INET , SOCK_STREAM , 0)) == 0)  
    {  
        perror("socket failed");  
        exit(EXIT_FAILURE);  
    }  
    
    //set master socket to allow multiple connections , 
    //this is just a good habit, it will work without this 
    if( setsockopt(master_socket, SOL_SOCKET, SO_REUSEADDR, (char *)&opt, 
          sizeof(opt)) < 0 )  
    {  
        perror("setsockopt");  
        exit(EXIT_FAILURE);  
    }  
    
    //type of socket created 
    address.sin_family = AF_INET;  
    address.sin_addr.s_addr = INADDR_ANY;  
    address.sin_port = htons( PORT );  
        
    //bind the socket to localhost port 8888 
    if (bind(master_socket, (struct sockaddr *)&address, sizeof(address))<0)  
    {  
        perror("bind failed");  
        exit(EXIT_FAILURE);  
    }  
    printf("Listener on port %d \n", PORT);  
        
    //try to specify maximum of 3 pending connections for the master socket 
    if (listen(master_socket, 3) < 0)  
    {  
        perror("listen");  
        exit(EXIT_FAILURE);  
    }  
        
    //accept the incoming connection 
    addrlen = sizeof(address);  
    puts("Waiting for connections ...");


    int index[6] = {1,1,1,1,1,1}; // 玩家身份 1=平民 2=警察 3=殺手
    char name[6][100] = {0}; // 儲存每個玩家所輸入的名字
    bool gameStart = false; // 是否開始遊戲
    bool alive[6]={1,1,1,1,1,1};//記錄誰還活著
    int killed; // 被殺的玩家編號
    int check; // 被警察指認的玩家編號

    // 等待所有玩家連線
    for (int i = 0; i < playersCount; ++i)
    {
      new_socket = accept(master_socket, (struct sockaddr *)&address, (socklen_t*)&addrlen);
      client_socket[i] = new_socket;
      message = "歡迎來到天黑請閉眼\n連線成功\n等待其他玩家連線中...\n\n\0";

      send(new_socket, message, strlen(message), 0);
    }
    // 所有玩家連線成功
    for (int i = 0; i < playersCount; ++i)
    {
       message = "所有玩家已連線\n遊戲即將開始\n請等待其他玩家輸入名字\n\n\0";

       send(client_socket[i], message, strlen(message), 0);
    }
    // 接收玩家輸入的名字並儲存
    for (int i = 0; i < playersCount; ++i)
    {
      message = "請輸入你的名字：\n\0";
      send(client_socket[i], message, strlen(message), 0);
      int k = read(client_socket[i], buffer, 1024);
      message = "名字輸入成功！\n正在等待所有玩家輸入名字...\n\n\0";
      send(client_socket[i], message, strlen(message), 0);
      
      buffer[k] = '\0';
      strcpy(name[i], buffer); // 儲存每個玩家的名字
    }

    for (int i = 0; i < playersCount; ++i)
    {
      for (int j = 0; j < playersCount; ++j)
      {
        send(client_socket[i], name[j], strlen(name[j]), 0);
        message = "是\0";
        send(client_socket[i], message, strlen(message), 0);
        if (j == 0)
        {
          message = "0號 \n\0";
        }
        else if (j == 1)
        {
          message = "1號 \n\0";
        }
        else if (j == 2)
        {
          message = "2號 \n\0";
        }
        else if (j == 3)
        {
          message = "3號 \n\0";
        }
        else if (j == 4)
        {
          message = "4號 \n\0";
        }
        else if (j == 5)
        {
          message = "5號 \n\0";
        }
        send(client_socket[i], message, strlen(message), 0);
      }
      // message = "quit";
      // send(client_socket[i], message, strlen(message), 0);
    }

    // 隨機分配
    srand(time(0));
    int rn = 0;
    // 隨機分配殺手的玩家編號
    rn = rand() % playersCount; // 0~5
    index[rn] = 3;
    // 隨機分配警察的玩家編號
    rn = rand() % playersCount;
    while(index[rn] == 3) // 如果是殺手就在隨機分配下一個玩家
      rn = rand() % playersCount;
    index[rn] = 2;

    // 傳給每個人自己的身份
    for (int i = 0; i < playersCount; ++i)
    {
      if (index[i] == 1)
      {
        message = "\n***你是平民！***\n\n\0";
        send(client_socket[i], message, strlen(message), 0);
      }
      else if (index[i] == 2)
      {
        message = "\n***你是警察！***\n\n\0";
        send(client_socket[i], message, strlen(message), 0);
      }
      else if (index[i] == 3)
      {
        message = "\n***你是殺手！***\n\n\0";
        send(client_socket[i], message, strlen(message), 0);
      }
    }
    int people = playersCount;
    while(1) // people > 3
    {
      // 天黑
      for (int i = 0; i < playersCount; ++i)
      {
        message = "天黑了\n\n\0";
        send(client_socket[i], message, strlen(message), 0);
        message = "請殺手輸入要殺的玩家編號\n\0";
        send(client_socket[i], message, strlen(message), 0);
      }
      // 殺手殺人
      for (int i = 0; i < playersCount; ++i)
      {
        if (index[i] == 3)
        {
          read(client_socket[i], buffer, 1024);
          killed = buffer[0] - '0';
          while (killed == i || alive[killed] == 0 || killed >= playersCount)
          {
            if (check == i)
              message = "不要殺自己喔～請再輸入一次\n\0";
            else if (alive[check] == 0)
              message = "他已經屎了喔～請再輸入一次\n\0";
            else if (check >= playersCount)
              message = "沒有這個人～請再投一次\n\0";
            send(client_socket[i], message, strlen(message), 0);
            read(client_socket[i], buffer, 1024);
            killed = buffer[0] - '0';
          }
          alive[killed] = 0;
          people--;
          break;
        }
      }

      message = "殺手殺完人了！請警察輸入要指認的玩家編號\n\0";
      for (int i = 0; i < playersCount; ++i)
        send(client_socket[i], message, strlen(message), 0);

      // 警察指認
      for (int i = 0; i < playersCount; ++i)
      {
        if (index[i] == 2) // 警察
        {
          if (alive[i] == 1)
          {
            check = read(client_socket[i], buffer, 1024);
            check = buffer[0] - '0';
            while (check == i || alive[check] == 0 || check >= playersCount)
            {
              if (check == i)
                message = "不要指認自己喔～請再輸入一次\n\0";
              else if (alive[check] == 0)
                message = "他已經屎了喔～請再輸入一次\n\0";
              else if (check >= playersCount)
                message = "沒有這個人～請再投一次\n\0";
              send(client_socket[i], message, strlen(message), 0);
              read(client_socket[i], buffer, 1024);
              check = buffer[0] - '0';
            }
            if (index[check] == 1)
              message = "\n***他是平民！***\n\n\0";
            else if (index[check] == 3)
              message = "\n***他是殺手！！！***\n\n\0";
            send(client_socket[i], message, strlen(message), 0);
            break;
          }
          else
          {
            message = "\n***你已經屎了***\n\n\0";
            send(client_socket[i], message, strlen(message), 0);
          }
        }
      }

      // 天亮
      for (int i = 0; i < playersCount; ++i)
      {
        message = "警察指認完！天亮了～\n\0";
        send(client_socket[i], message, strlen(message), 0);
        message = "昨天\0";
        send(client_socket[i], name[killed], strlen(name[killed]), 0);
        message = "被殺了！\n\0";
        send(client_socket[i], message, strlen(message), 0);
        message = "請討論誰是兇手\n\0";
        send(client_socket[i], message, strlen(message), 0);
        message = "等待其他玩家投票中...\n\0";
        send(client_socket[i], message, strlen(message), 0);
      }

      // 投票
      int vote[6] = {-1,-1,-1,-1,-1,-1};
      for (int i = 0; i < playersCount; ++i)
      {
        if (alive[i] == 1)
        {
          message = "請投票，你覺得誰是殺手？\n\0";
          send(client_socket[i], message, strlen(message), 0);
          read(client_socket[i], buffer, 1024);
          int k = buffer[0] - '0';
          while (alive[k] == 0 || k == i || k >= playersCount)
          {
            if (alive[k] == 0)
              message = "他已經屎了喔～請再投一次\n\0";
            else if (k == i)
              message = "不要投給自己喔～請再投一次\n\0";
            else if (k <= playersCount)
              message = "沒有這個人～請再投一次\n\0";
            send(client_socket[i], message, strlen(message), 0);
            read(client_socket[i], buffer, 1024);
            k = buffer[0] - '0';
          }
          vote[i] = k;
          message = "已成功投票！\n\n\0";
          send(client_socket[i], message, strlen(message), 0);
        }
        else
        {
          message = "\n***你已經屎了！***\n\n\0";
          send(client_socket[i], message, strlen(message), 0);
        }
      }

      // 公布投票結果
      for (int i = 0; i < playersCount; ++i)
      {
        for (int j = 0; j < playersCount; ++j)
        {
          if (alive[j] == 1)
          {
            if (j == 0)
              message = "0號投給\0";
            else if (j == 1)
              message = "1號投給\0";
            else if (j == 2)
              message = "2號投給\0";
            else if (j == 3)
              message = "3號投給\0";
            else if (j == 4)
              message = "4號投給\0";
            else if (j == 5)
              message = "5號投給\0";
            send(client_socket[i], message, strlen(message), 0);
            if (vote[j] == 0)
              message = "0號\n\0";
            else if (vote[j] == 1)
              message = "1號\n\0";
            else if (vote[j] == 2)
              message = "2號\n\0";
            else if (vote[j] == 3)
              message = "3號\n\0";
            else if (vote[j] == 4)
              message = "4號\n\0";
            else if (vote[j] == 5)
              message = "5號\n\0";
            send(client_socket[i], message, strlen(message), 0);
          }
          else
          {
            if (j == 0)
              message = "0號已經屎了！\n\0";
            else if (j == 1)
              message = "1號已經屎了！\n\0";
            else if (j == 2)
              message = "2號已經屎了！\n\0";
            else if (j == 3)
              message = "3號已經屎了！\n\0";
            else if (j == 4)
              message = "4號已經屎了！\n\0";
            else if (j == 5)
              message = "5號已經屎了！\n\0";
            send(client_socket[i], message, strlen(message), 0);
          }
        }
      }
      // 計算最高票
      int ticket[6] = {0};
      for (int i = 0; i < playersCount; ++i)
        if (vote[i] != -1)
          ticket[vote[i]]++;
      int max_tic = 0, max_player = -1;
      for (int i = 0; i < playersCount; ++i)
      {
        if (ticket[i] > max_tic)
        {
          max_tic = ticket[i];
          max_player = i;
        }
        // 同票出局任何人
        else if (ticket[i] == max_tic)
          max_player = -1;
      }
      if (max_player != -1)
      {
        alive[max_player] = 0;
        people--;
        if (max_player == 0)
          message = "最高票是0號，他被強制出局了！\n\0";
        else if (max_player == 1)
          message = "最高票是1號，他被強制出局了！\n\0";
        else if (max_player == 2)
          message = "最高票是2號，他被強制出局了！\n\0";
        else if (max_player == 3)
          message = "最高票是3號，他被強制出局了！\n\0";
        else if (max_player == 4)
          message = "最高票是4號，他被強制出局了！\n\0";
        else if (max_player == 5)
          message = "最高票是5號，他被強制出局了！\n\0";
        for (int i = 0; i < playersCount; ++i)
          send(client_socket[i], message, strlen(message), 0);
        if (index[max_player] == 3)
        {
          message = "殺手被強制出局了，恭喜正義的一方獲得勝利！\n\0";
          for (int i = 0; i < playersCount; ++i)
            send(client_socket[i], message, strlen(message), 0);
          break;
        }
      }
      else
      {
        message = "票數相同，沒有人出局！\n\0";
        for (int i = 0; i < playersCount; ++i)
          send(client_socket[i], message, strlen(message), 0);
      }
      if (people == 2)
      {
        message = "剩下2名玩家\n殺手獲得最後的勝利！遊戲結束\n\0";
        for (int i = 0; i < playersCount; ++i)
          send(client_socket[i], message, strlen(message), 0);
        sleep(5);
        break;
      }
    }

    close(master_socket);
        
    return 0;  
}  