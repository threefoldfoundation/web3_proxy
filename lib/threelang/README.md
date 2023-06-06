# Threelang Parser

## Usage

- For documentation on how to use Threelang, refer to this document [here](../../manual/src/threelang/parser.md)

## Development

- To add new books to the parser, follow these instructions:
  
  - Create a new module inside the threelang folder
  - Inside the new module, create a new handler for this book.
  - While creating a new Runner, the new handler should be initialized, then saved to the Runner's state.
  - The new handler should have its actions exposed in the Runner.run() method
  - The new handler must implement a handle_action method.
  - The handle_action method receives an actionparser.Action, and executes the action however it sees fit.
  - Handlers are responsible for logging their output, if any.
  - To add docuemntation on how to use the new book, create a new folder [here](../../manual/src/threelang/) with the book's name, and add all needed documentation files in this folder.
