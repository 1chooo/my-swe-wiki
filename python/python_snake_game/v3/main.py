from turtle import Screen
from snake import Snake
from food import Food
from scoreboard import Scoreboard
import time

def main() -> None:
    screen = Screen()
    screen.setup(600, 600)
    screen.bgcolor("black")
    screen.title("Snake game")
    screen.tracer(0)

    snake = Snake()
    food = Food()
    scoreboard = Scoreboard()
    screen.update()

    screen.listen()
    screen.onkey(snake.up, "Up")
    screen.onkey(snake.down, "Down")
    screen.onkey(snake.left, "Left")
    screen.onkey(snake.right, "Right")

    game_is_on = True
    while game_is_on:
        screen.update()
        time.sleep(0.1)
        snake.move()

        #detect collision with food
        if snake.head.distance(food) < 15:
            snake.extend()
            food.refresh()
            scoreboard.increase_score()

        #detect collision with wall
        if snake.head.xcor() > 280 or snake.head.xcor() < -280 or snake.head.ycor() > 280 or snake.head.ycor() < -280:
            scoreboard.reset()
            snake.reset()

        #detect collision with tale
        for segment in snake.segments[1:]:
            if snake.head.distance(segment) < 10:
                scoreboard.reset()
                snake.reset()

    screen.exitonclick()

if __name__ == "__main__":
    main()
