# Lab 01 - Testing

In this lab we will get more hands on practice writing unit tests, and learn how to test HTTP handlers.

## HTTP Handler Testing
Learning how to test your http handlers is dead simple, and will save you a lot of time and problems as the assignments get harder. It will also make you a better developer :)

Go into the `test_handlers` folder, there is a simple program that listens on `localhost:4242` and returns a tailored random compliment for any user. By passing a parameter to `/compliment?name=<name>`, it will respond with something like: `Brendan, you are bodacious`.

We want to test the ComplimentHandler to make sure that it is behaving as it should. 


Links: [Source](https://elithrar.github.io/article/testing-http-handlers-go/ "Testing Your (HTTP) Handlers in Go "), 

## Unit Testing Practice
Now that you've been properly complimented, let's get to more practical matters. You recently started working for the FCC on a team that automatically detects swear words on television. 

TODO: rewrite this. `They brought you on because they know that you have experience writing unit tests
Before you arrived this project had recently hit a snag, and the lead developer`

In `/swears`, is a simple program that reads in two text files: a `known_swears.txt` file that has every single swearword ever in it, _totally_, and a file passed as a command line argument, that is then scanned for any swear words. The program then reports the tally of each swear word that it finds in the text file.

Ethan was tired when he wrote this code, so there are _definitely_ no bugs in it. 

Your job is to finish writing the unit tests for the four function stubs in `swears_test.go`.
Try and think of weird edge cases as you write these tests! How are the files being read in? What swear words is it checking for, etc?