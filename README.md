# PublicLLM Search Project

PublicLLM is a Go project that utilizes the Censys API to search for the specific text ("Ollama is running") in HTTP response bodies. 

## Installation

To use this project, you need to have Go installed on your system. Follow these steps to set up the project:

1. Clone the repository:
    ```bash
    git clone https://github.com/S1monlol/publicLLM
    ```

2. Navigate to the project directory:
    ```bash
    cd publicLLM
    ```

3. Install dependencies:
    ```bash
    go get github.com/go-resty/resty/v2
    go get github.com/joho/godotenv
    ```

4. Set up your environment variables:
    - Rename `.env.example` to `.env`.
    - Add your Censys API key in the `.env` file.

## Usage

Run the program with the following command:

```bash
go run .
```

## Notes

- Ensure that you have a valid API key from Censys for the program to work correctly.
- The program outputs a large amount of data, so you might want to redirect it to a file for easier analysis.

## Acknowledgements

This project uses the following open-source packages:
- Resty (https://github.com/go-resty/resty)
- Godotenv (https://github.com/joho/godotenv)

---
