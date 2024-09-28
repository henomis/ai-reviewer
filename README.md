# AI Reviewer

AI Reviewer is a Go application that uses OpenAI's GPT-4o model to review code snippets. It allows users to submit a code snippet for review and generates feedback based on the code.

## Getting Started

These instructions will help you to build and run the project on your local machine.

### Prerequisites

You need to have Go installed on your machine. You can download it from the [official website](https://golang.org/dl/).

### Installing

1. Clone the repository:
```bash
git clone https://github.com/yourusername/ai-reviewer.git
```

2. Navigate to the project directory
```bash
cd ai-reviewer
```

3. Build the project:
```bash
make build
```

4. Run the project:
```bash
export OPENAI_API_KEY=your_openai_api_key

./build/ai-reviewer -h
```

## Usage

You can provide the following flags to the application:
```bash
-verbose: Enable verbose mode (default: false)
-temp: OpenAI temperature (default: 0)
-model: OpenAI model (gpt-3.5, gpt-4, gpt-4o) (default: gpt-4)
-maxTokens: OpenAI max tokens (default: 2000)
```

You can use the application to review code snippets by providing the code as input:
```bash
git diff main | ./build/ai-reviewer > review.md
```


## AI

This project uses [OpenAI](https://openai.com/) and [LinGoose](https://github.com/henomis/lingoose) AI Go framework.
