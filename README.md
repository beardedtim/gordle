## Gordle

A Wordle helper written in go

## Usage

This is to be used after you have guessed a word and have
been given feedback from the Wordle system.

### Example Wordle Board to Input

![Image of Wordle Board](https://media.npr.org/assets/img/2022/01/21/wordle-2_wide-6d11eb7fdae3cf829fb0650d92d10ea4c9f0c174-s1100-c50.jpg)

Given the Wordle Board above, we would put in the following inputs, followed by
the enter key

```sh
go run .
What is the letter for box 1? Enter "?" for unknown
t # enter in t for first letter

Is this in the right spot (is it green?) Y/N
n # enter in n as it is not green

You told me input 1 was t and you told me that it was not green
What is the letter for box 2? Enter "?" for unknown
o # enter in o for second letter 

Is this in the right spot (is it green?) Y/N
y # enter y as it is green

You told me input 2 was o and you told me that it was green
What is the letter for box 3? Enter "?" for unknown
? # enter ? since it was gray

You told me you don't know 3 character so we are going onto the next
What is the letter for box 4? Enter "?" for unknown
? # enter ? since it was gray

You told me you don't know 4 character so we are going onto the next
What is the letter for box 5? Enter "?" for unknown
? # enter ? since it was gray

You told me you don't know 5 character so we are going onto the next
What are the known "bad" letters (ones that are not valid for the solution)? Enter separated by a comma.
Example: 
a,d,t,f
g,u,e,s,h,w,d,a,y # enter in all gray letters on the board that are not green or yellow elsewhere

[moton vomit conto oobit volti noint joint motto lotic rozit ponto rotto potto pokit robot nooit motif potin roton molto potoo rotor lotto compt bortz coopt point cobot potro poort motor]
```

The list at the bottom are all possible words to be played
