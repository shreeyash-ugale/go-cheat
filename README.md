# Go Fiber Clipboard Manager

This project is a simple web application built with Go and the Fiber framework. It allows users to send data to the clipboard and view the current clipboard content through a web interface.

## Features

- Send data to the clipboard via a web form.
- View the current clipboard content.
- Copy the clipboard content to the system clipboard with a single click.


## Getting Started

### Prerequisites

- Go 1.22.2 or later
- [Fiber](https://gofiber.io/)
- [Clipboard](https://github.com/atotto/clipboard)

### Installation

1. Clone the repository:

```sh
git clone https://github.com/shreeyash-ugale/go-cheat.git
cd go-cheat
go mod tidy
go build && go-cheat
```

Allow go-cheat binary through firewall as we will access it from other device on the same LAN

### Endpoints

- `/` - Home page.
- `/send` - Form to send data to the clipboard.
- `/form` - Handles form submission and writes data to the clipboard.
- `/clip` - Displays the current clipboard content and provides a button to copy it to the clipboard.

### Usage

1. Make sure your mobile and laptop is on the same subnet
2. `ifconfig` for linux and mac `ipconfig` for Windows will provide you your internal IP, probably of the form `192.168.X.X`
3. Navigate to `http://<IP>:3000/clip` in your mobile to view clipboard content of laptop
4. Copy Question, Paste in You-know-what, get output
5. Navigate to `http://<IP>:3000/send` on your mobile to access the form.
6. Enter the data you want to send to the clipboard and submit the form.
7. If "could not paste" is seen, Before starting Drag and Drop (Allow Copy & Paste) to Bookmark Bar from the "Solution - Bookmarklet" section of [this website](https://www.adamsdesk.com/posts/enable-pasting-text-websites-block-it/). Once started, press esc to exit full-screen, click on the bookmark that was drag & dropped, click on enable-fullscreen to continue the test. Copy Paste should now work.

### TODO: Integrate Groq.com for Direct Question Input and Clipboard Update
