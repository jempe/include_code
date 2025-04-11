# Include Code
Small library to insert code of certain file in another file.

## Description

`include_code` is a Go library that allows you to include the content of one file into another file at specified locations. It searches for specific comment markers and replaces them with the content of the specified file.

## Usage

To use `include_code`, run the following command:

```sh
include_code <file>
```

## Example

Consider you have the following content in `main_file.js`:

```js
//This is the main file.
/*--include:include_file.js:--*/
/*--includeend--*/
//This is after the included file.
```

And the content of include_file.js is:

```js
//This is the content of the included file.
```

After running:
```sh
include_code main_file.txt
```

The `main_file.js` will be transformed into:

```js
//This is the main file.
/*--include:include_file.js:--*/
//This is the content of the included file.
/*--includeend--*/
//This is after the included file.
```

## Installation

To install the `include_code` library, use the following command:

```sh
go install github.com/jempe/include_code@latest
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or additions.

## License

This project is licensed under the Apache 2.0 License.
