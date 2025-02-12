# Website Monitor

This is a simple Go program for website monitoring. It checks the availability of sites listed in a file and logs the results.

## How to Use

1. Compile and run the program:
   ```sh
   go run main.go
   ```
2. Choose an option from the menu:
   - `1` to start monitoring websites.
   - `2` to view logs.
   - `3` to exit the program.

## Configuration

- **Site List:** Add URLs to the `sites.txt` file, one per line.
- **Logs:** Monitoring records are saved in the `log.txt` file.

## Dependencies

- No external dependencies required. The program uses only the Go standard library.

## License

This project is free to use.

