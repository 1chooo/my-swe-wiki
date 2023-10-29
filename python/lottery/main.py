# -*- coding: utf-8 -*-

from rule import Lottery

def main():

    confirm_or_not = 0

    while confirm_or_not != 1:

        print('___________________________________________________')
        print('|     Welcome to the lottery number generater     |')
        print('|  In case of all of you want to buy the lottery, |')
        print('|     but have no idea to pick up the number.     |')
        print('|     Here comes the supportive tool for you!     |')
        print('|_________________________________________________|')

        print()
        start, end = map(int, input('Please input the range of the lottery number: ').split())
        num = int(input('Please input the number of the lottery: '))
        sets = int(input('Please input how many sets you want to purchase: '))

        print()
        print('===================================================')

        print('\nHere are the informations we have gotten.')
        print(f'The range from the lottery is from {start} to {end}.')
        print(f'The number of the lottery we need to generate is {num}.')
        print(f'The sets you want to purchase is {sets}.')

        print()
        print('===================================================')

        print()
        print('___________________________________________________')
        print('|              Please type (1) that               |')
        print('|   we can make sure whole information are right  |')
        print('|        type (2) to correct the information      |')
        print('|   We are so appreciate that you give us info.   |')
        print('|_________________________________________________|')
        confirm_or_not = int(input('Please type in here: '))
        print()

        print('Hit the jackpot!')

    return start, end, num, sets


if __name__ == '__main__':


    main()

    lottery = Lottery()
