# Quilt HTML Pre-processor
Quilt (qlt) is a command-line HTML pre-processor that compiles .qtml files into HTML. It can run as a one-time compiler or watch for file changes and recompile automatically.

## Installation
To install Quilt globally, ensure you have Go installed on your system, then run:

```
go install github.com/brendan-mcmahon/Quilt@latest
```
Make sure your Go bin directory is in your PATH.

## Usage
### One-time Compilation

To compile .qtml files in a directory:
```
qlt [directory]
```
If no directory is specified, it uses the current working directory.

### Watch Mode
To start the file watcher and automatically recompile on changes:
```
qlt watch [directory]
```
Again, if no directory is specified, it watches the current working directory.

## File Structure
Quilt expects a specific file structure:
- `index.qtml`: The main file that will be compiled into `index.html`
- Other `.qtml` files: Can be included in the `index.qtml` file

## Syntax
Quilt uses a custom syntax for defining components and their attributes:

1. Component Tags:
   - Opening tag: `<[ComponentName]>`
   - Closing tag: `</[ComponentName]>`
   - Self-closing tag: `<[ComponentName]/>`

2. Component Attributes:
   - Syntax: `{attributeName}="value"`

## Example
Given the following files:

`index.qtml`:
```html
<!DOCTYPE html>
<html>
<head>
    <title>My Page</title>
</head>
<body>
    <[Header] {title}="Welcome"/>
    <main>
        <h1>Welcome to my page!</h1>
    </main>
    <[Footer] {year}="2024"/>
</body>
</html>
```
Header.qtml:
```html
<header>
    <h1>{title}</h1>
    <nav>
        <a href="/">Home</a>
        <a href="/about">About</a>
    </nav>
</header>
```
Footer.qtml:
```html
<footer>
    <p>&copy; {year} My Website</p>
</footer>
```
Running Quilt will create an index.html file with all the components resolved and attributes replaced:

index.html:
```html
<!DOCTYPE html>
<html>
  <head>
    <title>
      My Page
    </title>
  </head>
  <body>
    <main>
      <h1>
        Welcome to my page!
      </h1>
    </main>
  </body>
</html>
```