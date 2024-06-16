# Ascii-Art-Web

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of Ascii-Art.

Ascii-art is a Golang program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.
<sub>by [Cyril R.](https://zone01normandie.org/git/cramanan) and [Antoine M.](https://zone01normandie.org/git/anmarvin)</sub>

In this project, you will be able to download your ascii-art using the download button.

This program can handle an input with:
   -   `numbers`
   -   `letters`
   -   `spaces`
   -   `special characters`
   -   and `breack lines`

## How to run it ?
To run Ascii-Art :
   * Install or Update the latest version of [Go](https://go.dev/doc/install)
   * Clone the repo (Make sure you have the required authorizations)
   * Open the folder in an IDE
   * Open your terminal
   * Copy one of these commands followed by your input (myInput here):
```
$ go run main.go
```

And open the [localhost html page](http://localhost:8070)

### How does it work ?
The program runs a simple algorithm which uses the desired font/.txt file and process it to create ASCII letters.
The letters in the file being 9 lines long, it retrieve every 1st line of the input then the 2nd line etc...

Then the server handle the input and displays it in another page.

If any error reach the server, the user gets redirected into an error page.