import turtle
monty = turtle.Turtle()
monty.shape("turtle")
monty.color('green')
monty.fillcolor('blue')
monty.begin_fill()
for i in range(4):
  monty.forward(100)
  monty.left(90)
monty.end_fill()
