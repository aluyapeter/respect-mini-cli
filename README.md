# Respect CLI 🚀

A beginner-friendly command-line tool built with Go and Cobra that greets users by name — optionally in **shout mode**!

---

## ✨ Features

- Greet users by name: `--name` or `-n`
- Make your greeting loud with `--shout`
- Easy to install and use

---

## 🛠 Installation

To use this CLI, you need to have Go installed on your machine.

Then, run:

```bash
go install github.com/aluyapeter/respect-cli@latest
```

## 🚀 Usage

```bash
respect-cli greet
respect-cli greet --name Alice
respect-cli greet --name Alice --shout
respect-cli greet -n Alice --shout
```

| Command                                | Output      |
| -------------------------------------- | ----------- |
| `respect-cli greet`                    | Hello, !    |
| `respect-cli greet --name Bob`         | Hello, Bob! |
| `respect-cli greet --name Bob --shout` | HELLO, BOB! |

## ⚙️ Flags
--name, -n — Provide your name.
--shout — Print the greeting in uppercase.

## 🧠 Built With
Go
Cobra

## 🙌 Contributing
- Contributions are welcome! You can:

- Fork this repository

- Create your feature branch:
  ```bash
  git checkout -b my-new-feature

- Commit your changes:
  ```bash
  git commit -m 'Add new feature'

- Push to the branch:
  ```bash
  git push origin my-new-feature

- Open a pull request


