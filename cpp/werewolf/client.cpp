#include <iostream>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <thread>

using namespace std;

void inputAndSend(int sockfd)
{
  char message[100];
  while(1)
  {
    cin  >> message;
    send(sockfd, message, sizeof(message), 0);
  }
}

int main(int argc , char *argv[])
{

    //socket的建立
    int sockfd = 0;
    sockfd = socket(AF_INET , SOCK_STREAM , 0);

    if (sockfd == -1){
        printf("Fail to create a socket.");
    }

    //socket的連線

    struct sockaddr_in info;
    bzero(&info,sizeof(info));
    info.sin_family = PF_INET;

    //localhost test
    info.sin_addr.s_addr = inet_addr("127.0.0.1");
    info.sin_port = htons(8888);

    //Send a message to server
    char message[1000];
    char receiveMessage[1000];


    int sc = connect(sockfd,(struct sockaddr *)&info,sizeof(info));
    if(sc==-1){
        printf("Connection error");
        return 0;
    }

    thread input(inputAndSend, sockfd);
    input.detach();

    while (1)
    {
        int re = recv(sockfd, receiveMessage, sizeof(receiveMessage), 0);
        receiveMessage[re] = '\0';
        cout << receiveMessage;
        if (!strcmp(receiveMessage, "殺手獲得最後的勝利！遊戲結束\n\0") || !strcmp(receiveMessage, "殺手被強制出局了，恭喜正義的一方獲得勝利！\n\0"))
        {
          close(sockfd);
          cout << "ㄅㄅ～";
          return 0;
        }
    }


    printf("close Socket\n");
    close(sockfd);
    return 0;
}