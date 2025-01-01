# FILLLABSQUESTIONS 🚀
This project contains solutions to three different problems implemented in Golang. Each problem is structured as a separate module, making it modular and easy to maintain. 🛠️

## Project Structure 📂
FILLLABSQUESTIONS/
│

├── go.mod              # Go module definition

├── main.go             # Main file to run all questions

│

└── questions/          # Contains solutions for all questions
  
    ├── first_question/
   
    │   ├── first.go        # Solution for Question 1
   
    │   └── first_test.go   # Unit tests for Question 1
   
    │
   
    ├── second_question/
   
    │   ├── second.go       # Solution for Question 2
   
    │   └── second_test.go  # Unit tests for Question 2
   
    │
  
    └── third_question/
    
        ├── third.go        # Solution for Question 3
      
        └── third_test.go   # Unit tests for Question 3

## Questions Explanation 💡

1️⃣ Question 1: Sorting Words

Sorts a list of words based on the number of occurrences of the character "a" (in descending order). If two words have the same number of "a"s, they are sorted by length.

2️⃣ Question 2: Recursive Number Generation

Generates a sequence of numbers recursively based on a specific algorithm.

3️⃣ Question 3: Most Repeated Element

Finds the most repeated element in a given array.

## Setup & Usage ⚙️

1. Clone the Repository
   
Clone this project to your local machine:

```
git clone https://github.com/your-username/FILLLABSQUESTIONS.git
cd FILLLABSQUESTIONS
```

2. Install Dependencies
   
Use the go.mod file to download any necessary dependencies:

```
go mod tidy
```

3. Run the Project 🏃‍♂️
   
Run the main file to execute all the questions and see the outputs:

```
go run main.go
```

4. Run Tests ✅
   
Run all unit tests to verify the solutions:

```
go test ./...
```


