// - Specify how many set of Poker to be used/played
// - Specify how many players
// - Show out result of game

#include "poker.h"
#include <stdbool.h>
#define RUNNING 1

static char *FACE[] = {"1", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"};
static char *SUIT[] = {"H", "D", "C", "S"};

void print_rule(void){

    printf("\n\033[37m============================================================================================================\033[0m\n");
    printf("\033[37m|\033[0m\t\t\t\t \033[34m Welcome to the Poker-21 game! \033[0m\t\t\t\t\t\033[37m   |\033[0m\n");
    printf("\033[37m|\033[0m     21點是一種經典的撲克牌遊戲，每局的目的是讓手中的牌點數加起來等於或接近21點，但不能超過21點。         \033[37m|\033[0m\n");
    printf("\033[37m|\033[0m     遊戲開始後，玩家先行動，可以選擇要牌或停牌。若玩家選擇要牌，莊家就會從牌堆中給玩家一張牌             \033[37m|\033[0m\n");
    printf("\033[37m|\033[0m     ，玩家的牌點數會隨之增加。玩家可以一直要牌直到他們的牌點數等於或超過21點，此時玩家爆牌了，失去本局。 \033[37m|\033[0m\n");
    printf("\033[37m|\033[0m     最後，玩家和莊家的牌點數會被比較，如果玩家的牌點數比莊家高，玩家贏得本局；                           \033[37m|\033[0m\n");
    printf("\033[37m|\033[0m     如果玩家的牌點數等於莊家的牌點數，則是和局；否則莊家贏。                                             \033[37m|\033[0m\n");
    printf("\033[37m============================================================================================================\033[0m\n");
}

void print_station(int n, int player_t, float over21, float under21){

    printf("\n=========================================================================\n");
    printf("|\t\t\t\t \033[34m Information\033[0m \t\t\t\t|\n");
    printf("|\t\t\t\t\t\t\t\t\t|\n");
    printf("| You have used %d-set Poker cards.\t\t\t\t\t|\n", n);
    printf("| There are %d players.\t\t\t\t\t\t\t|\n", player_t+1);
    printf("| Betting odds under 21 is %2.2f.\t\t\t\t\t|\n", under21);
    printf("| Betting odds over 21 is %2.2f.\t\t\t\t\t\t|\n", over21);
    printf("=========================================================================\n\n");
    // printf("=========================================================================\n");

    // printf("=========================================================================\n");
}
int main(void){

    int n=0;        // # of card sets
    int player_n; // # of players
    bool restart = 1;
    int flag = 1; 
    float betting_odds_over21, betting_odds_under21; // betting ratio
    print_rule();

    

    // 1. Initialization
    printf("\t\t\033[43;30m Now we need to know some information \033[0m\n");
    printf("\tPlease specify how many players for this Poker-21: \t");
    scanf("%d", &player_n);
    Poker21 *player = malloc((player_n + 1) * sizeof(Poker21));
    while(restart == RUNNING)      // define stats
    {
        printf("\tPlease specify how may set of cards to play: \t");
        scanf("%d", &n);
        Card *deck = malloc(52 * n * sizeof(Card));


        printf("\tPlease enter the betting odds under 21: \t");
        scanf("%f", &betting_odds_under21);

        printf("\tPlease enter the betting odds over 21: \t");
        scanf("%f", &betting_odds_over21);
        print_station(n, player_n, betting_odds_over21, betting_odds_under21);
        // wallet money
        printf("\t\t\t\t\033[34m Wallet \033[0m\n");
        if (flag == 1){
            
            for(int i = 0; i < player_n + 1; i++){
                printf("\tPlayer %d,please enter how many money do you have: ",i);
                scanf("%f", &(player[i].purse));
                //錢包的錢不能為負數
                while ((player[i].purse < 0) == 1)
                {
                    printf("\tPurse shall be a positive number! Player %d,please enter how many money do you have: ",i);
                    scanf("%f", &(player[i].purse));
                }
            }
            flag = 0;
        }
        else if (flag == 0){
            for (int i = 0; i < player_n+1; i++){
                printf("\tPlayer %d,now your wallet have money : $%2.2f \n", i, player[i].purse);
            }
        }
        printf("=========================================================================\n");
        // bet money
        printf("\t\t\t\t\033[34m Lay A Bet \033[0m\n");
        for(int i = 0; i <= player_n; i++)
        {
            player[i].card_n = 0;
            player[i].sum = 0;
            printf("\tPlayer #%d, Please enter how much money you want to bet: ",i);
            scanf("%f", &(player[i].stake));

            //下注的錢不能超過錢包的錢
            while( (player[i].stake > player[i].purse)  == 1)
            {   
                printf("\tStake shall not exceed your wallet! Player #%d, Please enter how much money you want to bet: ",i);
                scanf("%f", &(player[i].stake)); 
            }
            
        }

        // 2. Fill n-set of Poker cards & shuffle cards.
        srand(time(NULL));
        fillDeck_n(deck, FACE, SUIT, n);
        shuffle_n(deck, n);

        // 3. Deal 1st-round cards (2 cards for the Banker & eacher Player)
        printf("\n\n\t\t\t\033[34m Now Beginning!! \n\n\033[0m");
        printf("\033[32m======================== 1st ROUND =====================\033[0m\n");
        int deal_begin = 0;
        for (int i = 1; i <= player_n; i++)
        {
            deal_Poker21(deck, &player[i], n, deal_begin, 2);
            deal_begin += 2;
            show_player(&player[i], i);
        }

        // 4. Deal 2nd-round cards for each player
        char YN;
        printf("\033[32m======================== 2nd ROUND =====================\033[0m\n");
        for (int i = 1; i <= player_n; i++)
        {
            show_player(&player[i], i);
            printf("Player %d, do you like to add another card (Y/N): ", i);
            scanf(" %c", &YN);
            while (YN == 'Y' | YN == 'y')
            {
                deal_Poker21(deck, &player[i], n, deal_begin, 1);
                deal_begin++;

                show_player(&player[i], i);

                if (*(&player[i].sum) > 21)
                {
                    printf("Player#%d's point total exceeds 21 points! \n",i);
                    break;
                }
                
                printf("Player %d, do you like to add another card (Y/N): ", i);
                scanf(" %c", &YN);
            }
        }

        deal_Poker21(deck, &player[0], n, deal_begin, 2);
        deal_begin += 2;

        // 5. Deal 2nd-round cards for the Banker
        show_player(&player[0], 0);
        while (player[0].sum < 16)
        {
            deal_Poker21(deck, &player[0], n, deal_begin, 1);
            deal_begin += 1;

            show_player(&player[0], 0);
        }

        // 6. Get the result of the game
        for (int i = 1; i <= player_n; i++)
        {
            //玩家和莊家的點數小於21點，且玩家點數大於莊家或是莊家點數大於21點，玩家即獲勝且贏錢。
            // if ((player[i].sum > player[0].sum & player[i].sum <= 21) | (player[i].sum <= 21 & player[0].sum > 21))
            if ((player[i].sum <= 21) & (player[0].sum > 21 | player[i].sum > player[0].sum) )
            {
                printf("Player #%d WON! Player #%d earns money: $%2.2f !\n", i, i, player[i].stake*betting_odds_under21);
                // printf("Player #%d earns money: $%d!\n", i, player[i].stake*betting_odds_under21);
                player[i].purse = player[i].purse + player[i].stake*betting_odds_under21;
                player[0].purse = player[0].purse - player[i].stake*betting_odds_under21;
            }
            //玩家點數等於莊家點數，平手以及賠錢        
            else if (player[i].sum == player[0].sum )
            {   
                printf("Player #%d is SAME with Player #0! Player #%d losses money: $%2.2f !\n", i,i ,player[i].stake*betting_odds_under21);
                // printf("Player #%d losses money: $%d!\n", i, player[i].stake*betting_odds_under21);
                player[i].purse = player[i].purse - player[i].stake*betting_odds_under21;
                player[0].purse = player[0].purse + player[i].stake*betting_odds_under21;

            }
            //玩家點數大於21點且莊家點數小於21點，或玩家點數小於莊家點數，輸以及賠錢
            else if ((player[i].sum > 21 & player[0].sum <=21) | (player[i].sum < player[0].sum))
            {
                printf("Player #%d is LOSS! Player #%d losses money: $%2.2f !\n", i, i, player[i].stake*betting_odds_over21);
                // printf("Player #%d losses money: $%d !\n", i, player[i].stake*betting_odds_under21);
                player[i].purse = player[i].purse - player[i].stake*betting_odds_over21;
                player[0].purse = player[0].purse + player[i].stake*betting_odds_over21;
            }
            else
            {
                printf("Something Wrong!!!");
            }       
            // player[0].purse = player[0].purse - player[0].stake*betting_odds_under21;
        }

        // 7. Restart
        //If someone runs out of money, the game will be forced to end
        for (int i = 0; i <= player_n; i++)
        {
            if (player[i].purse < 0)
            {
                restart = 0;
                printf("\nGame Over!!! Someone losts a lot of money in gamble!!!\n\n");
                printf("\t\t\t\t\033[34m Wallet \033[0m\n");

                for (int i = 0; i < player_n+1; i++){
                    printf("\tPlayer %d,now the money in your wallet is : $%2.2f \n", i, player[i].purse);
                }
                exit(0);
            }
            
        }
        //If all players still have money, will ask whether to continue to play
        printf("\n請問要再來一局嗎(Y/N):");
        scanf(" %c", &YN);
        while (YN == 'n' | YN == 'N'){
            restart = 0;
            printf("\nGame Over!!!\n\n");
            printf("\t\t\t\t\033[34m Wallet \033[0m\n");

            for (int i = 0; i < player_n+1; i++){
                printf("\tPlayer %d,now the money in your wallet is : $%2.2f \n", i, player[i].purse);
            }
            exit(1);
        }
        // free memory
        free(deck);
    }
    return 0;
}  
        //free memory
        //free(deck);

